package vector

type vecIter struct {
	pos int
	off int
	vec *vector
}

func newVecIter(vec *vector) *vecIter {
	return &vecIter{pos: 0, vec: vec}
}

func (vi *vecIter) next() bool {
	if vi.hasNext() {
		vi.off += int(vi.currSize())
		vi.pos++
		return true
	}
	return false
}

func (vi *vecIter) hasNext() bool {
	return (vi.pos + 1) < len(vi.vec.head)
}

func (vi *vecIter) currBase() uint16 {
	return vi.vec.head[vi.pos].base
}

func (vi *vecIter) currSize() uint8 {
	return vi.vec.head[vi.pos].size
}

func (vi *vecIter) currData() []uint8 {
	i := vi.off
	n := i + int(vi.currSize())
	return vi.vec.data[i:n]
}
