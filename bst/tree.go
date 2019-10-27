package bsarbol

import "fmt"

type BST struct {
	Root *Node
}

type Node struct {
	Data        int
	Left, Right *Node
}

//                 7
//              /    \
//            5       10
//         /    \    /  \
//        3      6  9   12
//      /  \
//     1    4

// Pretty print?
//                        3
//                       / \
//                      1   2
//
//            3                     3
//           / \                   / \
//          1   2                 1   2
//
//      3         3              3       3
//     / \       / \            / \     / \
//    1   2     1   2          1   2   1   2
//
//  3     3     3     3       3     3
// / \   / \   / \   / \     / \   / \
//1   2 1   2 1   2 1   2   1   2 1   2

// Traverse method traverses the tree applying a function
// to each Data element of the tree nodes.
// The tree can be traversed Pre Order, In Order and Post Order.
func (r *BST) Traverse(fn func(value int), traverseType string) {

	if traverseType == "" {
		traverseType = "inorder"
	}

	switch traverseType {
	case "preorder":
		r.Root.PreOrderTraversal(fn)
	case "postorder":
		r.Root.PostOrderTraversal(fn)
	default:
		r.Root.InOrderTraversal(fn)
	}
}

// PreOrderTraversal is called by the Traverse method
// to traverse the tree in Pre Order order and applies the supplied function.
func (n *Node) PreOrderTraversal(fn func(value int)) {
	if n == nil {
		return
	}
	fn(n.Data)
	n.Left.PreOrderTraversal(fn)
	n.Right.PreOrderTraversal(fn)
}

// InOrderTraversal is called by the Traverse method
// to traverse the tree in In Order order and applies the supplied function.
func (n *Node) InOrderTraversal(fn func(value int)) {
	if n == nil {
		return
	}
	n.Left.InOrderTraversal(fn)
	fn(n.Data)
	n.Right.InOrderTraversal(fn)
}

// PostOrderTraversal is called by the Traverse method
// to traverse the tree in Post Order order and applies the supplied function.
func (n *Node) PostOrderTraversal(fn func(value int)) {
	if n == nil {
		return
	}
	n.Left.PostOrderTraversal(fn)
	n.Right.PostOrderTraversal(fn)
	fn(n.Data)
}

// PrintNode is an auxiliary function to supply to Traverse method
// and print to stdout the Data element of each node.
func PrintNode(value int) {
	fmt.Printf("Node: %d\n", value)
}

// NaiveInsert inserts the supplied data value into the tree
// is very naive and makes the insertion without balancing the tree.
// Also allocates the Root node in case it's not present.
func (r *BST) NaiveInsert(data int) {
	if r.Root == nil {
		r.Root = &Node{Data: data}
		return
	}
	r.Root.NodeNaiveInsert(data)
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
