package vector

import (
	"github.com/freepk/iterator"
	"testing"
)

var maxValue = func(x int) int { return x }

func BenchmarkIntersectCurrent(b *testing.B) {
	a0 := randArray(firstArraySize, maxValue(firstArraySize))
	a1 := randArray(secondArraySize, maxValue(secondArraySize))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		it := iterator.NewInterIter(iterator.NewArrayIter(a0), iterator.NewArrayIter(a1))
		it.Reset()
		for {
			if _, ok := it.Next(); !ok {
				break
			}
		}
	}
}

func BenchmarkIntersectVec(b *testing.B) {
	a0 := randArray(firstArraySize, maxValue(firstArraySize))
	v0 := CreateVectorFromArray(a0)
	a1 := randArray(secondArraySize, maxValue(secondArraySize))
	v1 := CreateVectorFromArray(a1)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		it := NewIntersectVec(NewFetchVec(v0), NewFetchVec(v1))
		for {
			if _, _, ok := it.Next(); !ok {
				break
			}
		}
	}
}

func BenchmarkIntersectVec_WithUnpack(b *testing.B) {
	a0 := randArray(firstArraySize, maxValue(firstArraySize))
	v0 := CreateVectorFromArray(a0)
	a1 := randArray(secondArraySize, maxValue(secondArraySize))
	v1 := CreateVectorFromArray(a1)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		it := NewIntersectVec(NewFetchVec(v0), NewFetchVec(v1))
		unpack := NewUnpackVec(it)
		for {
			if _, ok := unpack.Next(); !ok {
				break
			}
		}
	}
}

func BenchmarkIntersectComplex(b *testing.B) {
	a0 := randArray(firstArraySize, maxValue(firstArraySize))
	v0 := CreateVectorFromArray(a0)
	a1 := randArray(secondArraySize, maxValue(secondArraySize))
	v1 := CreateVectorFromArray(a1)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		it := NewIntersectComplex(NewFetchVec(v0), NewFetchVec(v1))
		for {
			if _, _, ok := it.Next(); !ok {
				break
			}
		}
	}
}

func BenchmarkIntersectComplex_WithUnpack(b *testing.B) {
	a0 := randArray(firstArraySize, maxValue(firstArraySize))
	v0 := CreateVectorFromArray(a0)
	a1 := randArray(secondArraySize, maxValue(secondArraySize))
	v1 := CreateVectorFromArray(a1)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		it := NewIntersectComplex(NewFetchVec(v0), NewFetchVec(v1))
		unpack := NewUnpackVec(it)
		for {
			if _, ok := unpack.Next(); !ok {
				break
			}
		}
	}
}

var intsctSizeList = []int{firstArraySize, secondArraySize, firstArraySize, firstArraySize, secondArraySize, firstArraySize, secondArraySize, secondArraySize, secondArraySize, firstArraySize}

func BenchmarkIntersectCurrent_Many(b *testing.B) {
	arrs := make([][]int, len(intsctSizeList))
	for i, s := range intsctSizeList {
		arrs[i] = randArray(s, maxValue(s))
	}
	b.ResetTimer()
	var it iterator.Iterator
	for i := 0; i < b.N; i++ {
		it = iterator.NewArrayIter(arrs[0])
		for j := 1; j < len(arrs); j++ {
			it = iterator.NewInterIter(it, iterator.NewArrayIter(arrs[j]))
		}
		for {
			if _, ok := it.Next(); !ok {
				break
			}
		}
	}
}

func BenchmarkIntersectComplex_Many(b *testing.B) {
	vecs := make([]vector, len(intsctSizeList))
	for i, s := range intsctSizeList {
		for _, x := range randArray(s, maxValue(s)) {
			vecs[i].Add(uint32(x))
		}
	}
	fvs := make([]Iterator, len(intsctSizeList))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < len(vecs); j++ {
			fvs[j] = NewFetchVec(vecs[j])
		}
		ic := NewIntersectComplex(fvs...)
		for {
			if _, _, ok := ic.Next(); !ok {
				break
			}
		}
	}
}

func BenchmarkIntersectComplex_Many_WithUnpack(b *testing.B) {
	vecs := make([]vector, len(intsctSizeList))
	for i, s := range intsctSizeList {
		for _, x := range randArray(s, maxValue(s)) {
			vecs[i].Add(uint32(x))
		}
	}
	fvs := make([]Iterator, len(intsctSizeList))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < len(vecs); j++ {
			fvs[j] = NewFetchVec(vecs[j])
		}
		ic := NewIntersectComplex(fvs...)
		unpack := NewUnpackVec(ic)
		for {
			if _, ok := unpack.Next(); !ok {
				break
			}
		}
	}
}

func BenchmarkIntersectVec_Many(b *testing.B) {
	vecs := make([]vector, len(intsctSizeList))
	for i, s := range intsctSizeList {
		for _, x := range randArray(s, maxValue(s)) {
			vecs[i].Add(uint32(x))
		}
	}
	var iv Iterator
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		iv = NewFetchVec(vecs[0])
		for j := 1; j < len(vecs); j++ {
			iv = NewIntersectVec(iv, NewFetchVec(vecs[j]))
		}
		for {
			if _, _, ok := iv.Next(); !ok {
				break
			}
		}
	}
}

func BenchmarkIntersectVec_Many_WithUnpack(b *testing.B) {
	vecs := make([]vector, len(intsctSizeList))
	for i, s := range intsctSizeList {
		for _, x := range randArray(s, maxValue(s)) {
			vecs[i].Add(uint32(x))
		}
	}
	var iv Iterator
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		iv = NewFetchVec(vecs[0])
		for j := 1; j < len(vecs); j++ {
			iv = NewIntersectVec(iv, NewFetchVec(vecs[j]))
		}
		unpack := NewUnpackVec(iv)
		for {
			if _, ok := unpack.Next(); !ok {
				break
			}
		}
	}
}
