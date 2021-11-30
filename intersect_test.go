package vector

import (
	"github.com/freepk/iterator"
	"testing"
)

func TestIntersectVec(t *testing.T) {
	a1 := []int{1, 5, 10, 15, 20, 25, 30, 35, 40, 45, 50, 272}
	v1 := CreateVectorFromArray(a1)
	a2 := []int{1, 10, 20, 30, 40, 50, 60, 260, 272}
	v2 := CreateVectorFromArray(a2)
	intersect := NewIntersectVec(NewFetchVec(v1), NewFetchVec(v2))
	unpack := NewUnpackVec(intersect)
	for {
		v, ok := unpack.Next()
		if !ok {
			break
		}
		t.Log(v)
	}
}

func TestIntersectComplex(t *testing.T) {
	a1 := []int{1, 5, 10, 15, 20, 25, 30, 35, 40, 45, 50, 272}
	v1 := CreateVectorFromArray(a1)
	a2 := []int{1, 10, 20, 30, 40, 50, 60, 260, 272}
	v2 := CreateVectorFromArray(a2)
	intersect := NewIntersectComplex(NewFetchVec(v1), NewFetchVec(v2))
	unpack := NewUnpackVec(intersect)
	for {
		v, ok := unpack.Next()
		if !ok {
			break
		}
		t.Log(v)
	}
}

func TestIntersectsEqual(t *testing.T) {
	a0 := randArray(firstArraySize, maxValue(firstArraySize))
	v0 := CreateVectorFromArray(a0)
	a1 := randArray(secondArraySize, maxValue(secondArraySize))
	v1 := CreateVectorFromArray(a1)
	it := iterator.NewInterIter(iterator.NewArrayIter(a0), iterator.NewArrayIter(a1))
	it.Reset()
	c1 := 0
	for {
		if _, ok := it.Next(); !ok {
			break
		}
		c1++
	}

	iv := NewIntersectVec(NewFetchVec(v0), NewFetchVec(v1))
	unpack := NewUnpackVec(iv)
	c2 := 0
	for {
		if _, ok := unpack.Next(); !ok {
			break
		}
		c2++
	}

	ic := NewIntersectComplex(NewFetchVec(v0), NewFetchVec(v1))
	unpack = NewUnpackVec(ic)
	c3 := 0
	for {
		if _, ok := unpack.Next(); !ok {
			break
		}
		c3++
	}

	if !(c1 == c2 && c1 == c3) {
		t.Fatal("Not equal", c1, c2, c3)
	}
}

func TestIntersectsManyEqual(t *testing.T) {
	arrs := make([][]int, len(intsctSizeList))
	vecs := make([]vector, len(intsctSizeList))
	for i, s := range intsctSizeList {
		arrs[i] = randArray(s, maxValue(s))
		for _, x := range arrs[i] {
			vecs[i].Add(uint32(x))
		}
	}
	var it iterator.Iterator
	it = iterator.NewArrayIter(arrs[0])
	for j := 1; j < len(arrs); j++ {
		it = iterator.NewInterIter(it, iterator.NewArrayIter(arrs[j]))
	}
	c1 := 0
	for {
		if _, ok := it.Next(); !ok {
			break
		}
		c1++
	}

	fvs := make([]Iterator, len(intsctSizeList))
	for j := 0; j < len(vecs); j++ {
		fvs[j] = NewFetchVec(vecs[j])
	}
	ic := NewIntersectComplex(fvs...)
	unpack := NewUnpackVec(ic)
	c2 := 0
	for {
		if _, ok := unpack.Next(); !ok {
			break
		}
		c2++
	}

	var iv Iterator
	iv = NewFetchVec(vecs[0])
	for j := 1; j < len(vecs); j++ {
		iv = NewIntersectVec(iv, NewFetchVec(vecs[j]))
	}
	unpack = NewUnpackVec(iv)
	c3 := 0
	for {
		if _, ok := unpack.Next(); !ok {
			break
		}
		c3++
	}
	if !(c1 == c2 && c1 == c3) {
		t.Fatal("Not equal", c1, c2, c3)
	}
}

func TestIntersect_SameArrays(t *testing.T) {
	a := []int{1, 10, 20, 30, 40, 50, 60, 70}
	v := CreateVectorFromArray(a)
	it := NewUnpackVec(NewIntersectComplex(NewFetchVec(v), NewFetchVec(v)))
	i := 0
	for {
		k, ok := it.Next()
		if !ok {
			break
		}
		if k != int32(a[i]) {
			t.Fatal("IntersectComplex. Not equal", k, a[i])
		}
		i++
	}
	if i != len(a) {
		t.Fatal("IntersectComplex. Not to end", i, len(a))
	}

	it = NewUnpackVec(NewIntersectVec(NewFetchVec(v), NewFetchVec(v)))
	i = 0
	for {
		k, ok := it.Next()
		if !ok {
			break
		}
		if k != int32(a[i]) {
			t.Fatal("IntersectVec. Not equal", k, a[i])
		}
		i++
	}
	if i != len(a) {
		t.Fatal("IntersectVec. Not to end", i, len(a))
	}
}

func TestIntersect_Empty(t *testing.T) {
	a := []int{1, 10, 20, 30, 40, 50, 60, 70}
	v := CreateVectorFromArray(a)
	it := NewUnpackVec(NewIntersectComplex(NewFetchVec(v), NewFetchVec(CreateEmptyVector())))
	if _, ok := it.Next(); ok {
		t.Fatal("IntersectComplex. Not empty 1")
	}

	it = NewUnpackVec(NewIntersectComplex(NewFetchVec(CreateEmptyVector()), NewFetchVec(CreateEmptyVector())))
	if _, ok := it.Next(); ok {
		t.Fatal("IntersectComplex. Not empty 2")
	}

	it = NewUnpackVec(NewIntersectVec(NewFetchVec(v), NewFetchVec(CreateEmptyVector())))
	if _, ok := it.Next(); ok {
		t.Fatal("IntersectVec. Not empty 1")
	}

	it = NewUnpackVec(NewIntersectVec(NewFetchVec(CreateEmptyVector()), NewFetchVec(CreateEmptyVector())))
	if _, ok := it.Next(); ok {
		t.Fatal("IntersectVec. Not empty 2")
	}
}

func TestIntersectComplex_Reset(t *testing.T) {
	a := []int{1, 10, 20, 30, 40, 50, 60, 70}
	v0 := CreateVectorFromArray(a)
	b := []int{10, 33, 50, 80, 90}
	v1 := CreateVectorFromArray(b)
	expected := []int32{10, 50}

	it := NewUnpackVec(NewIntersectComplex(NewFetchVec(v0), NewFetchVec(v1)))
	i := 0
	for {
		k, ok := it.Next()
		if !ok {
			break
		}
		if k != expected[i] {
			t.Fatal("IntersectComplex. Not equal", k, expected[i])
		}
		i++
	}
	if i != len(expected) {
		t.Fatal("IntersectComplex. Not to end", i, len(a))
	}
	if _, ok := it.Next(); ok {
		t.Fatal("IntersectComplex. Next after end", i, len(a))
	}
	it.Reset()
	i = 0
	for {
		k, ok := it.Next()
		if !ok {
			break
		}
		if k != expected[i] {
			t.Fatal("IntersectComplex. After Reset. Not equal", k, a[i])
		}
		i++
	}
	if i != len(expected) {
		t.Fatal("IntersectComplex. After Reset. Not to end", i, len(a))
	}
}

func TestIntersectVec_Reset(t *testing.T) {
	a := []int{1, 10, 20, 30, 40, 50, 60, 70}
	v0 := CreateVectorFromArray(a)
	b := []int{10, 33, 50, 80, 90}
	v1 := CreateVectorFromArray(b)
	expected := []int32{10, 50}

	it := NewUnpackVec(NewIntersectVec(NewFetchVec(v0), NewFetchVec(v1)))
	i := 0
	for {
		k, ok := it.Next()
		if !ok {
			break
		}
		if k != expected[i] {
			t.Fatal("IntersectComplex. Not equal", k, a[i])
		}
		i++
	}
	if i != len(expected) {
		t.Fatal("IntersectComplex. Not to end", i, len(a))
	}
	if _, ok := it.Next(); ok {
		t.Fatal("IntersectComplex. Next after end", i, len(a))
	}
	it.Reset()
	i = 0
	for {
		k, ok := it.Next()
		if !ok {
			break
		}
		if k != expected[i] {
			t.Fatal("IntersectComplex. After Reset. Not equal", k, a[i])
		}
		i++
	}
	if i != len(expected) {
		t.Fatal("IntersectComplex. After Reset. Not to end", i, len(a))
	}
}
