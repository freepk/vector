package vector

func extract(n uint32) (base uint16, mask uint8, data uint64) {
	base = uint16(n >> 8)
	mask = 1 << (uint8(n) >> 6)
	data = 1 << ((uint8(n) << 2) >> 2)
	return
}

func encode(base uint16, mask uint8) (n uint64) {
	n = uint64(base)
	n |= uint64(mask) << 16
	return
}

func decode(n uint64) (base uint16, mask uint8) {
	base = uint16(n)
	mask = uint8(n >> 16)
	return
}

type Vector struct {
	last int
	data []uint64
}

func NewVector() *Vector {
	return &Vector{}
}

func (v *Vector) Clear() {
	v.last = 0
	v.data = v.data[:0]
}

func (v *Vector) Add(n uint32) {
	last := len(v.data)
	base, mask, data := extract(n)
	if last == 0 {
		v.last = last
		v.data = append(v.data, encode(base, mask), data)
	} else {
		_base, _mask := decode(v.data[v.last])
		if base > _base {
			v.last = last
			v.data = append(v.data, encode(base, mask), data)
		} else {
			if mask > _mask {
				v.data[v.last] = encode(_base, (_mask | mask))
				v.data = append(v.data, data)
			} else {
				v.data[last-1] |= data
			}
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

func (vi *VectorIter) Next() (base uint16, data [4]uint64, ok bool) {
	if !vi.hasNext() {
		return
	}
	var mask uint8
	base, mask = decode(vi.vec.data[vi.pos])
	vi.pos++
	switch mask & 0b1111 {
	case 0b0000:
	case 0b0001:
		data[0] = vi.vec.data[vi.pos]
		vi.pos++
	case 0b0010:
		data[1] = vi.vec.data[vi.pos]
		vi.pos++
	case 0b0011:
		data[0] = vi.vec.data[vi.pos]
		data[1] = vi.vec.data[vi.pos+1]
		vi.pos += 2
	case 0b0100:
		data[2] = vi.vec.data[vi.pos]
		vi.pos++
	case 0b0101:
		data[0] = vi.vec.data[vi.pos]
		data[2] = vi.vec.data[vi.pos+1]
		vi.pos += 2
	case 0b0110:
		data[1] = vi.vec.data[vi.pos]
		data[2] = vi.vec.data[vi.pos+1]
		vi.pos += 2
	case 0b0111:
		data[0] = vi.vec.data[vi.pos]
		data[1] = vi.vec.data[vi.pos+1]
		data[2] = vi.vec.data[vi.pos+2]
		vi.pos += 3
	case 0b1000:
		data[3] = vi.vec.data[vi.pos]
		vi.pos++
	case 0b1001:
		data[0] = vi.vec.data[vi.pos]
		data[3] = vi.vec.data[vi.pos+1]
		vi.pos += 2
	case 0b1010:
		data[1] = vi.vec.data[vi.pos]
		data[3] = vi.vec.data[vi.pos+1]
		vi.pos += 2
	case 0b1011:
		data[0] = vi.vec.data[vi.pos]
		data[1] = vi.vec.data[vi.pos+1]
		data[3] = vi.vec.data[vi.pos+2]
		vi.pos += 3
	case 0b1100:
		data[2] = vi.vec.data[vi.pos]
		data[3] = vi.vec.data[vi.pos+1]
		vi.pos += 2
	case 0b1101:
		data[0] = vi.vec.data[vi.pos]
		data[2] = vi.vec.data[vi.pos+1]
		data[3] = vi.vec.data[vi.pos+2]
		vi.pos += 3
	case 0b1110:
		data[1] = vi.vec.data[vi.pos]
		data[2] = vi.vec.data[vi.pos+1]
		data[3] = vi.vec.data[vi.pos+2]
		vi.pos += 3
	case 0b1111:
		data[0] = vi.vec.data[vi.pos]
		data[1] = vi.vec.data[vi.pos+1]
		data[2] = vi.vec.data[vi.pos+2]
		data[3] = vi.vec.data[vi.pos+3]
		vi.pos += 4
	}
	ok = true
	return
}
