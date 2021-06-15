package bitset

import "testing"

func TestBitset(t *testing.T) {
	t.Run("TestSetClearIsSetBit", func(t *testing.T) {
		b := New(100)
		for k := 0; k < 100; k++ {
			t.Logf("Case #%d", k)
			err := b.Set(k)
			if err != nil {
				t.Fatal("failed to set bit with error :", err)
			}
			if !b.IsSet(k) {
				t.Fatalf("expected %d-bit to be 1 got 0", k)
			}
			b.Clear(k)
			if b.IsSet(k) {
				t.Fatalf("expected %d-bit to be 0 got 1", k)
			}
		}
		err := b.Set(128)
		if err == nil {
			t.Fatal("expected set bit to fail for out of bands index but got :", err)
		}
	})
}

func BenchmarkBitSet(b *testing.B) {
	b.Run("BenchmarkSetBit", func(b *testing.B) {
		bitvec := New(b.N)
		for n := 0; n < b.N; n++ {
			bitvec.Set(n)
		}
	})
}
