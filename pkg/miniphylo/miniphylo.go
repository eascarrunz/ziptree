package minphylo

type Traversal int

const (
	PreOrder = iota
	PostOrder
)

func (node *Node) preTraverse(f func(*Node), parentNode *Node) {
	f(node)

	for _, link := range node.links {
		if link.node == node {
			continue
		}

		link.node.preTraverse(f, node)
	}
}

func (node *Node) postTraverse(f func(*Node), parentNode *Node) {
	for _, link := range node.links {
		if link.node == node {
			continue
		}

		link.node.preTraverse(f, node)
	}

	f(node)
}

func (node *Node) Traverse(f func(*Node), traversal Traversal) {
	switch traversal {
	case PreOrder:
		node.preTraverse(f, nil)
	case PostOrder:
		node.postTraverse(f, nil)
	}
}

// Return the capacity of nodes of a tree
func (tree *Tree) CapNode() int {
	return len(tree.Nodes)
}

func (node *Node) countNodes() int {
	nbNodes := 0

	visit := func(node *Node) {
		nbNodes += 1
	}

	node.preTraverse(visit, nil)

	return nbNodes
}
