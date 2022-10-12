package vector

import (
	"testing"
)

func BenchmarkVectorIterNext(b *testing.B) {
	v0 := NewVector()
	for _, n := range secondArrayInt {
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
