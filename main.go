// Maxwell Lysle
// COSC 3750
// February 2, 2026
package main

import (
	"fmt"
	"hw01/ds"
)

func main() {
	fmt.Println("Only here so the import doesn't leave an error")

	linkedlist := &ds.LinkedList{}
	linkedlist.InsertAt(0, "first")
	linkedlist.Insert("first")
	linkedlist.Insert("first")
	linkedlist.Insert("second")
	linkedlist.Insert("third")
	linkedlist.Insert("fourth")
	linkedlist.Insert("fifth")
	linkedlist.RemoveAt(4)
	linkedlist.PrintList()
	fmt.Println("The size of the linked list is:", linkedlist.GetSize())
	fmt.Println("-------------")
	linkedlist.RemoveAll("first")
	linkedlist.PrintList()
	fmt.Println("-------------")
	linkedlist.Reverse()
	linkedlist.PrintList()
	fmt.Println("The size of the linked list is:", linkedlist.GetSize())
	fmt.Println("-------------")
	linkedlist.InsertAt(100, "first") // shouldn't add anything
	linkedlist.InsertAt(-1, "first")
	linkedlist.InsertAt(3, "fifth")
	linkedlist.PrintList()

	stack := &ds.Stack{}
	stack.Push("first")
	stack.Push("second")
	stack.Push("third")
	data, _ := stack.Pop()
	println("Popped from stack:", data)

	queue := &ds.Queue{}
	queue.Push("first")
	queue.Push("second")
	queue.Push("third")
	data, _ = queue.Pop()
	println("Popped from queue:", data)
}
