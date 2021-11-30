package vector

type Iterator interface {
	Reset()
	Next() (uint16, [4]uint64, bool)
}

type FetchVec struct {
	pos int
	off int
	vec vector
}

func (f *FetchVec) Reset() {
	f.pos = 0
	f.off = 0
}

func (f *FetchVec) Next() (base uint16, bitset [4]uint64, ok bool) {
	if f.pos < len(f.vec.head) {
		base = f.vec.head[f.pos].base
		n := f.off + int(f.vec.head[f.pos].size) + 1
		cd := f.vec.data[f.off:n]
		for i := 0; i < len(cd); i++ {
			bitset[(cd[i] >> 6)] |= bitsTable[cd[i]]
		}
		ok = true
		f.pos++
		f.off = n
	} else {
		ok = false
	}
	return
}

func NewFetchVec(vec vector) *FetchVec {
	return &FetchVec{vec: vec}
}

type UnpackVec struct {
	base   uint16
	i      uint16
	values []uint8
	it     Iterator
}

func NewUnpackVec(it Iterator) *UnpackVec {
	u := &UnpackVec{
		base:   0,
		i:      0,
		values: make([]uint8, 0, 256),
		it:     it,
	}
	u.Reset()
	return u
}

func (u *UnpackVec) Reset() {
	u.it.Reset()
	u.i = 0
	u.values = u.values[:0]
}

func (u *UnpackVec) Next() (v int32, ok bool) {
	if int(u.i) >= len(u.values) {
		var r [4]uint64
		if u.base, r, ok = u.it.Next(); !ok {
			return
		}
		u.values = bitsToBytes(r, u.values[:256])
		u.i = 0
	}
	v = int32(u.base)<<8 | int32(u.values[u.i])
	u.i++
	ok = true
	return
}

type Empty struct{}

func (e Empty) Reset() {}

func (e Empty) Next() (base uint16, bits [4]uint64, ok bool) {
	return
}
