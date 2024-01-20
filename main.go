package main

func main() {
	// tree := NewRBTree(20)

	// for i := 0; i < 100; i++ {
	// 	num := rand.Intn(40000)
	// 	tree.Insert(num)
	// }

	// for _, value := range []int{50, 40} {
	// 	fmt.Println(value)
	// 	tree.Insert(value)
	// }

	// for value := range tree.InOrderTraversal() {
	// 	saver := -1
	// 	if saver < value {
	// 		saver = value
	// 	} else {
	// 		fmt.Println("bug", saver, value)
	// 	}
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

// func (tree *RBTree) InOrderTraversal() []int {
// 	result := make([]int, 0)
// 	stack := make([]*Node, 0)
// 	current := tree.root

// 	for current != nil || len(stack) > 0 {
// 		for current != nil {
// 			stack = append(stack, current)
// 			current = current.Left
// 		}

// 		current = stack[len(stack)-1]
// 		stack = stack[:len(stack)-1]

// 		result = append(result, current.Key)

// 		current = current.Right
// 	}

// 	return result
// }
