package main

import "fmt"

func main() {
	rb_tree := NewRBTree(20)

	rb_tree.Insert(11)
	rb_tree.Insert(12)
	rb_tree.Insert(14)

	fmt.Println(rb_tree.memory[0])
}

type Color int

const (
	RED   Color = iota //0
	BLACK              //1
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
		//!save new value :D
		return //if the find the same key
	}

	//making balanced tree
	for parentNode.Parent != nil && parentNode.Color == RED {
		grandparentNode := parentNode.Parent

		if parentNode == grandparentNode.Left {
			uncle := grandparentNode.Right

			if uncle != nil && uncle.Color == RED {
				grandparentNode.Color = RED
				parentNode.Color = BLACK
				uncle.Color = BLACK
				newNode = grandparentNode
			} else {
				if newNode == parentNode.Right {
					newNode = parentNode
					tree.rotateLeft(newNode)
				}

				parentNode.Color = BLACK
				grandparentNode.Color = RED
				tree.rotateRight(grandparentNode)
			}
		} else {
			uncle := grandparentNode.Left

			if uncle != nil && uncle.Color == RED {
				grandparentNode.Color = RED
				parentNode.Color = BLACK
				uncle.Color = BLACK
				newNode = grandparentNode
			} else {
				if newNode == parentNode.Left {
					newNode = parentNode
					tree.rotateRight(newNode)
				}

				parentNode.Color = BLACK
				grandparentNode.Color = RED
				tree.rotateLeft(grandparentNode)
			}
		}
		parentNode = newNode.Parent
	}

	tree.root.Color = BLACK
}

// here will exist one copy of element (one node more), why?
// we will set our node as root, but we don't know where is
// this element in memory slice, because of that we will old
// root set on end of slice (end is our pointer)
func (tree *RBTree) RootSwap(node *Node) {
	// tree.pointer++ //Provide space for a new node.
	// tree.memory[0], tree.memory[tree.capacity] = node, tree.memory[0]

	//?????
	tree.root, node = node, tree.root
}

// make left rotation in RBTree
func (tree *RBTree) rotateLeft(n *Node) {
	rChild := n.Right
	n.Right = rChild.Left
	if rChild.Left != nil {
		rChild.Left.Parent = n
	}
	rChild.Parent = n.Parent
	if n.Parent == nil {
		// n is root
		tree.RootSwap(n)
	} else if n == n.Parent.Left {
		n.Parent.Left = rChild
	} else {
		n.Parent.Right = rChild
	}
	rChild.Left = n
	n.Parent = rChild
}

// make right rotation in RBTree
func (tree *RBTree) rotateRight(n *Node) {
	lChild := n.Left
	n.Left = lChild.Right
	if lChild.Right != nil {
		lChild.Right.Parent = n
	}
	lChild.Parent = n.Parent
	if n.Parent == nil {
		// n is root
		tree.RootSwap(n)
	} else if n == n.Parent.Right {
		n.Parent.Right = lChild
	} else {
		n.Parent.Left = lChild
	}
	lChild.Right = n
	n.Parent = lChild
}
