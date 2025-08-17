package main

import (
	"fmt"

	"github.com/eascarrunz/ziptree/pkg/miniphylo"
)

func main() {
	taxonNames := []string{"A", "B", "C", "D", "E", "F", "G", "H"}
	taxa, _ := miniphylo.NewTaxonSet(taxonNames)
	tree := miniphylo.BalancedTree(taxa)

	fmt.Println(tree.Newick())
}
