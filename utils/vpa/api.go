package vpa

import (
	core "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/labels"
	vpa_types "k8s.io/autoscaler/vertical-pod-autoscaler/pkg/apis/autoscaling.k8s.io/v1"
	vpa_clientset "k8s.io/autoscaler/vertical-pod-autoscaler/pkg/client/clientset/versioned"
	vpa_api "k8s.io/autoscaler/vertical-pod-autoscaler/pkg/client/clientset/versioned/typed/autoscaling.k8s.io/v1"
	vpa_lister "k8s.io/autoscaler/vertical-pod-autoscaler/pkg/client/listers/autoscaling.k8s.io/v1"
)

// VpaWithSelector is a pair of VPA and its selector.
type VpaWithSelector struct {
	Vpa      *vpa_types.VerticalPodAutoscaler
	Selector labels.Selector
}

type patchRecord struct {
	Op    string      `json:"op,inline"`
	Path  string      `json:"path,inline"`
	Value interface{} `json:"value"`
}

/////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func NewVpasLister(vpaClient *vpa_clientset.Clientset, stopChannel <-chan struct{}, namespace string) vpa_lister.VerticalPodAutoscalerLister {
	return nil
}

func UpdateVpaStatusIfNeeded(vpaClient vpa_api.VerticalPodAutoscalerInterface, vpaName string, newStatus,
	oldStatus *vpa_types.VerticalPodAutoscalerStatus) (result *vpa_types.VerticalPodAutoscaler, err error) {
	// build a patchRecord struct, & call patchVpa if some difference found between old & newStatus
	return nil, nil
}

func GetControllingVPAForPod(pod *core.Pod, vpas []*VpaWithSelector) *VpaWithSelector {
	// PodMatchesVPA(), PodLabelsMatchVPA(), stronger() are the utility functions called from here
	return nil
}

func GetUpdateMode(vpa *vpa_types.VerticalPodAutoscaler) vpa_types.UpdateMode {
	// simple getter
	return ""
}

func GetContainerControlledValues(name string, vpaResourcePolicy *vpa_types.PodResourcePolicy) vpa_types.ContainerControlledValues {
	// simple getter
	return ""
}

func GetContainerResourcePolicy(containerName string, policy *vpa_types.PodResourcePolicy) *vpa_types.ContainerResourcePolicy {
	// just finds & returns a particular containerPolicy
	return nil
}

func CreateOrUpdateVpaCheckpoint(vpaCheckpointClient vpa_api.VerticalPodAutoscalerCheckpointInterface, vpaCheckpoint *vpa_types.VerticalPodAutoscalerCheckpoint) error {
	// traditional create-or-patch on VPACheckPoint object
	return nil
}

/////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

/*
UpdateVpaStatusIfNeeded called from routines/recommender.go
PodLabelsMatchVPA called from model/cluster.go
CreateOrUpdateVpaCheckpoint called from checkpoint_writer.go
*/
