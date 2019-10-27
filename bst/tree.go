package main

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

func (n *Node) PreOrderTraversal(fn func(value int)) {
	if n == nil {
		return
	}
	fn(n.Data)
	n.Left.PreOrderTraversal(fn)
	n.Right.PreOrderTraversal(fn)
}

func (n *Node) InOrderTraversal(fn func(value int)) {
	if n == nil {
		return
	}
	n.Left.InOrderTraversal(fn)
	fn(n.Data)
	n.Right.InOrderTraversal(fn)
}

func (n *Node) PostOrderTraversal(fn func(value int)) {
	if n == nil {
		return
	}
	n.Left.PostOrderTraversal(fn)
	n.Right.PostOrderTraversal(fn)
	fn(n.Data)
}

func PrintNode(value int) {
	fmt.Printf("Node: %d\n", value)
}

func (r *BST) NaiveInsert(data int) {
	if r.Root == nil {
		r.Root = &Node{Data: data}
		return
	}
	r.Root.NodeNaiveInsert(data)
}

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
