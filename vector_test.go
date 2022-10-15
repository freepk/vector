package vector

import (
	"os"
	"testing"
)

const (
	firstArraySize  = 256 * 1024
	secondArraySize = 1024 * 1024
	maxArrayValue   = 1 * 1024 * 1024
)

var (
	firstArrayInt  = randArray(firstArraySize, maxArrayValue)
	secondArrayInt = randArray(secondArraySize, maxArrayValue)
)

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
