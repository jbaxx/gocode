package bsarbol

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"unicode/utf8"
)

// ### Tree Print Ready prototype
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
//                      ┌─[18]
//               ┌─[16]─┘
//        ┌─[14]─┤
//                      ┌─[13]
//               └─[12]─┤
//                             ┌─[11]
//                      └─[10]─┤
//                             └─[ 9]
// [ 7]─┤
//               ┌─[ 6]
//        └─[ 5]─┤
//                      ┌─[ 4]
//               └─[ 3]─┤
//                      └─[ 1]
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
//      │         ┌ ─ [ 6]
//
//      └ ─ [ 5]─ ┤
//
//                │          ┌ ─ [ 4]
//
//                └ ─ [ 3] ─ ┤
//
//                           └ ─ [ 1]

// https://unicode.org/cldr/utility/character.jsp?a=2500&B1=Show
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

// Tracker interface defines the methods needed to keep track of node depth
type Tracker interface {
	SetDepth(i int)
	GetDepth() int
	Increment()
	Decrement()
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

func (n *nodeLevel) SetDepth(i int) {
	n.level = i
}

func (n *nodeLevel) GetDepth() int {
	return n.level
}

func NewNodeLevel() *nodeLevel {
	return &nodeLevel{level: 0}
}

// metaFormat holds metadata that will help format the printed tree
// metaFormat may implement the Tracker interface
type metaFormat struct {
	// Max tells the biggest number of the tree
	// This helps format the node block spaces
	Max int
	// depth keeps track of the current visited node depth
	depth    int
	level    map[int]int
	metadata map[int]*nodeMetadata
}

type nodeMetadata struct {
	depth int
	child string
	leaf  bool
}

func (m *metaFormat) Increment() {
	m.depth++
}

func (m *metaFormat) Decrement() {
	m.depth--
}

func (m *metaFormat) SetDepth(i int) {
	m.depth = i
}

func (m *metaFormat) GetDepth() int {
	return m.depth
}

func (m *metaFormat) MapNodeDepth(key, v int) {
	// old implementation
	if _, ok := m.level[key]; !ok {
		m.level[key] = v
	}
	// new implementation
	if _, ok := m.metadata[key]; !ok {
		m.metadata[key] = &nodeMetadata{}
		m.metadata[key].depth = v
	} else {
		m.metadata[key].depth = v
	}

}

func (m *metaFormat) SetChildSide(key int, s string) {
	if _, ok := m.metadata[key]; !ok {
		m.metadata[key] = &nodeMetadata{}
		m.metadata[key].child = s
	} else {
		m.metadata[key].child = s
	}
}

func NewMetaFormat() *metaFormat {
	return &metaFormat{
		Max:      20,
		depth:    0,
		level:    make(map[int]int),
		metadata: make(map[int]*nodeMetadata),
	}
}

func (b *BST) PrittyRoot() {

	mf := NewMetaFormat()
	mf.SetDepth(-1)

	b.Root.scanNode(mf)
	b.Root.prittyNode(mf)

	fmt.Println("Print Format: ")
	// for k, v := range mf.metadata {
	// 	fmt.Printf("Key: %d\tChildSide: %s\tDepth: %d\n", k, (*v).child, (*v).depth)
	// 	fmt.Printf("Key: %d\tChildSide: %q\tDepth: %v\n", k, v.child, v.depth)
	// 	fmt.Printf("Key: %d\tmetadata: %#v\n", k, v)
	// }
}

func (n *Node) Peek() int {
	return n.Data
}

func (n *Node) scanNode(m *metaFormat) {
	m.Increment()
	defer m.Decrement()

	if n.Right != nil {
		m.SetChildSide(n.Right.Peek(), "right")
		n.Right.scanNode(m)
	}

	m.MapNodeDepth(n.Data, m.GetDepth())

	if n.Left != nil {
		m.SetChildSide(n.Left.Peek(), "left")
		n.Left.scanNode(m)
	}

}

func (n *Node) prittyNode(m *metaFormat) {

	if n.Right != nil {
		n.Right.prittyNode(m)
	}

	// PRINT AREA
	prof := m.metadata[n.Data].depth * 7
	if prof < 0 {
		log.Fatalln("depth count is lower than 0")
	}
	if prof == 0 {
		prof = 2
	}
	spacetimes := bytes.Repeat([]byte(SPACE), prof)
	fmt.Printf("%s", string(spacetimes))
	if m.metadata[n.Data].child == "right" {
		fmt.Printf("%s", "┌─")
	}
	if m.metadata[n.Data].child == "left" {
		fmt.Printf("%s", "└─")
	}
	if err := PrintData(n.Data, m.Max); err != nil {
		log.Fatalln("func PrintData failed:", err)
	}
	if n.Left != nil && n.Right != nil {
		fmt.Printf("%s", "─┤")
	}
	if n.Left == nil && n.Right != nil {
		fmt.Printf("%s", "─┘")
	}
	if n.Left != nil && n.Right == nil {
		fmt.Printf("%s", "─┐")
	}
	// fmt.Println()
	fmt.Println()
	// PRINT AREA

	if n.Left != nil {
		n.Left.prittyNode(m)
	}

}

type ErrMaxRuneCount struct {
	nodeData int
}

func (e *ErrMaxRuneCount) Error() string {
	return fmt.Sprintf("node [%d]: max rune count smaller than data rune count", e.nodeData)
}

func PrintData(data, max int) error {
	return printData(os.Stdout, data, max)
}

func printData(writer io.Writer, data, max int) error {

	sdata := []byte(strconv.Itoa(data))
	maxRunes := utf8.RuneCountInString(strconv.Itoa(max))

	l := maxRunes - len(sdata)
	if l < 0 {
		return &ErrMaxRuneCount{nodeData: data}
	}

	if l > 0 {
		sdata = append(bytes.Repeat([]byte(SPACE), l), sdata...)
	}

	fmt.Fprintf(writer, "[%s]", string(sdata))
	return nil

}
