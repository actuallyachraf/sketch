package inthash

import (
	"encoding/binary"
	"hash/fnv"
)

// FNVHashInt32 hashes a 32-bit integer using FNV hash function
func FNVHashInt32(x int32) int32 {
	buf := make([]byte, 4)
	binary.LittleEndian.PutUint32(buf, uint32(x))
	h := fnv.New32a().Sum(buf)
	hInt := binary.LittleEndian.Uint32(h)
	return int32(hInt)
}
