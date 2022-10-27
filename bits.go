package vector

import "math/bits"

func bitsToBytes64(res *[256]uint8, data *[4]uint64) int {
	i := 0
	bs := data[0]
	for bs != 0 {
		res[i] = uint8(bits.TrailingZeros64(bs))
		bs &= (bs - 1)
		i++
	}
	bs = data[1]
	for bs != 0 {
		res[i] = uint8(bits.TrailingZeros64(bs)) + 64
		bs &= (bs - 1)
		i++
	}
	bs = data[2]
	for bs != 0 {
		res[i] = uint8(bits.TrailingZeros64(bs)) + 128
		bs &= (bs - 1)
		i++
	}
	bs = data[3]
	for bs != 0 {
		res[i] = uint8(bits.TrailingZeros64(bs)) + 192
		bs &= (bs - 1)
		i++
	}
	return i
}
