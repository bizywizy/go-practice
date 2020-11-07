package bloom_filter

import "testing"

func TestCountingBloomFilter(t *testing.T) {
	filter := NewCountingBloomFilter(10, 10000)
	var i, j int64

	for i = 0; i < 100; i++ {
		for j = 0; j < 10; j++ {
			filter.Add(j)
		}
	}

	for j = 0; j < 10; j++ {
		count := filter.Count(j)
		if count == 0 {
			t.Errorf("%d count is 0", j)
		}
	}

	for j = 10; j < 100; j++ {
		count := filter.Count(j)
		if count > 0 {
			t.Errorf("%d count is greater than 0", j)
		}
	}

}
