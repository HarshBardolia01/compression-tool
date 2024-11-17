package multiset

import (
	"fmt"

	rbt "github.com/HarshBardolia01/go-data-structures/redblacktree"
)

type Multiset[K comparable, V any] struct {
	tree *rbt.Tree[K, V]
	size int
}

func NewMultiset[K comparable, V any](comp func(k1, k2 K) int) *Multiset[K, V] {
	return &Multiset[K, V]{
		tree: rbt.NewTree[K, V](comp, true),
	}
}

func (ms *Multiset[K, V]) Insert(key K, val V) {
	ms.size++
	ms.tree.Insert(key, val)
}

func (ms *Multiset[K, V]) EraseFirst() {
	ms.size--
	node := ms.tree.Root.GetLeftmostNode()
	ms.tree.Erase(node.Key)
}

func (ms *Multiset[K, V]) Size() int {
	return ms.size
}

func (ms *Multiset[K, V]) Begin() V {
	return ms.tree.Root.GetLeftmostNode().Value
}

func (ms *Multiset[K, V]) GetTreeSize() int {
	return ms.tree.Len()
}

func (ms *Multiset[K, V]) Print() {
	var dfs func(node *rbt.Node[K, V])

	dfs = func(node *rbt.Node[K, V]) {
		if node.IsNil() {
			return
		}

		dfs(node.Left)
		fmt.Printf("%d: %v\n", node.Key, node.Value)
		dfs(node.Right)
	}

	if ms.tree.Len() == 0 {
		fmt.Println("Empty tree")
		return
	}

	dfs(ms.tree.Root)
}
