package main

import (
	"bytes"
	"math/rand"
	"testing"
)

func TestRootSwap(t *testing.T) {
	// rb_tree := NewRBTree(20)
}

func BenchmarkByteCompare(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bytes.Compare([]byte("Rotation are fundamental operations that help in maintaining the Red-Black Tree properties. There are two types of rotations: left rotation and right rotation."), []byte("Rotations are fundamental operations that help in maintaining the Red-Black Tree properties. There are two types of rotations: left rotation and right rotation."))
	}
}

func BenchmarkResetToZero(b *testing.B) {
	for i := 0; i < b.N; i++ {

	}
}

func BenchmarkRBTree(b *testing.B) {
	rb_tree := NewRBTree(b.N)

	for i := 0; i < b.N; i++ {
		rb_tree.Insert(rand.Intn(40000))
	}
}

func BenchmarkWriteAll(b *testing.B) {
	b.StopTimer()
	tree := NewRBTree(40000)

	for i := 0; i < 100; i++ {
		num := rand.Intn(40000)
		tree.Insert(num)
	}
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		tree.InOrderTraversal()
	}
}
