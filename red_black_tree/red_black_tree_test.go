package red_black_tree

import (
	"fmt"
	"math/rand"
	"testing"
)

const TEST_SAMPLE = 40_000

// test insert sorting ability of red black tree implementation
// func Test(t *testing.T) {
// 	//key is type of int, value is type of int
// 	RBTree := NewRBTree[int, int](TEST_SAMPLE)

// 	for index := 0; index < TEST_SAMPLE; index++ {
// 		RBTree.Insert(rand.Intn(100_000), 5) //Reduced possibility of data collision
// 	}

// 	var temp int
// 	for index, key := range RBTree.InOrderTraversal() {
// 		if key < temp {
// 			t.Errorf("index: %d | Got %v | expected %v", index, key, temp)
// 		}

// 		temp = key
// 	}
// }

// insert new key in red black tree
// 146ns
func BenchmarkRBTreeInsert(b *testing.B) {
	b.StopTimer()

	//key is type of int, value is type of int
	RBTree := NewRBTree[int, int](b.N)
	data := make([]int, b.N)

	for index := range data {
		data[index] = rand.Intn(TEST_SAMPLE)
	}

	b.StartTimer()

	for i := 0; i < b.N; i++ {
		RBTree.Insert(data[i], 5)
	}
}

// bug??? 500ns
func BenchmarkRBTreeInsertString(b *testing.B) {
	b.StopTimer()

	//key is type of int, value is type of int
	RBTree := NewRBTree[string, int](b.N)
	data := make([]string, b.N)

	for index := range data {
		data[index] = fmt.Sprintf("%d", rand.Intn(500_000))
	}

	b.StartTimer()

	for i := 0; i < b.N; i++ {
		RBTree.Insert(data[i], 5)
	}
}

// read 40_000 sorted elements in red black tree
// 2023ns
func BenchmarkReadAll(b *testing.B) {
	b.StopTimer()

	//key is type of int, value is type of int
	RBTree := NewRBTree[int, int](TEST_SAMPLE) //create capacity for 40_000 nods

	for i := 0; i < 100; i++ {
		RBTree.Insert(rand.Intn(TEST_SAMPLE), 5) //insert to our tree
	}
	b.StartTimer()

	//test our implementation
	for i := 0; i < b.N; i++ {
		RBTree.InOrderTraversal()
	}
}

func BenchmarkRBTreeInsert40_000(b *testing.B) {
	b.StopTimer()

	//key is type of int, value is type of int
	RBTree := NewRBTree[string, int](b.N)
	data := make([]string, 40_000)

	for index := range data {
		data[index] = fmt.Sprintf("%d", rand.Intn(500_000))
	}

	b.StartTimer()

	for i := 0; i < b.N; i++ {
		for j := 0; j < 40_000; j++ {
			RBTree.Insert(data[j], 5)
		}
	}
}
