package sbuf

import rbt "github.com/emirpasic/gods/trees/redblacktree"

type BufferTree struct {
	tree *rbt.Tree
}

func NewTree() *BufferTree {
	return &BufferTree{
		tree: rbt.NewWithIntComparator(),
	}
}

func (b *BufferTree) Size() int {
	return b.tree.Size()
}

func (b *BufferTree) Take(takeSize int) []byte {
	if node, ok := b.tree.Ceiling(takeSize); ok {
		buf := node.Value.([]byte)
		if cap(buf) > 2*takeSize {
			buf = make([]byte, takeSize)
			b.tree.Put(takeSize, buf)
		}
		return buf[:takeSize]
	} else {
		buf := make([]byte, takeSize)
		b.tree.Put(takeSize, buf)
		return buf[:takeSize]
	}
}
