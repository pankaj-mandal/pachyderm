package m

import (
	"context"
	"strconv"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/pachyderm/pachyderm/v2/src/internal/log"
	"github.com/pachyderm/pachyderm/v2/src/internal/pctx"
)

func TestImmediate(t *testing.T) {
	ctx := pctx.TestContext(t)
	Set(ctx, "gauge", 42)
	Inc(ctx, "counter", 1)
	Inc(ctx, "counter", 1)
	Inc(ctx, "counter", -1)
	Sample(ctx, "sample", "a")
	Sample(ctx, "sample", "b")
}

func TestAggregatedGauge(t *testing.T) {
	//ctx, h := log.TestWithCapture(t)
	ctx := pctx.TestContext(t)
	ctx, c := context.WithCancel(ctx)

	doneCh := make(chan struct{})
	ctx = newAggregatedGauge(ctx, "test", 0, withDoneCh(doneCh)) // No log expected.
	Set(ctx, "test", 42)                                         // No log expected.
	for i := 0; i < 1000; i++ {
		Set(ctx, "test", 43) // No log expected.
	}
	Set(ctx, "test", "string") // Log expected because "string" doesn't match the gauge type.
	Inc(ctx, "test", 42)       // Log expected because this is a non-gauge operation.
	Sample(ctx, "test", 123)   // Log expected because this is a non-gauge operation.
	for i := 0; i < 1000; i++ {
		Set(ctx, "test", 1) // 1 log expected after flush.
	}
	c()
	<-doneCh

	// var got []any
	// for _, l := range h.Logs() {
	// 	if l.Message == "metric: test" {
	// 		if x, ok := l.Keys["value"]; ok {
	// 			got = append(got, x)
	// 		}
	// 		if x, ok := l.Keys["delta"]; ok {
	// 			got = append(got, x)
	// 		}
	// 		if x, ok := l.Keys["sample"]; ok {
	// 			got = append(got, x)
	// 		}
	// 	} else {
	// 		t.Errorf("unexpected log line: %v", l.String())
	// 	}
	// }
	// want := []any{"string", float64(42), float64(123), float64(1)} // float64 because of JSON -> Go conversion in the log history.
	// if diff := cmp.Diff(got, want); diff != "" {
	// 	t.Errorf("diff logged metrics (-got +want):\n%s", diff)
	// }
}

func TestAggregatedCounter(t *testing.T) {
	ctx, h := log.TestWithCapture(t)
	ctx, c := context.WithCancel(ctx)

	doneCh := make(chan struct{})
	ctx = newAggregatedCounter(ctx, "test", 0, withDoneCh(doneCh)) // No log expected.
	for i := 0; i < 10000; i++ {
		Inc(ctx, "test", 1) // 1 log expected after flush.
	}
	c()
	<-doneCh

	var got []any
	for _, l := range h.Logs() {
		if l.Message == "metric: test" {
			if x, ok := l.Keys["delta"]; ok {
				got = append(got, x)
			}
		} else {
			t.Errorf("unexpected log line: %v", l.String())
		}
	}
	want := []any{float64(10000)}
	if diff := cmp.Diff(got, want); diff != "" {
		t.Errorf("diff logged metrics (-got +want):\n%s", diff)
	}

}

func BenchmarkGauge(b *testing.B) {
	ctx, w := log.NewBenchLogger(true) // log rate limiting = true is the best case
	for i := 0; i < b.N; i++ {
		Set(ctx, "bench", i)
	}
	if w.Load() == 0 {
		b.Fatal("no bytes added to logger")
	}
}

func BenchmarkAggregatedGauge(b *testing.B) {
	ctx, w := log.NewBenchLogger(false)
	ctx, c := context.WithCancel(ctx)
	doneCh := make(chan struct{})
	ctx = newAggregatedGauge(ctx, "bench", 0, withDoneCh(doneCh))
	go func() {
		for i := 0; ; i++ {
			select {
			case <-ctx.Done():
				b.Logf("%d garbage aggregators created", i)
				return
			default:
				// contend with the lock in the aggregatedMetrics.gauges structure
				newAggregatedGauge(ctx, strconv.Itoa(i), 0)
			}
		}
	}()
	writesDoneCh := make(chan struct{})
	// Write metrics from 10 goroutines, so the writes contend with each other.
	for i := 0; i < 10; i++ {
		go func() {
			for j := 0; j < 1+b.N/10; j++ {
				Set(ctx, "bench", j)
			}
			writesDoneCh <- struct{}{}
		}()
	}
	for i := 0; i < 10; i++ {
		<-writesDoneCh
	}
	close(writesDoneCh)
	c()
	<-doneCh
	if w.Load() == 0 {
		b.Fatal("no bytes added to logger")
	}
}
