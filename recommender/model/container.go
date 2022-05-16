package model

import (
	corev1 "k8s.io/api/core/v1"
	"time"
)

// ContainerUsageSample is a measure of resource usage of a container over some
// interval.
type ContainerUsageSample struct {
	// Start of the measurement interval.
	MeasureStart time.Time
	// Average CPU usage in cores or memory usage in bytes.
	Usage ResourceAmount
	// CPU or memory request at the time of measurment.
	Request ResourceAmount
	// Which resource is this sample for.
	Resource ResourceName
}

// ContainerState stores information about a single container instance.
// Each ContainerState has a pointer to the aggregation that is used for
// aggregating its usage samples.
// It holds the recent history of CPU and memory utilization.
//   Note: samples are added to intervals based on their start timestamps.
type ContainerState struct {
	// Current request.
	Request Resources
	// Start of the latest CPU usage sample that was aggregated.
	LastCPUSampleStart time.Time
	// Max memory usage observed in the current aggregation interval.
	memoryPeak ResourceAmount
	// Max memory usage estimated from an OOM event in the current aggregation interval.
	oomPeak ResourceAmount
	// End time of the current memory aggregation interval (not inclusive).
	WindowEnd time.Time
	// Start of the latest memory usage sample that was aggregated.
	lastMemorySampleStart time.Time
	// Aggregation to add usage samples to.
	aggregator ContainerStateAggregator
}

/////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func NewContainerState(request Resources, aggregator ContainerStateAggregator) *ContainerState {
	return nil
}

func (sample *ContainerUsageSample) isValid(expectedResource ResourceName) bool {
	// check if this sample is for our expected resource-type, & its usage is calculated
	return false
}

func (container *ContainerState) AddSample(sample *ContainerUsageSample) bool {
	// calls either addCPUSample() or addMemorySample() based on sample's resource type
	return false
}

func (container *ContainerState) addCPUSample(sample *ContainerUsageSample) bool {
	// Discard invalid, duplicate or out-of-order samples.
	// observeQualityMetrics()
	// container.aggregator.AddSample(sample)
	return false
}

func (container *ContainerState) addMemorySample(sample *ContainerUsageSample, isOOM bool) bool {
	// each container aggregates one peak per aggregation interval
	// If the timestamp of the current sample is earlier than the end of the current interval (WindowEnd):
	//   if current sample is larger the current peak : remove the old one, add current one
	// otherwise, add a new interval
	//
	// observeQualityMetrics()
	// adding a sample means , calling container.aggregator.addSample()
	return false
}

func (container *ContainerState) GetMaxMemoryPeak() ResourceAmount {
	// return max of memoryPeak & oomPeak of the container
	return 0
}

func (container *ContainerState) RecordOOM(timestamp time.Time, requestedMemory ResourceAmount) error {
	// discard old oom
	// call addMemorySample() with isOOM true
	return nil
}

func (container *ContainerState) observeQualityMetrics(usage ResourceAmount, isOOM bool, resource corev1.ResourceName) {
	// if there is no recommendation, or if recommended resource is 0, metrics_quality.ObserveQualityMetricsRecommendationMissing() & return
	// else metrics_quality.ObserveQualityMetrics()
}

/////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

// These are the main functions to look :
/*
>> called from model/cluster.go
AddSample
RecordOOM
*/
