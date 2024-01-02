package main

import (
	"bytes"
	"testing"
)

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
	rb_tree := NewRBTree(20000)

	for i := 0; i < b.N; i++ {
		rb_tree.Insert(12)
	}
}
