package vector

import (
	"testing"
)

func TestHashMatch(t *testing.T) {
	a := make([]uint8, 256)
	for i := 0; i < 256; i++ {
		a[i] = uint8(i)
	}
	hm := NewHashMatch()
	hm.Apply(a)
	if len(hm.InterZip(a)) != 256 {
		t.Fail()
	}
}