package vector

type IntersectIter struct {
	a, b Iterator
	c    [256]uint8
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

func (it *IntersectIter) Next() (uint16, []uint8, bool) {
	ab, at, ok := it.a.Next()
	if !ok {
		return 0, nil, false
	}
	bb, bt, ok := it.b.Next()
	if !ok {
		return 0, nil, false
	}
	for {
		if ab < bb {
			if ab, at, ok = it.a.Next(); !ok {
				return 0, nil, false
			}
			continue
		}
		if ab > bb {
			if bb, bt, ok = it.b.Next(); !ok {
				return 0, nil, false
			}
			continue
		}
		it.c = [256]uint8{}
		for _, v := range at {
			it.c[v] = 1
		}
		i := uint8(0)
		for _, v := range bt {
			j := it.c[v]
			it.c[i] = v
			i += j
		}
		if i > 0 {
			return ab, it.c[:i], true
		}
		if ab, at, ok = it.a.Next(); !ok {
			return 0, nil, false
		}
		if bb, bt, ok = it.b.Next(); !ok {
			return 0, nil, false
		}
	}
}
