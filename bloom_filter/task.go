package bloom_filter

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"hash/maphash"
	"math"
)

type CountingBloomFilter struct {
	counts []int
	h1     maphash.Hash
	h2     maphash.Hash
	k      uint
}

func NewCountingBloomFilter(itemCount int, size int) CountingBloomFilter {
	k := float64(itemCount) / float64(size) * math.Ln2
	k = math.RoundToEven(k)
	k = math.Max(k, 1)
	filter := CountingBloomFilter{
		counts: make([]int, size, size),
		k:      uint(k),
	}
	filter.h1.SetSeed(maphash.MakeSeed())
	filter.h2.SetSeed(maphash.MakeSeed())
	return filter
}

func (filter *CountingBloomFilter) hashes(key interface{}) []uint64 {
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.LittleEndian, key)
	if err != nil {
		fmt.Printf("%v", err)
		panic("could not convert key to bytes")
	}
	byteArray := buf.Bytes()
	_, _ = filter.h1.Write(byteArray)
	hash1 := filter.h1.Sum64()
	filter.h1.Reset()
	_, _ = filter.h2.Write(byteArray)
	hash2 := filter.h2.Sum64()
	filter.h2.Reset()

	hashes := make([]uint64, 0, filter.k)
	var i uint
	for i = 0; i < filter.k; i++ {
		hashes = append(hashes, hash1+uint64(i)*hash2)
	}
	return hashes
}

func (filter *CountingBloomFilter) indexes(key interface{}) []uint {
	hashes := filter.hashes(key)
	indexes := make([]uint, 0, len(hashes))
	for _, hash := range hashes {
		indexes = append(indexes, uint(hash%uint64(len(filter.counts))))
	}
	return indexes
}

func (filter *CountingBloomFilter) Add(item interface{}) {
	indexes := filter.indexes(item)
	for _, index := range indexes {
		filter.counts[index]++
	}
}

func (filter *CountingBloomFilter) Count(item interface{}) int {
	indexes := filter.indexes(item)
	var count int
	min := math.MaxInt32
	for _, index := range indexes {
		count = filter.counts[index]
		if count == 0 {
			return 0
		}
		if count < min {
			min = count
		}
	}
	return min
}
