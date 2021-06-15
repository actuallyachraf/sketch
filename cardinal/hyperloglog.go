// Package cardinal implements efficient cardinality estimation
// datastructures.
package cardinal

import (
	"errors"
	"hash"
	"hash/fnv"
	"math/bits"
	"math/rand"

	inthash "github.com/actuallyachraf/sketch/hash"
)

const (
	UpperCorrectionThreshold = 143165576 //  2^{32} / 30
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
	numCounters := 1 << precision
	return &HyperLogLog{
		precision:   precision,
		numCounters: uint32(numCounters),
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
func (h *HyperLogLog) Add(item int) {
	itemHash := inthash.FNVHashInt32(int32(item))
	value, counterIdx := divmod(itemHash, int32(h.numCounters))
	h.counter[counterIdx] = max(h.counter[counterIdx], h.Rank(uint32(value)))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func divmod(a, b int32) (quo, rem int32) {
	return a / b, a % b
}

// Weight computes the weight for the cardinality estimators
// based on the counters.
func Weight(numCounters int) float64 {
	if numCounters < 16 {
		return 0.673
	} else if numCounters < 32 {
		return 0.697
	} else if numCounters < 64 {
		return 0.709
	}
	return (0.7213 * float64(numCounters)) / (float64(numCounters) + 1.079)
}
