package cardinal

import (
	"testing"
)

func TestHyperLogLog(t *testing.T) {
	t.Run("TestHyperLogLogConsistency", func(t *testing.T) {
		_, err := NewHyperLogLog(10)
		if err != nil {
			t.Fatal("failed to create new hyperloglog instance with error", err)
		}
		_, err = NewHyperLogLog(2)
		if err == nil {
			t.Fatal("expected failure for low precision got nil instead")
		}
		h, _ := NewHyperLogLog(10)
		r := h.Rank(1)
		if r != 22 {
			t.Fatalf("expected rank(1) to be %d got %d", 22, r)
		}
		h.Add(1)
		h.Add(2)
		h.Add(3)
		h.Add(4)
	})
}
