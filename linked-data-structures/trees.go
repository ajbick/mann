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
	str := ""

	fmt.Println(strings.Repeat(indent, depth) + node.data)
	
	str = str + node.data
	
	if node.left != nil {
		depth++
		str += node.left.displayIndented(indent, depth)
		depth--
	}
	if node.right != nil {
		depth++
		str += node.right.displayIndented(indent, depth)
		depth--
	}
	
	return str
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

	aNode.left  = &bNode
	aNode.right = &cNode

	bNode.left  = &dNode
	bNode.right = &eNode

	cNode.right = &fNode

	eNode.left = &gNode

	fNode.left = &hNode

	hNode.left = &iNode
	hNode.right = &jNode

	return &aNode
}

	
func main() {
	x := buildTree()
	fmt.Println(x.displayIndented(" ", 0))
}
