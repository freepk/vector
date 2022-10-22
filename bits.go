package vector

import "math/bits"

func decodeBits(res *[64]int, mask uint64) int {
	i := 0
	for mask != 0 {
		res[i] = bits.TrailingZeros64(mask)
		mask &= (mask - 1)
		i++
	}
	return i
}
