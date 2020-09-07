package stupid

type PoolType uint8

const (
	SBufPool PoolType = iota
	ProtoRWPool
	SizedBufPool
	NoPool
)

