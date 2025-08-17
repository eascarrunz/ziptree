package minphylo

import (
	"math"
)

type Node struct {
	Id    int
	Taxon int
	Label string
	links []*Link
}

type Branch struct {
	Id     int
	Left   *Node
	Right  *Node
	Length float64
}

type Link struct {
	node   *Node
	branch *Branch
}

type Tree struct {
	anchorId int
	Nodes    []*Node
	Branches []*Branch
}

func (node *Node) Degree() int {
	return len(node.links)
}

func (node *Node) InDegree() int {
	d := 0
	for _, link := range node.links {
		d += bool2int(link.branch.Right == node)
	}

	return d
}

func (node *Node) OutDegree() int {
	d := 0
	for _, link := range node.links {
		d += bool2int(link.branch.Left == node)
	}

	return d
}

func NewUnassembledTree(taxa *TaxonSet) *Tree {
	nbTaxa := taxa.Len()
	nbNodes := 2*nbTaxa - 1
	nbBranches := nbTaxa - 2
	tree := Tree{
		anchorId: nbTaxa + 1,
		Nodes:    make([]*Node, nbNodes),
		Branches: make([]*Branch, 0, nbBranches),
	}

	for i := range nbTaxa {
		// Pre-allocate 3 links assuming by default that the tree will be binary
		tree.Nodes[i] = &Node{i, i, "", make([]*Link, 0, 3)}
	}

	for i := nbTaxa; i < nbNodes; i += 1 {
		// Pre-allocate 3 links assuming by default that the tree will be binary
		tree.Nodes[i] = &Node{i, -1, "", make([]*Link, 0, 3)}
	}

	// for i := range nbBranches {
	// 	tree.branches[i] = &Branch{i, nil, nil, math.NaN()}
	// }

	return &tree
}

func (tree *Tree) NewNode() *Node {
	i := len(tree.Nodes)
	tree.Nodes = append(tree.Nodes, &Node{i, i, "", make([]*Link, 0, 3)})

	return tree.Nodes[i]
}

func (tree *Tree) NewBranch() *Branch {
	i := len(tree.Branches)
	tree.Branches = append(tree.Branches, &Branch{i, nil, nil, math.NaN()})

	return tree.Branches[i]
}

func (leftNode *Node) addLink(rightNode *Node, branch *Branch) {
	leftNode.links = append(leftNode.links, &Link{rightNode, branch})
	rightNode.links = append(rightNode.links, &Link{leftNode, branch})
	branch.Left = leftNode
	branch.Right = rightNode
}
