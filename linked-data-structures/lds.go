package main

import (
	"fmt"
)

type Cell struct {
	data string
	next *Cell
}

func (c *Cell) addAfter(after *Cell) {
	after.next = c.next
	c.next = after
}

func (c *Cell) deleteAfter() *Cell {
	if c.next == nil {
		panic("No cell to delete.")
	}
	tmp := c.next
	c.next = tmp.next
	return tmp
}

type LinkedList struct {
	sentinel *Cell
}

func makeLinkedList() LinkedList {
	list := LinkedList{}
	list.sentinel = &Cell{"SENTINEL", nil}
	return list
}

func (list *LinkedList) addRange(values []string) {
	// move to last cell in list
	lastCell := list.sentinel
	for lastCell.next != nil {
		lastCell = lastCell.next
	}

	for _, s := range values {
		c := Cell{s, nil}
		lastCell.addAfter(&c)
		lastCell = &c
	}
}

func (list *LinkedList) toString(separator string) string {
	str := ""
	for c := list.sentinel.next; c != nil; c = c.next {
		str += c.data
		if c.next != nil {
			str += separator
		}
	}
	return str
}

func (list *LinkedList) length() int {
	llength := 0
	top := list.sentinel
	for top.next != nil {
		llength++
		top = top.next
	}
	return llength
}

func (list *LinkedList) isEmpty() bool {
	if list.sentinel.next != nil {
		return false
	} else {
		return true
	}
}

func (list *LinkedList) push(s string) {
	c := Cell{s, nil}
	list.sentinel.addAfter(&c)
}

func (list *LinkedList) pop() string {
	tmp := list.sentinel.deleteAfter()
	return tmp.data
}

func main() {
	//test()

	// smallListTest()

	// Make a list from an array of values.
	greekLetters := []string{
		"α", "β", "γ", "δ", "ε",
	}
	list := makeLinkedList()
	list.addRange(greekLetters)
	fmt.Println(list.toString(" "))
	fmt.Println()

	// Demonstrate a stack.
	stack := makeLinkedList()
	stack.push("Apple")
	stack.push("Banana")
	stack.push("Coconut")
	stack.push("Date")
	for !stack.isEmpty() {
		fmt.Printf("Popped: %-7s   Remaining %d: %s\n",
			stack.pop(),
			stack.length(),
			stack.toString(" "))
	}

}

func test() {
	ll := makeLinkedList()
	fmt.Println(ll.isEmpty())
	ll.addRange([]string{"qqq", "www", "zzz"})
	fmt.Println(ll.isEmpty())
	//printList(&ll)
	//fmt.Println(ll.toString("."))
	fmt.Println(ll.length())

	ll.push("test1")
	fmt.Println(ll.toString(" "))

	ll.push("test2")
	fmt.Println(ll.toString(" "))

	fmt.Println(ll.pop())
	fmt.Println(ll.toString(" "))
	fmt.Println(ll.pop())
	fmt.Println(ll.toString(" "))
	fmt.Println(ll.pop())
	fmt.Println(ll.toString(" "))

}
