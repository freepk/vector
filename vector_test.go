package vector

import (
	// "os"
	"testing"
)

const (
	firstArraySize  = 256 * 1024
	secondArraySize = 1024 * 1024
	maxArrayValue   = 16 * 1024 * 1024
)

var (
	firstArrayInt  = randArray(firstArraySize, secondArraySize)
	secondArrayInt = randArray(secondArraySize, secondArraySize)
)

func TestVectorAdd(t *testing.T) {
	v := NewVector()
	for i := 0; i < 256; i++ {
		v.Add(i)
	}
	// v.Clear()
	// for _, n := range firstArrayInt {
	// 	v.Add(n)
	// }
	// os.WriteFile("first.bin", v.data, 0666)
	// v.Clear()
	// for _, n := range secondArrayInt {
	// 	v.Add(n)
	// }
	// os.WriteFile("second.bin", v.data, 0666)
}
