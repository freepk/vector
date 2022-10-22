package vector

import "math/bits"

func bitsToBytes64(res *[64]uint8, mask uint64) int {
	i := 0
	for mask != 0 {
		res[i] = uint8(bits.TrailingZeros64(mask))
		mask &= (mask - 1)
		i++
	}
	return i
}
