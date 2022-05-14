package input

import (
	"k8s.io/apimachinery/pkg/labels"
	vpa_types "k8s.io/autoscaler/vertical-pod-autoscaler/pkg/apis/autoscaling.k8s.io/v1"
	corev1 "k8s.io/client-go/kubernetes/typed/core/v1"
	"k8s.io/client-go/rest"
	"time"

	vpa_api "k8s.io/autoscaler/vertical-pod-autoscaler/pkg/client/clientset/versioned/typed/autoscaling.k8s.io/v1"
	vpa_lister "k8s.io/autoscaler/vertical-pod-autoscaler/pkg/client/listers/autoscaling.k8s.io/v1"
	controllerfetcher "k8s.io/autoscaler/vertical-pod-autoscaler/pkg/recommender/input/controller_fetcher"
	"k8s.io/autoscaler/vertical-pod-autoscaler/pkg/recommender/input/history"
	"k8s.io/autoscaler/vertical-pod-autoscaler/pkg/recommender/input/metrics"
	"k8s.io/autoscaler/vertical-pod-autoscaler/pkg/recommender/input/oom"
	"k8s.io/autoscaler/vertical-pod-autoscaler/pkg/recommender/input/spec"
	"k8s.io/autoscaler/vertical-pod-autoscaler/pkg/recommender/model"
	"k8s.io/autoscaler/vertical-pod-autoscaler/pkg/target"
	kube_client "k8s.io/client-go/kubernetes"
	v1lister "k8s.io/client-go/listers/core/v1"
)

const (
	evictionWatchRetryWait                     = 10 * time.Second
	evictionWatchJitterFactor                  = 0.5
	scaleCacheLoopPeriod         time.Duration = 7 * time.Second
	scaleCacheEntryLifetime      time.Duration = time.Hour
	scaleCacheEntryFreshnessTime time.Duration = 10 * time.Minute
	scaleCacheEntryJitterFactor  float64       = 1.
	defaultResyncPeriod          time.Duration = 10 * time.Minute
	defaultRecommenderName                     = "default"
)

// ClusterStateFeeder can update state of ClusterState object.
type ClusterStateFeeder interface {
	// InitFromHistoryProvider loads historical pod spec into clusterState.
	InitFromHistoryProvider(historyProvider history.HistoryProvider)

	// InitFromCheckpoints loads historical checkpoints into clusterState.
	InitFromCheckpoints()

	// LoadVPAs updates clusterState with current state of VPAs.
	LoadVPAs()

	// LoadPods updates clusterState with current specification of Pods and their Containers.
	LoadPods()

	// LoadRealTimeMetrics updates clusterState with current usage metrics of containers.
	LoadRealTimeMetrics()

	// GarbageCollectCheckpoints removes historical checkpoints that don't have a matching VPA.
	GarbageCollectCheckpoints()
}

// ClusterStateFeederFactory makes instances of ClusterStateFeeder.
type ClusterStateFeederFactory struct {
	ClusterState        *model.ClusterState
	KubeClient          kube_client.Interface
	MetricsClient       metrics.MetricsClient
	VpaCheckpointClient vpa_api.VerticalPodAutoscalerCheckpointsGetter
	VpaLister           vpa_lister.VerticalPodAutoscalerLister
	PodLister           v1lister.PodLister
	OOMObserver         oom.Observer
	SelectorFetcher     target.VpaTargetSelectorFetcher
	MemorySaveMode      bool
	ControllerFetcher   controllerfetcher.ControllerFetcher
}

type clusterStateFeeder struct {
	coreClient          corev1.CoreV1Interface
	specClient          spec.SpecClient
	metricsClient       metrics.MetricsClient
	oomChan             <-chan oom.OomInfo
	vpaCheckpointClient vpa_api.VerticalPodAutoscalerCheckpointsGetter
	vpaLister           vpa_lister.VerticalPodAutoscalerLister
	clusterState        *model.ClusterState
	selectorFetcher     target.VpaTargetSelectorFetcher
	memorySaveMode      bool
	controllerFetcher   controllerfetcher.ControllerFetcher
}

type condition struct {
	conditionType vpa_types.VerticalPodAutoscalerConditionType
	delete        bool
	message       string
}

/////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func NewClusterStateFeeder(config *rest.Config, clusterState *model.ClusterState, memorySave bool, namespace string) ClusterStateFeeder {
	// it calls some utility functions of this file : newMetricsClient, NewPodListerAndOOMObserver, WatchEvictionEventsWithRetries
	return nil
}

func (feeder *clusterStateFeeder) InitFromHistoryProvider(historyProvider history.HistoryProvider) {
	// get the cluster history
	// run feeder.ClusterState's AddOrUpdatePod for all pods of clusterHistory
	// run feeder.ClusterState's AddOrUpdateContainer for all pods of clusterHistory
	// run feeder.ClusterState's AddSample for all the container samples
}

func (feeder *clusterStateFeeder) InitFromCheckpoints() {
	// feeder.leadVPAs()
	// feeder.setVpaCheckpoint() for all the vpa checkpoints
}

func (feeder *clusterStateFeeder) setVpaCheckpoint(checkpoint *vpa_types.VerticalPodAutoscalerCheckpoint) error {
	// make new AggregationContainerState & call LoadFromCheckpoint() for it
	return nil
}

func (feeder *clusterStateFeeder) GarbageCollectCheckpoints() {
	// feeder.leadVPAs()
	// For all the checkpoints, if its referred VPA doesn't exist in feeder.clusterState.Vpas, delete that
}

func (feeder *clusterStateFeeder) LoadVPAs() {
	// for all the VPAs :
	// getSelector(), try clusterState.AddOrUpdateVPA
	// if successful, update feeder.clusterState.Vpas[vpaID].Conditions accordingly
	//
	// call feeder.clusterState.DeleteVpa if previous AddOrUpdate was unsuccessful
}

func filterVPAs(feeder *clusterStateFeeder, allVpaCRDs []*vpa_types.VerticalPodAutoscaler) []*vpa_types.VerticalPodAutoscaler {
	// Filter VPA objects whose specified recommender names are not 'default'
	// implicitDefaultRecommender & selectsRecommender are the two utility called from here
	return nil
}

func (feeder *clusterStateFeeder) getSelector(vpa *vpa_types.VerticalPodAutoscaler) (labels.Selector, []condition) {
	// fetch vpa selector
	// validateTargetRef()
	// return selector & condition based on the validation result
	return nil, nil
}

func (feeder *clusterStateFeeder) validateTargetRef(vpa *vpa_types.VerticalPodAutoscaler) (bool, condition) {
	// checks if vpa.Spec.TargetRef is the topLevel controller(who doesn't have any controller) or not
	return false, condition{}
}

func (feeder *clusterStateFeeder) LoadPods() {
	// call feeder.clusterState.DeletePod if a pod doesn't exist
	// try AddOrUpdatePod & AddOrUpdateContainer for the pods & containers if matchesVPA() returns true
}

func (feeder *clusterStateFeeder) matchesVPA(pod *spec.BasicPodSpec) bool {
	// check if this pod matches any of the VPAs
	return false
}

func (feeder *clusterStateFeeder) LoadRealTimeMetrics() {
	// for each of the resourceSample<cpu,memory> of all the ContainerMetrics :
	// try feeder.clusterState.AddSample
	// call feeder.clusterState.RecordOOM in a for-select loop
	// call metrics_recommender.RecordAggregateContainerStatesCount()
}

func newContainerUsageSamplesWithKey(metrics *metrics.ContainerMetricsSnapshot) []*model.ContainerUsageSampleWithKey {
	// changes the metrics format to 'ContainerUsageSample' format <of container.go>
	return nil
}

/////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

// These are the main functions to look :
/*
>> called from main.go
InitFromHistoryProvider
InitFromCheckpoints

>> called from routines/recommender.go
LoadVPAs
LoadPods
GarbageCollectCheckpoints
LoadRealTimeMetrics
*/
