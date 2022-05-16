package util

import (
	vpa_types "k8s.io/autoscaler/vertical-pod-autoscaler/pkg/apis/autoscaling.k8s.io/v1"
	"time"
)

// Histogram represents an approximate distribution of some variable.
type Histogram interface {
	// Returns an approximation of the given percentile of the distribution.
	// Note: the argument passed to Percentile() is a number between
	// 0 and 1. For example 0.5 corresponds to the median and 0.9 to the
	// 90th percentile.
	// If the histogram is empty, Percentile() returns 0.0.
	Percentile(percentile float64) float64

	// Add a sample with a given value and weight.
	AddSample(value float64, weight float64, time time.Time)

	// Remove a sample with a given value and weight. Note that the total
	// weight of samples with a given value cannot be negative.
	SubtractSample(value float64, weight float64, time time.Time)

	// Add all samples from another histogram. Requires the histograms to be
	// of the exact same type.
	Merge(other Histogram)

	// Returns true if the histogram is empty.
	IsEmpty() bool

	// Returns true if the histogram is equal to another one. The two
	// histograms must use the same HistogramOptions object (not two
	// different copies).
	// If the two histograms are not of the same runtime type returns false.
	Equals(other Histogram) bool

	// Returns a human-readable text description of the histogram.
	String() string

	// SaveToChekpoint returns a representation of the histogram as a
	// HistogramCheckpoint. During conversion buckets with small weights
	// can be omitted.
	SaveToChekpoint() (*vpa_types.HistogramCheckpoint, error)

	// LoadFromCheckpoint loads data from the checkpoint into the histogram
	// by appending samples.
	LoadFromCheckpoint(*vpa_types.HistogramCheckpoint) error
}

// Simple bucket-based implementation of the Histogram interface. Each bucket
// holds the total weight of samples that belong to it.
// Percentile() returns the upper bound of the corresponding bucket.
// Resolution (bucket boundaries) of the histogram depends on the options.
// There's no interpolation within buckets (i.e. one sample falls to exactly one
// bucket).
// A bucket is considered empty if its weight is smaller than options.Epsilon().
type histogram struct {
	// Bucketing scheme.
	options HistogramOptions
	// Cumulative weight of samples in each bucket.
	bucketWeight []float64
	// Total cumulative weight of samples in all buckets.
	totalWeight float64
	// Index of the first non-empty bucket if there's any. Otherwise index
	// of the last bucket.
	minBucket int
	// Index of the last non-empty bucket if there's any. Otherwise 0.
	maxBucket int
}

/////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func NewHistogram(options HistogramOptions) Histogram {
	return nil
}

func (h *histogram) AddSample(value float64, weight float64, time time.Time) {
	// find its corresponding Bucket
	// update bucketWeight, totalWeight, minBucket, maxBucket
}

func (h *histogram) SubtractSample(value float64, weight float64, time time.Time) {
	// find its corresponding Bucket
	// safeSubtract for totalWeight & bucketWeight
	// updateMinAndMaxBucket()
}

func safeSubtract(value, sub, epsilon float64) float64 {
	// if value-sub < epsilon , return 0
	// otherwise value-sub
	return 0
}

func (h *histogram) updateMinAndMaxBucket() {
	// Adjusts the value of minBucket and maxBucket after any operation that decreases weights.
}

func (h *histogram) Merge(other Histogram) {
	// merging two histogram by updating bucketWeight, totalWeight, minBucket, maxBucket
}

func (h *histogram) Percentile(percentile float64) float64 {
	// threshold := percentile * h.totalWeight
	// return the max bucketStart point, upto where the cumulative sum of weights is less than the threshold
	return 0
}

func (h *histogram) IsEmpty() bool {
	// checks if the weight of the minBucket is < epsilon
	return false
}

func (h *histogram) String() string {
	// return human-readable version of the histogram
	return ""
}

func (h *histogram) Equals(other Histogram) bool {
	// checks if two histogram are equal.
	// h.options & other.options have also to be same
	return false
}

func (h *histogram) SaveToChekpoint() (*vpa_types.HistogramCheckpoint, error) {
	// converts the h.BucketWeights to ratio
	// & return a new checkpoint to be set on VPACheckpointStatus
	return nil, nil
}

func (h *histogram) LoadFromCheckpoint(checkpoint *vpa_types.HistogramCheckpoint) error {
	// load the checkpoint into struct
	// exact reverse of SaveToChekpoint
	return nil
}

func (h *histogram) scale(factor float64) {
	// Multiplies all weights by a given factor.
	// calls updateMinAndMaxBucket()
}

/////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

/*
almost all the functions are called from input/aggregate_container_state.go
*/
