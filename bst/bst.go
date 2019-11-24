// Package bsarbol implements a Binary Search Tree
package bsarbol

// BST defines a Binary Search Tree structure
type BST struct {
	Root *Node
}

// Node defines the nodes for the Binary Search Tree
type Node struct {
	Data        int
	Left, Right *Node
}

type Tracker interface {
	SetDepth(i int)
	GetDepth() int
	Increment()
	Decrement()
}

//                 7
//              /    \
//            5       10
//         /    \    /  \
//        3      6  9   12
//      /  \
//     1    4

type nodeLevel struct {
	level int
}

func (n *nodeLevel) Increment() {
	n.level++
}

func (n *nodeLevel) Decrement() {
	n.level--
}

func (n *nodeLevel) SetDepth(i int) {
	n.level = i
}

func (n *nodeLevel) GetDepth() int {
	return n.level
}

func NewNodeLevel() *nodeLevel {
	return &nodeLevel{level: 0}
}

// NaiveInsert inserts the supplied data value into the tree
// is very naive and makes the insertion without balancing the tree.
// Also allocates the Root node in case it's not present.
func (b *BST) NaiveInsert(data int) {
	if b.Root == nil {
		b.Root = &Node{Data: data}
		return
	}
	b.Root.NodeNaiveInsert(data)
}

// NodeNaiveInsert is called by the NaiveInsert method to allocate
// the nodes in the tree in binary search order
func (n *Node) NodeNaiveInsert(data int) {
	if data < n.Data {
		if n.Left == nil {
			n.Left = &Node{Data: data}
			return
		}
		n.Left.NodeNaiveInsert(data)
	}
	if data > n.Data {
		if n.Right == nil {
			n.Right = &Node{Data: data}
			return
		}
		n.Right.NodeNaiveInsert(data)
	}
}
