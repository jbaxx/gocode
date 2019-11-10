// Package bsarbol implements a Binary Search Tree
package bsarbol

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strconv"
	"unicode/utf8"
)

// BST defines a Binary Search Tree structure
type BST struct {
	Root *Node
}

// Node defines the nodes for the Binary Search Tree
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

type nodeLevel struct {
	level int
}

func (n *nodeLevel) Increment() {
	n.level++
}

func (n *nodeLevel) Decrement() {
	n.level--
}

func (n *nodeLevel) SetLevel(i int) {
	n.level = i
}

func (n *nodeLevel) GetLevel() int {
	return n.level
}

func NewNodeLevel() *nodeLevel {
	return &nodeLevel{level: 0}
}

type Tracker interface {
	SetLevel(i int)
	GetLevel() int
	Increment()
	Decrement()
}

// TraverseLevel method traverses the tree applying a function
// to each Data element of the tree nodes.
func (b *BST) TraverseLevel(fnl func(value int, levelTracker Tracker), traverseType string) {

	if traverseType == "" {
		traverseType = "iot"
	}

	//var level *int
	//*level = 0
	lev := NewNodeLevel()
	lev.SetLevel(-1)

	switch traverseType {
	case "iot":
		b.Root.IOT(fnl, lev)
	default:
		b.Root.IOT(fnl, lev)
	}
}

// IOT is called by the TraverseLevel method
// to traverse the tree in In Order order and applies the supplied function.
func (n *Node) IOT(fnl func(value int, levelTracker Tracker), level Tracker) {

	level.Increment()
	defer level.Decrement()

	if n.Left != nil {
		n.Left.IOT(fnl, level)
	}

	fmt.Printf("Level: %d; Data: %d\n", level.GetLevel(), n.Data)
	fnl(n.Data, level)

	if n.Right != nil {
		n.Right.IOT(fnl, level)
	}

}

// PrintNode is an auxiliary function to supply to Traverse method
// and print to stdout the Data element of each node.
func PrintNode(value int) {
	fmt.Printf("Node: %d\n", value)
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

// ### Ready prototype
// 012345678900123456789001234567890
//                    ┌─[18]
//             ┌─[16]─┘
//      ┌─[14]─┤
//      │      │      ┌─[13]
//      │      └─[12]─┤
//      │             │      ┌─[11]
//      │             └─[10]─┤
//      │                    └─[ 9]
// [ 7]─┤
//      │
//      │
//      │
//      │
//      │      ┌─[ 6]
//      └─[ 5]─┤
//             │      ┌─[ 4]
//             └─[ 3]─┤
//                    └─[ 1]
//
//
// ### Ready prototype with blanks
// 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9
//
//                            ┌ ─ [18]
//
//                 ┌ ─ [16] ─ ┘
//
//      ┌ ─ [14] ─ ┤
//
//      │          │          ┌ ─ [13]
//
//      │          └ ─ [12] ─ ┤
//
//      │                     │          ┌ ─ [11]
//
//      │                     └ ─ [10] ─ ┤
//
//      │                                └ ─ [ 9]
//
// [ 7]─┤
//
//      │
//
//      │         ┌─ [ 6]
//
//      └ ─ [ 5]─ ┤
//
//                │         ┌ ─ [ 4]
//
//                └─ [ 3] ─ ┤
//
//                          └ ─ [ 1]

const (
	VERT = "│" // 2502 179

	VERT_LEFT  = "┤" // 2524 180
	HORIZONTAL = "─" // 2500 196
	UP_LEFT    = "┘" // 2518 217
	UP_RIGHT   = "└" // 2514 192
	DOWN_RIGHT = "┌" // 250C 218
	DOWN_LEFT  = "┐" // 2510 191
	SPACE      = " "
)

// Pretier holds metadata that will help
// format the printed tree
type Pretier struct {
	// Max tells the biggest number of the tree
	// This helps format the node block spaces
	Max int
}

func (b *BST) PrittyRoot() {

	// TODO: This hardcoded values need to be calculated
	pt := &Pretier{
		Max: 18,
	}

	b.Root.prittyNode(pt)
}

func (n *Node) prittyNode(p *Pretier) {

	if n.Right != nil {
		n.Right.prittyNode(p)
	}

	PrintData(n.Data, p.Max)
	fmt.Println()

	if n.Left != nil {
		n.Left.prittyNode(p)
	}

}

func PrintData(data, max int) {
	printData(os.Stdout, data, max)
}

func printData(writer io.Writer, data, max int) {

	sdata := []byte(strconv.Itoa(data))
	maxRunes := utf8.RuneCountInString(strconv.Itoa(max))

	if l := maxRunes - len(sdata); l > 0 {
		sdata = append(bytes.Repeat([]byte(SPACE), l), sdata...)
	}

	fmt.Fprintf(writer, "[%s]", string(sdata))

}
