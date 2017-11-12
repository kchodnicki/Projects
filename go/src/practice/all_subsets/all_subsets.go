package main

import (
	"bytes"
	"fmt"
)

type set struct {
	nums []int
}

func (s set) String() string {
	var b bytes.Buffer
	for i := range s.nums {
		if i == 0 {
			b.WriteString(fmt.Sprintf("(%d", s.nums[i]))
		} else {
			b.WriteString(fmt.Sprintf(", %d", s.nums[i]))
		}
		if i == len(s.nums)-1 {
			b.WriteString(")")
		}
	}
	return b.String()
}

func main() {
	input := []int{1, 2, 3, 4, 5}
	L := make([][][]*set, len(input))
	for i := range L {
		L[i] = make([][]*set, len(input))
		for j := range L[i] {
			L[i][j] = []*set{}
		}
	}

	//initialize sets on diagonal
	for i, j := 0, 0; i < len(input); i, j = i+1, j+1 {
		newSet := &set{nums: []int{input[i]}}
		L[i][j] = []*set{newSet}
	}

	for i := 0; i < len(input); i++ {
		for j := i + 1; j < len(input); j++ {
			for _, prevSet := range L[i][j-1] {
				L[i][j] = append(L[i][j], prevSet)
				prevNums := append([]int(nil), prevSet.nums...)
				newSet := &set{nums: append(prevNums, j+1)}
				L[i][j] = append(L[i][j], newSet)
			}
		}
		fmt.Println(L[i][len(input)-1])
	}
	fmt.Println("[()]")
}
