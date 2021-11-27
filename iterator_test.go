package vector

import "testing"

func TestFetchVec_Solid(t *testing.T) {
	v := CreateVectorRange(0, 255)
	unpack := NewUnpackVec(NewFetchVec(v))
	i := 0
	for {
		x, ok := unpack.Next()
		if !ok {
			break
		}
		if int32(i) != x {
			t.Fatal("Not equal", x, i)
		}
		i++
	}
	if i != 256 {
		t.Fatal("Not to end")
	}
}

func TestFetchVec_WithDuplicate(t *testing.T) {
	v := CreateEmptyVector()
	for i := 0; i < 6; i++ {
		v.Add(uint32(i))
		v.Add(uint32(i))
	}
	unpack := NewUnpackVec(NewFetchVec(v))
	i := 0
	for {
		x, ok := unpack.Next()
		if !ok {
			break
		}
		if int32(i) != x {
			t.Fatal("Not equal", x, i)
		}
		i++
	}
	if i != 6 {
		t.Fatal("Not to end")
	}
}

func TestFetchUnpackVector(t *testing.T) {
	a := []int32{8, 320, 1536, 1544, 266752, 266800, 791088}
	vec := CreateVectorFromArray(a)
	it := NewFetchVec(vec)
	unpack := NewUnpackVec(it)
	i := 0
	for {
		v, ok := unpack.Next()
		if !ok {
			break
		}
		if v != a[i] {
			t.Fatal("Not equal", v, a[i], i)
		}
		i++
	}
	if i != len(a) {
		t.Fatal("Not to end")
	}
	if _, ok := unpack.Next(); ok {
		t.Fatal("Next after end")
	}
	unpack.Reset()
	i = 0
	for {
		v, ok := unpack.Next()
		if !ok {
			break
		}
		if v != a[i] {
			t.Fatal("Not equal", v, a[i], i)
		}
		i++
	}
	if i != len(a) {
		t.Fatal("Not to end")
	}
}

func TestFetchUnpackVector_Empty(t *testing.T) {
	it := NewUnpackVec(NewFetchVec(CreateEmptyVector()))
	if _, ok := it.Next(); ok {
		t.Fatal("Next ok 1")
	}
	it.Reset()
	if _, ok := it.Next(); ok {
		t.Fatal("Next ok 2")
	}
}
