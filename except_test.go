package vector

import (
	"github.com/freepk/iterator"
	"testing"
)

func TestExceptionVec(t *testing.T) {
	a1 := []int{1, 5, 10, 15, 20, 25, 30, 35, 40, 45, 50, 272}
	v1 := CreateVectorFromArray(a1)
	a2 := []int{1, 10, 20, 30, 40, 50, 60, 260, 272}
	v2 := CreateVectorFromArray(a2)
	intersect := NewExceptVec(NewFetchVec(v1), NewFetchVec(v2))
	unpack := NewUnpackVec(intersect)
	for {
		v, ok := unpack.Next()
		if !ok {
			break
		}
		t.Log(v)
	}
}

func TestExceptEquals(t *testing.T) {
	a0 := randArray(firstArraySize, maxValue(firstArraySize))
	v0 := CreateVectorFromArray(a0)
	a1 := randArray(secondArraySize, maxValue(secondArraySize))
	v1 := CreateVectorFromArray(a1)
	it := iterator.NewExceptIter(iterator.NewArrayIter(a0), iterator.NewArrayIter(a1))
	it.Reset()
	c1 := 0
	for {
		if _, ok := it.Next(); !ok {
			break
		}
		c1++
	}
	t.Log(c1)

	iv := NewExceptVec(NewFetchVec(v0), NewFetchVec(v1))
	unpack := NewUnpackVec(iv)
	c2 := 0
	for {
		if _, ok := unpack.Next(); !ok {
			break
		}
		c2++
	}
	t.Log(c2)
}

func TestExcept_SameArrays(t *testing.T) {
	a := []int{1, 10, 20, 30, 40, 50, 60, 70}
	v := CreateVectorFromArray(a)
	it := NewUnpackVec(NewExceptVec(NewFetchVec(v), NewFetchVec(v)))
	if _, ok := it.Next(); ok {
		t.Fatal("Not empty")
	}
}

func TestExcept_Empty(t *testing.T) {
	a := []int{1, 10, 20, 30, 40, 50, 60, 70}
	v := CreateVectorFromArray(a)
	it := NewUnpackVec(NewExceptVec(NewFetchVec(v), NewFetchVec(CreateEmptyVector())))
	i := 0
	for {
		k, ok := it.Next()
		if !ok {
			break
		}
		if k != int32(a[i]) {
			t.Fatal("Not equal", k, a[i])
		}
		i++
	}
	if i != len(a) {
		t.Fatal("Not to end", i, len(a))
	}

	it = NewUnpackVec(NewExceptVec(NewFetchVec(CreateEmptyVector()), NewFetchVec(v)))
	if _, ok := it.Next(); ok {
		t.Fatal("Not empty 1")
	}

	it = NewUnpackVec(NewExceptVec(NewFetchVec(CreateEmptyVector()), NewFetchVec(CreateEmptyVector())))
	if _, ok := it.Next(); ok {
		t.Fatal("Not empty 2")
	}
}

func TestExcept_Reset(t *testing.T) {
	a := []int{1, 10, 20, 30, 40, 50, 60, 70}
	v0 := CreateVectorFromArray(a)
	b := []int{10, 70, 98, 190}
	v1 := CreateVectorFromArray(b)
	expected := []uint32{1, 20, 30, 40, 50, 60}

	it := NewUnpackVec(NewExceptVec(NewFetchVec(v0), NewFetchVec(v1)))
	i := 0
	for {
		k, ok := it.Next()
		if !ok {
			break
		}
		if uint32(k) != expected[i] {
			t.Fatal("Not equal", k, a[i])
		}
		i++
	}
	if i != len(expected) {
		t.Fatal("Not to end", i, len(a))
	}
	if _, ok := it.Next(); ok {
		t.Fatal("Next after end", i, len(a))
	}
	it.Reset()
	i = 0
	for {
		k, ok := it.Next()
		if !ok {
			break
		}
		if uint32(k) != expected[i] {
			t.Fatal("After Reset. Not equal", k, a[i])
		}
		i++
	}
	if i != len(expected) {
		t.Fatal("After Reset. Not to end", i, len(a))
	}
}
