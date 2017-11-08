package main

import (
	"bytes"
	"fmt"
	"strconv"
)

type node struct {
	value      int
	childNodes []*node
}

func (n node) String() string {
	var b bytes.Buffer
	b.WriteString(strconv.Itoa(n.value))
	b.WriteString(" -> (")
	for _, child := range n.childNodes {
		b.WriteString(" " + strconv.Itoa(child.value))
	}
	b.WriteString(" )")
	return b.String()
}

func bfs(root *node) string {
	var currNode *node
	visitedNodes := []*node{root}
	queue := []*node{root}
	for len(queue) > 0 {
		//take last element in queue
		currNode, queue = popRight(queue)
		fmt.Println("currNode: ", strconv.Itoa(currNode.value))
		for _, childNode := range currNode.childNodes {
			if notVisited(childNode, visitedNodes) {
				fmt.Println("childNode: ", strconv.Itoa(childNode.value))
				visitedNodes = append(visitedNodes, childNode)
				queue = append(queue, childNode)
			}
		}
	}
	return ""
}

func notVisited(n *node, queue []*node) bool {
	for _, node := range queue {
		if node.value == n.value {
			return false
		}
	}
	return true
}

func popRight(queue []*node) (*node, []*node) {
	node := queue[len(queue)-1]
	queue = queue[:len(queue)-1]
	return node, queue
}

func main() {
	//Create graph structure
	n0 := &node{value: 0}
	n1 := &node{value: 1}
	n2 := &node{value: 2}
	n3 := &node{value: 3}
	n0.childNodes = []*node{n1, n2}
	n1.childNodes = []*node{n2}
	n2.childNodes = []*node{n0, n3}
	n3.childNodes = []*node{n3}
	fmt.Println(n0)
	fmt.Println(n1)
	fmt.Println(n2)
	fmt.Println(n3)

	fmt.Println(bfs(n2))
}
