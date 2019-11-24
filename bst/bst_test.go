package bsarbol

import (
	"reflect"
	"testing"
)

func TestNaiveInsert(t *testing.T) {

	t.Run("NaiveInsert function", func(t *testing.T) {
		var got []int
		want := []int{1, 3, 4, 5, 6, 7, 9, 10, 12}

		bst := &BST{}

		for _, elem := range want {
			bst.NaiveInsert(elem)
		}

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
		dups := []int{1, 3, 4, 4, 4, 5, 6, 7, 7, 9, 10, 12}
		want := []int{1, 3, 4, 5, 6, 7, 9, 10, 12}

		bst := &BST{}

		for _, elem := range dups {
			bst.NaiveInsert(elem)
		}

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
		wantReversed := []int{12, 10, 9, 7, 6, 5, 4, 3, 1}

		bst := &BST{}

		for _, elem := range wantReversed {
			bst.NaiveInsert(elem)
		}

		bst.Traverse(func(v int) {
			got = append(got, v)
		},
			"inorder")

		if !reflect.DeepEqual(want, got) {
			t.Errorf("got %d; want %d", got, want)
		}
	})

}
