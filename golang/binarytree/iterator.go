package binarytree

type NodeIterator interface {
	HasNext() bool
	Next() *Node
}

type BFSNodeIterator struct {
	stack []*Node
}

func NewBFSNodeIterator(root RootNode) *BFSNodeIterator {
	return &BFSNodeIterator{stack: []*Node{root.Root()}}
}

func (i *BFSNodeIterator) HasNext() bool {
	return len(i.stack) > 0
}

func (i *BFSNodeIterator) Next() *Node {
	next := i.stack[0]
	i.stack = i.stack[1:]
	if next.Left != nil {
		i.stack = append(i.stack, next.Left)
	}
	if next.Right != nil {
		i.stack = append(i.stack, next.Right)
	}
	return next
}

type OrderedNodeIterator struct {
	stack []*Node
}

func NewOrderedNodeIterator(root RootNode) *OrderedNodeIterator {
	it := &OrderedNodeIterator{}
	it.pushLeft(root.Root())
	return it
}

func (i *OrderedNodeIterator) HasNext() bool {
	return len(i.stack) > 0
}

func (i *OrderedNodeIterator) Next() *Node {
	if !i.HasNext() {
		panic("No more elements")
	}

	next := i.stack[len(i.stack)-1]
	i.stack = i.stack[:len(i.stack)-1]

	if next.Right != nil {
		i.pushLeft(next.Right)
	}

	return next
}

func (i *OrderedNodeIterator) pushLeft(n *Node) {
	for n != nil {
		i.stack = append(i.stack, n)
		n = n.Left
	}
}

type PeakingNodeIterator struct {
	it     OrderedNodeIterator
	peaked *Node
}

func (p *PeakingNodeIterator) Peak() *Node {
	if p.peaked != nil {
		next := p.it.Next()
		p.peaked = next
	}

	return p.peaked
}

func (p *PeakingNodeIterator) HasNext() bool {
	return p.peaked != nil || p.it.HasNext()
}

func (p *PeakingNodeIterator) Next() *Node {
	if p.peaked != nil {
		next := p.peaked
		p.peaked = nil
		return next
	}
	return p.it.Next()
}
