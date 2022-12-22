package testutil

import (
	"context"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/pachyderm/pachyderm/v2/src/internal/backoff"
	"github.com/pachyderm/pachyderm/v2/src/internal/errors"
	"github.com/pachyderm/pachyderm/v2/src/internal/require"

	apps "k8s.io/api/apps/v1"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"
	kube "k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

var (
	zero int64
)

// GetKubeClient connects to the Kubernetes API server either from inside the
// cluster or from a test binary running on a machine with kubectl (it will
// connect to the same cluster as kubectl)
func GetKubeClient(t testing.TB) *kube.Clientset {
	var config *rest.Config
	var err error
	host := os.Getenv("KUBERNETES_SERVICE_HOST")
	if host != "" {
		config, err = rest.InClusterConfig()
	} else {
		rules := clientcmd.NewDefaultClientConfigLoadingRules()
		kubeConfig := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(rules,
			&clientcmd.ConfigOverrides{})
		config, err = kubeConfig.ClientConfig()
	}
	require.NoError(t, err)
	k, err := kube.NewForConfig(config)
	require.NoError(t, err)
	return k
}

// Deletes a pod and then blocks until there is a delete event, and the pod is recreated and running.
// However, the pod may not be in a "ready" state when control is unblocked. Use WaitForFirstPodReady
// to block until the pod to begins accepting requests.
func DeletePod(t testing.TB, app, ns string) {
	kubeClient := GetKubeClient(t)
	podList, err := kubeClient.CoreV1().Pods(ns).List(
		context.Background(),
		metav1.ListOptions{
			LabelSelector: metav1.FormatLabelSelector(metav1.SetAsLabelSelector(
				map[string]string{"app": app, "suite": "pachyderm"},
			)),
		})
	require.NoError(t, err)
	require.Equal(t, 1, len(podList.Items))
	require.NoError(t, kubeClient.CoreV1().Pods(ns).Delete(
		context.Background(),
		podList.Items[0].ObjectMeta.Name, metav1.DeleteOptions{}))

	// Make sure the pod goes down
	startTime := time.Now()
	require.NoError(t, backoff.Retry(func() error {
		podList, err := kubeClient.CoreV1().Pods(ns).List(
			context.Background(),
			metav1.ListOptions{
				LabelSelector: metav1.FormatLabelSelector(metav1.SetAsLabelSelector(
					map[string]string{"app": app, "suite": "pachyderm"},
				)),
			})
		waitForPodEvent(t, app, ns, watch.Deleted)
		if err != nil {
			return errors.EnsureStack(err)
		}

		if len(podList.Items) == 0 {
			return nil
		}
		if time.Since(startTime) > 60*time.Second {
			return nil
		}
		return errors.Errorf("waiting for old %v pod to be killed", app)
	}, backoff.NewTestingBackOff()))

	// Make sure the pod comes back up
	require.NoErrorWithinTRetry(t, 30*time.Second, func() error {
		podList, err := kubeClient.CoreV1().Pods(ns).List(
			context.Background(),
			metav1.ListOptions{
				LabelSelector: metav1.FormatLabelSelector(metav1.SetAsLabelSelector(
					map[string]string{"app": app, "suite": "pachyderm"},
				)),
			})
		if err != nil {
			return errors.EnsureStack(err)
		}
		if len(podList.Items) == 0 {
			return errors.Errorf("no %v pod up yet", app)
		}
		return nil
	})

	require.NoErrorWithinTRetry(t, 120*time.Second, func() error {
		podList, err := kubeClient.CoreV1().Pods(ns).List(
			context.Background(),
			metav1.ListOptions{
				LabelSelector: metav1.FormatLabelSelector(metav1.SetAsLabelSelector(
					map[string]string{"app": app, "suite": "pachyderm"},
				)),
			})
		if err != nil {
			return errors.EnsureStack(err)
		}
		if len(podList.Items) == 0 {
			return errors.Errorf("no %v pod up yet", app)
		}
		if podList.Items[0].Status.Phase != v1.PodRunning {
			return errors.Errorf("%v not running yet", app)
		}
		return nil
	})
}

// DeletePipelineRC deletes the RC belonging to the pipeline 'pipeline'. This
// can be used to test PPS's robustness
func DeletePipelineRC(t testing.TB, pipeline, namespace string) {
	kubeClient := GetKubeClient(t)
	rcs, err := kubeClient.CoreV1().ReplicationControllers(namespace).List(
		context.Background(),
		metav1.ListOptions{
			LabelSelector: metav1.FormatLabelSelector(metav1.SetAsLabelSelector(
				map[string]string{"pipelineName": pipeline},
			)),
		})
	require.NoError(t, err)
	require.Equal(t, 1, len(rcs.Items))
	require.NoError(t, kubeClient.CoreV1().ReplicationControllers(namespace).Delete(
		context.Background(),
		rcs.Items[0].ObjectMeta.Name, metav1.DeleteOptions{
			GracePeriodSeconds: &zero,
		}))
	require.NoErrorWithinTRetry(t, 30*time.Second, func() error {
		rcs, err := kubeClient.CoreV1().ReplicationControllers(namespace).List(
			context.Background(),
			metav1.ListOptions{
				LabelSelector: metav1.FormatLabelSelector(metav1.SetAsLabelSelector(
					map[string]string{"pipelineName": pipeline},
				)),
			})
		if err != nil {
			return errors.EnsureStack(err)
		}
		if len(rcs.Items) != 0 {
			return errors.Errorf("RC %q not deleted yet", pipeline)
		}
		return nil
	})
}

// PachdDeployment finds the corresponding deployment for pachd in the
// kubernetes namespace and returns it.
func PachdDeployment(t testing.TB, namespace string) *apps.Deployment {
	return GetDeployment(t, "pachd", namespace)
}

// GetDeployment finds the deployment with the corresponding app name like "pachd"
func GetDeployment(t testing.TB, appName string, namespace string) *apps.Deployment {
	k := GetKubeClient(t)
	result, err := k.AppsV1().Deployments(namespace).Get(context.Background(), appName, metav1.GetOptions{})
	require.NoError(t, err)
	return result
}

// podRunning AndReady takes a watch event returned by v1.PodInterface.Watch and
// returns true iff the pod is in Running phase, all of the pod's containers statuses are Ready,
// and the pod condition "Ready" is true.
func podRunningAndReady(e watch.Event) (bool, error) {
	if e.Type == watch.Deleted {
		return false, errors.New("received DELETE while watching pods")
	}
	pod, ok := e.Object.(*v1.Pod)
	if !ok {
		return false, errors.Errorf("unexpected object type in watch.Event")
	}
	containersReady := true
	for _, cs := range pod.Status.ContainerStatuses {
		if !cs.Ready {
			containersReady = false
			break
		}
	}
	podReady := true
	for _, c := range pod.Status.Conditions {
		if c.Type == v1.PodReady && c.Status != v1.ConditionTrue {
			podReady = false
			break
		}
	}
	return pod.Status.Phase == v1.PodRunning && containersReady && podReady, nil
}

// WaitForPachdReady finds the pachd pods within the kubernetes namespace and
// blocks until they are all ready.
// WaitForPachdReady finds the pachd pods within the kubernetes namespace and
// blocks until they are all ready.
func WaitForPachdReady(t testing.TB, namespace string) {
	WaitForDeploymentReady(t, "pachd", namespace)
}

func WaitForDeploymentReady(t testing.TB, appName string, namespace string) {
	k := GetKubeClient(t)
	deployment := GetDeployment(t, appName, namespace)
	for {
		newDeployment, err := k.AppsV1().Deployments(namespace).Get(context.Background(), deployment.Name, metav1.GetOptions{})
		require.NoError(t, err)
		if newDeployment.Status.ObservedGeneration >= deployment.Generation && newDeployment.Status.Replicas == *newDeployment.Spec.Replicas {
			break
		}
		time.Sleep(time.Second * 5)
	}
	watch, err := k.CoreV1().Pods(namespace).Watch(context.Background(), metav1.ListOptions{
		LabelSelector: fmt.Sprintf("app=%s", appName),
	})
	defer watch.Stop()
	require.NoError(t, err)
	readyPods := make(map[string]bool)
	for event := range watch.ResultChan() {
		ready, err := podRunningAndReady(event)
		require.NoError(t, err)
		if ready {
			pod, ok := event.Object.(*v1.Pod)
			if !ok {
				t.Fatal("event.Object should be an object")
			}
			readyPods[pod.Name] = true
			if len(readyPods) == int(*deployment.Spec.Replicas) {
				break
			}
		}
	}
}

// Wait deployment only works for deployments and checks every pod.
// This function is simpler and just unblocks as soon as 1 pod is read.
// It can be used on non-Deployment objects like StatefulSets.
func WaitForFirstPodReady(t testing.TB, appName string, namespace string) {
	k := GetKubeClient(t)
	watch, err := k.CoreV1().Pods(namespace).Watch(context.Background(), metav1.ListOptions{
		LabelSelector: fmt.Sprintf("app=%s", appName),
	})
	defer watch.Stop()
	require.NoError(t, err)
	for event := range watch.ResultChan() {
		ready, err := podRunningAndReady(event)
		require.NoError(t, err)
		if ready {
			break
		}
	}
}

func waitForPodEvent(t testing.TB, appName string, namespace string, eventType watch.EventType) {
	k := GetKubeClient(t)
	watch, err := k.CoreV1().Pods(namespace).Watch(context.Background(), metav1.ListOptions{
		LabelSelector: fmt.Sprintf("app=%s", appName),
	})
	defer watch.Stop()
	require.NoError(t, err)
	for event := range watch.ResultChan() {
		if event.Type == eventType {
			break
		}
	}
}
