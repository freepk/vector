package vector

import (
	"github.com/freepk/iterator"
	iterator2 "github.com/iGordienko/iterator"
	"testing"
)

const (
	firstArraySize  = 128 * 1024
	secondArraySize = 1024 * 1024
	maxArrayValue   = 16 * 1024 * 1024
)

func init() {
	println(`firstArraySize`, firstArraySize,
		`secondArraySize`, secondArraySize,
		`maxValue`, maxArrayValue)
}

func intersectArr(a, b []int) {
	an := len(a)
	bn := len(b)
	i := 0
	j := 0
	for (i < an) && (j < bn) {
		if a[i] < b[j] {
			i++
		} else if a[i] > b[j] {
			j++
		} else {
			i++
			j++
		}
	}
}

func unionArr(a, b []int) {
	an := len(a)
	bn := len(b)
	i := 0
	j := 0
	for (i < an) && (j < bn) {
		if a[i] < b[j] {
			i++
		} else if a[i] > b[j] {
			j++
		} else {
			i++
			j++
		}
	}
	for i < an {
		i++
	}
	for j < bn {
		j++
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

func unionVec(v0, v1 *vector) {
	vi0 := newVecIter(v0)
	vi1 := newVecIter(v1)
	rx := [4]uint64{0, 0, 0, 0}
	for vi0.hasNext() && vi1.hasNext() {
		if vi0.currBase() < vi1.currBase() {
			rx[0], rx[1], rx[2], rx[3] = 0, 0, 0, 0
			bytesToBits(&rx, vi0.currData())
			vi0.next()
		} else if vi0.currBase() > vi1.currBase() {
			rx[0], rx[1], rx[2], rx[3] = 0, 0, 0, 0
			bytesToBits(&rx, vi0.currData())
			vi1.next()
		} else {
			rx[0], rx[1], rx[2], rx[3] = 0, 0, 0, 0
			bytesToBits(&rx, vi0.currData())
			bytesToBits(&rx, vi1.currData())
			vi0.next()
			vi1.next()
		}
	}
	for vi0.hasNext() {
		rx[0], rx[1], rx[2], rx[3] = 0, 0, 0, 0
		bytesToBits(&rx, vi0.currData())
		vi0.next()
	}
	for vi1.hasNext() {
		rx[0], rx[1], rx[2], rx[3] = 0, 0, 0, 0
		bytesToBits(&rx, vi0.currData())
		vi1.next()
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

func BenchmarkIGogoIterIntersect(b *testing.B) {
	a0 := randArray(firstArraySize, maxArrayValue)
	a1 := randArray(secondArraySize, maxArrayValue)
	it := iterator2.NewFasterIntersectionIterator(iterator2.NewArrayIter(a0), iterator2.NewArrayIter(a1))
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
	v0 := CreateEmptyVector()
	for _, x := range a0 {
		v0.Add(uint32(x))
	}
	a1 := randArray(secondArraySize, maxArrayValue)
	v1 := CreateEmptyVector()
	for _, x := range a1 {
		v1.Add(uint32(x))
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		intersectVec(&v0, &v1)
	}
}

func BenchmarkIterUnion(b *testing.B) {
	a0 := randArray(firstArraySize, maxArrayValue)
	a1 := randArray(secondArraySize, maxArrayValue)
	it := iterator.NewUnionIter(iterator.NewArrayIter(a0), iterator.NewArrayIter(a1))
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

func BenchmarkIGogoUnionIntersect(b *testing.B) {
	a0 := randArray(firstArraySize, maxArrayValue)
	a1 := randArray(secondArraySize, maxArrayValue)
	it := iterator2.NewFasterUnionIterator(iterator2.NewArrayIter(a0), iterator2.NewArrayIter(a1))
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

func BenchmarkArrayUnion(b *testing.B) {
	a0 := randArray(firstArraySize, maxArrayValue)
	a1 := randArray(secondArraySize, maxArrayValue)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		unionArr(a0, a1)
	}
}

func BenchmarkVectorUnion(b *testing.B) {
	a0 := randArray(firstArraySize, maxArrayValue)
	v0 := CreateEmptyVector()
	for _, x := range a0 {
		v0.Add(uint32(x))
	}
	a1 := randArray(secondArraySize, maxArrayValue)
	v1 := CreateEmptyVector()
	for _, x := range a1 {
		v1.Add(uint32(x))
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		unionVec(&v0, &v1)
	}
}
