package binarytree

type CompareFunc func(a interface{}, b interface{}) bool

type BinaryTree struct {
	Node    interface{}
	Left    *BinaryTree
	Right   *BinaryTree
	Compare CompareFunc
}

func New(fn CompareFunc) *BinaryTree {
	return &BinaryTree{
		Node:    nil,
		Compare: fn,
	}
}

func (tree *BinaryTree) Search(value interface{}) *BinaryTree {
	if tree.Node == nil {
		return nil
	}

	if tree.Node == value {
		return tree
	}
	if tree.Compare(value, tree.Node) {
		return tree.Left.Search(value)
	} else {
		return tree.Right.Search(value)
	}
}

func (tree *BinaryTree) Insert(value interface{}) {
	if tree.Node == nil {
		tree.Node = value
		tree.Right = New(tree.Compare)
		tree.Left = New(tree.Compare)
		return
	}
	if tree.Compare(value, tree.Node) {
		tree.Left.Insert(value)
	} else {
		tree.Right.Insert(value)
	}
}
