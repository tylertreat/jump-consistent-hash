package jump

// ConsistentHash is an implementation of the Jump Consistent Hash algorithm
// as presented by Lamping and Veach in A Fast, Minimal Memory, Consistent Hash
// Algorithm (http://arxiv.org/pdf/1406.2294.pdf). It takes a key and number of
// buckets and returns a bucket in the range [0, numBuckets).
func ConsistentHash(key uint64, numBuckets int32) int32 {
	var (
		b = int64(-1)
		j = int64(0)
	)

	for j < int64(numBuckets) {
		b = j
		key = key*2862933555777941757 + 1
		j = int64(float64(b+1) * (float64(int64(1)<<31) / float64((key>>33)+1)))
	}

	return int32(b)
}
