package vector

import "testing"

func BenchmarkBytesToBits(b *testing.B) {
	buf := []uint8{1, 4, 128, 200, 240}
	r := &[4]uint64{0, 0, 0, 0}
	for i := 0; i < b.N; i++ {
		bytesToBits(r, buf)
	}
}

func BenchmarkBitsToBitesDump_100K(b *testing.B) {
	s := make([][4]uint64, 100000)
	for i := 0; i < len(s); i++ {
		s[i][0] = uint64(i)
		s[i][1] = uint64(i)
		s[i][2] = uint64(i)
		s[i][3] = uint64(i)
	}
	for i := 0; i < b.N; i++ {
		x := make([]byte, 256)
		for k := 0; k < len(s); k++ {
			_ = bitsToBytesDump(s[i], x)
		}
	}
}

func BenchmarkBitsToBites_100K(b *testing.B) {
	s := make([][4]uint64, 100000)
	for i := 0; i < len(s); i++ {
		s[i][0] = uint64(i)
		s[i][1] = uint64(i)
		s[i][2] = uint64(i)
		s[i][3] = uint64(i)
	}
	for i := 0; i < b.N; i++ {
		x := make([]byte, 256)
		for k := 0; k < len(s); k++ {
			_ = bitsToBytes(s[i], x)
		}
	}
}

func BenchmarkBitsToBitesDump_500K(b *testing.B) {
	s := make([][4]uint64, 500000)
	for i := 0; i < len(s); i++ {
		s[i][0] = uint64(i)
		s[i][1] = uint64(i)
		s[i][2] = uint64(i)
		s[i][3] = uint64(i)
	}
	for i := 0; i < b.N; i++ {
		x := make([]byte, 256)
		for k := 0; k < len(s); k++ {
			_ = bitsToBytesDump(s[i], x)
		}
	}
}

func BenchmarkBitsToBites_500K(b *testing.B) {
	s := make([][4]uint64, 500000)
	for i := 0; i < len(s); i++ {
		s[i][0] = uint64(i)
		s[i][1] = uint64(i)
		s[i][2] = uint64(i)
		s[i][3] = uint64(i)
	}
	for i := 0; i < b.N; i++ {
		x := make([]byte, 256)
		for k := 0; k < len(s); k++ {
			_ = bitsToBytes(s[i], x)
		}
	}
}
