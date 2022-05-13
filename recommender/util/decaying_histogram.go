package util

import (
	vpa_types "k8s.io/autoscaler/vertical-pod-autoscaler/pkg/apis/autoscaling.k8s.io/v1"
	"time"
)

var (
	// When the decay factor exceeds 2^maxDecayExponent the histogram is
	// renormalized by shifting the decay start time forward.
	maxDecayExponent = 100
)

// A histogram that gives newer samples a higher weight than the old samples,
// gradually decaying ("forgetting") the past samples. The weight of each sample
// is multiplied by the factor of 2^((sampleTime - referenceTimestamp) / halfLife).
// This means that the sample loses half of its weight ("importance") with
// each halfLife period.
// Since only relative (and not absolute) weights of samples matter, the
// referenceTimestamp can be shifted at any time, which is equivalent to multiplying all
// weights by a constant. In practice the referenceTimestamp is shifted forward whenever
// the exponents become too large, to avoid floating point arithmetics overflow.
type decayingHistogram struct {
	histogram
	// Decay half life period.
	halfLife time.Duration
	// Reference time for determining the relative age of samples.
	// It is always an integer multiple of halfLife.
	referenceTimestamp time.Time
}

/////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func NewDecayingHistogram(options HistogramOptions, halfLife time.Duration) Histogram {
	return nil
}

func (h *decayingHistogram) Percentile(percentile float64) float64 {
	// call h.histogram.Percentile()
	return 0
}

func (h *decayingHistogram) AddSample(value float64, weight float64, time time.Time) {
	// call h.histogram.AddSample(), just change the weight accordingly, weight *= decayFactor(time)
}

func (h *decayingHistogram) SubtractSample(value float64, weight float64, time time.Time) {
	// call h.histogram.SubtractSample(), just change the weight accordingly, weight *= decayFactor(time)
}

func (h *decayingHistogram) Merge(other Histogram) {
	// shift the referenceTimeStamp of the newerOne
	//call h.histogram.Merge()
}

func (h *decayingHistogram) Equals(other Histogram) bool {
	// checks equality
	return false
}

func (h *decayingHistogram) IsEmpty() bool {
	// calls h.histogram.IsEmpty()
	return false
}

func (h *decayingHistogram) String() string {
	// human-readable string
	return ""
}

func (h *decayingHistogram) SaveToChekpoint() (*vpa_types.HistogramCheckpoint, error) {
	// call h.histogram.SaveToChekpoint() + set ReferenceTimestamp
	return nil, nil
}

func (h *decayingHistogram) LoadFromCheckpoint(checkpoint *vpa_types.HistogramCheckpoint) error {
	// call h.histogram.LoadFromCheckpoint() + set ReferenceTimestamp
	return nil
}

func (h *decayingHistogram) shiftReferenceTimestamp(newreferenceTimestamp time.Time) {
	// Scale all weights by 2^exponent. where exponent = (current referenceTimeStamp - new referenceTimestamp) / halfLife
	// update h.referenceTimestamp
}

func (h *decayingHistogram) decayFactor(timestamp time.Time) float64 {
	// if need to reconmalize the histogram, call shiftReferenceTimestamp
	// return (current referenceTimeStamp - timestamp) / halfLife
	return 0
}
