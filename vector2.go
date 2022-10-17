package vector

import "unsafe"

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

func extract(n uint32) (base uint16, mask uint8, data uint32) {
	base = uint16(n >> 8)
	mask = 1 << (uint8(n) >> 5)
	data = 1 << ((uint8(n) << 3) >> 3)
	return
}

func encode(base uint16, size, mask uint8) (n uint32) {
	n = uint32(base)
	n |= uint32(size) << 16
	n |= uint32(mask) << 24
	return
}

func decode(n uint32) (base uint16, size, mask uint8) {
	base = uint16(n)
	size = uint8(n >> 16)
	mask = uint8(n >> 24)
	return
}

func (v *Vector2) Add(n uint32) {
	last := len(v.data)
	base, mask, data := extract(n)
	if last == 0 {
		v.last = last
		v.data = append(v.data, encode(base, 0, mask), data)
	} else {
		_base, _size, _mask := decode(v.data[v.last])
		if base > _base {
			v.last = last
			v.data = append(v.data, encode(base, 0, mask), data)
		} else {
			if mask > _mask {
				v.data[v.last] = encode(_base, (_size + 1), (_mask | mask))
				v.data = append(v.data, data)
			} else {
				v.data[(last - 1)] |= data
			}
		}
	}
}

func (v *Vector2) unpack(data *[4]uint64, pos int) (base uint16, next int) {
	var size, mask uint8
	_ = mask
	base, size, mask = decode(v.data[pos])
	*data = *(*[4]uint64)(unsafe.Pointer(&v.data[(pos + 1)]))
	next = pos + 1 + int(size) + 1
	return
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

func (vi *Vector2Iter) Next() (base uint16, data [4]uint64, ok bool) {
	if !vi.hasNext() {
		return
	}
	base, vi.pos = vi.vec.unpack(&data, vi.pos)
	ok = true
	return
}
