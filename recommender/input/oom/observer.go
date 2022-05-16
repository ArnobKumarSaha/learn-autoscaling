package oom

import (
	apiv1 "k8s.io/api/core/v1"
	"k8s.io/autoscaler/vertical-pod-autoscaler/pkg/recommender/model"
	"k8s.io/client-go/tools/cache"
	"time"
)

// OomInfo contains data of the OOM event occurrence
type OomInfo struct {
	Timestamp   time.Time
	Memory      model.ResourceAmount
	ContainerID model.ContainerID
}

// Observer can observe pod resource update and collect OOM events.
type Observer interface {
	GetObservedOomsChannel() chan OomInfo
	OnEvent(*apiv1.Event)
	cache.ResourceEventHandler
}

// observer can observe pod resource update and collect OOM events.
type observer struct {
	observedOomsChannel chan OomInfo
}

/////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func NewObserver() *observer {
	return nil
}

func (o *observer) GetObservedOomsChannel() chan OomInfo {
	return o.observedOomsChannel
}

func (o *observer) OnEvent(event *apiv1.Event) {
	for _, oomInfo := range parseEvictionEvent(event) {
		o.observedOomsChannel <- oomInfo
	}
}

func parseEvictionEvent(event *apiv1.Event) []OomInfo {
	// If the event is not about Pod eviction, Or
	// if in AnnotationArray: offending_containers, offending_containers_usage or starved_resource length doesn't match one another : return
	//
	// For each of the offending containers who is starving for memory :
	// create OOMInfo by setting the fields appropriately
	return nil
}

func (o *observer) OnUpdate(oldObj, newObj interface{}) {
	// if lastTerminationReason is 'OOMKilled' for any of the container of newPod
	// if that container has more RestartCount than its old version
	// update o.observedOomsChannel <- by updated oomInfo
}

/////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
