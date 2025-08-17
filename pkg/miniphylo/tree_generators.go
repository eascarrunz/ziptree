package minphylo

func BalancedTree(taxset *TaxonSet) *Tree {
	tree := NewUnassembledTree(taxset)
	root := tree.Nodes[tree.anchorId]
	nbOuter := taxset.Len()
	tree.bifurcate(root, nbOuter, nbOuter, 0)

	for i := range nbOuter {
		tree.Nodes[i].Label = taxset.GetName(i)
	}

	return tree
}

/*
Recursively bifurcate nodes until a desired number of outer nodes is reached

Arguments:
- node:	Node to bifurcate
- nbOuter:	Target number of outer nodes descendant of this node
- nextIdInner:	ID for the next inner node
- nextIdOuter:	ID for the next outer node
*/
func (tree *Tree) bifurcate(node *Node, nbOuter int, nextIdInner int, nextIdOuter int) (int, int) {
	nbOuterRight := nbOuter / 2
	nbOuterLeft := nbOuter - nbOuterRight

	var id int

	// Left side
	if nbOuterLeft == 1 {
		id = nextIdOuter
		nextIdOuter += 1
	} else {
		id = nextIdInner
		nextIdInner += 1
	}
	leftChildNode := tree.Nodes[id]
	node.addLink(leftChildNode, tree.NewBranch())

	if nbOuterLeft > 1 {
		nextIdInner, nextIdOuter = tree.bifurcate(leftChildNode, nbOuterLeft, nextIdInner, nextIdOuter)
	}

	// Right side
	if nbOuterRight == 1 {
		id = nextIdOuter
		nextIdOuter += 1
	} else {
		id = nextIdInner
		nextIdInner += 1
	}
	rightChildNode := tree.Nodes[id]
	node.addLink(rightChildNode, tree.NewBranch())

	if nbOuterRight > 1 {
		nextIdInner, nextIdOuter = tree.bifurcate(rightChildNode, nbOuterRight, nextIdInner, nextIdOuter)
	}

	return nextIdInner, nextIdOuter
}
