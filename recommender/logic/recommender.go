package logic

import (
	"flag"
	"k8s.io/autoscaler/vertical-pod-autoscaler/pkg/recommender/model"
)

var (
	safetyMarginFraction = flag.Float64("recommendation-margin-fraction", 0.15, `Fraction of usage added as the safety margin to the recommended request`)
	podMinCPUMillicores  = flag.Float64("pod-recommendation-min-cpu-millicores", 25, `Minimum CPU recommendation for a pod`)
	podMinMemoryMb       = flag.Float64("pod-recommendation-min-memory-mb", 250, `Minimum memory recommendation for a pod`)
)

// PodResourceRecommender computes resource recommendation for a Vpa object.
type PodResourceRecommender interface {
	GetRecommendedPodResources(containerNameToAggregateStateMap model.ContainerNameToAggregateStateMap) RecommendedPodResources
}

// RecommendedPodResources is a Map from container name to recommended resources.
type RecommendedPodResources map[string]RecommendedContainerResources

// RecommendedContainerResources is the recommendation of resources for a
// container.
type RecommendedContainerResources struct {
	// Recommended optimal amount of resources.
	Target model.Resources
	// Recommended minimum amount of resources.
	LowerBound model.Resources
	// Recommended maximum amount of resources.
	UpperBound model.Resources
}

type podResourceRecommender struct {
	targetEstimator     ResourceEstimator
	lowerBoundEstimator ResourceEstimator
	upperBoundEstimator ResourceEstimator
}

/////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func CreatePodResourceRecommender() PodResourceRecommender {
	// it returns a PercentileEstimator with safetyMargin & confidenceMultiplier
	// defaults :
	// targetPercentile = 0.9, lowerBoundPercentile = 0.5, upperBoundPercentile = 0.95
	// safetyMargin = 0.15 // resources will be scaled up with this margin amount
	// upperBound => multiplier = 1.0, exponent = 1.0
	// lowerBound => multiplier = 0.001, exponent = -2.0 // it means, the lower bound is multiplied by the factor (1 + 0.001/history-length-in-days)^-2
	return nil
}

func (r *podResourceRecommender) GetRecommendedPodResources(containerNameToAggregateStateMap model.ContainerNameToAggregateStateMap) RecommendedPodResources {
	// make a podResourceRecommender struct with MinResources
	// call recommender.estimateContainerResources() for each of the aggregationStateMap, & set that as recommendation
	return nil
}

func (r *podResourceRecommender) estimateContainerResources(s *model.AggregateContainerState) RecommendedContainerResources {
	// calls FilterControlledResources for target, lower & upperBound estimator
	return RecommendedContainerResources{}
}

func FilterControlledResources(estimation model.Resources, controlledResources []model.ResourceName) model.Resources {
	// exclude the items from 'estimation' who are not in 'controlledResources'
	return nil
}

/////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

/*
CreatePodResourceRecommender & GetRecommendedPodResources are called from routines/recommender.go file
*/