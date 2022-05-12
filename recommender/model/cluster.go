package model

import (
	apiv1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/labels"
	vpa_types "k8s.io/autoscaler/vertical-pod-autoscaler/pkg/apis/autoscaling.k8s.io/v1"
	"time"
)

// ClusterState holds all runtime information about the cluster required for the
// VPA operations, i.e. configuration of resources (pods, containers,
// VPA objects), aggregated utilization of compute resources (CPU, memory) and
// events (container OOMs).
// All input to the VPA Recommender algorithm lives in this structure.
type ClusterState struct {
	// Pods in the cluster.
	Pods map[PodID]*PodState
	// VPA objects in the cluster.
	Vpas map[VpaID]*Vpa
	// VPA objects in the cluster that have no recommendation mapped to the first
	// time we've noticed the recommendation missing or last time we logged
	// a warning about it.
	EmptyVPAs map[VpaID]time.Time
	// Observed VPAs. Used to check if there are updates needed.
	ObservedVpas []*vpa_types.VerticalPodAutoscaler

	// All container aggregations where the usage samples are stored.
	aggregateStateMap aggregateContainerStatesMap
	// Map with all label sets used by the aggregations. It serves as a cache
	// that allows to quickly access labels.Set corresponding to a labelSetKey.
	labelSetMap labelSetMap

	lastAggregateContainerStateGC time.Time
	gcInterval                    time.Duration
}

// AggregateStateKey determines the set of containers for which the usage samples
// are kept aggregated in the model.
type AggregateStateKey interface {
	Namespace() string
	ContainerName() string
	Labels() labels.Labels
}

// String representation of the labels.LabelSet. This is the value returned by
// labelSet.String(). As opposed to the LabelSet object, it can be used as a map key.
type labelSetKey string

// Map of label sets keyed by their string representation.
type labelSetMap map[labelSetKey]labels.Set

// AggregateContainerStatesMap is a map from AggregateStateKey to AggregateContainerState.
type aggregateContainerStatesMap map[AggregateStateKey]*AggregateContainerState

// PodState holds runtime information about a single Pod.
type PodState struct {
	// Unique id of the Pod.
	ID PodID
	// Set of labels attached to the Pod.
	labelSetKey labelSetKey
	// Containers that belong to the Pod, keyed by the container name.
	Containers map[string]*ContainerState
	// PodPhase describing current life cycle phase of the Pod.
	Phase apiv1.PodPhase
}

// ContainerUsageSampleWithKey holds a ContainerUsageSample together with the
// ID of the container it belongs to.
type ContainerUsageSampleWithKey struct {
	ContainerUsageSample
	Container ContainerID
}

// Implementation of the AggregateStateKey interface. It can be used as a map key.
type aggregateStateKey struct {
	namespace     string
	containerName string
	labelSetKey   labelSetKey
	// Pointer to the global map from labelSetKey to labels.Set.
	// Note: a pointer is used so that two copies of the same key are equal.
	labelSetMap *labelSetMap
}

/////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func NewClusterState() {}

func (cluster *ClusterState) AddOrUpdatePod(podID PodID, newLabels labels.Set, phase apiv1.PodPhase) {
	// if cluster.pods doesn't contain the pod with given podID {
	//		create a new podState, & assign into that
	//      if the labels of the pod have changed, it updates the links between the containers and the aggregations
	//      for each of the container < cluster.findOrCreateAggregateContainerState(containerID) >
	//      cluster.addPodToItsVpa(pod)
	// }
	// else {
	//      cluster.removePodFromItsVpa(pod)
	// }
}

func (cluster *ClusterState) getLabelSetKey(labelSet labels.Set) labelSetKey {
	// cluster.labelSetMap[labelSetKey] = labelSet's-stringify-version
	return ""
}

func (cluster *ClusterState) addPodToItsVpa(pod *PodState) {
	// if podLabel matches vpa (from cluster.VPAs), vpa.PodCount++
}

func (cluster *ClusterState) removePodFromItsVpa(pod *PodState) {
	// if podLabel matches vpa (from cluster.VPAs), vpa.PodCount--
}

func (cluster *ClusterState) findOrCreateAggregateContainerState(containerID ContainerID) *AggregateContainerState {
	// key = cluster.aggregateStateKeyForContainerID(containerID)
	// if key exists in cluster.aggregateStateMap , return that
	// otherwise, NewAggregateContainerState() &
	// vpa.UseAggregationIfMatching() for all cluster.VPAs // TODO
	return nil
}

func (cluster *ClusterState) aggregateStateKeyForContainerID(containerID ContainerID) AggregateStateKey {
	// pod = cluster.Pods[containerID.PodID]
	// cluster.MakeAggregateStateKey(pod, containerID.ContainerName)
	return nil
}

func (cluster *ClusterState) MakeAggregateStateKey(pod *PodState, containerName string) AggregateStateKey {
	// returns a new AggregateStateKey with given data
	return nil
}

func (cluster *ClusterState) DeletePod(podID PodID) {
	// cluster.removePodFromItsVpa(pod)
	// delete podId from cluster.pods
}

func (cluster *ClusterState) GetContainer(containerID ContainerID) *ContainerState {
	// pod = cluster.Pods[containerID.PodID]
	// get container object from pod.Containers
	return nil
}

func (cluster *ClusterState) AddOrUpdateContainer(containerID ContainerID, request Resources) error {
	// if the pod, which contains containerID, exists in cluster.Pods ->
	// get the containerState & assign the request to it
	//
	// otherwise,
	// cluster.findOrCreateAggregateContainerState(containerID)
	// set the NewContainerState(request) in pod.Containers
	return nil
}

