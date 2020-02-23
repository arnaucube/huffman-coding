package huffman

import (
	"fmt"
	"sort"
)

const (
	LEAF = 0
	MID  = 1
)

type Node struct {
	typ  int64
	key  int64
	leaf byte
	l    *Node
	r    *Node
}

func (n *Node) String() string {
	if n.typ == LEAF {
		return fmt.Sprintf("%s (%v)", string(n.leaf), n.key)
	}
	return fmt.Sprintf("%v", n.key)
}

func newLeaf(b byte, k int64) *Node {
	return &Node{
		typ:  LEAF,
		key:  k,
		leaf: b,
	}
}

func newNode(a, b *Node) *Node {
	var l, r *Node
	if a.key < b.key {
		l = a
		r = b
	} else {
		l = b
		r = a
	}
	n := &Node{
		typ: MID,
		key: a.key + b.key,
		l:   l,
		r:   r,
	}
	return n
}

func leafsFromValues(values map[byte]int64) []*Node {
	var nodes []*Node
	for k, v := range values {
		n := newLeaf(k, v)
		nodes = append(nodes, n)
	}
	return nodes
}

func rmFromNodes(i int, nodes []*Node) []*Node {
	if i == len(nodes) {
		return nodes[:i]
	}
	return append(nodes[:i], nodes[i+1:]...)
}

// sortNodes returns a []*Node sorted
func sortNodes(nodes []*Node) []*Node {
	sort.Slice(nodes, func(i, j int) bool {
		return nodes[i].key < nodes[j].key
	})
	return nodes
}

func buildTree(nodes []*Node) *Node {
	// build binary tree
	for len(nodes) > 1 {
		// pair two smallest nodes
		nodes = sortNodes(nodes)
		// create a new node with them
		n := newNode(nodes[0], nodes[1])
		// remove the picked nodes from the array of nodes
		nodes = rmFromNodes(0, nodes)
		nodes = rmFromNodes(0, nodes)
		// add the new node (the two old paired) to the node array
		nodes = append(nodes, n)
	}
	return nodes[0]
}

func descendTable(n *Node, path string, table map[byte]string) map[byte]string {
	if n.typ == MID {
		table = descendTable(n.l, path+"0", table)
		table = descendTable(n.r, path+"1", table)
	} else if n.typ == LEAF {
		table[n.leaf] = path
	}
	return table
}

func generateTable(n *Node) map[byte]string {
	table := make(map[byte]string)
	table = descendTable(n, "", table)
	return table
}

func Huffman(b []byte) error {
	// compute frequencies
	values := make(map[byte]int64)
	for _, v := range b {
		values[v]++
	}
	leafs := leafsFromValues(values)

	// sort frequencies
	sortedLeafs := sortNodes(leafs)

	// build binary tree
	n := buildTree(sortedLeafs)
	fmt.Println(n)

	// get the table (binary code for each value)
	table := generateTable(n)
	fmt.Println(table)

	// WIP

	return nil
}
