package bsarbol

import (
	"bytes"
	"reflect"
	"strings"
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

	t.Run("Breadth First Search Traversal with root-node struct", func(t *testing.T) {

		var got []int
		want := []int{7, 5, 10, 3, 6, 9, 12, 1, 4}

		bst.Traverse(func(v int) {
			got = append(got, v)
		},
			"bfs")

		if !reflect.DeepEqual(want, got) {
			t.Errorf("got %d; want %d", got, want)
		}

	})

}

func TestTraverseLevel(t *testing.T) {

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

	t.Run("Basic case", func(t *testing.T) {
		var valueGot []int
		valueWant := []int{1, 3, 4, 5, 6, 7, 9, 10, 12}
		var levelGot []int
		levelWant := []int{3, 2, 3, 1, 2, 0, 2, 1, 2}

		bst.TraverseLevel(func(v int, l *int) {
			valueGot = append(valueGot, v)
			levelGot = append(levelGot, *l)
		},
			"iot")

		if !reflect.DeepEqual(valueWant, valueGot) {
			t.Errorf("Value: got %d; want %d", valueGot, valueWant)
		}

		if !reflect.DeepEqual(levelWant, levelGot) {
			t.Errorf("Level: got %d; want %d", levelGot, levelWant)
		}
	})

}

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

func TestPrintData(t *testing.T) {

	cases := []struct {
		name string
		data int
		max  int
		want string
	}{
		{
			name: "Data: 2 runes, max: 2 runes",
			data: 12,
			max:  12,
			want: "[12]",
		},
		{
			name: "Data: 1 runes, max: 2 runes",
			data: 7,
			max:  12,
			want: "[ 7]",
		},
		{
			name: "Data: 1 runes, max: 3 runes",
			data: 3,
			max:  124,
			want: "[  3]",
		},
		{
			name: "Data: 1 runes, max: 6 runes",
			data: 3,
			max:  124497,
			want: "[     3]",
		},
		{
			name: "Data: 2 runes, max: 6 runes",
			data: 32,
			max:  124497,
			want: "[    32]",
		},
	}

	buffer := bytes.Buffer{}

	for _, test := range cases {
		t.Run(test.name, func(t *testing.T) {

			printDataB(&buffer, test.data, test.max)
			got := buffer.String()
			got = strings.TrimSuffix(got, "\n")
			defer buffer.Reset()

			if got != test.want {
				t.Errorf("got %s; want %s", got, test.want)
			}

		})
	}

}

func TestPrintDataDos(t *testing.T) {

	cases := []struct {
		name string
		data int
		max  int
		want string
	}{
		{
			name: "Data: 2 runes, max: 2 runes",
			data: 12,
			max:  12,
			want: "[12]",
		},
		{
			name: "Data: 1 runes, max: 2 runes",
			data: 7,
			max:  12,
			want: "[ 7]",
		},
		{
			name: "Data: 1 runes, max: 3 runes",
			data: 3,
			max:  124,
			want: "[  3]",
		},
		{
			name: "Data: 1 runes, max: 6 runes",
			data: 3,
			max:  124497,
			want: "[     3]",
		},
		{
			name: "Data: 2 runes, max: 6 runes",
			data: 32,
			max:  124497,
			want: "[    32]",
		},
	}

	builder := strings.Builder{}

	for _, test := range cases {
		t.Run(test.name, func(t *testing.T) {

			printData(&builder, test.data, test.max)
			got := builder.String()
			got = strings.TrimSuffix(got, "\n")
			defer builder.Reset()

			if got != test.want {
				t.Errorf("got %s; want %s", got, test.want)
			}

		})
	}

}

func BenchmarkPrintData(b *testing.B) {
	builder := strings.Builder{}

	data := 32
	max := 124497

	for i := 0; i < b.N; i++ {
		printData(&builder, data, max)
		builder.Reset()
	}

}

func BenchmarkPrintDataB(b *testing.B) {
	builder := strings.Builder{}

	data := 32
	max := 124497

	for i := 0; i < b.N; i++ {
		printDataB(&builder, data, max)
		builder.Reset()
	}

}
