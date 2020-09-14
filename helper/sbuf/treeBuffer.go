package sbuf

import rbt "github.com/emirpasic/gods/trees/redblacktree"

type TreeBuffer struct {
	tree       *rbt.Tree
	concurSafe bool
}

type ConcurSafeNode struct {
	buf []byte

}
func NewTree(concurSafe bool) *TreeBuffer {
	return &TreeBuffer{
		tree:       rbt.NewWithIntComparator(),
		concurSafe: concurSafe,
	}
}

func (b *TreeBuffer) Size() int {
	return b.tree.Size()
}

func (b *TreeBuffer) Take(takeSize int) []byte {
	size := ceilingSize(takeSize)
	if node, ok := b.tree.Ceiling(size); ok {
		buf := node.Value.([]byte)
		if cap(buf) > 2*size {
			buf = make([]byte, size)
			b.tree.Put(size, buf)
		}
		return buf[:takeSize]
	} else {
		buf := make([]byte, size)
		b.tree.Put(size, buf)
		return buf[:takeSize]
	}
}

//ceil size to integer multiple of 64
func ceilingSize(size int) int {
	return size + (64 - size%64)
}
