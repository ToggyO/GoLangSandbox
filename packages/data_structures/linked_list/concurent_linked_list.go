package linked_list

import (
	"sync"
)

type ConcurrentLinkedList[T any] struct {
	list *LinkedList[T]

	lock sync.RWMutex
}

func NewConcurrentLinkedList[T any]() *ConcurrentLinkedList[T] {
	return &ConcurrentLinkedList[T]{
		list: NewLinkedList[T](),
	}
}

func (cl *ConcurrentLinkedList[T]) Len() int {
	cl.lock.RLock()
	defer cl.lock.RUnlock()
	return cl.list.len
}

func (cl *ConcurrentLinkedList[T]) Head() T {
	cl.lock.Lock()
	defer cl.lock.Unlock()
	return cl.list.head.Data()
}

func (cl *ConcurrentLinkedList[T]) Append(data T) {
	cl.lock.Lock()
	cl.list.Append(data)
	cl.lock.Unlock()
}

func (cl *ConcurrentLinkedList[T]) AppendFirst(data T) {
	cl.lock.Lock()
	cl.list.AppendFirst(data)
	cl.lock.Unlock()
}

func (cl *ConcurrentLinkedList[T]) RemoveHead() {
	cl.lock.Lock()
	cl.list.RemoveHead()
	defer cl.lock.Unlock()
}

func (cl *ConcurrentLinkedList[T]) Remove(data T) bool {
	cl.lock.Lock()
	defer cl.lock.Unlock()
	return cl.list.Remove(data)
}

func (cl *ConcurrentLinkedList[T]) ForEach(action func(data T)) {
	cl.lock.RLock()
	cl.list.ForEach(action)
	cl.lock.RUnlock()
}
