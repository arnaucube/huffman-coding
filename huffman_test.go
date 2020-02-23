package huffman

import (
	"bytes"
	"fmt"
	"testing"

	"gopkg.in/go-playground/assert.v1"
)

var debug = true

func TestBuildTree0(t *testing.T) {
	values := make(map[byte]int64)
	values[65] = 3
	values[66] = 5
	values[67] = 6
	values[68] = 4
	values[69] = 2

	leafs := leafsFromValues(values)
	sortedLeafs := sortNodes(leafs)
	n := buildTree(sortedLeafs)

	/*
		Expected result (both valid):
			         20
				/  \
			       9    \
			      / \    \
			     /   5    11
		            /  	/ \   / \
			   4   2   3  5  6

			         20
				/  \
			       /    11
			      /     / \
			     9     5   \
		            / \   / \   \
			   4   5 2  3    6
		The second tree is what is genereated in this test
	*/
	assert.Equal(t, int64(20), n.key)
	assert.Equal(t, int64(9), n.l.key)
	assert.Equal(t, int64(4), n.l.l.key)
	assert.Equal(t, int64(5), n.l.r.key)
	assert.Equal(t, int64(11), n.r.key)
	assert.Equal(t, int64(5), n.r.l.key)
	assert.Equal(t, int64(2), n.r.l.l.key)
	assert.Equal(t, int64(3), n.r.l.r.key)
	assert.Equal(t, int64(6), n.r.r.key)

	w := bytes.NewBufferString("")
	printTree(w, n)
	if debug {
		fmt.Println(w)
	}
}

func TestBuildTree1(t *testing.T) {
	values := make(map[byte]int64)
	values[97] = 10
	values[101] = 15
	values[105] = 12
	values[115] = 3
	values[116] = 4
	values[112] = 13
	values[10] = 1

	leafs := leafsFromValues(values)
	sortedLeafs := sortNodes(leafs)
	n := buildTree(sortedLeafs)

	assert.Equal(t, int64(58), n.key)
	assert.Equal(t, int64(25), n.l.key)
	assert.Equal(t, int64(33), n.r.key)

	w := bytes.NewBufferString("")
	printTree(w, n)
	if debug {
		fmt.Println(w)
	}
}

func TestGenerateTable(t *testing.T) {
	values := make(map[byte]int64)
	values[65] = 3
	values[66] = 5
	values[67] = 6
	values[68] = 4
	values[69] = 2

	leafs := leafsFromValues(values)
	sortedLeafs := sortNodes(leafs)
	n := buildTree(sortedLeafs)

	table := generateTable(n)

	fmt.Println(table)
}
