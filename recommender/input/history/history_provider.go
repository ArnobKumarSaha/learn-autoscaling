package history

import (
	prometheusv1 "github.com/prometheus/client_golang/api/prometheus/v1"
	prommodel "github.com/prometheus/common/model"
	"k8s.io/autoscaler/vertical-pod-autoscaler/pkg/recommender/model"
	"time"
)

// PrometheusHistoryProviderConfig allow to select which metrics
// should be queried to get real resource utilization.
type PrometheusHistoryProviderConfig struct {
	Address                                          string
	QueryTimeout                                     time.Duration
	HistoryLength, HistoryResolution                 string
	PodLabelPrefix, PodLabelsMetricName              string
	PodNamespaceLabel, PodNameLabel                  string
	CtrNamespaceLabel, CtrPodNameLabel, CtrNameLabel string
	CadvisorMetricsJobName                           string
	Namespace                                        string
}

// PodHistory represents history of usage and labels for a given pod.
type PodHistory struct {
	// Current samples if pod is still alive, last known samples otherwise.
	LastLabels map[string]string
	LastSeen   time.Time
	// A map for container name to a list of its usage samples, in chronological
	// order.
	Samples map[string][]model.ContainerUsageSample
}

func newEmptyHistory() *PodHistory {
	return &PodHistory{LastLabels: map[string]string{}, Samples: map[string][]model.ContainerUsageSample{}}
}

// HistoryProvider gives history of all pods in a cluster.
// TODO(schylek): this interface imposes how history is represented which doesn't work well with checkpoints.
// Consider refactoring to passing ClusterState and create history provider working with checkpoints.
type HistoryProvider interface {
	GetClusterHistory() (map[model.PodID]*PodHistory, error)
}

type prometheusHistoryProvider struct {
	prometheusClient  prometheusv1.API
	config            PrometheusHistoryProviderConfig
	queryTimeout      time.Duration
	historyDuration   prommodel.Duration
	historyResolution prommodel.Duration
}

/////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func NewPrometheusHistoryProvider(config PrometheusHistoryProviderConfig) (HistoryProvider, error) {
	return nil, nil
}

func (p *prometheusHistoryProvider) GetClusterHistory() (map[model.PodID]*PodHistory, error) {
	// construct podSelector, historicalCPUQuery & historicalMemoryQuery from the historyProvider's config
	// call readResourceHistory() for both cpu & memory
	// sort the containers resourceHistory of each pod, on basis of MeasureStart
	// call readLastLabels()
	return nil, nil
}

func (p *prometheusHistoryProvider) readResourceHistory(res map[model.PodID]*PodHistory, query string, resource model.ResourceName) error {
	// execute the query using prometheus Client
	// getContainerIDFromLabels()
	// getContainerUsageSamplesFromSamples()
	// set the samples on the corresponding pod
	return nil
}

func (p *prometheusHistoryProvider) getContainerIDFromLabels(metric prommodel.Metric) (*model.ContainerID, error) {
	// convert the metric to map
	// struct the containerId with containerName, podName & namespace , by taking them from that map
	return nil, nil
}

func getContainerUsageSamplesFromSamples(samples []prommodel.SamplePair, resource model.ResourceName) []model.ContainerUsageSample {
	// construct a ContainerUsageSample from each item of 'samples', & return the list
	// call resourceAmountFromValue() when constructing
	return nil
}

func resourceAmountFromValue(value float64, resource model.ResourceName) model.ResourceAmount {
	// just converts the type from float to 'ResourceAmount', using the conversion function from model/types.go
	return 0
}

func (p *prometheusHistoryProvider) readLastLabels(res map[model.PodID]*PodHistory, query string) error {
	// execute the query using prometheus Client
	// getPodIDFromLabels()
	// getPodLabelsMap()
	// update LastSeen & LastLabel variables on each of the PodHistory
	return nil
}

func (p *prometheusHistoryProvider) getPodIDFromLabels(metric prommodel.Metric) (*model.PodID, error) {
	// convert the metric to map
	// struct the PodId with podName & namespace , by taking them from that map
	return nil, nil
}

func (p *prometheusHistoryProvider) getPodLabelsMap(metric prommodel.Metric) map[string]string {
	// list all the podLabels whose key starts with p.config.PodLabelPrefix
	return nil
}

/*
NewPrometheusHistoryProvider is called from main.go, config-parameters are taken from tha main.go's flags

GetClusterHistory ia called from cluster_feeder.go
|
|-- readResourceHistory -- |
|                          |-- getContainerIDFromLabels
|                          |-- getContainerUsageSamplesFromSamples |-- resourceAmountFromValue
|-- readLastLabels -- |
|                     |-- getPodIDFromLabels
|                     |-- getPodLabelsMap
*/
