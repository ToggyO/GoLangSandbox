package linked_list

import (
    "hello/packages/data_structures/models"
)

// ONLY FOR TESTS
type iTestLinkedList[T any] interface {
    Append(data T)
    AppendFirst(data T)
    RemoveHead()
    Remove(data T) bool
    NewIterator() iIterator[T]
}

type iIterator[T any] interface {
    Current() *models.Node[T]
    Next() *models.Node[T]
    HasNext() bool
}
