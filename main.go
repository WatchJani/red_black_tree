package main

import (
	"fmt"
	"math/rand"
	"root/red_black_tree"
)

func main() {
	tree := red_black_tree.NewRBTree(300)

	for i := 0; i < 300; i++ {
		tree.Insert(rand.Intn(5000))
	}

	fmt.Println(tree.InOrderTraversal())
}
