package vector

import "unsafe"

type hashmatch struct {
	temp [256]uint8
}

func (hm *hashmatch) clear() {
	p := (*[32]uint64)(unsafe.Pointer(&hm.temp[0]))
	p[0x00], p[0x01], p[0x02], p[0x03], p[0x04], p[0x05], p[0x06], p[0x07] = 0, 0, 0, 0, 0, 0, 0, 0
	p[0x08], p[0x09], p[0x0a], p[0x0b], p[0x0c], p[0x0d], p[0x0e], p[0x0f] = 0, 0, 0, 0, 0, 0, 0, 0
	p[0x10], p[0x11], p[0x12], p[0x13], p[0x14], p[0x15], p[0x16], p[0x17] = 0, 0, 0, 0, 0, 0, 0, 0
	p[0x18], p[0x19], p[0x1a], p[0x1b], p[0x1c], p[0x1d], p[0x1e], p[0x1f] = 0, 0, 0, 0, 0, 0, 0, 0
}

func (hm *hashmatch) apply(b []uint8) {
	n := len(b)
	var p *[8]uint8
	for n >= 8 {
		n -= 8
		p = (*[8]uint8)(unsafe.Pointer(&b[n]))
		hm.temp[p[0]] = 1
		hm.temp[p[1]] = 1
		hm.temp[p[2]] = 1
		hm.temp[p[3]] = 1
		hm.temp[p[4]] = 1
		hm.temp[p[5]] = 1
		hm.temp[p[6]] = 1
		hm.temp[p[7]] = 1
	}
	for n >= 2 {
		n -= 2
		p = (*[8]uint8)(unsafe.Pointer(&b[n]))
		hm.temp[p[0]] = 1
		hm.temp[p[1]] = 1
	}
	if n == 1 {
		hm.temp[b[0]] = 1
	}
}

func (hm *hashmatch) inter(b []uint8) (r int) {
	i := 8
	n := len(b)
	var x uint8
	var p *[8]uint8
	for i <= n {
		p = (*[8]uint8)(unsafe.Pointer(&b[i-8]))
		x = hm.temp[p[0]]
		hm.temp[r] = p[0]
		r += int(x)
		x = hm.temp[p[1]]
		hm.temp[r] = p[1]
		r += int(x)
		x = hm.temp[p[2]]
		hm.temp[r] = p[2]
		r += int(x)
		x = hm.temp[p[3]]
		hm.temp[r] = p[3]
		r += int(x)
		x = hm.temp[p[4]]
		hm.temp[r] = p[4]
		r += int(x)
		x = hm.temp[p[5]]
		hm.temp[r] = p[5]
		r += int(x)
		x = hm.temp[p[6]]
		hm.temp[r] = p[6]
		r += int(x)
		x = hm.temp[p[7]]
		hm.temp[r] = p[7]
		r += int(x)
		i += 8
	}
	i -= 8
	i += 2
	for i <= n {
		p = (*[8]uint8)(unsafe.Pointer(&b[i-2]))
		x = hm.temp[p[0]]
		hm.temp[r] = p[0]
		r += int(x)
		x = hm.temp[p[1]]
		hm.temp[r] = p[1]
		r += int(x)
		i += 2
	}
	i -= 2
	i += 1
	if i == n {
		x = hm.temp[b[n-1]]
		hm.temp[r] = b[n-1]
		r += int(x)
	}
	return
}
