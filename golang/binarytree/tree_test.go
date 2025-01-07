package binarytree

import (
	"testing"
)

func TestTree_Search(t *testing.T) {
	in := []int{10, 1, 9, 2, 8, 3, 7, 4, 6, 5}
	tree := NewTree()
	for _, v := range in {
		tree.Insert(v)
	}

	node, ok := tree.Search(10)
	if !ok {
		t.Errorf("should have found the node for 10")
	} else if node.Value != 10 {
		t.Errorf("Wrong node. wanted 10, got %v", node.Value)
	}

	node, ok = tree.Search(100)
	if ok {
		t.Errorf("should have found the node for 100")
	}
}
