package vector

type header struct {
	base uint16
	size uint8
	mask uint8
}

type vector struct {
	head []header
	data []uint8
}

func newVector() *vector {
	return new(vector)
}

func (v *vector) addUint24(n uint32) {
	base := uint16(n >> 8)
	if len(v.head) > 0 {
		last := len(v.head) - 1
		if v.head[last].base == base {
			v.head[last].size++
		} else if v.head[last].base < base {
			v.head = append(v.head, header{base: base, size: 1})
		} else {
			panic(`wrong data`)
		}
	} else {
		v.head = append(v.head, header{base: uint16(n >> 8), size: 1})
	}
	v.data = append(v.data, uint8(n))
}

func (v *vector) size() int {
	return cap(v.head)*4 + cap(v.data)
}

func (v *vector) len() int {
	return len(v.head)*4 + len(v.data)
}
