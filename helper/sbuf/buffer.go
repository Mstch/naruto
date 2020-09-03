package sbuf

type Buffer struct {
	buf []byte
	off int //take at off,do not export the fk read or write api like bytes.Buffer
}

func NewBuffer() *Buffer{
	return &Buffer{
		off: 0,
	}
}

func (b *Buffer) Take(len int) []byte {
	c := cap(b.buf)
	if len+b.off > c {
		b.buf = make([]byte, len)
		b.off = 0
	}
	b.off += len
	return b.buf[b.off-len : b.off]
}

func (b *Buffer) Reset() {
	b.off = 0
}
