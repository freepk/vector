package vector

type Iterator interface {
	Reset()
	Next() (base uint16, tail []uint8, ok bool)
}
