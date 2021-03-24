package vector

import "testing"

func BenchmarkBitsSet(b *testing.B) {
	buf := []uint8{1, 4, 128, 200, 240}
	r := &[4]uint64{0, 0, 0, 0}
	for i := 0; i < b.N; i++ {
		bitsSet(r, buf)
	}
}

func BenchmarkBitsOr(b *testing.B) {
	buf := []uint8{1, 4, 128, 200, 240}
	r := &[4]uint64{0, 0, 0, 0}
	for i := 0; i < b.N; i++ {
		bitsOr(r, buf)
	}
}

func BenchmarkBitsAnd(b *testing.B) {
	buf := []uint8{1, 4, 128, 200, 240}
	r := &[4]uint64{0, 0, 0, 0}
	for i := 0; i < b.N; i++ {
		bitsAnd(r, buf)
	}
}
