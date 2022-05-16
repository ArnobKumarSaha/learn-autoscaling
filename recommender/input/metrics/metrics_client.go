package metrics

import (
	k8sapiv1 "k8s.io/api/core/v1"
	"k8s.io/metrics/pkg/apis/metrics/v1beta1"
	"time"

	"k8s.io/autoscaler/vertical-pod-autoscaler/pkg/recommender/model"
	resourceclient "k8s.io/metrics/pkg/client/clientset/versioned/typed/metrics/v1beta1"
)

// ContainerMetricsSnapshot contains information about usage of certain container within defined time window.
type ContainerMetricsSnapshot struct {
	// ID identifies a specific container those metrics are coming from.
	ID model.ContainerID
	// End time of the measurement interval.
	SnapshotTime time.Time
	// Duration of the measurement interval, which is [SnapshotTime - SnapshotWindow, SnapshotTime].
	SnapshotWindow time.Duration
	// Actual usage of the resources over the measurement interval.
	Usage model.Resources
}

// MetricsClient provides simple metrics on resources usage on containter level.
type MetricsClient interface {
	// GetContainersMetrics returns an array of ContainerMetricsSnapshots,
	// representing resource usage for every running container in the cluster
	GetContainersMetrics() ([]*ContainerMetricsSnapshot, error)
}

type metricsClient struct {
	metricsGetter resourceclient.PodMetricsesGetter
	namespace     string
}

/////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func NewMetricsClient(metricsGetter resourceclient.PodMetricsesGetter, namespace string) MetricsClient {
	return nil
}

func (c *metricsClient) GetContainersMetrics() ([]*ContainerMetricsSnapshot, error) {
	// for each of the listed PodMetrics : call createContainerMetricsSnapshots()
	return nil, nil
}

func createContainerMetricsSnapshots(podMetrics v1beta1.PodMetrics) []*ContainerMetricsSnapshot {
	// just a plural version of newContainerMetricsSnapshot()
	return nil
}

func newContainerMetricsSnapshot(containerMetrics v1beta1.ContainerMetrics, podMetrics v1beta1.PodMetrics) *ContainerMetricsSnapshot {
	// create & return a ContainerMetricsSnapshot struct by properly set the fields
	return nil
}

func calculateUsage(containerUsage k8sapiv1.ResourceList) model.Resources {
	// just converts the resource type
	return nil
}

/////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
