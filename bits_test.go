package vector

import "testing"

func TestBitsToBytes64(t *testing.T) {
	var res [256]uint8
	var data = [4]uint64{0x1, 0x1, 0x1, 0x1}
	bitsToBytes64(&res, &data)
	t.Log(res)
}

func BenchmarkBitsToBytes64(b *testing.B) {
	var res [256]uint8
	var data = [4]uint64{0x1, 0x1, 0x1, 0x1}
	for i := 0; i < b.N; i++ {
		bitsToBytes64(&res, &data)
	}
}
