package spec

import (
	"k8s.io/api/core/v1"
	"k8s.io/autoscaler/vertical-pod-autoscaler/pkg/recommender/model"
	v1lister "k8s.io/client-go/listers/core/v1"
)

// BasicPodSpec contains basic information defining a pod and its containers.
type BasicPodSpec struct {
	// ID identifies a pod within a cluster.
	ID model.PodID
	// Labels of the pod. It is used to match pods with certain VPA opjects.
	PodLabels map[string]string
	// List of containers within this pod.
	Containers []BasicContainerSpec
	// PodPhase describing current life cycle phase of the Pod.
	Phase v1.PodPhase
}

// BasicContainerSpec contains basic information defining a container.
type BasicContainerSpec struct {
	// ID identifies the container within a cluster.
	ID model.ContainerID
	// Name of the image running within the container.
	Image string
	// Currently requested resources for this container.
	Request model.Resources
}

//SpecClient provides information about pods and containers Specification
type SpecClient interface {
	// Returns BasicPodSpec for each pod in the cluster
	GetPodSpecs() ([]*BasicPodSpec, error)
}

type specClient struct {
	podLister v1lister.PodLister
}

/////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func NewSpecClient(podLister v1lister.PodLister) SpecClient {
	return nil
}

func (client *specClient) GetPodSpecs() ([]*BasicPodSpec, error) {
	// calls newBasicPodSpec, newContainerSpecs, newContainerSpec & calculateRequestedResources are called internally
	return nil, nil
}

/////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
