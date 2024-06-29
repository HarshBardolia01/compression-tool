package node

type Node struct {
	value       rune
	weight      int
	left, right *Node
	code        byte
}

func New(value rune, weight int, left *Node, right *Node) *Node {
	return &Node{
		value:  value,
		weight: weight,
		left:   left,
		right:  right,
	}
}

func (node *Node) isLeaf() bool {
	if node == nil {
		return false
	}

	return node.left == nil && node.right == nil
}
