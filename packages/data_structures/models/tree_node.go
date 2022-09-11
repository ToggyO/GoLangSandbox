package models

import "math"

type TreeNode[T interface{}] struct {
	value         T
	left          *TreeNode[T]
	right         *TreeNode[T]
	subtreeHeight int
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
	n.left = node // TODO использовать инсерт для добавления потомка
}

func (n *TreeNode[T]) UpdateRightNode(node *TreeNode[T]) {
	n.right = node // TODO использовать инсерт для добавления потомка
}

func (n *TreeNode[T]) GetSubtreeHeight() int {
	return n.subtreeHeight
}

func (n *TreeNode[T]) RecalculateSubtreeHeight(reduceHeight bool) {
	count := 1
	if reduceHeight {
		count = -1
	}

	if n.left == nil && n.right == nil {
		n.subtreeHeight = 1
		return
	}

	if n.left != nil && n.right != nil {
		maxSubtreeHeight := math.Max(float64(n.left.subtreeHeight), float64(n.right.subtreeHeight))
		n.subtreeHeight = int(maxSubtreeHeight) + count
		return
	}

	if n.left != nil {
		n.subtreeHeight = n.left.subtreeHeight + count
		return
	}

	if n.right != nil {
		n.subtreeHeight = n.right.subtreeHeight + count
		return
	}
}
