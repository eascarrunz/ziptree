package main

import (
	"fmt"
)

func main() {
	taxonNames := []string{"A", "B", "C", "D", "E", "F", "G", "H"}
	taxa, _ := NewTaxonSet(taxonNames)
	tree := BalancedTree(taxa)

	fmt.Println(tree.Newick())
}
