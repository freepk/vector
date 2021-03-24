package vector

import "testing"

func BenchmarkBytesToBits(b *testing.B) {
	buf := []uint8{1, 4, 128, 200, 240}
	r := &[4]uint64{0, 0, 0, 0}
	for i := 0; i < b.N; i++ {
		bytesToBits(r, buf)
	}
}
