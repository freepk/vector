package vector

import "unsafe"

type HashMatch struct {
	h [256]uint8
}

func NewHashMatch() *HashMatch {
	return &HashMatch{}
}

func (hm *HashMatch) Clear() {
	p := (*[32]uint64)(unsafe.Pointer(&hm.h[0]))
	p[0x00], p[0x01], p[0x02], p[0x03], p[0x04], p[0x05], p[0x06], p[0x07] = 0, 0, 0, 0, 0, 0, 0, 0
	p[0x08], p[0x09], p[0x0a], p[0x0b], p[0x0c], p[0x0d], p[0x0e], p[0x0f] = 0, 0, 0, 0, 0, 0, 0, 0
	p[0x10], p[0x11], p[0x12], p[0x13], p[0x14], p[0x15], p[0x16], p[0x17] = 0, 0, 0, 0, 0, 0, 0, 0
	p[0x18], p[0x19], p[0x1a], p[0x1b], p[0x1c], p[0x1d], p[0x1e], p[0x1f] = 0, 0, 0, 0, 0, 0, 0, 0
}

func (hm *HashMatch) Apply(b []uint8) {
	n := len(b)
	var p *[8]uint8
	for n >= 8 {
		n -= 8
		p = (*[8]uint8)(unsafe.Pointer(&b[n]))
		hm.h[p[0]] = 1
		hm.h[p[1]] = 1
		hm.h[p[2]] = 1
		hm.h[p[3]] = 1
		hm.h[p[4]] = 1
		hm.h[p[5]] = 1
		hm.h[p[6]] = 1
		hm.h[p[7]] = 1
	}
	for n >= 2 {
		n -= 2
		p = (*[8]uint8)(unsafe.Pointer(&b[n]))
		hm.h[p[0]] = 1
		hm.h[p[1]] = 1
	}
	if n == 1 {
		hm.h[b[0]] = 1
	}
}

func (hm *HashMatch) InterZip(b []uint8) []uint8 {
	i := 8
	n := len(b)
	var x int
	var r int
	var p *[8]uint8
	for i <= n {
		p = (*[8]uint8)(unsafe.Pointer(&b[i-8]))
		x = int(hm.h[p[0]])
		hm.h[r] = p[0]
		r += x
		x = int(hm.h[p[1]])
		hm.h[r] = p[1]
		r += x
		x = int(hm.h[p[2]])
		hm.h[r] = p[2]
		r += x
		x = int(hm.h[p[3]])
		hm.h[r] = p[3]
		r += x
		x = int(hm.h[p[4]])
		hm.h[r] = p[4]
		r += x
		x = int(hm.h[p[5]])
		hm.h[r] = p[5]
		r += x
		x = int(hm.h[p[6]])
		hm.h[r] = p[6]
		r += x
		x = int(hm.h[p[7]])
		hm.h[r] = p[7]
		r += x
		i += 8
	}
	i -= 8
	i += 2
	for i <= n {
		p = (*[8]uint8)(unsafe.Pointer(&b[i-2]))
		x = int(hm.h[p[0]])
		hm.h[r] = p[0]
		r += x
		x = int(hm.h[p[1]])
		hm.h[r] = p[1]
		r += x
		i += 2
	}
	i -= 2
	i += 1
	if i == n {
		x = int(hm.h[b[n-1]])
		hm.h[r] = b[n-1]
		r += x
	}
	return hm.h[:r]
}
