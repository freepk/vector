package vector

import "math"

type header struct {
	base uint16
	size uint8
	mask uint8
}

type vector struct {
	head []header
	data []uint8
	last uint32
}

func CreateEmptyVector() (v vector) {
	v.Reset()
	return
}

func CreateVectorFromArray(a []int) (v vector) {
	v.Reset()
	for _, x := range a {
		v.Add(uint32(x))
	}
	return
}

func CreateVectorRange(a, b int) (v vector) {
	v.Reset()
	for i := a; i <= b; i++ {
		v.Add(uint32(i))
	}
	return
}

func (v *vector) Reset() {
	v.head = v.head[:0]
	v.data = v.data[:0]
	v.last = math.MaxUint32
}

func (v *vector) Add(n uint32) {
	if n == v.last {
		return
	}
	base := uint16(n >> 8)
	if len(v.head) > 0 {
		last := len(v.head) - 1
		if v.head[last].base == base {
			v.head[last].size++
		} else if v.head[last].base < base {
			v.head = append(v.head, header{base: base, size: 0})
		} else {
			panic(`wrong data`)
		}
	} else {
		v.head = append(v.head, header{base: uint16(n >> 8), size: 0})
	}
	v.data = append(v.data, uint8(n))
	v.last = n
}

func (v *vector) size() int {
	return cap(v.head)*4 + cap(v.data)
}

func (v *vector) len() int {
	return len(v.head)*4 + len(v.data)
}
