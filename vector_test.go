package vector

import (
	"github.com/freepk/iterator"
	"testing"
)

const (
	firstArraySize  = 128 * 1024
	secondArraySize = 1024 * 1024
	maxArrayValue   = 16 * 1024 * 1024
)

func intersectArr(a0, a1 []int) {
	i := 0
	j := 0
	for i < len(a0) && j < len(a1) {
		if a0[i] < a1[j] {
			i++
		} else if a0[i] > a1[j] {
			j++
		} else {
			i++
			j++
		}
	}
}

func intersectArr8(a0, a1 []uint8) {
	i := 0
	j := 0
	for i < len(a0) && j < len(a1) {
		if a0[i] < a1[j] {
			i++
		} else if a0[i] > a1[j] {
			j++
		} else {
			i++
			j++
		}
	}
}

func intersectVec(v0, v1 *vector) {
	vi0 := newVecIter(v0)
	vi1 := newVecIter(v1)
	r0 := [4]uint64{0, 0, 0, 0}
	r1 := [4]uint64{0, 0, 0, 0}
	for vi0.hasNext() && vi1.hasNext() {
		if vi0.currBase() < vi1.currBase() {
			vi0.next()
		} else if vi0.currBase() > vi1.currBase() {
			vi1.next()
		} else {
			r0[0], r0[1], r0[2], r0[3] = 0, 0, 0, 0
			r1[0], r1[1], r1[2], r1[3] = 0, 0, 0, 0
			bytesToBits(&r0, vi0.currData())
			bytesToBits(&r1, vi1.currData())
			r0[0] &= r1[0]
			r0[1] &= r1[1]
			r0[2] &= r1[2]
			r0[3] &= r1[3]
			vi0.next()
			vi1.next()
		}
	}
}

func BenchmarkIterIntersect(b *testing.B) {
	a0 := randArray(firstArraySize, maxArrayValue)
	a1 := randArray(secondArraySize, maxArrayValue)
	it := iterator.NewInterIter(iterator.NewArrayIter(a0), iterator.NewArrayIter(a1))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		it.Reset()
		for {
			if _, ok := it.Next(); !ok {
				break
			}
		}
	}
}

func BenchmarkArrayIntersect(b *testing.B) {
	a0 := randArray(firstArraySize, maxArrayValue)
	a1 := randArray(secondArraySize, maxArrayValue)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		intersectArr(a0, a1)
	}
}

func BenchmarkVectorIntersect(b *testing.B) {
	a0 := randArray(firstArraySize, maxArrayValue)
	v0 := newVector()
	for _, x := range a0 {
		v0.addUint24(uint32(x))
	}
	a1 := randArray(secondArraySize, maxArrayValue)
	v1 := newVector()
	for _, x := range a1 {
		v1.addUint24(uint32(x))
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		intersectVec(v0, v1)
	}
}
