// Package cardinal implements efficient cardinality estimation
// datastructures.
package cardinal

import (
	"errors"
	"hash"
	"hash/fnv"
	"math"
	"math/bits"
	"math/rand"

	"github.com/actuallyachraf/sketch/inthash"
)

// HyperLogLog data structure does cardinality estimation
// by ranking items according to the LSB bits index.
type HyperLogLog struct {
	precision   int
	numCounters uint32
	size        int
	counter     []int
	seed        uint32
	alpha       float64
	h           hash.Hash
}

// NewHyperLogLog creates a new instance of hyperloglog
// precision takes values in (4..16).
func NewHyperLogLog(precision int) (*HyperLogLog, error) {
	if precision < 4 || precision > 16 {
		return nil, errors.New("precision has to be in range 4..16")
	}
	var numCounters uint32 = 1 << precision
	return &HyperLogLog{
		precision:   precision,
		numCounters: numCounters,
		size:        32 - precision,
		counter:     make([]int, numCounters),
		seed:        rand.Uint32(),
		alpha:       Weight(numCounters),
		h:           fnv.New32a(),
	}, nil
}

// Rank which is the least significant bit position.
func (h *HyperLogLog) Rank(value uint32) int {
	return h.size - bits.Len32(value) + 1
}

// Add indexes an element into the counter.
func (h *HyperLogLog) Add(item int32) {
	itemHash := inthash.FNVHashInt32(int32(item))
	value := 32 - h.precision
	r := leftmostActiveBit(itemHash << uint32(h.precision))
	j := itemHash >> uint(value)

	if r > h.counter[j] {
		h.counter[j] = r
	}
}

// Cardinal approximately counts the number of unique elements
// indexed by the HyperLogLog counter.
func (h *HyperLogLog) Cardinal() int {
	sum := 0.
	m := float64(h.numCounters)
	for _, v := range h.counter {
		sum += math.Pow(math.Pow(2, float64(v)), -1)
	}
	estimate := h.alpha * m * m / sum
	return int(estimate)
}
func leftmostActiveBit(x uint32) int {
	return 1 + bits.LeadingZeros32(x)
}

// Weight computes the weight for the cardinality estimators
// based on the counters.
func Weight(numCounters uint32) float64 {
	if numCounters < 16 {
		return 0.673
	} else if numCounters < 32 {
		return 0.697
	} else if numCounters < 64 {
		return 0.709
	}
	return (0.7213 * float64(numCounters)) / (float64(numCounters) + 1.079)
}
