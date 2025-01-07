package binarytree

type Tree struct {
	root *Node
}

func NewTree() *Tree {
	return &Tree{}
}

func (t *Tree) Root() *Node {
	return t.root
}

func (t *Tree) Insert(value int) {
	t.root = t.root.Insert(value)
}

func (t *Tree) Search(value int) (*Node, bool) {
	node := t.root
	for node != nil {
		if node.Value == value {
			return node, true
		} else if value < node.Value {
			node = node.Left
		} else {
			node = node.Right
		}
	}
	return nil, false
}

type Node struct {
	Value       int
	Left, Right *Node
	height      int
	count       int
}

func NewNode(value int) *Node {
	return &Node{Value: value, count: 1}
}

func (n *Node) Insert(value int) *Node {
	if n == nil {
		return NewNode(value)
	}

	if value < n.Value {
		n.Left = n.Left.Insert(value)
	} else if value > n.Value {
		n.Right = n.Right.Insert(value)
	} else {
		n.count++
		return n // duplicate
	}
	n.updateHeight()

	balance := n.getBalance()

	if balance >= -1 && balance <= 1 {
		return n
	}

	// Case 1: Left Left (LL)
	if balance > 1 && value < n.Left.Value {
		return leftRotate(n)
	}

	// Case 2: Right Right (RR)
	if balance < -1 && value > n.Right.Value {
		return rightRotate(n)
	}

	// Case 3: Left Right (LR)
	if balance > 1 && value > n.Left.Value {
		n.Left = leftRotate(n.Left)
		return rightRotate(n)
	}

	// Case 4: Right Left (RL)
	if balance < -1 && value < n.Right.Value {
		n.Right = rightRotate(n.Right)
		return leftRotate(n)
	}

	return n
}

func (n *Node) Height() int {
	if n == nil {
		return -1
	}
	return n.height
}

func (n *Node) updateHeight() {
	n.height = 1 + max(n.Left.Height(), n.Right.Height())
}

func (n *Node) getBalance() int {
	if n == nil {
		return 0
	}
	return n.Left.Height() - n.Right.Height()
}

func leftRotate(n *Node) *Node {
	x := n.Left
	xR := x.Right

	x.Right = n
	n.Left = xR

	x.updateHeight()
	n.updateHeight()

	return x
}

func rightRotate(n *Node) *Node {
	x := n.Right
	xL := x.Left

	x.Left = n
	n.Right = xL

	x.updateHeight()
	n.updateHeight()

	return x
}

type RootNode interface {
	Root() *Node
}
