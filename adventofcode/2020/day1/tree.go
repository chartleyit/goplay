package main

import (
	"fmt"
	"strconv"
)

// TODO could this be handled better by using a heap?
type Node struct {
	Key   int
	Left  *Node
	Right *Node
}

// This is DUMB
// Insert data
func (n *Node) Insert(k int) {
	if n.Key < k {
		if n.Right == nil {
			n.Right = &Node{Key: k}
		} else {
			n.Right.Insert(k)
		}
	} else if n.Key > k {
		if n.Left == nil {
			n.Left = &Node{Key: k}
		} else {
			n.Left.Insert(k)
		}
	}
}

// Search data
func (n *Node) Search(k int) {

}

// Tree
func Tree(input *[]string) {
	node := &Node{}
	for _, x := range *input {
		s, err := strconv.Atoi(x)
		if err != nil {
			fmt.Println("Failed to convert string to int", s)
			fmt.Println(err)
		}
		node.Insert(s)
		// node.Insert(strconv.Atoi(x))
	}

	fmt.Println(node)
}
