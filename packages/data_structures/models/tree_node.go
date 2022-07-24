package models

type TreeNode[T interface{}] struct {
	value T
	left  *TreeNode[T]
	right *TreeNode[T]
}

func NewTreeNode[T interface{}](value T) *TreeNode[T] {
	return &TreeNode[T]{value: value}
}

func (n *TreeNode[T]) InsertLeft(value T) {
	node := TreeNode[T]{value: value}
	n.left = &node
}

func (n *TreeNode[T]) InsertRight(value T) {
	node := TreeNode[T]{value: value}
	n.right = &node
}

func (n *TreeNode[T]) GetValue() T {
	return n.value
}

func (n *TreeNode[T]) GetLeftNode() *TreeNode[T] {
	return n.left
}

func (n *TreeNode[T]) GetRightNode() *TreeNode[T] {
	return n.right
}

func (n *TreeNode[T]) UpdateLeftNode(node *TreeNode[T]) {
	n.left = node
}

func (n *TreeNode[T]) UpdateRightNode(node *TreeNode[T]) {
	n.right = node
}
