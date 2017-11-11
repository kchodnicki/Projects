package main

import "fmt"

type listNode struct {
	value    int
	nextNode *listNode
}

func (node listNode) String() string {
	if node.nextNode != nil {
		return fmt.Sprintf("%d -> %d ", node.value, node.nextNode.value)
	}
	return fmt.Sprintf("%d -> nil", node.value)
}

func reverseList(root *listNode) {
	var currNode, prevNode, tmpNode *listNode
	currNode = root
	prevNode = nil
	for {
		if currNode.nextNode == nil {
			currNode.nextNode = prevNode
			break
		}
		tmpNode = currNode.nextNode
		currNode.nextNode = prevNode
		prevNode = currNode
		currNode = tmpNode
	}
}

func main() {
	n3 := &listNode{value: 3}
	n2 := &listNode{value: 2, nextNode: n3}
	n1 := &listNode{value: 1, nextNode: n2}

	fmt.Println(n1, n2, n3)
	reverseList(n1)
	fmt.Println(n3, n2, n1)
}
