package util

// HistogramOptions define the number and size of buckets of a histogram.
type HistogramOptions interface {
	// Returns the number of buckets in the histogram.
	NumBuckets() int
	// Returns the index of the bucket to which the given value falls.
	// If the value is outside of the range covered by the histogram, it
	// returns the closest bucket (either the first or the last one).
	FindBucket(value float64) int
	// Returns the start of the bucket with a given index. If the index is
	// outside the [0..NumBuckets() - 1] range, the result is undefined.
	GetBucketStart(bucket int) float64
	// Returns the minimum weight for a bucket to be considered non-empty.
	Epsilon() float64
}

type linearHistogramOptions struct {
	numBuckets int
	bucketSize float64
	epsilon    float64
}

type exponentialHistogramOptions struct {
	numBuckets      int
	firstBucketSize float64
	ratio           float64
	epsilon         float64
}

/////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func NewLinearHistogramOptions(maxValue float64, bucketSize float64, epsilon float64) (HistogramOptions, error) {
	return nil, nil
}

func (o *linearHistogramOptions) NumBuckets() int {
	return o.numBuckets
}
func (o *linearHistogramOptions) FindBucket(value float64) int {
	return int(value / o.bucketSize)
}

func (o *linearHistogramOptions) GetBucketStart(bucket int) float64 {
	return float64(bucket) * o.bucketSize
}

func (o *linearHistogramOptions) Epsilon() float64 {
	return o.epsilon
}

func NewExponentialHistogramOptions(maxValue float64, firstBucketSize float64, ratio float64, epsilon float64) (HistogramOptions, error) {
	return nil, nil
}

func (o *exponentialHistogramOptions) NumBuckets() int {
	return o.numBuckets
}
func (o *exponentialHistogramOptions) FindBucket(value float64) int {
	// x = value*(o.ratio-1)/o.firstBucketSize +1
	// return o.ratio_base_log(x)
	//
	// this is exact opposite of GetBucketStart formula, as
	// if x^y=n, then log_x(n)=y
	return 0
}

func (o *exponentialHistogramOptions) GetBucketStart(bucket int) float64 {
	// formula => a * (r^n -1) / (r-1) <= sum of first n element of a multiplicative series
	// here a=firstBucketSize, r=ratio, n=bucket
	return 0
}

func (o *exponentialHistogramOptions) Epsilon() float64 {
	return o.epsilon
}
