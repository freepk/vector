package vector

import "testing"

func TestBitsToBytes64(t *testing.T) {
	var res [64]uint8
	bitsToBytes64(&res, 0xffaaffaaffaaffaa)
	t.Log(res)
}

func BenchmarkBitsToBytes64(b *testing.B) {
	var res [64]uint8
	for i := 0; i < b.N; i++ {
		bitsToBytes64(&res, 0xffaaffaaffaaffaa)
	}
}
