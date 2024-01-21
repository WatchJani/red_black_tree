package red_black_tree

import (
	"math/rand"
	"testing"
)

const TEST_SAMPLE = 40_000

func BenchmarkRBTree(b *testing.B) {
	b.StopTimer()

	rb_tree := NewRBTree(b.N)
	data := make([]int, b.N)

	for index := range data {
		data[index] = rand.Intn(TEST_SAMPLE)
	}
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		rb_tree.Insert(data[i])
	}
}

func BenchmarkReadAll(b *testing.B) {
	b.StopTimer()
	tree := NewRBTree(TEST_SAMPLE) //create capacity for 40_000 nods

	for i := 0; i < 100; i++ {
		tree.Insert(rand.Intn(TEST_SAMPLE)) //insert to our tree
	}
	b.StartTimer()

	//test our implementation
	for i := 0; i < b.N; i++ {
		tree.InOrderTraversal()
	}
}
