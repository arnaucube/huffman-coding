package huffman

import (
	"fmt"
	"io"
)

func printNode(w io.Writer, n *Node) {
	if n.typ == MID {
		fmt.Fprintf(w, "\"%v\" -> {\"%v\" \"%v\"}\n", n.String(), n.l.String(), n.r.String())
		printNode(w, n.l)
		printNode(w, n.r)
	} else if n.typ == LEAF {
		fmt.Fprintf(w, "\"%v\" [style=filled];\n", n.String())
	}
}
func printTree(w io.Writer, n *Node) {
	fmt.Fprintf(w, `digraph hierarchy {
		node [fontname=Monospace,fontsize=10,shape=box]
		`)

	printNode(w, n)
	fmt.Fprintf(w, "}\n")
}
