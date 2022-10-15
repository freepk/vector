package vector

type Vector struct {
	last int
	test int
	data []uint8
}

func NewVector() *Vector {
	return &Vector{}
}

func NewVectorEx(data []uint8, last, test int) *Vector {
	return &Vector{
		data: data,
		last: last,
		test: test}
}

func (v *Vector) Add(n int) {
	m := len(v.data)
	a := uint8(n >> 16)
	b := uint8(n >> 8)
	c := uint8(n)
	if m == 0 {
		v.last = m
		v.test = n
		v.data = append(v.data, b, a, 0, c)
		return
	}
	if n > v.test {
		i := v.last
		if b == v.data[i] {
			i++
			if a == v.data[i] {
				i++
				v.test = n
				v.data[i]++
				v.data = append(v.data, c)
				return
			}
		}
		v.last = m
		v.test = n
		v.data = append(v.data, b, a, 0, c)
		return
	}
}

func (v *Vector) Clear() {
	v.last = 0
	v.test = 0
	v.data = v.data[:0]
}

func (v *Vector) Bytes() []uint8 {
	return v.data
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

func (vi *VectorIter) Next() (base uint16, tail []uint8, ok bool) {
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
	tail = vi.vec.data[i:vi.pos]
	ok = true
	return
}
