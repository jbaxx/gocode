package bsarbol

import (
	"reflect"
	"testing"
)

func TestTraverse(t *testing.T) {

	bst := &BST{
		Root: &Node{
			Data: 7,
			Left: &Node{
				Data: 5,
				Left: &Node{
					Data:  3,
					Left:  &Node{Data: 1},
					Right: &Node{Data: 4},
				},
				Right: &Node{Data: 6},
			},
			Right: &Node{
				Data:  10,
				Left:  &Node{Data: 9},
				Right: &Node{Data: 12},
			},
		},
	}

	t.Run("In Order Traversal with root-node struct", func(t *testing.T) {

		var got []int
		want := []int{1, 3, 4, 5, 6, 7, 9, 10, 12}

		bst.Traverse(func(v int) {
			got = append(got, v)
		},
			"inorder")

		if !reflect.DeepEqual(want, got) {
			t.Errorf("got %d; want %d", got, want)
		}

	})

	t.Run("In Order Traversal with empty BST", func(t *testing.T) {

		emptyBst := &BST{}
		var got []int
		var want []int

		emptyBst.Traverse(func(v int) {
			got = append(got, v)
		},
			"inorder")

		if !reflect.DeepEqual(want, got) {
			t.Errorf("got %d; want %d", got, want)
		}

	})

	t.Run("Pre Order Traversal with root-node struct", func(t *testing.T) {

		var got []int
		want := []int{7, 5, 3, 1, 4, 6, 10, 9, 12}

		bst.Traverse(func(v int) {
			got = append(got, v)
		},
			"preorder")

		if !reflect.DeepEqual(want, got) {
			t.Errorf("got %d; want %d", got, want)
		}

	})

	t.Run("Post Order Traversal with root-node struct", func(t *testing.T) {

		var got []int
		want := []int{1, 4, 3, 6, 5, 9, 12, 10, 7}

		bst.Traverse(func(v int) {
			got = append(got, v)
		},
			"postorder")

		if !reflect.DeepEqual(want, got) {
			t.Errorf("got %d; want %d", got, want)
		}

	})

}

func TestNaiveInsert(t *testing.T) {

	t.Run("NaiveInsert function", func(t *testing.T) {
		var got []int
		want := []int{1, 3, 4, 5, 6, 7, 9, 10, 12}

		bst := &BST{}

		bst.NaiveInsert(1)
		bst.NaiveInsert(3)
		bst.NaiveInsert(4)
		bst.NaiveInsert(5)
		bst.NaiveInsert(6)
		bst.NaiveInsert(7)
		bst.NaiveInsert(9)
		bst.NaiveInsert(10)
		bst.NaiveInsert(12)

		bst.Traverse(func(v int) {
			got = append(got, v)
		},
			"inorder")

		if !reflect.DeepEqual(want, got) {
			t.Errorf("got %d; want %d", got, want)
		}
	})

	t.Run("NaiveInsert with duplicates", func(t *testing.T) {
		var got []int
		want := []int{1, 3, 4, 5, 6, 7, 9, 10, 12}

		bst := &BST{}

		bst.NaiveInsert(1)
		bst.NaiveInsert(3)
		bst.NaiveInsert(4)
		bst.NaiveInsert(4)
		bst.NaiveInsert(4)
		bst.NaiveInsert(5)
		bst.NaiveInsert(6)
		bst.NaiveInsert(7)
		bst.NaiveInsert(7)
		bst.NaiveInsert(9)
		bst.NaiveInsert(10)
		bst.NaiveInsert(12)

		bst.Traverse(func(v int) {
			got = append(got, v)
		},
			"inorder")

		if !reflect.DeepEqual(want, got) {
			t.Errorf("got %d; want %d", got, want)
		}
	})

	t.Run("NaiveInsert in reverse order", func(t *testing.T) {
		var got []int
		want := []int{1, 3, 4, 5, 6, 7, 9, 10, 12}

		bst := &BST{}

		bst.NaiveInsert(12)
		bst.NaiveInsert(10)
		bst.NaiveInsert(9)
		bst.NaiveInsert(7)
		bst.NaiveInsert(6)
		bst.NaiveInsert(5)
		bst.NaiveInsert(4)
		bst.NaiveInsert(3)
		bst.NaiveInsert(1)

		bst.Traverse(func(v int) {
			got = append(got, v)
		},
			"inorder")

		if !reflect.DeepEqual(want, got) {
			t.Errorf("got %d; want %d", got, want)
		}
	})

}
