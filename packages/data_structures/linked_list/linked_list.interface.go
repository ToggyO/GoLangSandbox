package linked_list

type ILinkedList[T any] interface {
	Len() int
	Head() T
	Append(data T)
	AppendFirst(data T)
	RemoveHead()
	Remove(data T) bool
	ForEach(action func(data T))
}
