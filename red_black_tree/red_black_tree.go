package red_black_tree

import (
	"math"
	q "root/quick_store"
)

type Color int

const (
	RED   Color = iota //0
	BLACK              //1
)

type Node[K string | int, V any] struct {
	Key    K
	Value  V
	Color  Color
	Parent *Node[K, V]
	Left   *Node[K, V]
	Right  *Node[K, V]
}

type RBTree[K string | int, V any] struct {
	root   *Node[K, V]
	memory q.QuickStore[*Node[K, V]]

	result []K //ReadBlackTree
	stack  q.QuickStore[*Node[K, V]]
}

func NewRBTree[K string | int, V any](capacity int) RBTree[K, V] {
	return RBTree[K, V]{
		memory: q.New[*Node[K, V]](capacity),

		result: make([]K, 0, capacity), //0 => append
		stack:  q.New[*Node[K, V]](int(math.Log2(float64(capacity)))),
	}
}

func NewNode[K string | int, V any](key K, value V) *Node[K, V] {
	return &Node[K, V]{
		Key:   key,
		Color: RED,
	}
}

func (tree *RBTree[K, V]) AddNode(key K, value V) *Node[K, V] {
	tree.memory.Append(NewNode(key, value))

	return tree.memory.Get()
}

func (tree *RBTree[K, V]) Insert(key K, value V) {
	newNode := tree.AddNode(key, value)

	if tree.root == nil {
		tree.root = newNode
	} else {
		currentNode := tree.root
		var parentNode *Node[K, V]

		//find parent for our newNode
		for currentNode != nil {
			parentNode = currentNode
			if newNode.Key < currentNode.Key {
				currentNode = currentNode.Left
			} else if newNode.Key > currentNode.Key {
				currentNode = currentNode.Right
			} else {
				newNode.Key = key
				return //if the find the same key
			}
		}

		//Set parentNode to our NewNode
		newNode.Parent = parentNode
		if newNode.Key < parentNode.Key {
			parentNode.Left = newNode
		} else if newNode.Key > parentNode.Key {
			parentNode.Right = newNode
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
func (tree *RBTree[K, V]) rotateLeft(n *Node[K, V]) {
	rChild := n.Right
	n.Right = rChild.Left
	if rChild.Left != nil {
		rChild.Left.Parent = n
	}
	rChild.Parent = n.Parent
	if n.Parent == nil {
		// lChild is new root
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
func (tree *RBTree[K, V]) rotateRight(n *Node[K, V]) {
	lChild := n.Left
	n.Left = lChild.Right
	if lChild.Right != nil {
		lChild.Right.Parent = n
	}
	lChild.Parent = n.Parent
	if n.Parent == nil {
		// lChild is new root
		tree.root = lChild
	} else if n == n.Parent.Right {
		n.Parent.Right = lChild
	} else {
		n.Parent.Left = lChild
	}
	lChild.Right = n
	n.Parent = lChild
}

func (tree *RBTree[K, V]) InOrderTraversal() []K {
	current := tree.root

	for current != nil || tree.stack.Len() > 0 { //len
		for current != nil {
			tree.stack.Append(current) //add
			current = current.Left
		}

		current = tree.stack.Get() //get
		tree.stack.Delete()        //delete

		tree.result = append(tree.result, current.Key)

		current = current.Right
	}

	return tree.result
}
