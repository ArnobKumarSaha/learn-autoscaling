package logic

import (
	"k8s.io/autoscaler/vertical-pod-autoscaler/pkg/recommender/model"
)

// ResourceEstimator is a function from AggregateContainerState to
// model.Resources, e.g. a prediction of resources needed by a group of
// containers.
type ResourceEstimator interface {
	GetResourceEstimation(s *model.AggregateContainerState) model.Resources
}

// Implementation of ResourceEstimator that returns constant amount of
// resources. This can be used as by a fake recommender for test purposes.
type constEstimator struct {
	resources model.Resources
}

// Simple implementation of the ResourceEstimator interface. It returns specific
// percentiles of CPU usage distribution and memory peaks distribution.
type percentileEstimator struct {
	cpuPercentile    float64
	memoryPercentile float64
}

type marginEstimator struct {
	marginFraction float64
	baseEstimator  ResourceEstimator
}

type minResourcesEstimator struct {
	minResources  model.Resources
	baseEstimator ResourceEstimator
}

type confidenceMultiplier struct {
	multiplier    float64
	exponent      float64
	baseEstimator ResourceEstimator
}

/////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func NewConstEstimator(resources model.Resources) ResourceEstimator {
	return &constEstimator{resources}
}

func (e *constEstimator) GetResourceEstimation(s *model.AggregateContainerState) model.Resources {
	return e.resources
}

func NewPercentileEstimator(cpuPercentile float64, memoryPercentile float64) ResourceEstimator {
	return &percentileEstimator{cpuPercentile, memoryPercentile}
}

func (e *percentileEstimator) GetResourceEstimation(s *model.AggregateContainerState) model.Resources {
	// Returns specific percentiles of CPU and memory peaks distributions.
	// it uses histogram.Percentile() for calculation
	return nil
}

func WithMargin(marginFraction float64, baseEstimator ResourceEstimator) ResourceEstimator {
	return &marginEstimator{marginFraction, baseEstimator}
}

func (e *marginEstimator) GetResourceEstimation(s *model.AggregateContainerState) model.Resources {
	// calculate e.baseEstimator.GetResourceEstimation(s)
	// then scaleUp all with the marginFraction
	return nil
}

func WithMinResources(minResources model.Resources, baseEstimator ResourceEstimator) ResourceEstimator {
	// calculate e.baseEstimator.GetResourceEstimation(s)
	// then set the minResources if resourceAmount is less than minResource
	return nil
}

func WithConfidenceMultiplier(multiplier, exponent float64, baseEstimator ResourceEstimator) ResourceEstimator {
	// getConfidence()
	// calculate e.baseEstimator.GetResourceEstimation(s)
	// then scaleUp all resources with (1.+e.multiplier/confidence)^e.exponent
	return nil
}

func getConfidence(s *model.AggregateContainerState) float64 {
	// returns the minimum of lifeSpanDays & amount
	// where, lifeSpanDays = Distance between the first and the last observed sample time, measured in days.
	// & imaginaryDays = number of days if the frequency was 1 sample-per-minute
	//
	// suppose, we are running estimator for 3 days, & taking 2 samples per minute.
	// so, lifeSpanDays=3, totalSamples = (2*60*24*3), imaginaryDays = totalSamples/(60*24) = 6
	// confidence is min(3,6) = 3
	return 0
}

/////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

/*
estimation related functions are called from recommender.go internally
*/