package linked_list

import (
	"fmt"
	"github.com/google/go-cmp/cmp"
	"hello/packages/data_structures/common/iterator"
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
	}
}

func (l *LinkedList[T]) Len() int {
	return l.len
}

func (l *LinkedList[T]) Head() T {
	return l.head.Data()
}

func (l *LinkedList[T]) Append(data T) {
	// TODO: delete
	fmt.Printf("Append data: %v\n", data)
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
	l.len--
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

func (l *LinkedList[T]) ForEach(action func(data T)) {
	iter := iterator.NewIterator(l.head)
	for iter.HasNext() {
		action(iter.Current().Data())
		iter.Next()
	}
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
