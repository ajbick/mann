package main

import (
	"fmt"
	"strings"
)

type Node struct {
	data  string
	left  *Node
	right *Node
}

func (node *Node) displayIndented(indent string, depth int) string {
	result := ""

	result = result + strings.Repeat(indent, depth) + node.data + "\n"

	if node.left != nil {
		depth++
		result += node.left.displayIndented(indent, depth)
		depth--
	}
	if node.right != nil {
		depth++
		result += node.right.displayIndented(indent, depth)
		depth--
	}

	return result
}

func (node *Node) preorder() string {
	result := ""

	result = node.data + " "

	if node.left != nil {
		result += node.left.preorder()
	}

	if node.right != nil {
		result += node.right.preorder()
	}

	return result
}

func (node *Node) inorder() string {
	result := ""

	if node.left != nil {
		result += node.left.inorder()
	}

	result += node.data + " "

	if node.right != nil {
		result += node.right.inorder()
	}

	return result
}

func (node *Node) postorder() string {
	result := ""

	if node.left != nil {
		result += node.left.postorder()
	}

	if node.right != nil {
		result += node.right.postorder()
	}

	result += node.data + " "

	return result
}

func (node *Node) breadthFirst() string {
	ll := makeDoublyLinkedList()
	ll.enqueue(node)

	result := ""
	for !ll.isEmpty() {
		n := ll.dequeue()
		result += n.data

		if n.left != nil {
			ll.enqueue(n.left)
		}
		if n.right != nil {
			ll.enqueue(n.right)
		}

		if !ll.isEmpty() {
			result += " "
		}
	}

	return result
}

func buildTree() *Node {
	aNode := Node{"A", nil, nil}
	bNode := Node{"B", nil, nil}
	cNode := Node{"C", nil, nil}
	dNode := Node{"D", nil, nil}
	eNode := Node{"E", nil, nil}
	fNode := Node{"F", nil, nil}
	gNode := Node{"G", nil, nil}
	hNode := Node{"H", nil, nil}
	iNode := Node{"I", nil, nil}
	jNode := Node{"J", nil, nil}

	aNode.left = &bNode
	aNode.right = &cNode

	bNode.left = &dNode
	bNode.right = &eNode

	cNode.right = &fNode

	eNode.left = &gNode

	fNode.left = &hNode

	hNode.left = &iNode
	hNode.right = &jNode

	return &aNode
}

func main() {
	// Build a tree.
	aNode := buildTree()

	// Display with indentation.
	fmt.Println(aNode.displayIndented("  ", 0))

	// Display traversals.
	fmt.Println("Preorder:     ", aNode.preorder())
	fmt.Println("Inorder:      ", aNode.inorder())
	fmt.Println("Postorder:    ", aNode.postorder())
	fmt.Println("Breadth first:", aNode.breadthFirst())
}

// Cell
type Cell struct {
	data *Node
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
	topSentinel    *Cell
	bottomSentinel *Cell
}

func makeDoublyLinkedList() DoublyLinkedList {
	list := DoublyLinkedList{}
	list.topSentinel = &Cell{nil, nil, nil}
	list.bottomSentinel = &Cell{nil, nil, nil}
	list.topSentinel.next = list.bottomSentinel
	list.topSentinel.prev = nil
	list.bottomSentinel.prev = list.topSentinel
	list.bottomSentinel.next = nil
	return list
}

func (list *DoublyLinkedList) isEmpty() bool {
	return list.topSentinel.next == list.bottomSentinel
}

func (stack *DoublyLinkedList) push(value *Node) {
	c := Cell{value, nil, nil}
	stack.topSentinel.addAfter(&c)
}

func (stack *DoublyLinkedList) pop() *Node {
	if stack.isEmpty() {
		panic("Stack is empty.")
	}
	c := stack.topSentinel.next.delete()
	return c.data
}

// queue
func (queue *DoublyLinkedList) enqueue(value *Node) {
	queue.push(value)
}

func (queue *DoublyLinkedList) dequeue() *Node {
	if queue.isEmpty() {
		panic("Stack is empty.")
	}
	c := queue.bottomSentinel.prev.delete()
	return c.data
}
