package vector

type IntersectIter struct {
	hashmatch
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

func (it *IntersectIter) Next() (ab uint16, at []uint8, ok bool) {
	var bb uint16
	var bt []uint8
	if ab, at, ok = it.a.Next(); !ok {
		return
	}
	if bb, bt, ok = it.b.Next(); !ok {
		return
	}
	for {
		if ab < bb {
			if ab, at, ok = it.a.Next(); !ok {
				return
			}
		} else if ab > bb {
			if bb, bt, ok = it.b.Next(); !ok {
				return
			}
		} else {
			it.clear()
			it.apply(at)
			n := it.inter(bt)
			if n > 0 {
				at = it.temp[:n]
				ok = true
				return
			}
			if ab, at, ok = it.a.Next(); !ok {
				return
			}
			if bb, bt, ok = it.b.Next(); !ok {
				return
			}
		}
	}
}
