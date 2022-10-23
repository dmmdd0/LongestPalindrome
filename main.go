package main

import "fmt"

func main() {
	s := "abcccccdd"
	fmt.Println(longestPalindrome(s))
}

func longestPalindrome(s string) int {
	uneven := 0
	lenth := len(s)

	if lenth < 2 {
		return lenth
	}

	if lenth%2 != 0 {
		uneven = 1
	}

	var table [60]int

	for _, v := range s {
		table[v-65] += 1
	}
	return 3 + uneven
}
