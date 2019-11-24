package main

import (
	"fmt"
	"github.com/jbaxx/gocode/bst"
)

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

func main() {

	bst := &bsarbol.BST{
		Root: &bsarbol.Node{
			Data: 7,
			Right: &bsarbol.Node{
				Data: 14,
				Right: &bsarbol.Node{
					Data:  16,
					Right: &bsarbol.Node{Data: 18},
				},
				Left: &bsarbol.Node{
					Data:  12,
					Right: &bsarbol.Node{Data: 13},
					Left: &bsarbol.Node{
						Data:  10,
						Right: &bsarbol.Node{Data: 11},
						Left:  &bsarbol.Node{Data: 9},
					},
				},
			},
			Left: &bsarbol.Node{
				Data:  5,
				Right: &bsarbol.Node{Data: 6},
				Left: &bsarbol.Node{
					Data:  3,
					Right: &bsarbol.Node{Data: 4},
					Left:  &bsarbol.Node{Data: 1},
				},
			},
		},
	}

	bst2 := &bsarbol.BST{
		Root: &bsarbol.Node{
			Data: 7,
			Right: &bsarbol.Node{
				Data: 14,
				Right: &bsarbol.Node{
					Data: 16,
					Right: &bsarbol.Node{
						Data:  18,
						Right: &bsarbol.Node{Data: 23},
					},
					Left: &bsarbol.Node{Data: 17},
				},
				Left: &bsarbol.Node{
					Data:  12,
					Right: &bsarbol.Node{Data: 13},
					Left: &bsarbol.Node{
						Data:  10,
						Right: &bsarbol.Node{Data: 11},
						Left:  &bsarbol.Node{Data: 9},
					},
				},
			},
			Left: &bsarbol.Node{
				Data:  5,
				Right: &bsarbol.Node{Data: 6},
				Left: &bsarbol.Node{
					Data:  3,
					Right: &bsarbol.Node{Data: 4},
					Left:  &bsarbol.Node{Data: 1},
				},
			},
		},
	}

	fmt.Println("In Order Traversal")
	bst.Traverse(func(p int) {
		fmt.Println(p)
	}, "inorder")

	fmt.Println()
	fmt.Println("Pre Order Traversal")
	bst.Traverse(func(p int) {
		fmt.Println(p)
	}, "preorder")

	fmt.Println()
	fmt.Println("Pritty Traversal")
	bst.PrittyRoot()

	fmt.Println()
	fmt.Println()
	fmt.Println("Pritty Traversal")
	bst2.PrittyRoot()
}
