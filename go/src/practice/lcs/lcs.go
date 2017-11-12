package main

import "fmt"

func lcsCalc(X string, Y string) {
	lcsTab := make([][]int, len(X)+1)
	for i := range lcsTab {
		lcsTab[i] = make([]int, len(Y)+1)
	}

	for i := 0; i <= len(X); i++ {
		for j := 0; j <= len(Y); j++ {
			if i == 0 || j == 0 {
				lcsTab[i][j] = 0
			} else if X[i-1] == Y[j-1] {
				lcsTab[i][j] = 1 + lcsTab[i-1][j-1]
			} else {
				lcsTab[i][j] = max(lcsTab[i-1][j], lcsTab[i][j-1])
			}
		}
	}
	printLCS(lcsTab, X, Y)
}

func max(x int, y int) int {
	if x < y {
		return y
	}
	return x
}

func printLCS(lcs [][]int, X string, Y string) {
	idx := lcs[len(X)][len(Y)]
	lcsString := make([]rune, idx)
	i := len(X)
	j := len(Y)
	for i > 0 && j > 0 {
		if X[i-1] == Y[j-1] {
			lcsString[idx-1] = rune(X[i-1])
			i--
			j--
			idx--
		} else if lcs[i-1][j] > lcs[i][j-1] {
			i--
		} else {
			j--
		}
	}
	fmt.Println(string(lcsString))
}

func main() {
	s1 := "AGGTAB"
	s2 := "GXTXAYB"
	lcsCalc(s1, s2)
}
