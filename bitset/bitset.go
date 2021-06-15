// Package bitset provides a fast implementation of Bitsets
// also called Bitarray, Bitvector...
package bitset

import (
	"fmt"
)

// Bitset represents a bitset of fixed length
type Bitset struct {
	bitvec []int32
	length int
}

// New creates a new bitset instance with length l.
func New(l int) Bitset {
	return Bitset{
		bitvec: make([]int32, l),
		length: (l / 32) + 1,
	}
}

// BitLength
// Set will set a bit at pos.
func (b *Bitset) Set(pos int) error {
	var flag int32 = 1
	if pos < 0 || pos >= b.length*32 {
		return fmt.Errorf("invalid position for bitset of length %d", b.length)
	}
	// Pos will take a value between 0 and length
	// each chunck of the bitset is a 4 byte integer
	// so for e.g Set(10) will need to set the 10th
	// bit which is found in the first chunck i.e bitvec[0].
	// By reducing modulo 32 we find the local position in the bitvec.
	rpos := pos / 32 // (relative position inside the integer slice)
	bpos := pos % 32 // (local bit position inside bitvec[rpos])

	flag = 1 << bpos
	b.bitvec[rpos] = b.bitvec[rpos] | flag

	return nil
}

// Clear will clear the bit at pos.
func (b *Bitset) Clear(pos int) error {
	var flag int32 = 1
	if pos < 0 || pos >= b.length*32 {
		return fmt.Errorf("invalid position for bitset of length %d", b.length)
	}
	rpos := int32(pos) / 32 // (relative position inside the integer slice)
	bpos := int32(pos) % 32 // (local bit position inside bitvec[rpos])
	flag = flag << bpos
	flag = ^flag
	b.bitvec[rpos] = b.bitvec[rpos] & flag

	return nil
}

// IsSet checks if bit at pos is set.
func (b *Bitset) IsSet(pos int) bool {
	var flag int32 = 1

	if pos < 0 || pos >= b.length*32 {
		return false
	}
	rpos := int32(pos) / 32 // (relative position inside the integer slice)
	bpos := int32(pos) % 32 // (local bit position inside bitvec[rpos])

	flag = flag << int32(bpos)
	return (b.bitvec[rpos] & flag) != 0
}
