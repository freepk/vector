package vector

type vectorHeader struct {
	base uint16
	mask uint16
	link uint32
}

type Vector struct {
	head []vectorHeader
	data []uint64
}

func NewVector() *Vector {
	return &Vector{}
}

func (vec *Vector) Clear() {
	vec.head = vec.head[:0]
	vec.data = vec.data[:0]
}

func extract(n uint32) (base uint16, mask uint16, data uint64) {
	base = uint16(n >> 8)
	mask = 1 << (uint8(n) >> 6)
	data = 1 << ((uint8(n) << 2) >> 2)
	return
}

func (vec *Vector) Add(n uint32) {
	last := len(vec.head)
	base, mask, data := extract(n)
	if last == 0 {
		vec.head = append(vec.head, vectorHeader{
			base: base,
			mask: mask,
			link: uint32(len(vec.data))})
		vec.data = append(vec.data, data)
	} else {
		last--
		if base > vec.head[last].base {
			vec.head = append(vec.head, vectorHeader{
				base: base,
				mask: mask,
				link: uint32(len(vec.data))})
			vec.data = append(vec.data, data)
		} else {
			if mask > vec.head[last].mask {
				vec.head[last].mask |= mask
				vec.data = append(vec.data, data)
			} else {
				link := vec.head[last].link
				vec.data[link] |= data
			}
		}
	}
}

type VectorIter struct {
	pos int
	vec *Vector
}

func NewVectorIter(vec *Vector) *VectorIter {
	return &VectorIter{vec: vec}
}

func (vi *VectorIter) Reset() {
	vi.pos = 0
}

func (vi *VectorIter) hasNext() bool {
	last := len(vi.vec.head)
	if last == 0 {
		return false
	}
	if vi.pos > (last - 1) {
		return false
	}
	return true
}

func (vi *VectorIter) Next() (base uint16, data [4]uint64, ok bool) {
	if !vi.hasNext() {
		return
	}
	base = vi.vec.head[vi.pos].base
	link := vi.vec.head[vi.pos].link
	switch vi.vec.head[vi.pos].mask {
	case 0b0000:
	case 0b0001:
		data[0] = vi.vec.data[link]
	case 0b0010:
		data[1] = vi.vec.data[link]
	case 0b0011:
		data[0] = vi.vec.data[link]
		data[1] = vi.vec.data[link+1]
	case 0b0100:
		data[2] = vi.vec.data[link]
	case 0b0101:
		data[0] = vi.vec.data[link]
		data[2] = vi.vec.data[link+1]
	case 0b0110:
		data[1] = vi.vec.data[link]
		data[2] = vi.vec.data[link+1]
	case 0b0111:
		data[0] = vi.vec.data[link]
		data[1] = vi.vec.data[link+1]
		data[2] = vi.vec.data[link+2]
	case 0b1000:
		data[3] = vi.vec.data[link]
	case 0b1001:
		data[0] = vi.vec.data[link]
		data[3] = vi.vec.data[link+1]
	case 0b1010:
		data[1] = vi.vec.data[link]
		data[3] = vi.vec.data[link+1]
	case 0b1011:
		data[0] = vi.vec.data[link]
		data[1] = vi.vec.data[link+1]
		data[3] = vi.vec.data[link+2]
	case 0b1100:
		data[2] = vi.vec.data[link]
		data[3] = vi.vec.data[link+1]
	case 0b1101:
		data[0] = vi.vec.data[link]
		data[2] = vi.vec.data[link+1]
		data[3] = vi.vec.data[link+2]
	case 0b1110:
		data[1] = vi.vec.data[link]
		data[2] = vi.vec.data[link+1]
		data[3] = vi.vec.data[link+2]
	case 0b1111:
		data[0] = vi.vec.data[link]
		data[1] = vi.vec.data[link+1]
		data[2] = vi.vec.data[link+2]
		data[3] = vi.vec.data[link+3]
	}
	ok = true
	vi.pos++
	return
}

// func (vi *VectorIter) Advance(n uint32) (base uint16, data [4]uint64, ok bool) {
// 	if !vi.hasNext() {
// 		return
// 	}
// 	_base := uint16(n >> 8)
// 	one := vi.pos
// 	two := len(vi.vec.head)
// 	for one != two {
// 		mid := (one + two) / 2
// 		base = vi.vec.head[mid].base
// 		if base == _base {
// 			break
// 		} else if base > _base {
// 			one = mid + 1
// 		} else {
// 			two = mid - 1
// 		}
// 	}
// 	ok = true
// 	vi.pos++
// 	return
// }
