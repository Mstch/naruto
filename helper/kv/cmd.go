package kv

type Cmd struct {
	Opt   string
	Key   []byte
	Value []byte
}

func (c *Cmd) Apply() {

}
