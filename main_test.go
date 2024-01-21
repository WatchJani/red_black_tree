package main

import (
	"math/rand"
	"testing"
)

// 182.9 ns/op
func BenchmarkRBTree(b *testing.B) {
	b.StopTimer()

	rb_tree := NewRBTree(b.N)
	data := make([]int, b.N)

	for index := range data {
		data[index] = rand.Intn(40000)
	}
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		rb_tree.Insert(data[i])
	}
}

// func BenchmarkWriteAll(b *testing.B) {
// 	b.StopTimer()
// 	tree := NewRBTree(40000) //create capacity for 40_000 nods

// 	for i := 0; i < 100; i++ {
// 		tree.Insert(rand.Intn(400000)) //insert to our tree
// 	}
// 	b.StartTimer()

// 	//test our implementation
// 	for i := 0; i < b.N; i++ {
// 		tree.InOrderTraversal()
// 	}
// }

func BenchmarkReadAll2(b *testing.B) {
	b.StopTimer()
	tree := NewRBTree(40000) //create capacity for 40_000 nods

	for i := 0; i < 100; i++ {
		tree.Insert(rand.Intn(400000)) //insert to our tree
	}
	b.StartTimer()

	//test our implementation
	for i := 0; i < b.N; i++ {
		tree.InOrderTraversal2()
	}
}
