package cardinal

import (
	"math"
	"math/rand"
	"testing"
)

// Testing Linear Counter using an almost normal distribution of items.
func TestLinearCounter(t *testing.T) {
	t.Run("TestLinearCounterConsistency", func(t *testing.T) {
		lc := NewLinearCounter(8000)
		lc.Add(3)
		lc.Add(2)
		lc.Add(1)
		if lc.Cardinal() != 3 {
			t.Fatalf("failed to validate consistency for small cardinality expected %d got %f", 3, lc.Cardinal())
		}
	})
	t.Run("TestLinearCounterFNV-1a-Estimated-Biased", func(t *testing.T) {
		lc := NewLinearCounter(268)
		dataset := make([]int, 10000)
		for k := 0; k < 10000; k++ {
			m := rand.Int31n(100)
			dataset[k] = int(m)
			err := lc.Add(int(m))

			if err != nil {
				t.Fatal("failed to add item to LC with error :", err)
			}
		}
		estimate := lc.Cardinal()
		t.Log("Estimated number of unique entries", estimate)
	})
	t.Run("TestLinearCounterFNV-1a-Estimated-NonBiased", func(t *testing.T) {
		lc := NewLinearCounter(100000)
		cardinality := 0
		misses := make([]float64, 0)
		sumMisses := 0.
		for k := 0; k < 1000; k++ {
			cardinality++
			lc.Add(k)
			miss := math.Abs(float64(cardinality)-lc.Cardinal()) / float64(cardinality)
			misses = append(misses, miss)
			sumMisses += miss
		}
		avgError := sumMisses / float64(len(misses))
		if avgError < 0. || avgError >= 0.1 {
			t.Fatalf("failed to assert positive average error got %f", avgError)
		}
		t.Log("Average Error :", avgError)
	})
	t.Run("TestLinearCounterFNV-1a-Exact", func(t *testing.T) {
		lc := NewLinearCounter(5329)
		dataset := make([]int, 10000)
		for k := 0; k < 10000; k++ {
			m := rand.Int31n(100)
			dataset[k] = int(m)
			err := lc.Add(int(m))

			if err != nil {
				t.Fatal("failed to add item to LC with error :", err)
			}
		}
		estimate := lc.Cardinal()
		t.Log("Estimated number of unique entries", estimate)
	})
}
