package main

import "fmt"

func main() {
	rb_tree := NewRBTree(20)

	rb_tree.Insert(11)

	fmt.Println(rb_tree)
}

type Color int

const (
	RED Color = iota
	BLACK
)

type Node struct {
	Key    int
	Color  Color
	Parent *Node
	Left   *Node
	Right  *Node
}

type RBTree struct {
	root     *Node
	memory   []*Node
	pointer  int
	capacity int
}

func NewRBTree(capacity int) *RBTree {
	return &RBTree{
		memory:   make([]*Node, capacity),
		pointer:  -1,
		capacity: capacity,
	}
}

func NewNode(key int) *Node {
	return &Node{
		Key:   key,
		Color: RED,
	}
}

func (tree *RBTree) AddNode(key int) *Node {
	tree.pointer++

	if tree.pointer == tree.capacity {
		tree.pointer = 0
	}

	tree.memory[tree.pointer] = NewNode(key)

	return tree.memory[tree.pointer]
}

// 100ns
func (tree *RBTree) Insert(key int) {
	newNode := tree.AddNode(key) //Done

	if tree.root == nil { //Done
		tree.root = newNode
		return
	}

	currentNode := tree.root
	var parentNode *Node

	//find parent for our newNode
	for currentNode != nil {
		parentNode = currentNode
		if newNode.Key < currentNode.Key {
			currentNode = currentNode.Left
		} else {
			currentNode = currentNode.Right
		}
	}

	//Set parentNode to our NewNode
	newNode.Parent = parentNode
	if newNode.Key < parentNode.Key {
		parentNode.Left = newNode
	} else if newNode.Key > parentNode.Key {
		parentNode.Right = newNode
	} else {
		//save new value :D
		return //if the find the same key
	}
}

// func (n *Node) rotateLeft() *Node {
// 	rChild := n.Right
// 	n.Right = rChild.Left
// 	if rChild.Left != nil {
// 		rChild.Left.Parent = n
// 	}
// 	rChild.Parent = n.Parent
// 	if n.Parent == nil {
// 		// n is root
// 		root = rChild
// 	} else if n == n.Parent.Left {
// 		n.Parent.Left = rChild
// 	} else {
// 		n.Parent.Right = rChild
// 	}
// 	rChild.Left = n
// 	n.Parent = rChild
// 	return rChild
// }

// func (n *Node) rotateRight() *Node {
// 	lChild := n.Left
// 	n.Left = lChild.Right
// 	if lChild.Right != nil {
// 		lChild.Right.Parent = n
// 	}
// 	lChild.Parent = n.Parent
// 	if n.Parent == nil {
// 		// n is root
// 		root = lChild
// 	} else if n == n.Parent.Right {
// 		n.Parent.Right = lChild
// 	} else {
// 		n.Parent.Left = lChild
// 	}
// 	lChild.Right = n
// 	n.Parent = lChild
// 	return lChild
// }
