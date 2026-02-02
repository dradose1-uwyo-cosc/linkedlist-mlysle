package ds

import (
	"fmt"
)

type LinkedList struct {
	Head *Node
	Tail *Node
	Size int
}

// insert at the tail
func (l *LinkedList) Insert(value string) {
	new_node := &Node{
		data: value,
		Next: nil,
	}

	if l.IsEmpty() { // in a list with one node, head and tail are the same
		l.Head = new_node
		l.Tail = new_node
	} else {
		l.Tail.Next = new_node
		l.Tail = new_node // update tail
	}
	l.Size++
}

// inserts at a position, returns true if success or false if not, like if pos doesn't exist
// comment: The prototype had return type error but your description suggests you want a bool
func (l *LinkedList) InsertAt(position int, value string) bool {
	if position < 0 || position > l.Size {
		return false
	}

	new_node := &Node{
		data: value,
		Next: nil,
	}

	// special case: insert at head
	if position == 0 {
		new_node.Next = l.Head
		l.Head = new_node
		// update tail if it doesn't exist
		if l.Tail == nil {
			l.Tail = new_node
		}
		l.Size++
		return true
	}

	curr := l.Head
	// traverse until just before the position
	for i := 0; i < position-1; i++ {
		curr = curr.Next
	}
	new_node.Next = curr.Next
	curr.Next = new_node

	// special case: update tail
	if curr.Next == nil {
		l.Tail = new_node
	}

	l.Size++
	return true
}

// remove first occurrence of the value
func (l *LinkedList) Remove(value string) bool {
	// special case: empty list
	if l.IsEmpty() {
		return false
	}

	// special case: remove at head
	if l.Head.data == value {
		l.Head = l.Head.Next
		return true
	}

	prev := l.Head
	curr := l.Head.Next

	for curr != nil {
		if curr.data == value {
			prev.Next = curr.Next
			return true
		}
		prev = curr
		curr = curr.Next
	}
	return false
}

// remove all occurrences of a value
// not working
func (l *LinkedList) RemoveAll(value string) bool {
	if l.Head.data == value {
		l.Head = l.Head.Next
	}

	prev := l.Head
	curr := l.Head.Next

	for curr != nil {
		if curr.data == value {
			prev.Next = curr.Next
		}
		prev = curr
		curr = curr.Next
	}
	return false
}

// remove at a position, if index exists
func (l *LinkedList) RemoveAt(pos int) bool {
	if pos < 0 || pos >= l.Size {
		return false
	}

	// special case: remove at head
	if pos == 0 {
		l.Head = l.Head.Next
		l.Size--
		return true
	}

	curr := l.Head
	for i := 0; i < pos-1; i++ {
		curr = curr.Next
	}

	// special case: remove tail
	if curr.Next == l.Tail {
		l.Tail = curr
	}

	curr.Next = curr.Next.Next
	l.Size--
	return true
}

// checks if the linked list is empty
func (l *LinkedList) IsEmpty() bool {
	return l.Head == nil
}

// get size of ll
func (l *LinkedList) GetSize() int {
	return l.Size
}

// reverse the list
func (l *LinkedList) Reverse() {
	var prev *Node
	curr := l.Head
	var next *Node

	for curr != nil {
		next = curr.Next
		curr.Next = prev

		prev = curr
		curr = next
	}
	l.Head = prev
}

// print the list
func (l *LinkedList) PrintList() {
	curr := l.Head
	for curr != nil {
		fmt.Println(curr.data)
		curr = curr.Next
	}
}
