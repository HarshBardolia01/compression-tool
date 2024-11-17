package multiset

import (
	rbt "github.com/HarshBardolia01/go-data-structures/redblacktree"
)

type Multiset[K comparable, V any] struct {
	tree *rbt.Tree[K, V]
}

func NewMultiset[K comparable, V any](comp func(k1, k2 K) int) *Multiset[K, V] {
	return &Multiset[K, V]{
		tree: rbt.NewTree[K, V](comp, true),
	}
}

func (ms *Multiset[K, V]) Insert(key K, val V) {
	ms.tree.Insert(key, val)
}

func (ms *Multiset[K, V]) EraseFirst() {
	node := ms.tree.Root.GetLeftmostNode()
	ms.tree.Erase(node.Key)
}
