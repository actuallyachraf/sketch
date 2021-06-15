package cardinal

import (
	"math/rand"
	"testing"
)

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
	t.Run("TestLinearCounterFNV-1a-Estimated", func(t *testing.T) {
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
