package binarytree

import "testing"

func compare(x interface{}, y interface{}) bool {
	return x.(int) < y.(int)
}

func Test_binaryTree(t *testing.T) {
	tree := New(compare)

	tree.Insert(1)
	tree.Insert(2)
	tree.Insert(3)

	findTree := tree.Search(2)
	if findTree.Node != 2 {
		t.Error("[Error] Search error")
	}

	findNilTree := tree.Search(100)

	if findNilTree != nil {
		t.Error("[Error] 2. Search error")
	}
}
