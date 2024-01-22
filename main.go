package main

import (
	"fmt"
	"math/rand"
	"root/red_black_tree"
)

func main() {
	tree := red_black_tree.NewRBTree[int, int](300)

	for i := 0; i < 300; i++ {
		tree.Insert(rand.Intn(5000), 5)
	}

	fmt.Println(tree.InOrderTraversal())

	// liteStore := quick_store.New[int](20)

	// num := rand.Intn(20)

	// for index := 0; index < num; index++ {
	// 	liteStore.Append(index)
	// }

	// fmt.Println(num)
	// liteStore.Delete()
	// fmt.Println(liteStore.Len())
	// fmt.Println(liteStore.GetStore())

	// buff := smart_buffer.New(10)

	// buff.Buff([]byte("ja"))
	// fmt.Println(string(buff.GetStore()))

	// buff.Buff([]byte("jankoe"))
	// fmt.Println(string(buff.GetStore()))

	// buff.Buff([]byte("janko"))
	// fmt.Println(string(buff.GetStore()))
}
