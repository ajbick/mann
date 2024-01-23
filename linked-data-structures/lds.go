package main

import (
	"fmt"
)

// Cell -------------------------------------------
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

// LinkedList --------------------------------------
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

func (list *LinkedList) toStringMax(separator string, max int) string {
	str := ""
	cnt := 1
	for c := list.sentinel.next; c != nil; c = c.next {
		str += c.data
		cnt++
		if cnt > max {
			return str
		}
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

func (list *LinkedList) hasLoop() bool {
	tort := list.sentinel
	hare := list.sentinel
	for {
		for i := 0; i < 2; i++ {
			if hare.next == nil {
				return false
			} else {
				hare = hare.next
			}
			if hare == tort {
				return true
			}
		}
		if tort.next == nil {
			return false
		}
		tort = tort.next
	}
	return true
}

// -----------------------------------------------

func test() {
	ll := makeLinkedList()
	//fmt.Println(ll.isEmpty())
	ll.addRange([]string{"aaa", "bbb", "ccc", "ddd", "eee"})

	//fmt.Println(ll.isEmpty())

	fmt.Println(ll.toStringMax(" ", 3))
	fmt.Println(ll.hasLoop())

	//fmt.Println(ll.length())

	//ll.push("test1")
	//fmt.Println(ll.toString(" "))

	//ll.push("test2")
	//fmt.Println(ll.toString(" "))

	//fmt.Println(ll.pop())
	//fmt.Println(ll.toString(" "))
	//fmt.Println(ll.pop())
	//fmt.Println(ll.toString(" "))
	//fmt.Println(ll.pop())
	//fmt.Println(ll.toString(" "))

	ll.sentinel.next.next.next.next.next = ll.sentinel.next.next
	fmt.Println(ll.hasLoop())

}

func main() {
	// Make a list from an array of values.
	values := []string{
		"0", "1", "2", "3", "4", "5",
	}
	list := makeLinkedList()
	list.addRange(values)

	fmt.Println(list.toString(" "))
	if list.hasLoop() {
		fmt.Println("Has loop")
	} else {
		fmt.Println("No loop")
	}
	fmt.Println()

	// Make cell 5 point to cell 2.
	list.sentinel.next.next.next.next.next.next = list.sentinel.next.next

	fmt.Println(list.toStringMax(" ", 10))
	if list.hasLoop() {
		fmt.Println("Has loop")
	} else {
		fmt.Println("No loop")
	}
	fmt.Println()

	// Make cell 4 point to cell 2.
	list.sentinel.next.next.next.next.next = list.sentinel.next.next

	fmt.Println(list.toStringMax(" ", 10))
	if list.hasLoop() {
		fmt.Println("Has loop")
	} else {
		fmt.Println("No loop")
	}
}

/*
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
*/
