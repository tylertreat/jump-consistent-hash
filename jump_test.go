package jump

import "testing"

func TestConsistentHash(t *testing.T) {
	for i := 0; i < 500; i++ {
		h := ConsistentHash(uint64(i), 100)
		if h < 0 || h >= 100 {
			t.Errorf("Expected hash in range [0, 100), got %d", h)
		}
	}
}

func BenchmarkConsistentHash(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ConsistentHash(uint64(i), int32(b.N))
	}
}
