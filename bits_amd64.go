//go:build amd64

package vector

import "unsafe"

//go:noescape
func _bitsToBytes64(res unsafe.Pointer, data unsafe.Pointer) int

func bitsToBytes64(res *[256]uint8, data *[4]uint64) int {
	return _bitsToBytes64(unsafe.Pointer(&res[0]), unsafe.Pointer(&data[0]))
}
