package main

import "fmt"

type t *int

func addElem(input []t, subset []t, i int) {
	if i == len(input) {
		printSubset(subset)
	} else {
		subset[i] = nil
		addElem(input, subset, i+1)
		subset[i] = input[i]
		addElem(input, subset, i+1)
	}
}

func printSubset(subset []t) {
	for i := range subset {
		if subset[i] != nil {
			fmt.Print(fmt.Sprintf("%d.", *subset[i]))
		}
	}
	fmt.Println("")
}

func main() {
	input := []t{}
	for i := 0; i < 5; i++ {
		val := i + 1
		input = append(input, &val)
	}
	addElem(input, make([]t, len(input)), 0)
	fmt.Println("_")
}
