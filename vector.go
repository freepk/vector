package vector

import "unsafe"

type Vector struct {
	last int
	test int
	data []uint8
}

func NewVector() *Vector {
	return &Vector{}
}

func (v *Vector) Clear() {
	v.last = 0
	v.data = v.data[:0]
}

func (v *Vector) Bytes() []uint8 {
	return v.data
}

// func NewVectorEx(data []uint8, last, test int) *Vector {
// 	return &Vector{
// 		data: data,
// 		last: last}
// }

type vectorElem struct {
	base uint16
	size uint8
}

func newVectorElem(base uint16, size uint8) *vectorElem {
	return &vectorElem{base, size}
}

func (ve *vectorElem) bytes() []uint8 {
	ptr := (*[3]uint8)(unsafe.Pointer(ve))
	return ptr[:3]
}

func (v *Vector) elem(n int) *vectorElem {
	return (*vectorElem)(unsafe.Pointer(&v.data[n]))
}

func (v *Vector) lastElem() *vectorElem {
	return v.elem(v.last)
}

func (v *Vector) Add(n uint32) {
	last := len(v.data)
	base := uint16(n >> 8)
	data := uint8(n)
	if last == 0 {
		v.last = last
		v.data = append(v.data, newVectorElem(base, 0).bytes()...)
		v.data = append(v.data, data)
		return
	} else {
		elem := v.lastElem()
		if base > elem.base {
			v.last = last
			v.data = append(v.data, newVectorElem(base, 0).bytes()...)
			v.data = append(v.data, data)
		} else {
			elem.size++
			v.data = append(v.data, data)
		}
	}
}

type VectorIter struct {
	pos int
	vec *Vector
}

func NewVectorIter(v *Vector) *VectorIter {
	return &VectorIter{vec: v}
}

func (vi *VectorIter) Reset() {
	vi.pos = 0
}

func (vi *VectorIter) hasNext() bool {
	if len(vi.vec.data) == 0 {
		return false
	}
	if vi.pos > vi.vec.last {
		return false
	}
	return true
}

func (vi *VectorIter) Next() (base uint16, data []uint8, ok bool) {
	if !vi.hasNext() {
		return
	}
	i := vi.pos
	base = uint16(vi.vec.data[i])
	i++
	base += uint16(vi.vec.data[i]) << 8
	i++
	size := int(vi.vec.data[i]) + 1
	i++
	vi.pos = i + size
	data = vi.vec.data[i:vi.pos]
	ok = true
	return
}
