package vector

import (
	// "os"
	"testing"
)

const (
	firstArraySize  = 128 * 1024
	secondArraySize = 1024 * 1024
	maxArrayValue   = 2 * 1024 * 1024
)

var (
	firstArrayInt  = randArray(firstArraySize, maxArrayValue)
	secondArrayInt = randArray(secondArraySize, maxArrayValue)
)

func TestVectorAdd(t *testing.T) {
	vec := NewVector()
	for _, v := range secondArrayInt {
		vec.Add(uint32(v))
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
}

func TestVector2Add(t *testing.T) {
	vec := NewVector2()
	for _, v := range secondArrayInt {
		vec.Add(uint32(v))
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
}

// func TestVectorCompareBaseCount(t *testing.T) {
// 	n := len(firstArrayInt)
// 	v1 := NewVector()
// 	v2 := NewVector2()
// 	it1 := NewVectorIter(v1)
// 	it2 := NewVector2Iter(v2)
// 	for n > 0 {
// 		v1.Clear()
// 		v2.Clear()
// 		for _, v := range firstArrayInt[:n] {
// 			v1.Add(uint32(v))
// 			v2.Add(uint32(v))
// 		}
// 		it1.Reset()
// 		it2.Reset()
// 		for {
// 			base1, _, ok1 := it1.Next()
// 			base2, _, ok2 := it2.Next()
// 			if base1 != base2 || ok1 != ok2 {
// 				t.Fatal("Broken data", base1, base2)
// 			}
// 			if !ok1 {
// 				break
// 			}
// 		}
// 		n--
// 	}
// }

// func TestVectorDump(t *testing.T) {
// 	vec := NewVector()
// 	for _, v := range secondArrayInt {
// 		vec.Add(uint32(v))
// 	}
// 	os.WriteFile("vector.bin", vec.Bytes(), 0666)
// }

// func TestVector2Dump(t *testing.T) {
// 	vec := NewVector2()
// 	for _, v := range secondArrayInt {
// 		vec.Add(uint32(v))
// 	}
// 	os.WriteFile("vector2.bin", vec.Bytes(), 0666)
// }

func BenchmarkVectorNext(b *testing.B) {
	vec := NewVector()
	for _, v := range secondArrayInt {
		vec.Add(uint32(v))
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
	for _, v := range secondArrayInt {
		vec.Add(uint32(v))
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
