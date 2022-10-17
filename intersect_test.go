package vector

// import (
// 	"testing"
// )

// func TestIntersectIter(t *testing.T) {
// 	a0 := []int{0, 1, 5, 10, 15, 20, 25, 30, 35, 40, 45, 50, 220, 256, 272}
// 	a1 := []int{0, 1, 10, 20, 30, 40, 50, 60, 256, 260, 272}
// 	v0 := NewVector()
// 	v1 := NewVector()
// 	for _, n := range a0 {
// 		v0.Add(n)
// 	}
// 	for _, n := range a1 {
// 		v1.Add(n)
// 	}
// 	it := NewIntersectIter(
// 		NewVectorIter(v0),
// 		NewVectorIter(v1),
// 	)
// 	it.Reset()
// 	for {
// 		base, tail, ok := it.Next()
// 		if !ok {
// 			break
// 		}
// 		_ = base
// 		_ = tail
// 	}
// }

// func BenchmarkVectorIterIntersect(b *testing.B) {
// 	v0 := NewVector()
// 	v1 := NewVector()
// 	for _, n := range firstArrayInt {
// 		v0.Add(n)
// 	}
// 	for _, n := range secondArrayInt {
// 		v1.Add(n)
// 	}
// 	it := NewIntersectIter(
// 		NewVectorIter(v0),
// 		NewVectorIter(v1),
// 	)
// 	b.ResetTimer()
// 	for i := 0; i < b.N; i++ {
// 		it.Reset()
// 		for {
// 			if _, _, ok := it.Next(); !ok {
// 				break
// 			}
// 		}
// 	}
// }
