package vector

import "math"

type unionItem struct {
	vec  Iterator
	base uint16
	bits [4]uint64
	ok   bool
}

type UnionVec struct {
	a, b unionItem
}

func NewUnionVec(a, b Iterator) *UnionVec {
	u := &UnionVec{
		a: unionItem{
			vec: a,
		},
		b: unionItem{
			vec: b,
		},
	}
	u.Reset()
	return u
}

func (u *UnionVec) Reset() {
	u.a.vec.Reset()
	u.b.vec.Reset()
	u.a.base, u.a.bits, u.a.ok = u.a.vec.Next()
	u.b.base, u.b.bits, u.b.ok = u.b.vec.Next()
}

func (u *UnionVec) Next() (base uint16, bits [4]uint64, ok bool) {
	if ok = u.a.ok || u.b.ok; !ok {
		return
	}
	if u.a.ok && u.b.ok {
		if u.a.base < u.b.base {
			base = u.a.base
			bits = u.a.bits
			u.a.base, u.a.bits, u.a.ok = u.a.vec.Next()
			return
		}
		if u.a.base > u.b.base {
			base = u.b.base
			bits = u.b.bits
			u.b.base, u.b.bits, u.b.ok = u.b.vec.Next()
			return
		}
		base = u.a.base
		bits = u.a.bits
		bits[0] |= u.b.bits[0]
		bits[1] |= u.b.bits[1]
		bits[2] |= u.b.bits[2]
		bits[3] |= u.b.bits[3]
		u.a.base, u.a.bits, u.a.ok = u.a.vec.Next()
		u.b.base, u.b.bits, u.b.ok = u.b.vec.Next()
		return
	}
	if u.a.ok {
		base = u.a.base
		bits = u.a.bits
		u.a.base, u.a.bits, u.a.ok = u.a.vec.Next()
		return
	}
	if u.b.ok {
		base = u.b.base
		bits = u.b.bits
		u.b.base, u.b.bits, u.b.ok = u.b.vec.Next()
		return
	}
	return
}

type UnionComplex struct {
	origin      []Iterator
	items       []unionItem
	currentBase uint16
}

func NewUnionComplex(source ...Iterator) *UnionComplex {
	u := &UnionComplex{origin: source, items: make([]unionItem, 0, len(source))}
	u.Reset()
	return u
}

func (u *UnionComplex) Reset() {
	u.items = u.items[:0]
	min := uint16(math.MaxUint16)
	for i := 0; i < len(u.origin); i++ {
		u.origin[i].Reset()
		if base, bits, ok := u.origin[i].Next(); ok {
			if base < min {
				min = base
			}
			u.items = append(u.items, unionItem{
				vec:  u.origin[i],
				base: base,
				bits: bits,
			})
		}
	}
	u.currentBase = min
}

func (u *UnionComplex) Next() (base uint16, bits [4]uint64, ok bool) {
	if len(u.items) == 0 {
		return
	}
	min := uint16(math.MaxUint16)
	for i := 0; i < len(u.items); {
		x := u.items[i]
		if x.base == u.currentBase {
			bits[0] |= x.bits[0]
			bits[1] |= x.bits[1]
			bits[2] |= x.bits[2]
			bits[3] |= x.bits[3]
			u.items[i].base, u.items[i].bits, ok = x.vec.Next()
			if !ok {
				u.items = append(u.items[:i], u.items[i+1:]...)
				continue
			}
			if u.items[i].base < min {
				min = u.items[i].base
			}
		} else {
			if x.base < min {
				min = x.base
			}
		}
		i++
	}
	ok = true
	base = u.currentBase
	u.currentBase = min
	return
}

type UnionBuilder struct {
	source []Iterator
}

func (b *UnionBuilder) Add(vec vector) {
	b.source = append(b.source, NewFetchVec(vec))
}

func (b *UnionBuilder) Build() (it Iterator) {
	switch len(b.source) {
	case 0:
		return Empty{}
	case 1:
		return b.source[0]
	case 2:
		return NewUnionVec(b.source[0], b.source[1])
	default:
		return NewUnionComplex(b.source...)
	}
}
