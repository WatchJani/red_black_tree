package main

import "fmt"

func main() {
	tree := NewRBTree(20)

	for _, value := range []int{18, 11, 6, 50, 40} {
		tree.Insert(value)
	}

	fmt.Println(tree.memory[1]) //root
	fmt.Println(tree.memory[2]) //left
	fmt.Println(tree.memory[0]) //right
	fmt.Println(tree.memory[3]) //right
	fmt.Println(tree.memory[4]) //right

	// for _, value := range []int{50, 40} {
	// 	fmt.Println(value)
	// 	tree.Insert(value)
	// }

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

func (tree *RBTree) Reset() {
	tree.pointer = 0
}

func (tree *RBTree) AddNode(key int) *Node {
	tree.pointer++

	if tree.pointer == tree.capacity {
		tree.Reset()
	}

	tree.memory[tree.pointer] = NewNode(key)

	return tree.memory[tree.pointer]
}

// 100ns
func (tree *RBTree) Insert(key int) {
	newNode := tree.AddNode(key) //Done

	if tree.root == nil { //Done
		tree.root = newNode
	} else {
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
	}

	parentNode := newNode.Parent //get the parent of new node

	//making balanced tree
	for parentNode != nil && parentNode.Color == RED {
		grandparentNode := parentNode.Parent // get parent of parent, that we can came to uncle

		//Check uncle, if parent left of grandparent then uncle is right
		if parentNode == grandparentNode.Left {
			uncle := grandparentNode.Right //get the uncle

			if uncle != nil && uncle.Color == RED { //if uncle exist and if red then recolor
				grandparentNode.Color = RED
				parentNode.Color = BLACK
				uncle.Color = BLACK
				newNode = grandparentNode
			} else { //if uncle not exist then rotate and recolor
				if newNode == parentNode.Right {
					tree.rotateLeft(parentNode)
					newNode, parentNode = parentNode, newNode
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
					tree.rotateRight(parentNode)
					newNode, parentNode = parentNode, newNode
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
		tree.root = rChild
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
		tree.root = lChild
	} else if n == n.Parent.Right {
		n.Parent.Right = lChild
	} else {
		n.Parent.Left = lChild
	}
	lChild.Right = n
	n.Parent = lChild
}
