package vector

import (
	"github.com/RoaringBitmap/roaring"
	"os"
	"testing"
)

func TestRoaringDump(t *testing.T) {
	rb1 := roaring.New()
	for _, x := range firstArrayInt {
		rb1.Add(uint32(x))
	}
	rb1.RunOptimize()
	buf1, _ := rb1.ToBytes()
	os.WriteFile("roaring.bin", buf1, 0666)
}

func BenchmarkRoaringIntersectSmall(b *testing.B) {
	rb1 := roaring.New()
	for _, x := range randArray(60, 256) {
		rb1.Add(uint32(x))
	}
	rb1.RunOptimize()
	rb2 := roaring.New()
	for _, x := range randArray(252, 256) {
		rb2.Add(uint32(x))
	}
	rb2.RunOptimize()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		roaring.And(rb1, rb2)
	}
}

func BenchmarkRoaringIntersect(b *testing.B) {
	rb1 := roaring.New()
	for _, x := range firstArrayInt {
		rb1.Add(uint32(x))
	}
	rb1.RunOptimize()
	rb2 := roaring.New()
	for _, x := range secondArrayInt {
		rb2.Add(uint32(x))
	}
	rb2.RunOptimize()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		roaring.And(rb1, rb2)
	}
}
