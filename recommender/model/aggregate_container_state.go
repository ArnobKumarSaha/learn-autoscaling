package model

import (
	corev1 "k8s.io/api/core/v1"
	vpa_types "k8s.io/autoscaler/vertical-pod-autoscaler/pkg/apis/autoscaling.k8s.io/v1"
	"k8s.io/autoscaler/vertical-pod-autoscaler/pkg/recommender/util"
	"time"
)

// ContainerNameToAggregateStateMap maps a container name to AggregateContainerState
// that aggregates state of containers with that name.
type ContainerNameToAggregateStateMap map[string]*AggregateContainerState

// ContainerStateAggregator is an interface for objects that consume and
// aggregate container usage samples.
type ContainerStateAggregator interface {
	// AddSample aggregates a single usage sample.
	AddSample(sample *ContainerUsageSample)
	// SubtractSample removes a single usage sample. The subtracted sample
	// should be equal to some sample that was aggregated with AddSample()
	// in the past.
	SubtractSample(sample *ContainerUsageSample)
	// GetLastRecommendation returns last recommendation calculated for this
	// aggregator.
	GetLastRecommendation() corev1.ResourceList
	// NeedsRecommendation returns true if this aggregator should have
	// a recommendation calculated.
	NeedsRecommendation() bool
	// GetUpdateMode returns the update mode of VPA controlling this aggregator,
	// nil if aggregator is not autoscaled.
	GetUpdateMode() *vpa_types.UpdateMode
}

// AggregateContainerState holds input signals aggregated from a set of containers.
// It can be used as an input to compute the recommendation.
// The CPU and memory distributions use decaying histograms by default
// (see NewAggregateContainerState()).
// Implements ContainerStateAggregator interface.
type AggregateContainerState struct {
	// AggregateCPUUsage is a distribution of all CPU samples.
	AggregateCPUUsage util.Histogram
	// AggregateMemoryPeaks is a distribution of memory peaks from all containers:
	// each container should add one peak per memory aggregation interval (e.g. once every 24h).
	AggregateMemoryPeaks util.Histogram
	// Note: first/last sample timestamps as well as the sample count are based only on CPU samples.
	FirstSampleStart  time.Time
	LastSampleStart   time.Time
	TotalSamplesCount int
	CreationTime      time.Time

	// Following fields are needed to correctly report quality metrics
	// for VPA. When we record a new sample in an AggregateContainerState
	// we want to know if it needs recommendation, if the recommendation
	// is present and if the automatic updates are on (are we able to
	// apply the recommendation to the pods).
	LastRecommendation  corev1.ResourceList
	IsUnderVPA          bool
	UpdateMode          *vpa_types.UpdateMode
	ScalingMode         *vpa_types.ContainerScalingMode
	ControlledResources *[]ResourceName
}

// ContainerStateAggregatorProxy is a wrapper for ContainerStateAggregator
// that creates ContainerStateAgregator for container if it is no longer
// present in the cluster state.
type ContainerStateAggregatorProxy struct {
	containerID ContainerID
	cluster     *ClusterState
}

// AggregateContainerState & ContainerStateAggregatorProxy both implements ContainerStateAggregator interface
// calls from container.go first comes to the Proxy then the Proxy calls AggregateContainerState's methods

/////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

// simple getters ::

func (a *AggregateContainerState) GetLastRecommendation() corev1.ResourceList {
	return a.LastRecommendation
}

func (a *AggregateContainerState) NeedsRecommendation() bool {
	return a.IsUnderVPA && a.ScalingMode != nil && *a.ScalingMode != vpa_types.ContainerScalingModeOff
}

func (a *AggregateContainerState) GetUpdateMode() *vpa_types.UpdateMode {
	return a.UpdateMode
}

func (a *AggregateContainerState) GetScalingMode() *vpa_types.ContainerScalingMode {
	return a.ScalingMode
}

func (a *AggregateContainerState) GetControlledResources() []ResourceName {
	if a.ControlledResources != nil {
		return *a.ControlledResources
	}
	return []ResourceName{ResourceCPU, ResourceMemory}
}

// actual functions ::

func (a *AggregateContainerState) MarkNotAutoscaled() {
	// set IsUnderVPA, LastRecommendation, UpdateMode, ScalingMode, ControlledResources = nil
}

func (a *AggregateContainerState) MergeContainerState(other *AggregateContainerState) {
	// calls historam.Merge() function for CPU & Memory
	// set FirstSampleStart, LastSampleStart, TotalSamplesCount
}

func NewAggregateContainerState() *AggregateContainerState {
	return nil
}

func (a *AggregateContainerState) AddSample(sample *ContainerUsageSample) {
	// if resource is CPU, calls addCPUSample()
	// if memory , calls historam.AddSample()
}

func (a *AggregateContainerState) SubtractSample(sample *ContainerUsageSample) {
	// similar to AddSample, but
	// if CPU resource substracting is not possible yet
}

func (a *AggregateContainerState) addCPUSample(sample *ContainerUsageSample) {
	// calls histogram.AddSample
	// set FirstSampleStart, LastSampleStart, TotalSamplesCount
}

func (a *AggregateContainerState) SaveToCheckpoint() (*vpa_types.VerticalPodAutoscalerCheckpointStatus, error) {
	// histogram.SaveToCheckpoint()
	// return the VPACheckpointStatus object
	return nil, nil
}

func (a *AggregateContainerState) LoadFromCheckpoint(checkpoint *vpa_types.VerticalPodAutoscalerCheckpointStatus) error {
	// Set AggregateContainerState's fields from checkpoint fields
	// histogram.LoadFromCheckpoint()
	return nil
}

func (a *AggregateContainerState) isExpired(now time.Time) bool {
	// check if the diff between now & lastSampleStart is <= AggregationWindowLength
	return false
}

func (a *AggregateContainerState) isEmpty() bool {
	// check if TotalSamplesCount == 0
	return false
}

func (a *AggregateContainerState) UpdateFromPolicy(resourcePolicy *vpa_types.ContainerResourcePolicy) {
	// updates containerState's scaling mode and controlledResources based on resourcePolicy of the VPA
}

func AggregateStateByContainerName(aggregateContainerStateMap aggregateContainerStatesMap) ContainerNameToAggregateStateMap {
	// converts a map[AggregateStateKey]*AggregateContainerState to map[containerName]*AggregateContainerState
	// by grouping them by the container name
	return nil
}

// methods of ContainerStateAggregatorProxy
// it has all the methods of ContainerStateAggregator interface, as it implements that.

// each of the methods has same structure like below:
/*
func (p *ContainerStateAggregatorProxy) AddSample(sample *ContainerUsageSample) {
	aggregator := p.cluster.findOrCreateAggregateContainerState(p.containerID)
	aggregator.AddSample(sample)
}
*/

/////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

// These are the main functions to look :
/*
>> called from model/container.go
AddSample
SubtractSample
GetLastRecommendation
NeedsRecommendation
GetUpdateMode

>> called from model/vpa.go
UpdateFromPolicy
MarkNotAutoscaled
MergeContainerState

>> called from input/cluster_feeder.go
LoadFromCheckpoint

>> called from checkpoint/checkpoint_writer.go
SaveToCheckpoint
*/
