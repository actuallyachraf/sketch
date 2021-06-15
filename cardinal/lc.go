// Package cardinal implements efficient cardinality estimation
// datastructures.
package cardinal

import (
	"hash"
	"hash/fnv"
	"math"

	"github.com/actuallyachraf/sketch/bitset"
	inthash "github.com/actuallyachraf/sketch/hash"
)

// LinearCounter is a simple map(hash(value) => bit) and
// cardinality can be estimated using n = ~m*ln(V) where
// m is a paramter choosen based on how much entires we expect
// and V is the number of 0-set bits.
type LinearCounter struct {
	m      int
	h      hash.Hash
	bitvec bitset.Bitset
}

// NewLinearCounter creates a new linear counter.
func NewLinearCounter(m int) LinearCounter {
	return LinearCounter{
		m:      m,
		h:      fnv.New32a(),
		bitvec: bitset.New(m),
	}
}

// Add an item to the linear counter
func (lc *LinearCounter) Add(item int) error {

	pos := inthash.FNVHashInt32(int32(item))
	if !lc.bitvec.IsSet(int(pos)) {
		return lc.bitvec.Set(int(pos))
	}
	return nil
}

// Cardinal returns the estimated cardinal of the dataset
func (lc *LinearCounter) Cardinal() float64 {
	var Z float64
	var m float64 = float64(lc.m)
	for i := 0; i < lc.m; i++ {
		if !lc.bitvec.IsSet(i) {
			Z++
		}
	}
	estimate := -m * math.Log(Z/m)
	return math.Floor(estimate)
}
