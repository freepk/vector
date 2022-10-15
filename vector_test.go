package vector

import (
	"os"
	"testing"
)

const (
	firstArraySize  = 128 * 1024
	secondArraySize = 1024 * 1024
	maxArrayValue   = 16 * 1024 * 1024
)

var (
	firstArrayInt  = randArray(firstArraySize, maxArrayValue)
	secondArrayInt = randArray(secondArraySize, maxArrayValue)
)

func TestVectorAdd(t *testing.T) {
	vec := NewVector()
	for _, n := range firstArrayInt[:256] {
		vec.Add(n)
	}
	it := NewVectorIter(vec)
	for {
		base, data, ok := it.Next()
		if !ok {
			break
		}
		t.Log(base, data)
	}
}

func TestVector2Add(t *testing.T) {
	vec := NewVector2()
	for _, n := range firstArrayInt[:256] {
		vec.Add(n)
	}
	it := NewVector2Iter(vec)
	for {
		base, mask, data, ok := it.Next()
		if !ok {
			break
		}
		t.Log(base, mask, data)
	}
}

func TestVectorDump(t *testing.T) {
	vec := NewVector()
	for _, n := range firstArrayInt {
		vec.Add(n)
	}
	os.WriteFile("first.vector.bin", vec.Bytes(), 0666)
	vec.Clear()
	for _, n := range secondArrayInt {
		vec.Add(n)
	}
	os.WriteFile("second.vector.bin", vec.Bytes(), 0666)
}

func TestVector2Dump(t *testing.T) {
	vec := NewVector2()
	for _, n := range firstArrayInt {
		vec.Add(n)
	}
	os.WriteFile("first.vector2.bin", vec.Bytes(), 0666)
	vec.Clear()
	for _, n := range secondArrayInt {
		vec.Add(n)
	}
	os.WriteFile("second.vector2.bin", vec.Bytes(), 0666)
}
