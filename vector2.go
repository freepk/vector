package vector

import "unsafe"

func extract2(n uint32) (base uint16, mask uint8, data uint64) {
	base = uint16(n >> 8)
	mask = 1 << (uint8(n) >> 6)
	data = 1 << ((uint8(n) << 2) >> 2)
	return
}

func encode2(base uint16, size, mask uint8) (n uint64) {
	n = uint64(base)
	n |= uint64(size) << 16
	n |= uint64(mask) << 24
	return
}

func decode2(n uint64) (base uint16, size, mask uint8) {
	base = uint16(n)
	size = uint8(n >> 16)
	mask = uint8(n >> 24)
	return
}

type Vector2 struct {
	last int
	data []uint64
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
	n := len(v.data) * 8
	p := (*[0x7fffffff]uint8)(unsafe.Pointer(&v.data[0]))
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
	var mask, size uint8
	base, size, mask = decode2(vi.vec.data[vi.pos])
	vi.pos++
	switch mask & 0b1111 {
	case 0b0000:
	case 0b0001:
		data[0] = vi.vec.data[vi.pos]
	case 0b0010:
		data[1] = vi.vec.data[vi.pos]
	case 0b0011:
		data[0] = vi.vec.data[vi.pos]
		data[1] = vi.vec.data[vi.pos+1]
	case 0b0100:
		data[2] = vi.vec.data[vi.pos]
	case 0b0101:
		data[0] = vi.vec.data[vi.pos]
		data[2] = vi.vec.data[vi.pos+1]
	case 0b0110:
		data[1] = vi.vec.data[vi.pos]
		data[2] = vi.vec.data[vi.pos+1]
	case 0b0111:
		data[0] = vi.vec.data[vi.pos]
		data[1] = vi.vec.data[vi.pos+1]
		data[2] = vi.vec.data[vi.pos+2]
	case 0b1000:
		data[3] = vi.vec.data[vi.pos]
	case 0b1001:
		data[0] = vi.vec.data[vi.pos]
		data[3] = vi.vec.data[vi.pos+1]
	case 0b1010:
		data[1] = vi.vec.data[vi.pos]
		data[3] = vi.vec.data[vi.pos+1]
	case 0b1011:
		data[0] = vi.vec.data[vi.pos]
		data[1] = vi.vec.data[vi.pos+1]
		data[3] = vi.vec.data[vi.pos+2]
	case 0b1100:
		data[2] = vi.vec.data[vi.pos]
		data[3] = vi.vec.data[vi.pos+1]
	case 0b1101:
		data[0] = vi.vec.data[vi.pos]
		data[2] = vi.vec.data[vi.pos+1]
		data[3] = vi.vec.data[vi.pos+2]
	case 0b1110:
		data[1] = vi.vec.data[vi.pos]
		data[2] = vi.vec.data[vi.pos+1]
		data[3] = vi.vec.data[vi.pos+2]
	case 0b1111:
		data[0] = vi.vec.data[vi.pos]
		data[1] = vi.vec.data[vi.pos+1]
		data[2] = vi.vec.data[vi.pos+2]
		data[3] = vi.vec.data[vi.pos+3]
	}
	vi.pos += int(size) + 1
	ok = true
	return
}
