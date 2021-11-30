package vector

type exceptItem struct {
	vec  Iterator
	base uint16
	bits [4]uint64
	ok   bool
}

type ExceptVec struct {
	it     Iterator
	except exceptItem
}

func (ex *ExceptVec) Reset() {
	ex.it.Reset()
	ex.except.vec.Reset()
	ex.except.base, ex.except.bits, ex.except.ok = ex.except.vec.Next()
}

func (ex *ExceptVec) Next() (base uint16, bits [4]uint64, ok bool) {
	for {
		if base, bits, ok = ex.it.Next(); !ok {
			return
		}
		if !ex.except.ok {
			return
		}
		for {
			if base < ex.except.base {
				return
			} else if base > ex.except.base {
				if ex.except.base, ex.except.bits, ex.except.ok = ex.except.vec.Next(); !ex.except.ok {
					return
				}
				continue
			} else {
				bits[0] ^= bits[0] & ex.except.bits[0]
				bits[1] ^= bits[1] & ex.except.bits[1]
				bits[2] ^= bits[2] & ex.except.bits[2]
				bits[3] ^= bits[3] & ex.except.bits[3]
				if bits[0]|bits[1]|bits[2]|bits[3] > 0 {
					return
				}
				break
			}
		}
	}
}

func NewExceptVec(b, except Iterator) *ExceptVec {
	ex := &ExceptVec{
		it:     b,
		except: exceptItem{vec: except},
	}
	except.Reset()
	return ex
}
