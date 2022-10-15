package vector

import (
	"testing"
)

func TestHashMatch(t *testing.T) {
	a := make([]uint8, 256)
	for i := 0; i < 256; i++ {
		a[i] = uint8(i)
	}
	hm := NewHashMatch()
	hm.Apply(a)
	if len(hm.InterZip(a)) != 256 {
		t.Fail()
	}
}

func BenchmarkHashMatchIntersect(b *testing.B) {
	var a0 []uint8
	var a1 []uint8
	for _, v := range randArray(60, 256) {
		a0 = append(a0, uint8(v))
	}
	for _, v := range randArray(252, 256) {
		a1 = append(a1, uint8(v))
	}
	hm := NewHashMatch()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		hm.Clear()
		hm.Apply(a0)
		hm.InterZip(a1)
	}
}
