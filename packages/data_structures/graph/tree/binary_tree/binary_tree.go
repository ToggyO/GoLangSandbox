package binary_tree

import "hello/packages/data_structures/models"

type ComparerFunc[K interface{}] func(obj1, obj2 K) int

type BinaryTree[T interface{}] struct {
	root      *models.TreeNode[T]
	comparer  ComparerFunc[T]
	traverser *TreeTraverser[T]
}

// TODO
//  1. добавить поле, содержащее высоту поддерева данной вершины
//  2. создать метод, который пересчитывает высоту затронутого поддерева, при операциях изменяющих высоту дерева

func NewBinaryTree[T interface{}](comparer ComparerFunc[T]) *BinaryTree[T] {
	tree := &BinaryTree[T]{comparer: comparer}
	tree.traverser = NewTreeTraverser[T](tree)
	return tree
}

func (b *BinaryTree[T]) Traverse() {
	//b.traverser.Postorder()
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
	if b.root == nil {
		return false
	}

	current := b.root
	var parent *models.TreeNode[T]
	var isCurrentNodeLeft bool

	// метод, возвращающий высоту дерева, поддерева
	//for current != nil {
	//compareResult := b.comparer(value, current.GetValue())
	//var nextNode *models.TreeNode[T]
	//switch compareResult {
	//case 0:
	//	b.handleRemoveNodes(current, parent, isCurrentNodeLeft)
	//	return true
	//case 1:
	//	isCurrentNodeLeft = false
	//	nextNode = current.GetRightNode()
	//	break
	//case -1:
	//	isCurrentNodeLeft = true
	//	nextNode = current.GetLeftNode()
	//	break
	//}
	//
	//parent = current
	//current = nextNode
	//}
	b.remove(current, parent, value, isCurrentNodeLeft)

	return false
}

func (b *BinaryTree[T]) remove(
	current *models.TreeNode[T],
	parent *models.TreeNode[T],
	value T,
	isCurrentNodeLeft bool,
) {
	compareResult := b.comparer(value, current.GetValue())
	var nextNode *models.TreeNode[T]
	switch compareResult {
	case 0:
		b.handleRemoveNodes(current, parent, isCurrentNodeLeft)
		b.traverser.Postorder(func(node *models.TreeNode[T]) {
			node.RecalculateSubtreeHeight(false)
		})
		// обойти всех потомков parent и пересчитать высоты поддеревьев

		//if isCurrentNodeLeft {
		//
		//}
		parent.RecalculateSubtreeHeight(true)
		return
		//return true
	case 1:
		isCurrentNodeLeft = false
		nextNode = current.GetRightNode()
		break
	case -1:
		isCurrentNodeLeft = true
		nextNode = current.GetLeftNode()
		break
	}

	b.remove(nextNode, current, value, isCurrentNodeLeft)

	if parent != nil {
		parent.RecalculateSubtreeHeight(true)
	}
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
			current.RecalculateSubtreeHeight(false)
			return
		}
		break
	case -1: // меньше
		nextNode = current.GetLeftNode()
		if nextNode == nil {
			current.InsertLeft(value)
			current.RecalculateSubtreeHeight(false)
			return
		}
		break
	}

	b.insert(nextNode, value)
	current.RecalculateSubtreeHeight(false)
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
