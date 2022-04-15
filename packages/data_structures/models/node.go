package models

type Node[T any] struct {
	data T
	next *Node[T]
}

func NewNode[T any](data T, node *Node[T]) *Node[T] {
	return &Node[T]{
		data: data,
		next: node,
	}
}

func (n *Node[T]) Data() T {
	return n.data
}

func (n *Node[T]) Next() *Node[T] {
	return n.next
}

func (n *Node[T]) SetNext(node *Node[T]) {
	n.next = node
}
