package binarytree

import (
	"fmt"
	"reflect"
	"testing"
)

var (
	testRoot *Tree
)

func init() {
	testRoot = &Tree{}

	testRoot.Insert(20)
	testRoot.Insert(15)
	testRoot.Insert(13)

	testRoot.Insert(18)

	testRoot.Insert(25)
	testRoot.Insert(45)
	testRoot.Insert(50)
	testRoot.Insert(48)
	testRoot.Insert(55)

	printPrettyTree(testRoot.Root(), "", true, false)
}

// printPrettyTree prints the tree in a "pretty" tree format
func printPrettyTree(node *Node, prefix string, isRoot, isLeft bool) {
	if node == nil {
		return
	}

	// Print the right subtree first (so it appears on top when printed)
	if node.Right != nil {
		printPrettyTree(node.Right, prefix+"        ", false, false)
	}

	// Print the current node
	fmt.Printf("%s", prefix)
	if !isRoot {
		if isLeft {
			fmt.Printf("└──")
		} else {
			fmt.Printf("┌──")
		}
	}
	fmt.Printf("%d (%d)\n", node.Value, node.Height())

	// Print the left subtree (so it appears at the bottom)
	if node.Left != nil {
		printPrettyTree(node.Left, prefix+"        ", false, true)
	}
}

func TestBFSNodeIterator(t *testing.T) {
	expected := []int{45, 20, 50, 15, 25, 48, 55, 13, 18}

	it := NewBFSNodeIterator(testRoot.Root())
	for _, exp := range expected {
		n := it.Next()
		if n.Value != exp {
			t.Errorf("expected %d, got %d", exp, n.Value)
		}
	}

	if it.HasNext() {
		t.Errorf("should have reached the end of the tree")
	}
}

func TestOrderedNodeIterator(t *testing.T) {
	tests := []struct {
		in  []int
		out []int
	}{
		{
			[]int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1},
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
		{
			[]int{10, 1, 9, 2, 8, 3, 7, 4, 6, 5},
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("case[%d]", i), func(t *testing.T) {
			tree := NewTree()
			for _, v := range tt.in {
				tree.Insert(v)
			}

			var res []int
			it := NewOrderedNodeIterator(tree)
			for it.HasNext() {
				res = append(res, it.Next().Value)
			}

			if !reflect.DeepEqual(res, tt.out) {
				t.Errorf("OrderedNodeIterator = %v, want %v", res, tt.out)
			}
		})
	}
}
