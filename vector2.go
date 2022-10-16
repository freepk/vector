package vector

import "unsafe"

type Vector2 struct {
	last uint32
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

type vectorElem struct {
	base uint16
	size uint8
	mask uint8
}

func newVectorElem(base uint16, size, mask uint8) *vectorElem {
	return &vectorElem{base, size, mask}
}

func (ve *vectorElem) uint32() *uint32 {
	return (*uint32)(unsafe.Pointer(ve))
}

func (v *Vector2) elem(n uint32) *vectorElem {
	return (*vectorElem)(unsafe.Pointer(&v.data[n]))
}

func (v *Vector2) lastElem() *vectorElem {
	return v.elem(v.last)
}

func (v *Vector2) Add(n uint32) {
	last := uint32(len(v.data))
	base, mask, data := extract(n)
	if last == 0 {
		elem := newVectorElem(base, 1, mask)
		v.last = last
		v.data = append(v.data, *elem.uint32(), data)
	} else {
		elem := v.lastElem()
		if base == elem.base {
			if mask > elem.mask {
				elem.size++
				elem.mask |= mask
				v.data = append(v.data, data)
			} else {
				v.data[(last - 1)] |= data
			}
		} else {
			elem := newVectorElem(base, 1, mask)
			v.last = last
			v.data = append(v.data, *elem.uint32(), data)
		}
	}
}

func (v *Vector2) Bytes() []uint8 {
	n := len(v.data) * 4
	p := (*[0xffffffff]uint8)(unsafe.Pointer(&v.data[0]))
	return p[:n]
}

type Vector2Iter struct {
	pos uint32
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

func (vi *Vector2Iter) Next() (base uint16, mask uint8, data []uint32, ok bool) {
	if !vi.hasNext() {
		return
	}
	elem := vi.vec.lastElem()
	base = elem.base
	mask = elem.mask
	vi.pos++
	pos := vi.pos
	vi.pos += uint32(elem.size)
	data = vi.vec.data[pos:vi.pos]
	ok = true
	return
}
