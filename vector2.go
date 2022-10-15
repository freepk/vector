package vector

import "unsafe"
import "math/bits"

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
				v.data[(size - 1)] |= data
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

type Vector2Iter struct {
	pos int
	vec *Vector2
}

func NewVector2Iter(v *Vector2) *Vector2Iter {
	return &Vector2Iter{vec: v}
}

func (vi *Vector2Iter) Reset() {
	vi.pos = 0
}

func (vi *Vector2Iter) hasNext() bool {
	if len(vi.vec.data) == 0 {
		return false
	}
	if vi.pos > vi.vec.last {
		return false
	}
	return true
}

func (vi *Vector2Iter) Next() (base uint16, mask uint16, data []uint16, ok bool) {
	if !vi.hasNext() {
		return
	}
	i := vi.pos
	base = vi.vec.data[i]
	i++
	mask = vi.vec.data[i]
	i++
	vi.pos = i + bits.OnesCount16(mask)
	data = vi.vec.data[i:vi.pos]
	ok = true
	return
}
