package vector

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
