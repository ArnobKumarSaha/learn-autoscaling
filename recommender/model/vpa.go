package model

import (
	autoscaling "k8s.io/api/autoscaling/v1"
	"k8s.io/apimachinery/pkg/labels"
	vpa_types "k8s.io/autoscaler/vertical-pod-autoscaler/pkg/apis/autoscaling.k8s.io/v1"
	"time"
)

// Map from VPA annotation key to value.
type vpaAnnotationsMap map[string]string

// Map from VPA condition type to condition.
type vpaConditionsMap map[vpa_types.VerticalPodAutoscalerConditionType]vpa_types.VerticalPodAutoscalerCondition

// Vpa (Vertical Pod Autoscaler) object is responsible for vertical scaling of
// Pods matching a given label selector.
type Vpa struct {
	ID VpaID
	// Labels selector that determines which Pods are controlled by this VPA
	// object. Can be nil, in which case no Pod is matched.
	PodSelector labels.Selector
	// Map of the object annotations (key-value pairs).
	Annotations vpaAnnotationsMap
	// Map of the status conditions (keys are condition types).
	Conditions vpaConditionsMap
	// Most recently computed recommendation. Can be nil.
	Recommendation *vpa_types.RecommendedPodResources
	// All container aggregations that contribute to this VPA.
	// TODO: Garbage collect old AggregateContainerStates.
	aggregateContainerStates aggregateContainerStatesMap
	// Pod Resource Policy provided in the VPA API object. Can be nil.
	ResourcePolicy *vpa_types.PodResourcePolicy
	// Initial checkpoints of AggregateContainerStates for containers.
	// The key is container name.
	ContainersInitialAggregateState ContainerNameToAggregateStateMap
	// UpdateMode describes how recommendations will be applied to pods
	UpdateMode *vpa_types.UpdateMode
	// Created denotes timestamp of the original VPA object creation
	Created time.Time
	// CheckpointWritten indicates when last checkpoint for the VPA object was stored.
	CheckpointWritten time.Time
	// IsV1Beta1API is set to true if VPA object has labelSelector defined as in v1beta1 api.
	IsV1Beta1API bool
	// TargetRef points to the controller managing the set of pods.
	TargetRef *autoscaling.CrossVersionObjectReference
	// PodCount contains number of live Pods matching a given VPA object.
	PodCount int
}

/////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func (conditionsMap *vpaConditionsMap) Set(conditionType vpa_types.VerticalPodAutoscalerConditionType, status bool, reason string, message string) *vpaConditionsMap {
	// just set a condition in vpa.Status
	return nil
}

func (conditionsMap *vpaConditionsMap) AsList() []vpa_types.VerticalPodAutoscalerCondition {
	// return a sorted list of conditions
	return nil
}

func (conditionsMap *vpaConditionsMap) ConditionActive(conditionType vpa_types.VerticalPodAutoscalerConditionType) bool {
	// check if the status of conditionType is true
	return false
}

func NewVpa(id VpaID, selector labels.Selector, created time.Time) *Vpa {
	return nil
}

func (vpa *Vpa) UseAggregationIfMatching(aggregationKey AggregateStateKey, aggregation *AggregateContainerState) {
	// check if this given aggregation matches this VPA
	// if yes, add it. < vpa.aggregateContainerStates[aggregationKey] = aggregation >
}

func (vpa *Vpa) UsesAggregation(aggregationKey AggregateStateKey) bool {
	// returns true iff an aggregation with the given key contributes to the VPA.
	return false
}

func (vpa *Vpa) matchesAggregation(aggregationKey AggregateStateKey) bool {
	// returns true iff the VPA matches the given aggregation key.
	return false
}

func (vpa *Vpa) DeleteAggregation(aggregationKey AggregateStateKey) {
	// state.MarkNotAutoscaled()
	// delete that from vpa.aggregateContainerStates
}

func (vpa *Vpa) UpdateRecommendation(recommendation *vpa_types.RecommendedPodResources) {
	// metrics_quality.ObserveRecommendationChange for each of containerRecommendations
	// vpa.Recommendation = recommendation
}

func (vpa *Vpa) HasRecommendation() bool {
	// if the VPA object contains any recommendation
	return false
}

func (vpa *Vpa) AggregateStateByContainerName() ContainerNameToAggregateStateMap {
	// calls AggregateStateByContainerName func from aggregate_container_state.go
	// vpa.MergeCheckpointedState()
	return nil
}

func (vpa *Vpa) MergeCheckpointedState(aggregateContainerStateMap ContainerNameToAggregateStateMap) {
	// merge vpa.ContainersInitialAggregateState with the given parameter
}

func (vpa *Vpa) SetResourcePolicy(resourcePolicy *vpa_types.PodResourcePolicy) {
	// update vpa.ResourcePolicy
	// calls AggregateContainerState's UpdateFromPolicy()
}

func (vpa *Vpa) SetUpdateMode(updatePolicy *vpa_types.PodUpdatePolicy) {
	// update vpa.UpdateMode & AggregateContainerState's UpdateMode
}

func (vpa *Vpa) UpdateConditions(podsMatched bool) {
	// set podsMatched condition & RecommendationProvided condition in the vpa status
}

func (vpa *Vpa) AsStatus() *vpa_types.VerticalPodAutoscalerStatus {
	// returns the vpaStatus, made from the vpa struct
	return nil
}

func (vpa *Vpa) HasMatchedPods() bool {
	// checks if any matching pods found for this vpa
	return false
}

/////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

// These are the main functions to look :
/*
>> called from model/cluster.go
UseAggregationIfMatching
SetResourcePolicy

>> called from routines/recommender.go
UpdateRecommendation
HasRecommendation
UpdateConditions

>> called from checkpoint/checkpoint_writer.go
AggregateStateByContainerName
*/
