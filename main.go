package main

import "fmt"

func main() {
	s := "ccc"
	fmt.Println(longestPalindrome(s))
}

func longestPalindrome(s string) int {
	uneven := 0
	lenth := len(s)
	paliandrome := 0

	if lenth <= 1 {
		return lenth
	}

	var table [60]int

	for _, v := range s {
		table[v-65] += 1
	}

	for i := range table {
		paliandrome += table[i] - table[i]%2
	}

	if paliandrome < lenth {
		uneven = 1
	}

	return paliandrome + uneven
}
