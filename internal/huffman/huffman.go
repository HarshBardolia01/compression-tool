package huffman

import (
	"github.com/HarshBardolia01/compression-tool/internal/multiset"
)

type TreeNode struct {
	freq  int
	char  string
	left  *TreeNode
	right *TreeNode
}

func (node *TreeNode) IsNil() bool {
	return node == nil
}

func (node *TreeNode) IsLeaf() bool {
	if node.IsNil() {
		return false
	}

	return node.left.IsNil() && node.right.IsNil()
}

func BuildTree(freqMap map[byte]int) *TreeNode {
	ms := multiset.NewMultiset[int, *TreeNode](func(k1, k2 int) int {
		return k1 - k2
	})

	for ch, cnt := range freqMap {
		curNode := &TreeNode{
			freq:  cnt,
			char:  string(ch),
			left:  nil,
			right: nil,
		}

		ms.Insert(cnt, curNode)
	}

	for ms.Size() > 1 {
		left := ms.Begin()
		ms.EraseFirst()

		right := ms.Begin()
		ms.EraseFirst()

		top := &TreeNode{
			freq:  left.freq + right.freq,
			char:  left.char + right.char,
			left:  left,
			right: right,
		}

		ms.Insert(left.freq+right.freq, top)
	}

	return ms.Begin()
}

func (node *TreeNode) GetEncoding() map[string]string {
	enc := make(map[string]string)
	var helper func(cur *TreeNode, path string)

	helper = func(cur *TreeNode, path string) {
		if cur.IsNil() {
			return
		}

		if len(cur.char) == 1 {
			enc[cur.char] = path
		}

		helper(cur.left, path+"0")
		helper(cur.right, path+"1")
	}

	helper(node, "")
	return enc
}
