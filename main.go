package main

import (
	"fmt"
	"root/quick_store"
)

func main() {
	store := quick_store.New[int](5)

	for i := 0; i < 45; i++ {
		store.Append(i)
	}

	fmt.Println(store.GetStore(), store.Cap())
}
