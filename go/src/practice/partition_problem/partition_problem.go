package main

import (
	"fmt"
)

func canPartition(input []int) bool {
	if len(input) == 0 {
		return true
	}

	sum := 0
	for i := range input {
		sum += input[i]
	}
	if sum%2 != 0 {
		return false
	}

	L := make([][]bool, (sum/2)+1)
	for i := range L {
		L[i] = make([]bool, len(input))
	}
	for i := 0; i <= sum/2; i++ {
		for j := 0; j < len(input); j++ {
			if i == 0 {
				L[i][j] = true
			} else if i == input[j] {
				L[i][j] = true
			} else if j == 0 {
				L[i][j] = false
			} else if input[j] > i {
				L[i][j] = L[i][j-1]
			} else {
				L[i][j] = L[i][j-1] || L[i-input[j]][j-1]
			}
		}
	}
	return L[sum/2][len(input)-1]
}

func main() {
	input := []int{5, 1, 8, 2, 1, 5}
	fmt.Println(canPartition(input))
}
