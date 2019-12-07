package narytree

import (
	"fmt"
	"reflect"
	"testing"
)

func TestPrintTree(t *testing.T) {

	h := &NaryTree{
		Root: &Node{
			Key:  1,
			Data: "uno",
			Nodes: []*Node{
				&Node{Key: 3, Data: "tres", Nodes: []*Node{
					&Node{Key: 5, Data: "cinco"},
					&Node{Key: 8, Data: "ocho"}},
				},
				&Node{Key: 6, Data: "seis", Nodes: []*Node{
					&Node{Key: 17, Data: "diecisiete"},
					&Node{Key: 18, Data: "dieciocho"},
					&Node{Key: 19, Data: "diecinueve", Nodes: []*Node{
						&Node{Key: 25, Data: "veinticinco", Nodes: []*Node{&Node{Key: 26, Data: "veintiseis", Nodes: []*Node{
							&Node{Key: 306, Data: "trescientos seis", Nodes: []*Node{
								&Node{Key: 317, Data: "trescientos diecisiete"},
								&Node{Key: 318, Data: "trescientos dieciocho"},
								&Node{Key: 319, Data: "trescientos diecinueve", Nodes: []*Node{
									&Node{Key: 325, Data: "trescientos veinticinco"},
									&Node{Key: 329, Data: "trescientos veintinueve"},
								}},
								&Node{Key: 339, Data: "trescientos treinta y nueve"},
								&Node{Key: 349, Data: "trescientos cuarenta y nueve"},
							},
							}}}}},
						&Node{Key: 29, Data: "veintinueve"},
					}},
					&Node{Key: 39, Data: "treinta y nueve"},
					&Node{Key: 49, Data: "cuarenta y nueve"},
				},
				},
				&Node{Key: 73, Data: "setenta y tres", Nodes: []*Node{
					&Node{Key: 85, Data: "ochenta y cinco"},
					&Node{Key: 88, Data: "ochenta y ocho", Nodes: []*Node{
						&Node{Key: 91, Data: "noventa y uno", Nodes: []*Node{&Node{Key: 94, Data: "noventa y cuatro"}}},
					}}},
				},
				&Node{Key: 7, Data: "siete"},
			},
		},
	}

	fmt.Println("PrintTree: ")
	h.PrintTree(func(v int) { fmt.Println(v) })
	fmt.Println()
	fmt.Println()
	fmt.Println("PrintPaths: ")
	h.PrintPaths()
	fmt.Println()
	fmt.Println()
	fmt.Println("PrintBlock: ")
	h.PrintBlock()
	fmt.Println()

	want := []int{1, 3, 5, 8, 6, 17, 18, 19, 25, 26, 306, 317, 318, 319, 325, 329, 339, 349, 29, 39, 49, 73, 85, 88, 91, 94, 7}
	// want := []int{1, 3, 5, 8, 6, 17, 18, 19, 39, 49, 7}
	got := []int{}
	h.PrintTree(func(v int) {
		got = append(got, v)
	})

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v; want %v", got, want)
	}

}

func TestInsert(t *testing.T) {
	h := &NaryTree{
		Root: &Node{
			Key:  1,
			Data: "uno",
			Nodes: []*Node{
				&Node{Key: 3, Data: "tres", Nodes: []*Node{
					&Node{Key: 5, Data: "cinco"},
					&Node{Key: 8, Data: "ocho"}},
				},
			},
		},
	}
	h.PrintBlock()

	q := NewNaryTree()
	q.Insert([]int{1, 3, 5})
	q.Insert([]int{1, 3, 8})
	q.Insert([]int{9, 3, 5})
	q.PrintBlock()
}
