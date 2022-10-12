package vector

import (
	"testing"
)

func BenchmarkVectorIterIntersect(b *testing.B) {
	v0 := NewVector()
	v1 := NewVector()
	for _, n := range firstArrayInt {
		v0.Add(n)
	}
	for _, n := range secondArrayInt {
		v1.Add(n)
	}
	it := NewIntersectIter(
		NewVectorIter(v0),
		NewVectorIter(v1),
	)
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
