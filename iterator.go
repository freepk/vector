package vector

type Iterator interface {
	Reset()
	Advance(base uint16)
	Next() (base uint16, data [4]uint64, ok bool)
}
