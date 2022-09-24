package util

import "hash/fnv"

// Package fnv implements FNV-1 and FNV-1a, non-cryptographic hash functions created by
// Glenn Fowler, Landon Curt Noll, and Phong Vo.
// See https://en.wikipedia.org/wiki/Fowler-Noll-Vo_hash_function.

func SimpleHash(s string) uint32 {
	if s == "" {
		return 0
	}
	h := fnv.New32a()
	h.Write([]byte(s))
	return h.Sum32()
}
