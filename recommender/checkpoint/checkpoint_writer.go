package checkpoint

import (
	"context"
	vpa_api "k8s.io/autoscaler/vertical-pod-autoscaler/pkg/client/clientset/versioned/typed/autoscaling.k8s.io/v1"
	"k8s.io/autoscaler/vertical-pod-autoscaler/pkg/recommender/model"
	"time"
)

// CheckpointWriter persistently stores aggregated historical usage of containers
// controlled by VPA objects. This state can be restored to initialize the model after restart.
type CheckpointWriter interface {
	// StoreCheckpoints writes at least minCheckpoints if there are more checkpoints to write.
	// Checkpoints are written until ctx permits or all checkpoints are written.
	StoreCheckpoints(ctx context.Context, now time.Time, minCheckpoints int) error
}

type checkpointWriter struct {
	vpaCheckpointClient vpa_api.VerticalPodAutoscalerCheckpointsGetter
	cluster             *model.ClusterState
}

/////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func NewCheckpointWriter(cluster *model.ClusterState, vpaCheckpointClient vpa_api.VerticalPodAutoscalerCheckpointsGetter) CheckpointWriter {
	return nil
}

func (writer *checkpointWriter) StoreCheckpoints(ctx context.Context, now time.Time, minCheckpoints int) error {
	// getVpasToCheckpoint()
	// buildAggregateContainerStateMap()
	// for each of the aggregateContainerState :
	// SaveToCheckpoint()
	// construct the checkPoint object
	// CreateOrUpdateVpaCheckpoint()
	return nil
}

func getVpasToCheckpoint(clusterVpas map[model.VpaID]*model.Vpa) []*model.Vpa {
	// filter the VPAs who is currently not fetching history
	// sort them according to lastCheckpointWritten time
	return nil
}

func buildAggregateContainerStateMap(vpa *model.Vpa, cluster *model.ClusterState, now time.Time) map[string]*model.AggregateContainerState {
	// for all the containers who are connected to any of the VPAs :
	// call subtractCurrentContainerMemoryPeak
	return nil
}

func subtractCurrentContainerMemoryPeak(a *model.AggregateContainerState, container *model.ContainerState, now time.Time) {
	// call histogram.SubtractSample() to subtract the MaxMemoryPeak
}

/////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

/*
NewCheckpointWriter & StoreCheckpoints are called from routines/recommender.go
*/
