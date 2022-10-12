package vector

import (
	"testing"
)

func TestVectorAdd(t *testing.T) {
	v := NewVector()
	for i := 0; i < 256; i++ {
		v.Add(i)
	}
}
