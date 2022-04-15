package linked_list

import (
	"github.com/google/go-cmp/cmp"
	"hello/packages/data_structures/common"
	"hello/packages/data_structures/models"
)

type LinkedList[T any] struct {
	head *models.Node[T]
	tail *models.Node[T]
	len  int
}

func NewLinkedList[T any]() *LinkedList[T] {
	return &LinkedList[T]{
		head: nil,
		tail: nil,
		len:  0,
	}
}

func (l *LinkedList[T]) Append(data T) {
	node := models.NewNode(data, nil)

	if l.head == nil {
		l.head = node
	} else {
		l.tail.SetNext(node)
	}

	l.tail = node
	l.len++
}

func (l *LinkedList[T]) AppendFirst(data T) {
	node := models.NewNode(data, l.head)
	if l.len == 0 {
		l.tail = node
	}

	l.head = node
	l.len++
}

func (l *LinkedList[T]) RemoveHead() {
	l.head = l.head.Next()
}

func (l *LinkedList[T]) Remove(data T) bool {
	current := l.head
	previous := new(models.Node[T])

	for current != nil {
		if cmp.Equal(current.Data(), data) {
			l.handleRemove(current, previous)
			return true
		}

		previous = current
		current = current.Next()
	}

	return false
}

func (l *LinkedList[T]) Iterator() *common.Iterator[T] {
	return common.NewIterator(l.head)
}

func (l *LinkedList[T]) handleRemove(current, previous *models.Node[T]) {
	if previous == nil {
		l.head = current.Next()
		if l.head == nil {
			l.tail = nil
		}
	} else {
		next := current.Next()
		previous.SetNext(next)
		if next == nil {
			l.tail = previous
		}
	}

	l.len--
}
