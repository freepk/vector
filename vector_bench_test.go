package vector

import (
	"testing"
)

func BenchmarkVectorIterNext(b *testing.B) {
	a0 := randArray(firstArraySize, secondArraySize)
	v0 := NewVector()
	for _, n := range a0 {
		v0.Add(n)
	}
	it := NewVectorIter(v0)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		it.Reset()
		for {
			if _, _, ok := it.Next(); !ok {
				break
			}
		}
	}
}
