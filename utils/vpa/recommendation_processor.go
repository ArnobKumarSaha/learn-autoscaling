package vpa

import (
	"k8s.io/api/core/v1"
	vpa_types "k8s.io/autoscaler/vertical-pod-autoscaler/pkg/apis/autoscaling.k8s.io/v1"
)

// ContainerToAnnotationsMap contains annotations per container.
type ContainerToAnnotationsMap = map[string][]string

// RecommendationProcessor post-processes recommendation adjusting it to limits and environment context
type RecommendationProcessor interface {
	// Apply processes and updates recommendation for given pod, based on container limits,
	// VPA policy and possibly other internal RecommendationProcessor context.
	// Must return a non-nil pointer to RecommendedPodResources or error.
	Apply(podRecommendation *vpa_types.RecommendedPodResources,
		policy *vpa_types.PodResourcePolicy,
		conditions []vpa_types.VerticalPodAutoscalerCondition,
		pod *v1.Pod) (*vpa_types.RecommendedPodResources, ContainerToAnnotationsMap, error)
}
