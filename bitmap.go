package vector

import "unsafe"

type Bitmap struct {
	last int
	data []uint64
}

func NewBitmap() *Bitmap {
	return &Bitmap{}
}

func (bm *Bitmap) Clear() {
	bm.last = 0
	bm.data = bm.data[:0]
}

func (bm *Bitmap) Add(n uint32) {
	m := len(bm.data)
	b := uint64(uint16(n >> 8))
	p := uint64(1) << (uint8(n) >> 5)
	d := uint64(1) << (uint8(n) << 3 >> 3)
	if m == 0 {
		d <<= 16
		d |= p
		d <<= 16
		d |= b
		bm.last = m
		bm.data = append(bm.data, d)
	} else {
		_d := bm.data[bm.last]
		_b := uint64(uint16(_d))
		_d >>= 16
		_p := uint64(uint8(_d))
		_d >>= 8
		_s := uint64(uint8(_d))
		_d >>= 8
		if b > _b {
			d <<= 16
			d |= p
			d <<= 16
			d |= b
			bm.last = m
			bm.data = append(bm.data, d)
		} else {
			if p > _p {
				_s++
				_p |= p
				_d <<= 8
				_d |= _s
				_d <<= 8
				_d |= _p
				_d <<= 16
				_d |= _b
				bm.data[bm.last] = _d
				switch _s {
				case 0, 2, 4, 6:
					d <<= 32
					bm.data[m-1] |= d
				case 1, 3, 5, 7:
					bm.data = append(bm.data, d)
				default:
					panic("unknown size")
				}
			} else {
				switch _s {
				case 0, 2, 4, 6:
					d <<= 32
					bm.data[m-1] |= d
				case 1, 3, 5, 7:
					bm.data[m-1] |= d
				default:
					panic("unknown size")
				}
			}
		}
	}
}

func (bm *Bitmap) Bytes() []uint8 {
	n := len(bm.data) * 8
	p := (*[0xffffffff]uint8)(unsafe.Pointer(&bm.data[0]))
	return p[:n]
}

type BitmapIter struct {
	p int
	b *Bitmap
}

func NewBitmapIter(b *Bitmap) *BitmapIter {
	return &BitmapIter{p: 0, b: b}
}

func (it *BitmapIter) Reset() {
	it.p = 0
}

func (it *BitmapIter) hasNext() bool {
	if len(it.b.data) == 0 {
		return false
	}
	if it.p > it.b.last {
		return false
	}
	return true
}

func (it *BitmapIter) Next() (base uint16, data [4]uint64, ok bool) {
	if !it.hasNext() {
		return
	}
	d := it.b.data[it.p]
	b := uint16(d)
	d >>= 16
	p := uint8(d)
	d >>= 8
	s := uint8(d)
	d >>= 8
	it.p += ((int(s) + 1) / 2) + 1
	base = b
	ok = true
	return
}
