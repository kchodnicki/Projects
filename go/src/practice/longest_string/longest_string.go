package main

import "fmt"

func solution(pString string) (string, int) {
	currLen, maxLen := 0, 0
	var prevLetter rune
	var maxLetter rune
	for _, currLetter := range pString {
		if prevLetter == currLetter {
			currLen++
			if currLen > maxLen {
				maxLen = currLen
				maxLetter = currLetter
			}
		} else {
			currLen = 1
		}
		prevLetter = currLetter
	}
	return string(maxLetter), maxLen
}

func main() {
	str := "AAABBCDAABCCCCDDEEACDD"
	fmt.Println(solution(str))
}
