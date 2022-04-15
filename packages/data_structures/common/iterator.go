package common

import "hello/packages/data_structures/models"

type Iterator[T any] struct {
	current *models.Node[T]
}

func NewIterator[T any](node *models.Node[T]) *Iterator[T] {
	return &Iterator[T]{
		current: node,
	}
}

func (i *Iterator[T]) Next() *models.Node[T] {
	next := i.current
	i.current = i.current.Next()
	return next
}

func (i *Iterator[T]) HasNext() bool {
	return i.current != nil
}

func (i *Iterator[T]) SetCurrent(node *models.Node[T]) {
	i.current = node
}
