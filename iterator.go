package vector

type Iterator interface {
	Reset()
	Next() (base uint16, data [4]uint64, ok bool)
	//Advance(n uint32) (base uint16, data [4]uint64, ok bool)
}
