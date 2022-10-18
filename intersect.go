package vector

type IntersectIter struct {
	ai Iterator
	bi Iterator
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

func (it *IntersectIter) Next() (ab uint16, ad [4]uint64, ok bool) {
	var bb uint16
	var bd [4]uint64
	if ab, ad, ok = it.ai.Next(); !ok {
		return
	}
	if bb, bd, ok = it.bi.Next(); !ok {
		return
	}
	for {
		if ab < bb {
			if ab, ad, ok = it.ai.Next(); !ok {
				return
			}
			continue
		}
		if ab > bb {
			if bb, bd, ok = it.bi.Next(); !ok {
				return
			}
			continue
		}
		ad[0] &= bd[0]
		ad[1] &= bd[1]
		ad[2] &= bd[2]
		ad[3] &= bd[3]
		if (ad[0] | ad[1] | ad[2] | ad[3]) > 0 {
			ok = true
			return
		}
		if ab, ad, ok = it.ai.Next(); !ok {
			return
		}
		if bb, bd, ok = it.bi.Next(); !ok {
			return
		}
	}
}
