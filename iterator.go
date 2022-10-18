package vector

type Iterator interface {
	Reset()
	Next() (base uint16, data []uint8, ok bool)
}

type Iterator2 interface {
	Reset()
	Next() (base uint16, data [4]uint64, ok bool)
}
