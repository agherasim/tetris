package main

import "math/bits"

func intMax(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func bitsLen(b byte) int {
	i := uint(b)
	return bits.Len(i)
}
