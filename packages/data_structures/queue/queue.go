package queue

import "hello/packages/data_structures/linked_list"

type Queue[T any] struct {
	list *linked_list.LinkedList[T]
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{
		list: linked_list.NewLinkedList[T](),
	}
}

func (q *Queue[T]) Enqueue(item T) {
	q.list.Append(item)
}

func (q *Queue[T]) Dequeue() {
	q.list.RemoveHead()
}
