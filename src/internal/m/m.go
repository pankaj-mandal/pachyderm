// Package m implements lightweight metrics.  A metric is a name, a set of key/value pairs (set by
// the underlying logger), and a value.  A value can be an absolute value ("gauge"), an incremental
// value ("counter"), or a set of samples (like "histogram", though lossless in this
// implementation).  Metrics are most useful when additional code analyzes the entire log of a
// program's run; this is called "analysis code" or "anaylsis software" throughout the
// documentation.
//
// Values are stored by writing the set of operations to build the value to the logs.  For example,
// each time a sample is added to a sampler with Sample, a log line is produced.  Reading the logs
// will allow analysis software to recover the entire set of samples.  Counters are similar; each
// increment event emits a log line, and and analysis code can add the deltas to see the final
// value, or the value at a particular point in time.
//
// Values are logically associated with key=value metadata.  Each value is uniquely identified by
// the set of key=value metadata; a metric foo with field foo=bar is logically a different value
// from the metric foo with field foo=baz.  The metadata is defined by the fields applied to the
// underlying logger.  For example, if you declare a counter `tx_bytes` on each HTTP request you
// handle, each time the count of transmitted bytes is updated, the log message will include fields
// inherited from the default HTTP logger, which includes a per-request ID.  Therefore, analysis
// code can calculate a tx_byte count for every request handled by the system by observing the value
// of the x-request-id field on log lines matching the tx_byte counter format.  And, of course, it
// can ignore the extra fields and add everything up to show a program-wide count of bytes
// transmitted.
//
// Normally it's considered "too expensive" to store metrics with such a high cardinality (per-user,
// per-IP, per-request), but this system has no such limitation.  Cardinality has no impact on write
// performance.  Analysis code can make the decision on which fields to discard to decrease the
// cardinality for long-term storage, if desired.
//
// Aggregating each operations on a metric over time recovers the value of the metric at a
// particular time.  Any sort of smartness or validation comes from the reader, not from this
// writing code.  If you want to treat a certain metric name as a string gauge, integer counter, and
// sampler of proto messages, that is OK.  The analysis code that processes the logs will need to be
// ready for that, or at least ready to ignore values it doesn't think are valid.
//
// To emit metrics, simply call these public functions in this package:
//
//	Set(ctx, "metric_name", value) // Set the current value of the metric to value.
//	Inc(ctx, "metric_name", delta) // Increment the current value of the metric by delta.
//	Sample(ctx, "metric_name", sample) // Add sample to the set of samples captured by the metric.
//
// It is safe to write to the same metric from different goroutines concurrently.
//
// Some minimal aggregation can be done in-process.  This is a compromise to reduce the "noise" in
// the logs.  If you were uploading a 1GB file, it would make sense to increment the byte count
// metric 1 billion times, as each byte of the file is passed to the TCP stack.  This, however,
// would be very noisy and make the logs difficult to read.  So, we offer aggregates to flush the
// value of gauges (Set) and counters (Inc) periodically, grouping many value-changing operations
// into a single log message.
//
// An aggregator can be registered on a context (with pctx.Child), causing all future writes to that
// metric on that context to be aggregated.  (The code that writes the metric need not be aware of
// the aggregated nature; the above public API automatically does the right thing.)

// Aggregated metrics are emitted to the logs based on a time interval set at registration time.  If
// a write occurs, and the metric hasn't been logged for that interval, then a log line will be
// produced showing the current value of the metric.  If unflushed data exists when the context is
// Done, a log line is also emitted.  From an analysis standpoint, nothing changes; the emitted
// aggregated metrics are indistinguishable from immediately emitted metrics.
//
// Aggregated metrics can only be one type of metric with one type of value; if you create a counter
// named tx_bytes and then call Set on it, no aggregation will be done on that data.  (Inc calls
// continue to be aggregated; it doesn't "break" the aggregation to accidentally call the wrong
// value-changing function.)  Similarly, the type of the value is set at registration time; calls
// writing values of a different type will not be aggregated, but again, do not break any ongoing
// aggregation.
package m

import (
	"context"
	"math"
	"sync"
	"time"

	"github.com/pachyderm/pachyderm/v2/src/internal/log"
	"go.uber.org/zap"
	"golang.org/x/exp/constraints"
)

// aggregatedMetricsKey is the context key storing metrics aggregators.
type aggregatedMetricsKey struct{}

// aggregatedMetrics is an object that stores aggregated metrics in the context.  This allows a new
// metric to begin being aggregated without creating a new child context.
type aggregatedMetrics struct {
	gauges   sync.Map // map from metric name to *gauge[T].
	counters sync.Map // map from metric name to *counter[T].
}

// valueType returns a type hint for parsers that need to know the value type to synthesize a metric
// in some other stricter monitoring system.  The parser will have to choose what it wants to do
// when different set/increment/sample operations emit different types; we don't care, but they
// might.
func valueType(val any) string {
	switch val.(type) {
	case int, uint, int8, uint8, int16, uint16, int32, uint32, int64, uint64:
		return "int"
	case float32, float64:
		return "float"
	case string, []byte:
		// rune is an alias for int32, byte is an alias for uint8, so those types can't be
		// (single-character) strings
		return "string"

	default:
		// TODO(jonathan): Consider dereferencing pointers in the emitter methods.  Inc(...,
		// &int(42)) should increment something by 42.  Inc(..., nil) should increment it by
		// 0.
		return "any"
	}
}

// Set sets the value of a metric.
func Set[T any](ctx context.Context, metric string, val T) {
	g := gaugeFromContext[T](ctx, metric)
	if g != nil {
		g.set(ctx, val)
		return
	}
	logGauge(ctx, metric, val)
}

func logGauge[T any](ctx context.Context, metric string, val T) {
	log.Debug(ctx, "metric: "+metric, zap.String("metric", metric), zap.String("type", valueType(val)), zap.Any("value", val))
}

// Monoid is a constraint matching types that have an "empty" value and "append" operation.
// Consider an integer; 0 is "empty", and "append" is addition.  Any Monoid can be a Counter metric
// value.
type Monoid interface {
	constraints.Integer | constraints.Float | constraints.Complex
}

// Inc changes the value of a metric by the provided delta.
func Inc[T Monoid](ctx context.Context, metric string, delta T) {
	c := counterFromContext[T](ctx, metric)
	if c != nil {
		c.inc(ctx, delta)
		return
	}
	// In immediate mode, a delta of 0 is a no-op; the absence of a log line also indicates that
	// the counter was incremented by 0.  In aggregate mode above, that's not quite true; it is
	// potentially an opportunity to flush the aggregated value to the logs.
	if delta == 0 {
		return
	}
	logCounter(ctx, metric, delta)
}

func logCounter[T Monoid](ctx context.Context, metric string, delta T) {
	log.Debug(ctx, "metric: "+metric, zap.String("metric", metric), zap.String("type", valueType(delta)), zap.Any("delta", delta))
}

// Sample adds a sample to the value of the metric.
func Sample[T any](ctx context.Context, metric string, val T) {
	log.Debug(ctx, "metric: "+metric, zap.String("metric", metric), zap.String("type", valueType(val)), zap.Any("sample", val))
}

// aggregateOption contains optional parameters for customizing an aggregating metric.
type aggregateOptions struct {
	flushInterval time.Duration // How long to wait, at a minimum, before reporting the value of the metric.
	doneCh        chan struct{} // only for testing
}

// Option supplies optional configuration to aggregated metrics.
type Option func(o *aggregateOptions)

// WithFlushInterval is an Option that sets the amount of time to aggregate a metric for before emitting.
func WithFlushInterval(interval time.Duration) Option {
	return func(o *aggregateOptions) {
		o.flushInterval = interval
	}
}

// Deferred sets a metric to be aggregated until the underlying context is Done.
func Deferred() Option {
	return func(o *aggregateOptions) {
		o.flushInterval = time.Duration(math.MaxInt64)
	}
}

// withDoneCh sets a channel to be closed when a metric is flushed for the last time.  It's only
// used for testing.
func withDoneCh(ch chan struct{}) Option {
	return func(o *aggregateOptions) {
		o.doneCh = ch
	}
}

const defaultFlushInterval = 10 * time.Second

// aggregatedMetric is a metric that emits on write only if flushInterval has passed.
type aggregatedMetric struct {
	sync.Mutex
	aggregateOptions           // Config options.
	metric           string    // The name of the metric.
	dirty            bool      // Whether or not the metric needs to be flushed.
	last             time.Time // When the metric was last flushed.
}

func (m *aggregatedMetric) init(metric string, options []Option) {
	m.Lock()
	defer m.Unlock()
	m.metric = metric
	m.last = time.Now()
	m.flushInterval = defaultFlushInterval
	for _, opt := range options {
		opt(&m.aggregateOptions)
	}
}

type gauge[T any] struct {
	aggregatedMetric
	value T
}

func (g *gauge[T]) flush(ctx context.Context, now bool) {
	g.Lock()
	defer g.Unlock()
	if g.dirty && (now || time.Since(g.last) > g.flushInterval) {
		logGauge(ctx, g.metric, g.value)
		g.dirty = false
		g.last = time.Now()
	}
}

func (g *gauge[T]) set(ctx context.Context, val T) {
	if g == nil {
		return
	}
	g.Lock()
	g.value = val
	g.dirty = true
	g.Unlock()
	g.flush(ctx, false)
}

func gaugeFromContext[T any](ctx context.Context, metric string) *gauge[T] {
	if m, ok := ctx.Value(aggregatedMetricsKey{}).(*aggregatedMetrics); ok {
		if maybeG, ok := m.gauges.Load(metric); ok {
			if g, ok := maybeG.(*gauge[T]); ok {
				return g
			}
		}
	}
	return nil
}

// cloneAggregatedMetrics creates a context with a new aggregatedMetrics-containing value context.
// this is necessary when creating a child context with new logging fields, so that aggregation is
// always unique per set of field=value pairs.
func cloneAggregatedMetrics(ctx context.Context) context.Context {
	// TODO(jonathan): Implement.
	return ctx
}

func newAggregatedGauge[T any](ctx context.Context, metric string, zero T, options ...Option) context.Context {
	m, ok := ctx.Value(aggregatedMetricsKey{}).(*aggregatedMetrics)
	if !ok {
		m = new(aggregatedMetrics)
		ctx = context.WithValue(ctx, aggregatedMetricsKey{}, m)
	}
	g := new(gauge[T])
	g.init(metric, options)
	go func() {
		<-ctx.Done()
		g.flush(ctx, true)
		if g.doneCh != nil {
			close(g.doneCh)
		}
	}()
	m.gauges.Store(metric, g)
	return ctx
}

type counter[T Monoid] struct {
	aggregatedMetric
	delta T
}

func (c *counter[T]) flush(ctx context.Context, now bool) {
	c.Lock()
	defer c.Unlock()
	if c.dirty && (now || time.Since(c.last) > c.flushInterval) {
		logCounter(ctx, c.metric, c.delta)
		c.dirty = false
		c.last = time.Now()
		c.delta = 0
	}
}

func (c *counter[T]) inc(ctx context.Context, val T) {
	if c == nil {
		return
	}
	c.Lock()
	c.delta += val // val == 0 means to try flushing
	c.dirty = true
	c.Unlock()
	c.flush(ctx, false)
}

func counterFromContext[T Monoid](ctx context.Context, metric string) *counter[T] {
	if m, ok := ctx.Value(aggregatedMetricsKey{}).(*aggregatedMetrics); ok {
		if maybeC, ok := m.counters.Load(metric); ok {
			if c, ok := maybeC.(*counter[T]); ok {
				return c
			}
		}
	}
	return nil
}

func NewAggregatedCounter[T Monoid](ctx context.Context, metric string, zero T, options ...Option) context.Context {
	m, ok := ctx.Value(aggregatedMetricsKey{}).(*aggregatedMetrics)
	if !ok {
		m = new(aggregatedMetrics)
		ctx = context.WithValue(ctx, aggregatedMetricsKey{}, m)
	}
	c := new(counter[T])
	c.init(metric, options)
	go func() {
		<-ctx.Done()
		c.flush(ctx, true)
		if c.doneCh != nil {
			close(c.doneCh)
		}
	}()
	m.counters.Store(metric, c)
	return ctx
}
