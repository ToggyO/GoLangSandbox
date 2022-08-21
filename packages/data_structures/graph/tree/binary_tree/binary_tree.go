package binary_tree

import "hello/packages/data_structures/models"

type ComparerFunc[K interface{}] func(obj1, obj2 K) int

type BinaryTree[T interface{}] struct {
	root     *models.TreeNode[T]
	comparer ComparerFunc[T]
}

// TODO
//  1. добавить поле, содержащее высоту поддерева данной вершины
//  2. создать метод, который пересчитывает высоту затронутого поддерева, при операциях изменяющих высоту дерева

func NewBinaryTree[T interface{}](comparer ComparerFunc[T]) *BinaryTree[T] {
	return &BinaryTree[T]{comparer: comparer}
}

func (b *BinaryTree[T]) Insert(value T) {
	if b.root == nil {
		b.root = models.NewTreeNode[T](value)
		return
	}
	b.insert(b.root, value)
}

func (b *BinaryTree[T]) Remove(value T) bool {
	// TODO: check on value existence
	parent := b.root
	current := b.root
	var isCurrentNodeLeft bool

	// метод, возвращающий высоту дерева, поддерева
	for current != nil {
		compareResult := b.comparer(value, current.GetValue())
		var nextNode *models.TreeNode[T]
		switch compareResult {
		case 0:
			b.handleRemoveNodes(current, parent, isCurrentNodeLeft)
			return true
		case 1:
			isCurrentNodeLeft = false
			nextNode = current.GetRightNode()
			break
		case -1:
			isCurrentNodeLeft = true
			nextNode = current.GetLeftNode()
			break
		}

		parent = current
		current = nextNode
	}

	return false
}

func (b *BinaryTree[T]) Find(value T) bool {
	if b.root == nil {
		return false
	}

	hasValue, _ := b.find(b.root, value)
	return hasValue
}

func (b *BinaryTree[T]) find(current *models.TreeNode[T], value T) (bool, *models.TreeNode[T]) {
	compareResult := b.comparer(value, current.GetValue())
	var nextNode *models.TreeNode[T]

	switch compareResult {
	case 0:
		return true, current
	case 1:
		nextNode = current.GetRightNode()
		break
	case -1:
		nextNode = current.GetLeftNode()
		break
	}

	if nextNode == nil {
		return false, nil
	}
	return b.find(nextNode, value)
}

func (b *BinaryTree[T]) insert(current *models.TreeNode[T], value T) {
	compareResult := b.comparer(value, current.GetValue())
	var nextNode *models.TreeNode[T]

	switch compareResult {
	case 0, 1:
		nextNode = current.GetRightNode()
		if nextNode == nil {
			current.InsertRight(value)
			current.RecalculateSubtreeHeight()
			//current.IncrementSubtreeHeight()

			//if current.GetSubtreeHeight() == 0 {
			//	current.IncrementSubtreeHeight()
			//	incremented = true
			//}
			return
		}
		break
	case -1: // меньше
		nextNode = current.GetLeftNode()
		if nextNode == nil {
			current.InsertLeft(value)
			current.RecalculateSubtreeHeight()
			//if current.GetSubtreeHeight() == 0 {
			//	current.IncrementSubtreeHeight()
			//	incremented = true
			//}
			return
		}
		break
	}

	current.RecalculateSubtreeHeight()
	//incremented = b.insert(nextNode, value)
	//if incremented {
	//	current.IncrementSubtreeHeight()
	//}
	//return
	b.insert(nextNode, value)
}

//             current
//        ?            nil <--- сюда вставляем новую ноду

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

//            30   + только один потомок
//    26                61 <----
//  .......       40           74*
//            38      52    67     77
//
// при удалении правого потомка (при удалении левого ставить самого младшего потомка из двух)
//            30   + только один потомок
//    26                74   вместо родителя ставим правого потомка
//                  67      77
//                  самому левому потомку всех потомков присваиваем левого брата
//
//  .......       40           74*
//            38      52    67     77
//
