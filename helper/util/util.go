package util

import (
	"encoding/binary"
)

func UInt32ToBytes(data uint32) []byte {
	buf := make([]byte, 4)
	binary.BigEndian.PutUint32(buf, data)
	return buf
}
func Int32ToBytes(data int32) []byte {
	return UInt32ToBytes(uint32(data))
}

func WriteUInt32ToBytes(data uint32, buf []byte) {
	binary.BigEndian.PutUint32(buf, data)
}

func BytesToUInt32(data []byte) uint32 {
	return binary.BigEndian.Uint32(data)
}

func BytesToInt32(data []byte) int32 {
	return int32(BytesToInt32(data))
}
