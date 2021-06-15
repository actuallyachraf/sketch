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

	inthash "github.com/actuallyachraf/sketch/hash"
)

const (
	UpperCorrectionThreshold = 143165576. //  2^{32} / 30
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

// Cardinal approximately counts the number of unique elements
// indexed by the HyperLogLog counter.
func (h *HyperLogLog) Cardinal() float64 {
	var R float64 = 0.
	var allZero bool = true
	var numCounters float64 = float64(h.numCounters)

	for counterIdx := 0; counterIdx < int(h.numCounters); counterIdx++ {
		if h.counter[counterIdx] > 0 {
			allZero = false
		}
		R += 1. / float64(uint32(1)<<h.counter[counterIdx])
	}

	if allZero {
		return 0
	}
	n := math.Round(h.alpha * numCounters * numCounters / R)
	Z := 0.
	if n < 2.5*numCounters {
		for counterIdx := 0; counterIdx < int(h.numCounters); counterIdx++ {
			if h.counter[counterIdx] == 0 {
				Z += 1
			}
		}
		if Z > 0 {
			return math.Round(numCounters*math.Log(numCounters) - math.Log(Z))
		}
	} else if n > UpperCorrectionThreshold {
		n = math.Round(-44294967296. * math.Log(1-(n/4294967296.)))
	}
	return n
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
