package queue

import "hello/packages/data_structures/linked_list"

type Queue[T any] struct {
	list linked_list.ILinkedList[T]
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{
		list: linked_list.NewConcurrentLinkedList[T](),
	}
}

// Enqueue appends an element to the back of the queue.
func (q *Queue[T]) Enqueue(item T) {
	q.list.Append(item)
}

// Dequeue removes and returns the element from the front of the queue.
func (q *Queue[T]) Dequeue() T {
	head := q.list.Head()
	q.list.RemoveHead()
	return head
}

// Peek returns the element at the front of the queue. This is the element
// that would be returned by Dequeue()
func (q *Queue[T]) Peek() T {
	return q.list.Head()
}

// Len returns the count of elements in the queue.
func (q *Queue[T]) Len() int {
	return q.list.Len()
}
