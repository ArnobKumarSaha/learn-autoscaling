package routines

import (
	"k8s.io/autoscaler/vertical-pod-autoscaler/pkg/recommender/model"
)

/////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func GetContainerNameToAggregateStateMap(vpa *model.Vpa) model.ContainerNameToAggregateStateMap {
	// vpa.AggregateStateByContainerName()
	// get the container resourcePolicy & call aggregatedContainerState.UpdateFromPolicy() if autoScaling enabled
	return nil
}
