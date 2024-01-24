package main

import (
	"fmt"
)

// Cell
type Cell struct {
	data string
	next *Cell
	prev *Cell
}

func (c *Cell) addAfter(after *Cell) {
	after.next = c.next
	after.prev = c
	c.next.prev = after
	c.next = after
}

func (c *Cell) addBefore(before *Cell) {
	tmp := c.prev
	tmp.addAfter(before)
}

func (c *Cell) delete() *Cell {
	c.prev.next = c.next
	c.next.prev = c.prev
	return c
}



// LL
type DoublyLinkedList struct {
	topSentinel *Cell
	bottomSentinel *Cell
}

func makeDoublyLinkedList() DoublyLinkedList {
	list := DoublyLinkedList{}
	list.topSentinel = &Cell{"TOPSENTINEL", nil, nil}
	list.bottomSentinel = &Cell{"BOTTOMSENTINEL", nil, nil}
	list.topSentinel.next = list.bottomSentinel
	list.topSentinel.prev = nil
	list.bottomSentinel.prev = list.topSentinel
	list.bottomSentinel.next = nil
	return list
}

func (list *DoublyLinkedList) addRange(values []string) {
	for _, s := range values {
		c := Cell{s, nil, nil}
		list.bottomSentinel.addBefore(&c)
	}
}

func (list *DoublyLinkedList) toString(separator string) string {
	str := ""
	for c := list.topSentinel.next; c != list.bottomSentinel; c = c.next {
		str += c.data
		if c != list.bottomSentinel {
			str += separator
		}
	}
	return str
}

func (list *DoublyLinkedList) toStringMax(separator string, max int) string {
	str := ""
	c := list.topSentinel.next	
	for i := 0; i < max; i++ {
		str += c.data
		if c != list.bottomSentinel {
			str += separator
		}
		c = c.next
	}
	return str
}

func (list *DoublyLinkedList) length() int {
	cnt := 0
	for c := list.topSentinel; c != list.bottomSentinel; c = c.next {
		cnt++
	}
	return cnt
}

func (list *DoublyLinkedList) isEmpty() bool {
	return list.topSentinel == list.bottomSentinel
}

func (stack *DoublyLinkedList) push(value string) {
	c := Cell{value, nil, nil}
	stack.topSentinel.addAfter(&c)
}

func (stack *DoublyLinkedList) pop() string {
	if stack.isEmpty() {
		panic("Stack is empty.")
	}
	c := stack.topSentinel.next.delete()
	return c.data
}

func main() {
	
}
