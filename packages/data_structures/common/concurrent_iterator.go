package common

import (
	"hello/packages/data_structures/models"
	"sync"
)

type ConcurrentIterator[T any] struct {
	iterator *Iterator[T]

	lock sync.RWMutex
}

func NewConcurrentIterator[T any](head *models.Node[T]) *ConcurrentIterator[T] {
	return &ConcurrentIterator[T]{
		iterator: NewIterator[T](head),
	}
}

func (i *ConcurrentIterator[T]) Current() *models.Node[T] {
	i.lock.RLock()
	defer i.lock.RUnlock()
	return i.iterator.Current()
}

func (i *ConcurrentIterator[T]) Next() *models.Node[T] {
	i.lock.RLock()
	defer i.lock.RUnlock()
	return i.iterator.Next()
}

func (i *ConcurrentIterator[T]) HasNext() bool {
	i.lock.RLock()
	defer i.lock.RUnlock()
	return i.iterator.HasNext()
}

