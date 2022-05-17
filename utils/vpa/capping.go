package vpa

import (
	"fmt"
	apiv1 "k8s.io/api/core/v1"
	vpa_types "k8s.io/autoscaler/vertical-pod-autoscaler/pkg/apis/autoscaling.k8s.io/v1"
	"k8s.io/autoscaler/vertical-pod-autoscaler/pkg/utils/limitrange"
)

type cappingAction string

var (
	cappedToMinAllowed             cappingAction = "capped to minAllowed"
	cappedToMaxAllowed             cappingAction = "capped to maxAllowed"
	cappedToLimit                  cappingAction = "capped to container limit"
	cappedProportionallyToMaxLimit cappingAction = "capped to fit Max in container LimitRange"
	cappedProportionallyToMinLimit cappingAction = "capped to fit Min in container LimitRange"
)

func toCappingAnnotation(resourceName apiv1.ResourceName, action cappingAction) string {
	return fmt.Sprintf("%s %s", resourceName, action)
}

type cappingRecommendationProcessor struct {
	limitsRangeCalculator limitrange.LimitRangeCalculator
}

/////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func NewCappingRecommendationProcessor(limitsRangeCalculator limitrange.LimitRangeCalculator) RecommendationProcessor {
	return nil
}

func (c *cappingRecommendationProcessor) Apply(podRecommendation *vpa_types.RecommendedPodResources, policy *vpa_types.PodResourcePolicy,
	conditions []vpa_types.VerticalPodAutoscalerCondition, pod *apiv1.Pod) (*vpa_types.RecommendedPodResources, ContainerToAnnotationsMap, error) {
	// capProportionallyToPodLimitRange()
	// for each of the containerRecommendation :
	// limitsRangeCalculator.GetContainerLimitRangeItem()
	// getCappedRecommendationForContainer()
	return nil, nil, nil
}

