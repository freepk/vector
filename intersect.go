package vector

type IntersectIter struct {
	a, b Iterator
}

func NewIntersectIter(a, b Iterator) *IntersectIter {
	it := &IntersectIter{a: a, b: b}
	it.Reset()
	return it
}

func (it *IntersectIter) Reset() {
	it.a.Reset()
	it.b.Reset()
}

func (it *IntersectIter) Advance(base uint16) {
	it.a.Advance(base)
	it.b.Advance(base)
}

func (it *IntersectIter) Next() (abase uint16, adata [4]uint64, ok bool) {
	var bbase uint16
	var bdata [4]uint64
	if abase, adata, ok = it.a.Next(); !ok {
		return
	}
	if bbase, bdata, ok = it.b.Next(); !ok {
		return
	}
	for {
		if abase < bbase {
			it.a.Advance(bbase)
			if abase, adata, ok = it.a.Next(); !ok {
				return
			}
			continue
		}
		if abase > bbase {
			it.b.Advance(abase)
			if bbase, bdata, ok = it.b.Next(); !ok {
				return
			}
			continue
		}
		adata[0] &= bdata[0]
		adata[1] &= bdata[1]
		adata[2] &= bdata[2]
		adata[3] &= bdata[3]
		if (adata[0] | adata[1] | adata[2] | adata[3]) > 0 {
			ok = true
			return
		}
		if abase, adata, ok = it.a.Next(); !ok {
			return
		}
		if bbase, bdata, ok = it.b.Next(); !ok {
			return
		}
	}
}
