package main

import (
	"fmt"
)

type node struct {
	value  int
	rank   int
	parent *node
}

func (node node) String() string {
	if node.parent != nil {
		return fmt.Sprintf("%d -> %d ", node.value, node.parent.value)
	}
	return fmt.Sprintf("%d -> nil", node.value)
}

func unionSets(n1 *node, n2 *node) {
	n1Parent := findSet(n1)
	n2Parent := findSet(n2)
	fmt.Println("BEFORE")
	fmt.Println(fmt.Sprintf("First node: %d (%d) Second node: %d (%d)", n1.value, n1.rank, n2.value, n2.rank))
	fmt.Println(fmt.Sprintf("First parent: %d (%d), Second parent: %d (%d)", n1Parent.value, n1Parent.rank, n2Parent.value, n2Parent.rank))
	if n1Parent.value == n2Parent.value {
		return
	}
	if n1Parent.rank > n2Parent.rank {
		n2Parent.parent = n1Parent
	} else if n1Parent.rank < n2Parent.rank {
		n1Parent.parent = n2Parent
	} else {
		n2Parent.parent = n1Parent
		n1Parent.rank++
	}
	fmt.Println("AFTER")
	fmt.Println(fmt.Sprintf("First node: %d (%d) Second node: %d (%d)", n1.value, n1.rank, n2.value, n2.rank))
	fmt.Println(fmt.Sprintf("First parent: %d (%d), Second parent: %d (%d)\n", n1.parent.value, n1.parent.rank, n2.parent.value, n2.parent.rank))
}

func findSet(n *node) *node {
	if n.parent == n {
		return n
	}
	n.parent = findSet(n.parent)
	return n.parent
}

func makeSet(v int) *node {
	n := &node{value: v, rank: 0}
	n.parent = n
	return n
}

func main() {
	n1 := makeSet(1)
	n2 := makeSet(2)
	n3 := makeSet(3)
	n4 := makeSet(4)
	n5 := makeSet(5)
	n6 := makeSet(6)
	n7 := makeSet(7)

	unionSets(n1, n2)
	unionSets(n2, n3)
	unionSets(n4, n5)
	unionSets(n6, n7)
	unionSets(n5, n6)
	unionSets(n3, n7)

	fmt.Println(n1, n2, n3, n4, n5, n6, n7)
}
