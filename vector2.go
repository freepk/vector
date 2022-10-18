package vector

import "unsafe"

func extract2(n uint32) (base uint16, mask uint8, data uint32) {
	base = uint16(n >> 8)
	mask = 1 << (uint8(n) >> 5)
	data = 1 << ((uint8(n) << 3) >> 3)
	return
}

func encode2(base uint16, size, mask uint8) (n uint32) {
	n = uint32(base)
	n |= uint32(size) << 16
	n |= uint32(mask) << 24
	return
}

func decode2(n uint32) (base uint16, size, mask uint8) {
	base = uint16(n)
	size = uint8(n >> 16)
	mask = uint8(n >> 24)
	return
}

type Vector2 struct {
	last int
	data []uint32
}

func NewVector2() *Vector2 {
	return &Vector2{}
}

func (v *Vector2) Clear() {
	v.last = 0
	v.data = v.data[:0]
}

func (v *Vector2) Add(n uint32) {
	last := len(v.data)
	base, mask, data := extract2(n)
	if last == 0 {
		v.last = last
		v.data = append(v.data, encode2(base, 0, mask), data)
	} else {
		_base, _size, _mask := decode2(v.data[v.last])
		if base > _base {
			v.last = last
			v.data = append(v.data, encode2(base, 0, mask), data)
		} else {
			if mask > _mask {
				v.data[v.last] = encode2(_base, (_size + 1), (_mask | mask))
				v.data = append(v.data, data)
			} else {
				v.data[(last - 1)] |= data
			}
		}
	}
}

func (v *Vector2) Bytes() []uint8 {
	n := len(v.data) * 4
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

func (vi *Vector2Iter) Next() (base uint16, data [8]uint32, ok bool) {
	if !vi.hasNext() {
		return
	}
	var mask uint8
	base, _, mask = decode2(vi.vec.data[vi.pos])
	vi.pos++
	switch mask & 0b00000011 {
	case 3:
		data[0] = vi.vec.data[vi.pos]
		vi.pos++
		data[1] = vi.vec.data[vi.pos]
		vi.pos++
	case 2:
		data[0] = vi.vec.data[vi.pos]
		vi.pos++
	case 1:
		data[1] = vi.vec.data[vi.pos]
		vi.pos++
	}
	switch mask & 0b00001100 {
	case (3 << 2):
		data[2] = vi.vec.data[vi.pos]
		vi.pos++
		data[3] = vi.vec.data[vi.pos]
		vi.pos++
	case (2 << 2):
		data[2] = vi.vec.data[vi.pos]
		vi.pos++
	case (1 << 2):
		data[3] = vi.vec.data[vi.pos]
		vi.pos++
	}
	switch mask & 0b00110000 {
	case (3 << 4):
		data[4] = vi.vec.data[vi.pos]
		vi.pos++
		data[5] = vi.vec.data[vi.pos]
		vi.pos++
	case (2 << 4):
		data[4] = vi.vec.data[vi.pos]
		vi.pos++
	case (1 << 4):
		data[5] = vi.vec.data[vi.pos]
		vi.pos++
	}
	switch mask & 0b11000000 {
	case (3 << 6):
		data[6] = vi.vec.data[vi.pos]
		vi.pos++
		data[7] = vi.vec.data[vi.pos]
		vi.pos++
	case (2 << 6):
		data[6] = vi.vec.data[vi.pos]
		vi.pos++
	case (1 << 6):
		data[7] = vi.vec.data[vi.pos]
		vi.pos++
	}
	ok = true
	return
}
