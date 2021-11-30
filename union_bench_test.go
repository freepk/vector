package vector

import (
	"github.com/freepk/iterator"
	"testing"
)

func BenchmarkUnionCurrent(b *testing.B) {
	a0 := randArray(firstArraySize, maxArrayValue)
	a1 := randArray(secondArraySize, maxArrayValue)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		it := iterator.NewUnionIter(iterator.NewArrayIter(a0), iterator.NewArrayIter(a1))
		for {
			if _, ok := it.Next(); !ok {
				break
			}
		}
	}
}

func BenchmarkUnionComplex(b *testing.B) {
	a0 := randArray(firstArraySize, maxArrayValue)
	v0 := CreateVectorFromArray(a0)
	a1 := randArray(secondArraySize, maxArrayValue)
	v1 := CreateVectorFromArray(a1)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		it := NewUnionComplex(NewFetchVec(v0), NewFetchVec(v1))
		for {
			if _, _, ok := it.Next(); !ok {
				break
			}
		}
	}
}

func BenchmarkUnionVec(b *testing.B) {
	a0 := randArray(firstArraySize, maxArrayValue)
	v0 := CreateVectorFromArray(a0)
	a1 := randArray(secondArraySize, maxArrayValue)
	v1 := CreateVectorFromArray(a1)
	b.ResetTimer()
	ok := false
	for i := 0; i < b.N; i++ {
		it := NewUnionVec(NewFetchVec(v0), NewFetchVec(v1))
		for {
			if _, _, ok = it.Next(); !ok {
				break
			}
		}
	}
}

//var unSizeList = []int{21506, 64316, 4065, 1313, 84, 173, 146, 974, 8369, 771, 555, 3983, 15693, 254, 1945, 11934, 1722, 6122, 1150, 893, 3, 234, 1670, 776, 2335, 1296, 150, 2215, 3518, 535, 1435, 561, 761, 1266, 278, 1347, 352, 1695, 8824, 609, 262, 998, 108, 832, 316, 818, 9, 233, 36, 24, 660, 214, 261, 903, 560, 34, 42, 9905, 25, 1, 1095, 258, 575, 861, 126, 535, 2025, 1064, 105, 1487, 485, 217, 345, 191, 1071, 220, 936, 96, 760, 305, 62, 2546, 79, 5, 65, 895, 38, 1926, 7, 77, 80, 27, 28, 295, 42, 284, 9569, 298, 39, 62, 78, 831, 56, 3, 66, 3, 18, 56, 497, 45, 548, 189, 63, 65, 12323, 72, 354, 497, 1006, 11511, 439, 2644, 95, 169, 514, 6, 159, 23, 63, 24, 18, 171, 8, 532, 36}
var unSizeList = []int{firstArraySize, 64316, secondArraySize, 1313, 84, 173, 146, 974, 8369, 771, 555, 3983, 15693, 254, 1945, 11934, 1722, 6122, 1150, 893, 3, 234, 1670, 776, 2335, 1296, 150, 2215, 3518, 535, 1435, 561, 761, 1266, 278, 1347, 352, 1695, 8824, 609, 262, 998, 108, 832, 316, 818, 9, 233, 36, 24, 660, 214, 261, 903, 560, 34, 42, 9905, 25, 1, 1095, 258, 575, 861, 126, 535, 2025, 1064, 105, 1487, 485, 217, 345, 191, 1071, 220, 936, 96, 760, 305, 62, 2546, 79, 5, 65, 895, 38, 1926, 7, 77, 80, 27, 28, 295, 42, 284, 9569, 298, 39, 62, 78, 831, 56, 3, 66, 3, 18, 56, 497, 45, 548, 189, 63, 65, 12323, 72, 354, 497, 1006, 11511, 439, 2644, 95, 169, 514, 6, 159, 23, 63, 24, 18, 171, 8, 532, 36}

// 135
const takeSizeList = 135

func BenchmarkUnionCurrent_Many(b *testing.B) {
	unSizeList = unSizeList[:takeSizeList]
	arrs := make([][]int, len(unSizeList))
	for i, s := range unSizeList {
		arrs[i] = randArray(s, maxArrayValue)
	}
	b.ResetTimer()
	var union iterator.Iterator
	for i := 0; i < b.N; i++ {
		union = iterator.NewArrayIter(arrs[0])
		for j := 1; j < len(arrs); j++ {
			union = iterator.NewUnionIter(union, iterator.NewArrayIter(arrs[j]))
		}
		for {
			if _, ok := union.Next(); !ok {
				break
			}
		}
	}
}

func BenchmarkUnionComplex_Many(b *testing.B) {
	unSizeList = unSizeList[:takeSizeList]
	vecs := make([]vector, len(unSizeList))
	for i, s := range unSizeList {
		for _, x := range randArray(s, maxArrayValue) {
			vecs[i].Add(uint32(x))
		}
	}
	fvs := make([]Iterator, len(unSizeList))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < len(vecs); j++ {
			fvs[j] = NewFetchVec(vecs[j])
		}
		union := NewUnionComplex(fvs...)
		for {
			if _, _, ok := union.Next(); !ok {
				break
			}
		}
	}
}

func BenchmarkUnionVec_Many(b *testing.B) {
	unSizeList = unSizeList[:takeSizeList]
	vecs := make([]vector, len(unSizeList))
	for i, s := range unSizeList {
		for _, x := range randArray(s, maxArrayValue) {
			vecs[i].Add(uint32(x))
		}
	}
	var union Iterator
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		union = NewFetchVec(vecs[0])
		for j := 1; j < len(vecs); j++ {
			union = NewUnionVec(union, NewFetchVec(vecs[j]))
		}
		for {
			if _, _, ok := union.Next(); !ok {
				break
			}
		}
	}
}
