package miniphylo

import (
	"fmt"
	"math"
	"strings"
)

func (link *Link) _newick(b *strings.Builder) {
	if len(link.node.links) > 0 {
		b.WriteString("(")
		isFirst := true
		for _, childLink := range link.node.links {
			if childLink.node == link.node {
				continue
			}
			childLink._newick(b)
			if !isFirst {
				b.WriteString(",")
			}
			isFirst = false
		}
		b.WriteString(")")
	}
	b.WriteString(link.node.Label)

	if !math.IsNaN(link.branch.Length) {
		b.WriteString(":")
		fmt.Fprint(b, link.branch.Length)
	}
}

func (node *Node) _newick(b *strings.Builder) {
	if len(node.links) > 0 {
		b.WriteString("(")
		isFirst := true
		for _, childLink := range node.links {
			childLink._newick(b)
			if !isFirst {
				b.WriteString(",")
			}
			isFirst = false
		}
		b.WriteString(")")
	}
	b.WriteString(node.Label)
}

// Return the Newick representation of the tree rooted on the node
func (node *Node) Newick() string {
	var b strings.Builder
	node._newick(&b)
	b.WriteString(";")

	return b.String()
}

// Return the Newick representation of the tree
func (tree *Tree) Newick() string {
	return tree.Nodes[tree.anchorId].Newick()
}
