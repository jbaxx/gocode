// Package narytree defines a simple n-ary Tree data structure
package narytree

import (
	"fmt"
	"strconv"
	"strings"
)

// NaryTree defines an n-ary Tree data structure
type NaryTree struct {
	Root *Node
}

// Node defines a node element of an NaryTree
type Node struct {
	Key   int
	Data  string
	Nodes []*Node
}

// NewNaryTree returns a NaryTree with a fixed Root node
func NewNaryTree() *NaryTree {
	return &NaryTree{
		Root: &Node{
			Key:   0,
			Data:  "i am root",
			Nodes: []*Node{},
		},
	}
}

// Insert implements a method to insert elements into the nary tree.
// It expects arrays that represent the path from root to leaf only.
// It does not support insertion based on a keyed element inside the tree
func (h *NaryTree) Insert(trace []int) {
	h.Root.insert(trace)
}

func (n *Node) insert(trace []int) {
	// does the current node exist?
	//     if not, append
	//     if yes, insert
	if len(trace) == 0 { // makes below pop safe
		return
	}

	current, trace := trace[0], trace[1:]

	index, ok := n.find(current)
	if !ok {
		n.Nodes = append(n.Nodes, &Node{Key: current})
		n.Nodes[len(n.Nodes)-1].insert(trace)
		return
	}

	n.Nodes[index].insert(trace)

}

func (n *Node) find(e int) (int, bool) {
	for i, w := range n.Nodes {
		if w.Key == e {
			return i, true
		}
	}
	return 0, false
}

// PrintTree:
// 1
// 3
// 5
// 8
// 6
// 17
// 18
// 19
// 25
// 26
// 306
// 317
// 318
// 319
// 325
// 329
// 339
// 349
// 29
// 39
// 49
// 73
// 85
// 88
// 91
// 94
// 7
//
//
// PrintPaths:
// [1 3 5]
// [1 3 8]
// [1 6 17]
// [1 6 18]
// [1 6 19 25 26 306 317]
// [1 6 19 25 26 306 318]
// [1 6 19 25 26 306 319 325]
// [1 6 19 25 26 306 319 329]
// [1 6 19 25 26 306 339]
// [1 6 19 25 26 306 349]
// [1 6 19 29]
// [1 6 39]
// [1 6 49]
// [1 73 85]
// [1 73 88 91 94]
// [1 7]
//
// adjacency list:
// [1 3]
// [1 6]
// [1 73]
// [1 7]
// [3 5]
// [3 8]
// [6 17]
// [6 18]
// [6 19]
// [6 39]
// [6 49]
// [73 85]
// [73 88]
// [19 25]
// [19 29]
// [25 26]
// [26 306]
// [306 317]
// [306 318]
// [306 319]
// [306 339]
// [306 349]
// [319 325]
// [319 329]
//
// root: 1
// subRoots:
// [3]
// [6]
// [73]
// [7]
// maps:
// [3]: [5 8]
// [6]: [17 18 19 39 49]
// [73]: [85 88]
// [19]: [25 29]
// [25]: [26]
// [26]: [306]
// [306]: [317 318 319 339 349]
// [319]: [325 329]

// String implements the stringer interface for the Node
// printing the Key and Data together
func (n *Node) String() string {
	return "[" + strconv.Itoa(n.Key) + "] " + n.Data
}

// PrintBlock prints the NaryTree in a format similar to the unix command: tree
func (h *NaryTree) PrintBlock() {
	lost := []bool{}
	if h.Root == nil {
		fmt.Printf("empty tree\n")
		return
	}
	h.Root.printBlock(lost)
}

func (t *Node) printBlock(sl stackLast) {

	fmt.Printf("%s%s\n", sl, t)

	// get len of node Nodes array, minus 1 to make len 0-based
	slong := len(t.Nodes) - 1
	for i, v := range t.Nodes {
		// compare len to index to know if this is the last elem
		isLast := i == slong
		// append (push) to the "is last" stack
		sl = append(sl, isLast)

		v.printBlock(sl)
		// when the recursive call returns, pop the last elemt
		sl = sl[:len(sl)-1]
	}

}

// stackLast is a stack to track for each ancestor
// if it's the last member of it's own slice, this way we know
// if it has to be represented with an elbow character as well as
// which one to not print when nested in a tree.
type stackLast []bool

// String implements the stringer interface, to pass the stackLast to the io Writer
// references: https://en.wikipedia.org/wiki/Box-drawing_character
// references: https://en.wikipedia.org/wiki/Code_page_437
// This is a section of a printed tree along with the stackLast trace
// for each node and its ancestors.
// If the node itself is last in its own []*Node slice we preceed it with an elbow: └───
// If the node isn't the last in its own []*Node slice we preceed it with a continue elbow: ├───
// For the stack trace, if an ancestor is not the last, we print its continuation trace: │
// Otherwise we print a blank space, below represented with: *
//
// ├─── 19                           [19]false
// │    ├─── 25                      [19]false [25]false
// │    │    └─── 26                 [19]false [25]false [26]true
// │    │    *    └─── 36            [19]false [25]false [26]true [36]true
// │    │    *    *    ├─── 37       [19]false [25]false [26]true [36]true [37]false
// │    │    *    *    ├─── 38       [19]false [25]false [26]true [36]true [38]false
// │    │    *    *    ├─── 39       [19]false [25]false [26]true [36]true [39]false
// │    │    *    *    │    ├─── 52  [19]false [25]false [26]true [36]true [39]false [52]false
func (s stackLast) String() string {
	if len(s) == 0 {
		return ""
	}
	ng := []string{}
	last, rest := s[len(s)-1], s[:len(s)-1]
	for _, r := range rest {
		if r {
			ng = append(ng, "     ")
		} else {
			ng = append(ng, "│    ")
		}
	}
	if last {
		ng = append(ng, "└─── ")
	} else {
		ng = append(ng, "├─── ")
	}
	ngj := strings.Join(ng, "")
	return ngj
}

// PrintTree prints a pre order traversal of the NaryTree
func (h *NaryTree) PrintTree(fn func(int)) {
	h.Root.printTurn(fn)
}

func (t *Node) printTurn(fn func(int)) {

	fn(t.Key)

	for _, v := range t.Nodes {
		v.printTurn(fn)
	}
}

// PrintPaths prints all the paths from root to each leaf node in the NaryTree
func (h *NaryTree) PrintPaths() {
	var q []int
	h.Root.printPaths(q)
}

func (t *Node) printPaths(q []int) {
	q = append(q, t.Key)
	if len(t.Nodes) == 0 {
		fmt.Println(q)
	}

	for _, v := range t.Nodes {
		v.printPaths(q)
	}
	q = q[:len(q)-1]
}
