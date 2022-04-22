package iterator

import "hello/packages/data_structures/models"

type IIterator[T any] interface {
	Current() *models.Node[T]
	Next() *models.Node[T]
	HasNext() bool
}
