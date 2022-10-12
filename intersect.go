package vector

type IntersectIter struct {
	ai Iterator
	bi Iterator
	hm HashMatch
}

func NewIntersectIter(a, b Iterator) *IntersectIter {
	it := &IntersectIter{ai: a, bi: b}
	it.Reset()
	return it
}

func (it *IntersectIter) Reset() {
	it.ai.Reset()
	it.bi.Reset()
}

func (it *IntersectIter) Next() (ab uint16, at []uint8, ok bool) {
	var bb uint16
	var bt []uint8
	if ab, at, ok = it.ai.Next(); !ok {
		return
	}
	if bb, bt, ok = it.bi.Next(); !ok {
		return
	}
	for {
		if ab < bb {
			if ab, at, ok = it.ai.Next(); !ok {
				return
			}
			continue
		}
		if ab > bb {
			if bb, bt, ok = it.bi.Next(); !ok {
				return
			}
			continue
		}
		it.hm.Clear()
		it.hm.Apply(at)
		at = it.hm.InterZip(bt)
		if len(at) > 0 {
			return
		}
		if ab, at, ok = it.ai.Next(); !ok {
			return
		}
		if bb, bt, ok = it.bi.Next(); !ok {
			return
		}
	}
}
