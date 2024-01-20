package main

import (
	"testing"
)

func TestRightRotation(t *testing.T) {
	tree := NewRBTree(20)

	for _, value := range []int{18, 11, 6} {
		tree.Insert(value)
	}

	if results, expected := tree.memory[1].Left, tree.memory[2]; results != expected {
		t.Errorf("this results %+v is not valid, expected value is %+v", results, expected)
	}

	if results, expected := tree.memory[1].Right, tree.memory[0]; results != expected {
		t.Errorf("this results %+v is not valid, expected value is %+v", results, expected)
	}

	// for _, value := range []int{50, 20, 30} {
	// 	fmt.Println(value)
	// 	tree.Insert(value)
	// }

	// if results, expected := tree.memory[2].Left, tree.memory[1]; results != expected {
	// 	t.Errorf("this results %+v is not valid, expected value is %+v", results, expected)
	// }

	// if results, expected := tree.memory[0].Left, tree.memory[2]; results != expected {
	// 	t.Errorf("this results %+v is not valid, expected value is %+v", results, expected)
	// }
}
