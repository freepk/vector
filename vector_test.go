package vector

import (
	"os"
	"testing"
)

const (
	firstArraySize  = 256 * 1024
	secondArraySize = 1024 * 1024
	maxArrayValue   = 2 * 1024 * 1024
)

var (
	firstArrayInt  = randArray(firstArraySize, maxArrayValue)
	secondArrayInt = randArray(secondArraySize, maxArrayValue)
)

func TestBitmapAdd(t *testing.T) {
	vec := NewBitmap()
	for _, n := range secondArrayInt {
		vec.Add(uint32(n))
	}
	n := 0
	it := NewBitmapIter(vec)
	for {
		base, data, ok := it.Next()
		if !ok {
			break
		}
		_ = base
		_ = data
		n++
	}
	t.Log(n)
}

func TestVectorAdd(t *testing.T) {
	vec := NewVector()
	for _, n := range secondArrayInt {
		vec.Add(n)
	}
	n := 0
	it := NewVectorIter(vec)
	for {
		base, data, ok := it.Next()
		if !ok {
			break
		}
		_ = base
		_ = data
		n++
	}
	t.Log(n)
}

func TestVector2Add(t *testing.T) {
	vec := NewVector2()
	for _, n := range secondArrayInt {
		vec.Add(uint32(n))
	}
	n := 0
	it := NewVector2Iter(vec)
	for {
		base, data, ok := it.Next()
		if !ok {
			break
		}
		_ = base
		_ = data
		n++
	}
	t.Log(n)
}

func TestVectorDump(t *testing.T) {
	vec := NewVector()
	for _, n := range secondArrayInt {
		vec.Add(n)
	}
	os.WriteFile("vector.bin", vec.Bytes(), 0666)
}

func TestVector2Dump(t *testing.T) {
	vec := NewVector2()
	for _, n := range secondArrayInt {
		vec.Add(uint32(n))
	}
	os.WriteFile("vector2.bin", vec.Bytes(), 0666)
}

func TestBitmap2Dump(t *testing.T) {
	vec := NewBitmap()
	for _, n := range secondArrayInt {
		vec.Add(uint32(n))
	}
	os.WriteFile("bitmap.bin", vec.Bytes(), 0666)
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
			base, data, ok := it.Next()
			if !ok {
				break
			}
			_ = base
			_ = data
		}
	}
}

func BenchmarkVector2Next(b *testing.B) {
	vec := NewVector2()
	for _, n := range secondArrayInt {
		vec.Add(uint32(n))
	}
	it := NewVector2Iter(vec)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		it.Reset()
		for {
			base, data, ok := it.Next()
			if !ok {
				break
			}
			_ = base
			_ = data
		}
	}
}

func BenchmarkBitmapNext(b *testing.B) {
	vec := NewBitmap()
	for _, n := range secondArrayInt {
		vec.Add(uint32(n))
	}
	it := NewBitmapIter(vec)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		it.Reset()
		for {
			base, data, ok := it.Next()
			if !ok {
				break
			}
			_ = base
			_ = data
		}
	}
}
