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
	os.WriteFile("vector.bin", vec.Bytes(), 0666)
}

func TestVector2Dump(t *testing.T) {
	vec := NewVector2()
	for _, n := range firstArrayInt {
		vec.Add(n)
	}
	os.WriteFile("vector2.bin", vec.Bytes(), 0666)
}

func BenchmarkVectorNext(b *testing.B) {
	vec := NewVector()
	for _, n := range secondArrayInt {
		vec.Add(n)
	}
	it := NewVectorIter(vec)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		it.Reset()
		for {
			_, _, ok := it.Next()
			if !ok {
				break
			}
		}
	}
}

func BenchmarkVector2Next(b *testing.B) {
	vec := NewVector2()
	for _, n := range secondArrayInt {
		vec.Add(n)
	}
	it := NewVector2Iter(vec)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		it.Reset()
		for {
			_, _, _, ok := it.Next()
			if !ok {
				break
			}
		}
	}
}
