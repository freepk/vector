package vector

import "unsafe"

type Vector2 struct {
	last int
	data []uint16
}

func NewVector2() *Vector2 {
	return &Vector2{}
}

func (v *Vector2) Clear() {
	v.last = 0
	v.data = v.data[:0]
}

func unpack(n int) (base uint16, mask uint16, data uint16) {
	return uint16(n >> 8),
		1 << (uint8(n) >> 4),
		1 << ((uint8(n) << 4) >> 4)
}

func (v *Vector2) Add(n int) {
	size := len(v.data)
	base, mask, data := unpack(n)
	if size == 0 {
		v.last = size
		v.data = append(v.data, base, mask, data)
	} else {
		i := v.last
		if base == v.data[i] {
			i++
			if mask > v.data[i] {
				v.data[i] |= mask
				v.data = append(v.data, data)
			} else {
				i++
				v.data[i] |= data
			}
		} else {
			v.last = size
			v.data = append(v.data, base, mask, data)
		}
	}
}

func (v *Vector2) Bytes() []uint8 {
	n := len(v.data) * 2
	p := (*[0xffffffff]uint8)(unsafe.Pointer(&v.data[0]))
	return p[:n]
}
