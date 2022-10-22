package vector

import "testing"

func TestDecodeBits(t *testing.T) {
	var res [64]int
	decodeBits(&res, 0xffaaffaaffaaffaa)
	t.Log(res)
}

func BenchmarkDecode(b *testing.B) {
	var res [64]int
	for i := 0; i < b.N; i++ {
		decodeBits(&res, 0xffaaffaaffaaffaa)
	}
}
