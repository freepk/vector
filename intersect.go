package vector

type IntersectVec struct {
	a, b Iterator
}

func NewIntersectVec(a, b Iterator) *IntersectVec {
	i := &IntersectVec{
		a: a,
		b: b,
	}
	i.Reset()
	return i
}

func (i *IntersectVec) Reset() {
	i.a.Reset()
	i.b.Reset()
}

func (i *IntersectVec) Next() (base uint16, bits [4]uint64, ok bool) {
	var (
		base2 uint16
		bits2 [4]uint64
	)
	if base, bits, ok = i.a.Next(); !ok {
		return
	}
	if base2, bits2, ok = i.b.Next(); !ok {
		return
	}
	for {
		if base < base2 {
			if base, bits, ok = i.a.Next(); !ok {
				return
			}
			continue
		}
		if base > base2 {
			if base2, bits2, ok = i.b.Next(); !ok {
				return
			}
			continue
		}
		bits[0] &= bits2[0]
		bits[1] &= bits2[1]
		bits[2] &= bits2[2]
		bits[3] &= bits2[3]
		if bits[0]|bits[1]|bits[2]|bits[3] > 0 {
			break
		}
		if base, bits, ok = i.a.Next(); !ok {
			return
		}
		if base2, bits2, ok = i.b.Next(); !ok {
			return
		}
	}
	return
}

type intersectItem struct {
	vec  Iterator
	base uint16
	bits [4]uint64
}

type IntersectComplex struct {
	items []intersectItem
}

func (ic *IntersectComplex) Reset() {
	for i := 0; i < len(ic.items); i++ {
		ic.items[i].vec.Reset()
		if i > 1 {
			ic.items[i].base, ic.items[i].bits, _ = ic.items[i].vec.Next()
		}
	}
}

func (ic *IntersectComplex) Next() (base uint16, bits [4]uint64, ok bool) {
	var (
		base2 uint16
		bits2 [4]uint64
	)
	for {
		if base, bits, ok = ic.items[0].vec.Next(); !ok {
			return
		}
		if base2, bits2, ok = ic.items[1].vec.Next(); !ok {
			return
		}
		for {
			if base < base2 {
				if base, bits, ok = ic.items[0].vec.Next(); !ok {
					return
				}
				continue
			}
			if base > base2 {
				if base2, bits2, ok = ic.items[1].vec.Next(); !ok {
					return
				}
				continue
			}
			bits[0] &= bits2[0]
			bits[1] &= bits2[1]
			bits[2] &= bits2[2]
			bits[3] &= bits2[3]
			ok = bits[0]|bits[1]|bits[2]|bits[3] != 0
			break
		}
		if !ok {
			continue
		}
		for i := 2; i < len(ic.items); i++ {
			for ic.items[i].base < base {
				ic.items[i].base, ic.items[i].bits, ok = ic.items[i].vec.Next()
				if !ok {
					return
				}
			}
			if ic.items[i].base > base {
				ok = false
				break
			}
			bits[0] &= ic.items[i].bits[0]
			bits[1] &= ic.items[i].bits[1]
			bits[2] &= ic.items[i].bits[2]
			bits[3] &= ic.items[i].bits[3]
			if ok = bits[0]|bits[1]|bits[2]|bits[3] != 0; !ok {
				break
			}
		}
		if ok {
			return
		}
	}
}

func NewIntersectComplex(source ...Iterator) *IntersectComplex {
	ic := &IntersectComplex{items: make([]intersectItem, len(source))}
	for i, v := range source {
		ic.items[i].vec = v
	}
	return ic
}
