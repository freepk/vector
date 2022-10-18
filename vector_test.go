package vector

import (
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
