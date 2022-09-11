package binary_tree

import "hello/packages/data_structures/models"

func (b *BinaryTree[T]) handleRemoveNodes(current, parent *models.TreeNode[T], isCurrentNodeLeft bool) {
	rightNode := current.GetRightNode()
	leftNode := current.GetLeftNode()

	if rightNode == nil && leftNode == nil {
		if isCurrentNodeLeft {
			parent.UpdateLeftNode(nil)
		} else {
			parent.UpdateRightNode(nil)
		}
		return
	}

	if rightNode != nil && leftNode != nil {
		b.removeWhenHasBothChildren(current, parent, isCurrentNodeLeft)
		return
	}

	// case if only left node or only right node is not nil
	node := rightNode
	if node == nil {
		node = leftNode
	}

	if isCurrentNodeLeft {
		parent.UpdateLeftNode(node)
	} else {
		parent.UpdateRightNode(node)
	}
}

func (b *BinaryTree[T]) removeWhenHasBothChildren(current, parent *models.TreeNode[T], isCurrentNodeLeft bool) {
	leftChild := current.GetLeftNode()
	rightChild := current.GetRightNode()

	if isCurrentNodeLeft {
		parent.UpdateLeftNode(leftChild)
		b.maxValue(leftChild).UpdateRightNode(rightChild)
	} else {
		parent.UpdateRightNode(rightChild)
		b.minValue(rightChild).UpdateLeftNode(leftChild)
	}
}

// поиск минимального значения в дереве
func (b *BinaryTree[T]) minValue(parent *models.TreeNode[T]) *models.TreeNode[T] {
	if parent == nil {
		return nil
	}

	var previous *models.TreeNode[T]
	current := parent.GetLeftNode()
	for current != nil {
		previous = current
		current = current.GetLeftNode()
	}

	return previous
}

// поиск максимального значения в дереве
func (b *BinaryTree[T]) maxValue(parent *models.TreeNode[T]) *models.TreeNode[T] {
	if parent == nil {
		return nil
	}

	var previous *models.TreeNode[T]
	current := parent.GetRightNode()
	for current != nil {
		previous = current
		current = current.GetRightNode()
	}

	return previous
}
