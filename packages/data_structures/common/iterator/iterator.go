package iterator

import "hello/packages/data_structures/models"

type Iterator[T any] struct {
	current *models.Node[T]
}

func NewIterator[T any](node *models.Node[T]) IIterator[T] {
	return &Iterator[T]{
		current: node,
	}
}

func (i *Iterator[T]) Current() *models.Node[T] {
	return i.current
}

func (i *Iterator[T]) Next() *models.Node[T] {
	next := i.current.Next()
	i.current = next
	return next
}

func (i *Iterator[T]) HasNext() bool {
	return i.current != nil
}
