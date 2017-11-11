package detectCycle

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

func isCycle(n *node) bool {
	visited, queue := []*node{n}, []*node{n}
	var currNode *node
	for {
		currNode = queue[len(queue)-1]
		queue = queue[:len(queue)-1]
		for _, childNode := range currNode.childNodes {
			if isVisited(childNode, visited) {
				return true
			}
			visited = append(visited, childNode)
			queue = append(queue, childNode)
		}
		if len(queue) == 0 {
			break
		}
	}
	return false
}

func isVisited(n *node, queue []*node) bool {
	for _, qNode := range queue {
		if n.value == qNode.value {
			return true
		}
	}
	return false
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
	fmt.Println(isCycle(n0))
}
