package generics

import "fmt"

// type LinkedList[T any] struct { // without func Contains, T can be any
type LinkedList[T comparable] struct {
	value T
	next  *LinkedList[T]
}

func (ll *LinkedList[T]) Len() int {
	count := 0
	for node := ll; node != nil; node = node.next {
		count++
	}
	return count
}

func (ll *LinkedList[T]) InsertAt(pos int, value T) *LinkedList[T] {
	if ll == nil || pos <= 0 {
		return &LinkedList[T]{
			value: value,
			next:  ll,
		}
	}
	ll.next = ll.next.InsertAt(pos-1, value)
	return ll
}

func (ll *LinkedList[T]) Append(value T) *LinkedList[T] {
	return ll.InsertAt(ll.Len(), value)
}

func (ll *LinkedList[T]) String() string {
	if ll == nil {
		return "nil"
	}
	return fmt.Sprintf("%v->%v", ll.value, ll.next.String())
}

func (ll *LinkedList[T]) Contains(value T) bool {
	for node := ll; node != nil; node = node.next {
		if node.value == value { // due to this line, comparator is required
			return true
		}
	}
	return false
}
