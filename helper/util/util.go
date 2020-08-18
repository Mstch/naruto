package util

import (
	"bytes"
	"encoding/binary"
)

func Int32ToBytes(data int32) ([]byte, error) {
	bytebuf := bytes.NewBuffer([]byte{})
	err := binary.Write(bytebuf, binary.BigEndian, data)
	return bytebuf.Bytes(), err
}

func BytesToInt32(bys []byte) (int32, error) {
	bytebuff := bytes.NewBuffer(bys)
	var data int32
	err := binary.Read(bytebuff, binary.BigEndian, &data)
	return data, err
}
