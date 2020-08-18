package kv

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
)

func (c *Cmd) ToGobBytes() []byte {
	buf := bytes.NewBuffer(make([]byte, 16))
	_ = gob.NewEncoder(buf).Encode(c)
	return buf.Bytes()
}

func (c *Cmd) ToJsonBytes() []byte {
	jBytes, _ := json.Marshal(c)
	return jBytes
}
