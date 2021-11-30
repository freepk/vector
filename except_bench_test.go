package vector

import (
	"github.com/freepk/iterator"
	"testing"
)

func BenchmarkExceptCurrent(b *testing.B) {
	a0 := randArray(firstArraySize, maxValue(firstArraySize))
	a1 := randArray(secondArraySize, maxValue(secondArraySize))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		it := iterator.NewExceptIter(iterator.NewArrayIter(a0), iterator.NewArrayIter(a1))
		it.Reset()
		for {
			if _, ok := it.Next(); !ok {
				break
			}
		}
	}
}

func BenchmarkExceptVec(b *testing.B) {
	a0 := randArray(firstArraySize, maxValue(firstArraySize))
	v0 := CreateVectorFromArray(a0)
	a1 := randArray(secondArraySize, maxValue(secondArraySize))
	v1 := CreateVectorFromArray(a1)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		it := NewExceptVec(NewFetchVec(v0), NewFetchVec(v1))
		for {
			if _, _, ok := it.Next(); !ok {
				break
			}
		}
	}
}

func BenchmarkExceptVec_WithUnpack(b *testing.B) {
	a0 := randArray(firstArraySize, maxValue(firstArraySize))
	v0 := CreateVectorFromArray(a0)
	a1 := randArray(secondArraySize, maxValue(secondArraySize))
	v1 := CreateVectorFromArray(a1)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		it := NewExceptVec(NewFetchVec(v0), NewFetchVec(v1))
		unpack := NewUnpackVec(it)
		for {
			if _, ok := unpack.Next(); !ok {
				break
			}
		}
	}
}
