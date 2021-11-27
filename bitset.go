package vector

import "math/bits"

var bitsTable [256]uint64

func init() {
	buildBitsTable(&bitsTable)
}

func buildBitsTable(r *[256]uint64) {
	for i := 0; i < 256; i++ {
		r[i] = 1 << (i & 0x3f)
	}
}

func bytesToBits(r *[4]uint64, b []uint8) {
	for i := 0; i < len(b); i++ {
		r[(b[i] >> 6)] |= bitsTable[b[i]]
	}
}

func bitsToBytesDump(r [4]uint64, b []uint8) []uint8 {
	b = b[:0]
	for i := 0; i < 4; i++ {
		if r[i] == 0 {
			continue
		}
		off := uint8(i * 64)
		for x := 0; x < 64; x++ {
			if r[i]&(1<<x) != 0 {
				b = append(b, uint8(x)+off)
			}
		}
	}
	return b
}

func bitsToBytes(r [4]uint64, b []uint8) []uint8 {
	n := 0
	for i := 0; i < 4; i++ {
		off := uint8(i << 6)
		v := 0
		rv := r[i]
		for {
			if rv == 0 {
				break
			}
			x := bits.TrailingZeros64(rv)
			b[n] = uint8(x+v) + off
			n++
			x++
			v += x
			rv >>= x
		}
	}
	return b[:n]
}
