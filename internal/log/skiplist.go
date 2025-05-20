package log

import (
	"math/rand"
	"sync/atomic"
)

type skiplist struct {
	rnd *rand.Rand
	head *Node
}

func NewSkipList() *skiplist {
	return &skiplist{}
}

type Node struct {
	value atomic.Uint64
	height uint16
}

func newNode() *Node {
	return &Node{}
}
