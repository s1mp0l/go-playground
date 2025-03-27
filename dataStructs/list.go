package dataStructs

import "fmt"

type Node struct {
	Value interface{}
	Next  *Node
}

type LinkedList struct {
	Head *Node
	Tail *Node
	Size int
}

func NewLinkedList() *LinkedList {
	return &LinkedList{
		Head: nil,
		Tail: nil,
		Size: 0,
	}
}

func (ll *LinkedList) IsEmpty() bool {
	return ll.Size == 0
}

func (ll *LinkedList) AddFirst(value interface{}) {
	newNode := &Node{
		Value: value,
		Next:  ll.Head,
	}

	ll.Head = newNode

	if ll.Size == 0 {
		ll.Tail = newNode
	}

	ll.Size++
}

func (ll *LinkedList) AddLast(value interface{}) {
	newNode := &Node{
		Value: value,
		Next:  nil,
	}

	if ll.IsEmpty() {
		ll.Head = newNode
		ll.Tail = newNode
	} else {
		ll.Tail.Next = newNode
		ll.Tail = newNode
	}

	ll.Size++
}

func (ll *LinkedList) RemoveFirst() (interface{}, bool) {
	if ll.IsEmpty() {
		return nil, false
	}

	value := ll.Head.Value
	ll.Head = ll.Head.Next
	ll.Size--

	if ll.Size == 0 {
		ll.Tail = nil
	}

	return value, true
}

func (ll *LinkedList) Contains(value interface{}) bool {
	current := ll.Head

	for current != nil {
		if current.Value == value {
			return true
		}
		current = current.Next
	}

	return false
}

func (ll *LinkedList) GetSize() int {
	return ll.Size
}

func (ll *LinkedList) String() string {
	if ll.IsEmpty() {
		return "[]"
	}

	result := "["
	current := ll.Head

	for current != nil {
		result += fmt.Sprintf("%v", current.Value)
		if current.Next != nil {
			result += ", "
		}
		current = current.Next
	}

	result += "]"
	return result
}

func (ll *LinkedList) GetAt(index int) (interface{}, bool) {
	if index < 0 || index >= ll.Size {
		return nil, false
	}

	current := ll.Head
	for i := 0; i < index; i++ {
		current = current.Next
	}

	return current.Value, true
}

func (ll *LinkedList) RemoveAt(index int) (interface{}, bool) {
	if index < 0 || index >= ll.Size {
		return nil, false
	}

	if index == 0 {
		return ll.RemoveFirst()
	}

	prev := ll.Head
	for i := 0; i < index-1; i++ {
		prev = prev.Next
	}

	value := prev.Next.Value

	if index == ll.Size-1 {
		ll.Tail = prev
	}

	prev.Next = prev.Next.Next
	ll.Size--

	return value, true
}
