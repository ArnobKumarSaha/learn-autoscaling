package routines

import (
	"context"
	"flag"
	"github.com/Arnobkumarsaha/learn-autoscaling/recommender/checkpoint"
	"github.com/Arnobkumarsaha/learn-autoscaling/recommender/input"
	"github.com/Arnobkumarsaha/learn-autoscaling/recommender/logic"
	vpa_types "k8s.io/autoscaler/vertical-pod-autoscaler/pkg/apis/autoscaling.k8s.io/v1"
	vpa_api "k8s.io/autoscaler/vertical-pod-autoscaler/pkg/client/clientset/versioned/typed/autoscaling.k8s.io/v1"
	controllerfetcher "k8s.io/autoscaler/vertical-pod-autoscaler/pkg/recommender/input/controller_fetcher"
	"k8s.io/autoscaler/vertical-pod-autoscaler/pkg/recommender/model"
	"k8s.io/client-go/rest"
	"time"
)

var (
	checkpointsWriteTimeout = flag.Duration("checkpoints-timeout", time.Minute, `Timeout for writing checkpoints since the start of the recommender's main loop`)
	minCheckpointsPerRun    = flag.Int("min-checkpoints", 10, "Minimum number of checkpoints to write per recommender's main loop")
	memorySaver             = flag.Bool("memory-saver", false, `If true, only track pods which have an associated VPA`)
)

// Recommender recommend resources for certain containers, based on utilization periodically got from metrics api.
type Recommender interface {
	// RunOnce performs one iteration of recommender duties followed by update of recommendations in VPA objects.
	RunOnce()
	// GetClusterState returns ClusterState used by Recommender
	GetClusterState() *model.ClusterState
	// GetClusterStateFeeder returns ClusterStateFeeder used by Recommender
	GetClusterStateFeeder() input.ClusterStateFeeder
	// UpdateVPAs computes recommendations and sends VPAs status updates to API Server
	UpdateVPAs()
	// MaintainCheckpoints stores current checkpoints in API Server and garbage collect old ones
	// MaintainCheckpoints writes at least minCheckpoints if there are more checkpoints to write.
	// Checkpoints are written until ctx permits or all checkpoints are written.
	MaintainCheckpoints(ctx context.Context, minCheckpoints int)
}

type recommender struct {
	clusterState                  *model.ClusterState
	clusterStateFeeder            input.ClusterStateFeeder
	checkpointWriter              checkpoint.CheckpointWriter
	checkpointsGCInterval         time.Duration
	controllerFetcher             controllerfetcher.ControllerFetcher
	lastCheckpointGC              time.Time
	vpaClient                     vpa_api.VerticalPodAutoscalersGetter
	podResourceRecommender        logic.PodResourceRecommender
	useCheckpoints                bool
	lastAggregateContainerStateGC time.Time
}

/////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func NewRecommender(config *rest.Config, checkpointsGCInterval time.Duration, useCheckpoints bool, namespace string) Recommender {
	// controllerfetcher.NewControllerFetcher, input.NewClusterStateFeeder, checkpoint.NewCheckpointWriter & logic.CreatePodResourceRecommender are called
	return nil
}

func (r *recommender) GetClusterState() *model.ClusterState {
	return nil
}

func (r *recommender) GetClusterStateFeeder() input.ClusterStateFeeder {
	return nil
}

func (r *recommender) RunOnce() {
	// call clusterFeeder's LoadVPAs(), LoadPods(), LoadRealTimeMetrics()
	// UpdateVPAs()
	// MaintainCheckpoints()
	// cluster's RateLimitedGarbageCollectAggregateCollectionStates()
}

func (r *recommender) UpdateVPAs() {
	// for each of the observed VPAs :
	// GetContainerNameToAggregateStateMap() & recommender.GetRecommendedPodResources
	// getCappedRecommendation(), vpa.UpdateRecommendation()
	// vpa.UpdateConditions()
	// cluster's RecordRecommendation()
	// UpdateVpaStatusIfNeeded()
}

func getCappedRecommendation(vpaID model.VpaID, resources logic.RecommendedPodResources, policy *vpa_types.PodResourcePolicy) *vpa_types.RecommendedPodResources {
	// converts the resources to vpa-CRD-status style
	// calls ApplyVPAPolicy, to get a recommendation adjusted to obey policy.
	return nil
}

func (r *recommender) MaintainCheckpoints(ctx context.Context, minCheckpointsPerRun int) {
	//  checkpointWriter.StoreCheckpoints()
	// call cluster_feeder's GarbageCollectCheckpoints() if needed
}

/////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

/*
RunOnce is the main function of this file which will be called from main.go
*/
