package vector

import (
	"github.com/freepk/iterator"
	"sort"
	"testing"
)

func TestUnionComplex_Fetch(t *testing.T) {
	expected := make([]int, 0)
	v1 := CreateEmptyVector()
	n := 30000
	for i := 0; i < n; i++ {
		x := 1 + i*4
		v1.Add(uint32(x))
		expected = append(expected, x)
	}
	v2 := CreateEmptyVector()
	n = 200
	for i := 0; i < n; i++ {
		x := 2 + i*4
		v2.Add(uint32(x))
		expected = append(expected, x)
	}
	v3 := CreateEmptyVector()
	n = 400
	for i := 0; i < n; i++ {
		x := 3 + i*4
		v3.Add(uint32(x))
		expected = append(expected, x)
	}
	v4 := CreateEmptyVector()
	n = 100
	for i := 0; i < n; i++ {
		x := 4 + i*4
		v4.Add(uint32(x))
		expected = append(expected, x)
	}
	sort.Ints(expected)
	union := NewUnionComplex(NewFetchVec(v4), NewFetchVec(v2), NewFetchVec(v3), NewFetchVec(v1))
	unpack := NewUnpackVec(union)
	i := 0
	for {
		v, ok := unpack.Next()
		if !ok {
			break
		}
		if expected[i] != int(v) {
			t.Fatal("unexpected", i, expected[i], v)
		}
		i++
	}
	if i != len(expected) {
		t.Fatal("not full fetched")
	}
}

func TestUnionVec_Fetch(t *testing.T) {
	expected := make([]int, 0)
	v1 := CreateEmptyVector()
	n := 300
	for i := 0; i < n; i++ {
		x := 1 + i*4
		v1.Add(uint32(x))
		expected = append(expected, x)
	}
	v2 := CreateEmptyVector()
	n = 200
	for i := 0; i < n; i++ {
		x := 2 + i*4
		v2.Add(uint32(x))
		expected = append(expected, x)
	}
	v3 := CreateEmptyVector()
	n = 400
	for i := 0; i < n; i++ {
		x := 3 + i*4
		v3.Add(uint32(x))
		expected = append(expected, x)
	}
	v4 := CreateEmptyVector()
	n = 100
	for i := 0; i < n; i++ {
		x := 4 + i*4
		v4.Add(uint32(x))
		expected = append(expected, x)
	}
	sort.Ints(expected)
	union := NewUnionVec(NewFetchVec(v1), NewFetchVec(v2))
	union = NewUnionVec(union, NewFetchVec(v3))
	union = NewUnionVec(union, NewFetchVec(v4))
	unpack := NewUnpackVec(union)
	i := 0
	for {
		v, ok := unpack.Next()
		if !ok {
			break
		}
		if expected[i] != int(v) {
			t.Fatal("unexpected", i, expected[i], v)
		}
		i++
	}
	if i != len(expected) {
		t.Fatal("not full fetched")
	}
}

func TestUnion_Equals(t *testing.T) {
	a0 := randArray(firstArraySize, maxArrayValue)
	v0 := CreateVectorFromArray(a0)
	a1 := randArray(secondArraySize, maxArrayValue)
	v1 := CreateVectorFromArray(a1)
	it := iterator.NewUnionIter(iterator.NewArrayIter(a0), iterator.NewArrayIter(a1))
	c1 := 0
	for {
		if _, ok := it.Next(); !ok {
			break
		}
		c1++
	}

	it2 := NewUnpackVec(NewUnionComplex(NewFetchVec(v0), NewFetchVec(v1)))
	c2 := 0
	for {
		if _, ok := it2.Next(); !ok {
			break
		}
		c2++
	}

	it3 := NewUnpackVec(NewUnionVec(NewFetchVec(v0), NewFetchVec(v1)))
	c3 := 0
	for {
		if _, ok := it3.Next(); !ok {
			break
		}
		c3++
	}
	if ok := c1 == c2 && c1 == c3; !ok {
		t.Fatal(c1, c2, c3)
	}
}

func TestUnion_SameArrays(t *testing.T) {
	a := []int{1, 10, 20, 30, 40, 50, 60, 70}
	v := CreateVectorFromArray(a)
	it := NewUnpackVec(NewUnionComplex(NewFetchVec(v), NewFetchVec(v)))
	i := 0
	for {
		k, ok := it.Next()
		if !ok {
			break
		}
		if k != int32(a[i]) {
			t.Fatal("UnionComplex. Not equal", k, a[i])
		}
		i++
	}
	if i != len(a) {
		t.Fatal("UnionComplex. Not to end", i, len(a))
	}

	it = NewUnpackVec(NewUnionVec(NewFetchVec(v), NewFetchVec(v)))
	i = 0
	for {
		k, ok := it.Next()
		if !ok {
			break
		}
		if k != int32(a[i]) {
			t.Fatal("UnionVec. Not equal", k, a[i])
		}
		i++
	}
	if i != len(a) {
		t.Fatal("UnionVec. Not to end", i, len(a))
	}
}

func TestUnion_Empty(t *testing.T) {
	a := []int{1, 10, 20, 30, 40, 50, 60, 70}
	v := CreateVectorFromArray(a)
	it := NewUnpackVec(NewUnionComplex(NewFetchVec(v), NewFetchVec(CreateEmptyVector())))
	i := 0
	for {
		k, ok := it.Next()
		if !ok {
			break
		}
		if k != int32(a[i]) {
			t.Fatal("UnionComplex. Not equal", k, a[i])
		}
		i++
	}
	if i != len(a) {
		t.Fatal("UnionComplex. Not to end", i, len(a))
	}

	it = NewUnpackVec(NewUnionComplex(NewFetchVec(CreateEmptyVector()), NewFetchVec(CreateEmptyVector())))
	if _, ok := it.Next(); ok {
		t.Fatal("UnionComplex. Not empty")
	}

	it = NewUnpackVec(NewUnionVec(NewFetchVec(v), NewFetchVec(CreateEmptyVector())))
	i = 0
	for {
		k, ok := it.Next()
		if !ok {
			break
		}
		if k != int32(a[i]) {
			t.Fatal("UnionVec. Not equal", k, a[i])
		}
		i++
	}
	if i != len(a) {
		t.Fatal("UnionVec. Not to end", i, len(a))
	}

	it = NewUnpackVec(NewUnionVec(NewFetchVec(CreateEmptyVector()), NewFetchVec(CreateEmptyVector())))
	if _, ok := it.Next(); ok {
		t.Fatal("UnionComplex. Not empty")
	}
}

func TestUnionComplex_Reset(t *testing.T) {
	a := []int{1, 10, 20, 30, 40, 50, 60, 70}
	v0 := CreateVectorFromArray(a)
	b := []int{80, 90}
	v1 := CreateVectorFromArray(b)
	expected := []uint32{1, 10, 20, 30, 40, 50, 60, 70, 80, 90}

	it := NewUnpackVec(NewUnionComplex(NewFetchVec(v0), NewFetchVec(v1)))
	i := 0
	for {
		k, ok := it.Next()
		if !ok {
			break
		}
		if uint32(k) != expected[i] {
			t.Fatal("UnionComplex. Not equal", k, a[i])
		}
		i++
	}
	if i != len(expected) {
		t.Fatal("UnionComplex. Not to end", i, len(a))
	}
	if _, ok := it.Next(); ok {
		t.Fatal("UnionComplex. Next after end", i, len(a))
	}
	it.Reset()
	i = 0
	for {
		k, ok := it.Next()
		if !ok {
			break
		}
		if uint32(k) != expected[i] {
			t.Fatal("UnionComplex. After Reset. Not equal", k, a[i])
		}
		i++
	}
	if i != len(expected) {
		t.Fatal("UnionComplex. After Reset. Not to end", i, len(a))
	}
}

func TestUnionVec_Reset(t *testing.T) {
	a := []int{1, 10, 20, 30, 40, 50, 60, 70}
	v0 := CreateVectorFromArray(a)
	b := []int{80, 90}
	v1 := CreateVectorFromArray(b)
	expected := []uint32{1, 10, 20, 30, 40, 50, 60, 70, 80, 90}

	it := NewUnpackVec(NewUnionVec(NewFetchVec(v0), NewFetchVec(v1)))
	i := 0
	for {
		k, ok := it.Next()
		if !ok {
			break
		}
		if uint32(k) != expected[i] {
			t.Fatal("UnionComplex. Not equal", k, a[i])
		}
		i++
	}
	if i != len(expected) {
		t.Fatal("UnionComplex. Not to end", i, len(a))
	}
	if _, ok := it.Next(); ok {
		t.Fatal("UnionComplex. Next after end", i, len(a))
	}
	it.Reset()
	i = 0
	for {
		k, ok := it.Next()
		if !ok {
			break
		}
		if uint32(k) != expected[i] {
			t.Fatal("UnionComplex. After Reset. Not equal", k, a[i])
		}
		i++
	}
	if i != len(expected) {
		t.Fatal("UnionComplex. After Reset. Not to end", i, len(a))
	}
}
