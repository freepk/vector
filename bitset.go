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

/*
func bitsOr(r *[4]uint64, b []uint8) {
	n := len(b)
	i := 0
	for i < (n - 4) {
		r[(b[i] >> 6)] |= bitsTable[b[i]]
		i++
		r[(b[i] >> 6)] |= bitsTable[b[i]]
		i++
		r[(b[i] >> 6)] |= bitsTable[b[i]]
		i++
		r[(b[i] >> 6)] |= bitsTable[b[i]]
		i++
	}
	switch n - i {
	case 3:
		r[(b[i] >> 6)] |= bitsTable[b[i]]
		i++
		r[(b[i] >> 6)] |= bitsTable[b[i]]
		i++
		r[(b[i] >> 6)] |= bitsTable[b[i]]
		i++
	case 2:
		r[(b[i] >> 6)] |= bitsTable[b[i]]
		i++
		r[(b[i] >> 6)] |= bitsTable[b[i]]
		i++
	case 1:
		r[(b[i] >> 6)] |= bitsTable[b[i]]
		i++
	}
}
*/

func bitsSet(r *[4]uint64, b []uint8) {
	for i := 0; i < len(b); i++ {
		r[(b[i] >> 6)] = bitsTable[b[i]]
	}
}

func bitsOr(r *[4]uint64, b []uint8) {
	for i := 0; i < len(b); i++ {
		r[(b[i] >> 6)] |= bitsTable[b[i]]
	}
}

func bitsAnd(r *[4]uint64, b []uint8) {
	for i := 0; i < len(b); i++ {
		r[(b[i] >> 6)] &= bitsTable[b[i]]
	}
}
