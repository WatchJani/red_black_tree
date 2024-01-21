package red_black_tree

import (
	"math/rand"
	"testing"
)

const TEST_SAMPLE = 40_000

// test insert sorting ability of red black tree implementation
func Test(t *testing.T) {
	RBTree := NewRBTree(TEST_SAMPLE)

	for index := 0; index < TEST_SAMPLE; index++ {
		RBTree.Insert(rand.Intn(100_000)) //Reduced possibility of data collision
	}

	var temp int
	for index, key := range RBTree.InOrderTraversal() {
		if key < temp {
			t.Errorf("index: %d | Got %v | expected %v", index, key, temp)
		}

		temp = key
	}
}

// insert new key in red black tree
// 175ns
func BenchmarkRBTreeInsert(b *testing.B) {
	b.StopTimer()

	RBTree := NewRBTree(b.N)
	data := make([]int, b.N)

	for index := range data {
		data[index] = rand.Intn(TEST_SAMPLE)
	}

	b.StartTimer()

	for i := 0; i < b.N; i++ {
		RBTree.Insert(data[i])
	}
}

// read 40_000 sorted elements in red black tree
// 2023ns
func BenchmarkReadAll(b *testing.B) {
	b.StopTimer()
	RBTree := NewRBTree(TEST_SAMPLE) //create capacity for 40_000 nods

	for i := 0; i < 100; i++ {
		RBTree.Insert(rand.Intn(TEST_SAMPLE)) //insert to our tree
	}
	b.StartTimer()

	//test our implementation
	for i := 0; i < b.N; i++ {
		RBTree.InOrderTraversal()
	}
}
