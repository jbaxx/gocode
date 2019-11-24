package bsarbol

import (
	"fmt"
)

// Traverse method traverses the tree applying a function
// to each Data element of the tree nodes.
// The tree can be traversed Pre Order, In Order and Post Order.
func (b *BST) Traverse(fn func(value int), traverseType string) {

	if traverseType == "" {
		traverseType = "inorder"
	}

	switch traverseType {
	case "bfs":
		b.Root.BFSTraversal(fn)
	case "preorder":
		b.Root.PreOrderTraversal(fn)
	case "postorder":
		b.Root.PostOrderTraversal(fn)
	default:
		b.Root.InOrderTraversal(fn)
	}
}

// BFSTraversal is called by the Traverse method
// to traverse the tree in Breadth First Search order and applies the supplied function.
func (n *Node) BFSTraversal(fn func(value int)) {
	var buffer []*Node
	var p *Node
	//counter keeps track of remaining elements in quere in current level
	var counter int
	//level keeps track of the depth of the tree
	//TODO: add level as tree/node attribute and write test
	var level int

	buffer = append([]*Node{n}, buffer...)

	for len(buffer) > 0 {
		//Pop on back
		p, buffer = buffer[len(buffer)-1], buffer[:len(buffer)-1]

		fn(p.Data)

		//Push on front
		if p.Left != nil {
			buffer = append([]*Node{p.Left}, buffer...)
		}
		if p.Right != nil {
			buffer = append([]*Node{p.Right}, buffer...)
		}

		if counter == 0 {
			level++
			counter = len(buffer)
		}
		counter--

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

// TraverseLevel method traverses the tree
// while keeps track of the current node depth
// Depth is tracked using the Tracker interface
func (b *BST) TraverseLevel(fnl func(value int, levelTracker Tracker), traverseType string) {

	if traverseType == "" {
		traverseType = "iot"
	}

	//var level *int
	//*level = 0
	lev := NewNodeLevel()
	lev.SetDepth(-1)

	switch traverseType {
	case "iot":
		b.Root.IOT(fnl, lev)
	default:
		b.Root.IOT(fnl, lev)
	}
}

// IOT is called by the TraverseLevel method to keep track of the current node depth,
// traverses the tree in In Order order and applies the supplied function.
// Depth is tracked using the Tracker interface
func (n *Node) IOT(fnl func(value int, levelTracker Tracker), level Tracker) {

	level.Increment()
	defer level.Decrement()

	if n.Left != nil {
		n.Left.IOT(fnl, level)
	}

	fmt.Printf("Level: %d; Data: %d\n", level.GetDepth(), n.Data)
	fnl(n.Data, level)

	if n.Right != nil {
		n.Right.IOT(fnl, level)
	}

}
