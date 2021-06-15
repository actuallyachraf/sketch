package cardinal

import (
	"math"
	"math/rand"
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
	t.Run("TestHyperLogLogEstimate-5", func(t *testing.T) {
		h, err := NewHyperLogLog(5)
		if err != nil {
			t.Fatal("failed to create new hyperloglog instance with error", err)
		}
		precShift := 1 << h.precision
		std := 1.04 / math.Sqrt(float64(precShift))

		cardinality := 100
		avgErr := 0.
		errRate := 0.
		errCount := 0.
		for i := 0; i < 100000; i++ {
			k := rand.Int31n(100)
			h.Add(k)
			errRate += math.Abs(float64(cardinality-h.Cardinal())) / float64(cardinality)
			errCount += 1
		}
		avgErr = errRate / errCount
		t.Log(avgErr)
		t.Log(std)
		t.Log(h.Cardinal())
	})
	t.Run("TestHyperLogLogEstimate-8", func(t *testing.T) {
		h, err := NewHyperLogLog(8)
		if err != nil {
			t.Fatal("failed to create new hyperloglog instance with error", err)
		}
		precShift := 1 << h.precision
		std := 1.04 / math.Sqrt(float64(precShift))

		cardinality := 100
		avgErr := 0.
		errRate := 0.
		errCount := 0.
		for i := 0; i < 100000; i++ {
			k := rand.Int31n(1000)
			h.Add(k)
			errRate += math.Abs(float64(cardinality-h.Cardinal())) / float64(cardinality)
			errCount += 1
		}
		avgErr = (errRate / errCount) / 100
		t.Log(avgErr)
		t.Log(std)
		t.Log(h.Cardinal())
	})
}
func BenchmarkHyperLogLog(b *testing.B) {
	b.Run("BenchmarkHyperLogLog-5", func(b *testing.B) {
		h, err := NewHyperLogLog(5)
		if err != nil {
			b.Fatal("failed to create new hyperloglog instance with error", err)
		}
		for i := 0; i < b.N; i++ {
			k := rand.Int31n(256)
			h.Add(k)
		}
	})
	b.Run("BenchmarkHyperLogLog-8", func(b *testing.B) {
		h, err := NewHyperLogLog(8)
		if err != nil {
			b.Fatal("failed to create new hyperloglog instance with error", err)
		}
		for i := 0; i < b.N; i++ {
			k := rand.Int31n(512)
			h.Add(k)
		}
	})
	b.Run("BenchmarkHyperLogLog-5", func(b *testing.B) {
		h, err := NewHyperLogLog(10)
		if err != nil {
			b.Fatal("failed to create new hyperloglog instance with error", err)
		}
		for i := 0; i < b.N; i++ {
			k := rand.Int31n(1024)
			h.Add(k)
		}
	})
}
