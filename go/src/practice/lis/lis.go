package main

import "fmt"

func lis(arr []int) []int {
	lis := []int{}
	for range arr {
		lis = append(lis, 1)
	}

	for i := range arr {
		for j := 0; j < i; j++ {
			if arr[i] > arr[j] && lis[i] < lis[j]+1 {
				lis[i] = lis[j] + 1
			}
		}
	}
	return lis
}

func main() {
	arr := []int{10, 22, 9, 33, 21, 50, 41, 60, 80}
	fmt.Println(lis(arr))
}
