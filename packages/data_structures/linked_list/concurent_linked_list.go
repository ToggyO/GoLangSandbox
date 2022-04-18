package linked_list

import (
	"github.com/google/go-cmp/cmp"
	"hello/packages/data_structures/common"
	"hello/packages/data_structures/models"
	"sync"
)

type ConcurrentLinkedList[T any] struct {
	head *models.Node[T]
	tail *models.Node[T]
	len  int

	lock sync.RWMutex
}

func NewConcurrentLinkedList[T any]() *ConcurrentLinkedList[T] {
	return &ConcurrentLinkedList[T]{
		head: nil,
		tail: nil,
		len:  0,
	}
}

func (l *ConcurrentLinkedList[T]) Append(data T) {
	l.lock.Lock()
	defer l.lock.Unlock()

	node := models.NewNode(data, nil)

	if l.head == nil {
		l.head = node
	} else {
		l.tail.SetNext(node)
	}

	l.tail = node
	l.len++
}

func (l *ConcurrentLinkedList[T]) AppendFirst(data T) {
	l.lock.Lock()
	defer l.lock.Unlock()

	node := models.NewNode(data, l.head)
	if l.len == 0 {
		l.tail = node
	}

	l.head = node
	l.len++
}

func (l *ConcurrentLinkedList[T]) RemoveHead() {
	l.lock.Lock()
	defer l.lock.Unlock()

	l.head = l.head.Next()
}

func (l *ConcurrentLinkedList[T]) Remove(data T) bool {
	l.lock.Lock()
	defer l.lock.Unlock()

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

func (l *ConcurrentLinkedList[T]) Iterator() *common.Iterator[T] {
	return common.NewIterator(l.head)
}

func (l *ConcurrentLinkedList[T]) handleRemove(current, previous *models.Node[T]) {
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
