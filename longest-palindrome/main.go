package main

import (
	"fmt"
)

func main() {
	fmt.Println(isPalindrome([]rune("aaaa")))
	fmt.Println(longestPalindrome("baaaggvdsbabbbb2faf"))
}

func longestPalindrome(s string) string {
	n, start, end := len(s), 0, 0
	dp := make([][]bool, n)
	for i := range dp {
		dp[i] = make([]bool, n)
	}

	for len := 0; len < n; len++ {
		for i := 0; i+len < n; i++ {
			dp[i][i+len] = s[i] == s[i+len] && (len < 2 || dp[i+1][i+len-1])
			if dp[i][i+len] && len > end-start {
				start = i
				end = i + len
			}
		}
	}
	return s[start : end+1]
}

func isPalindrome(runes []rune) bool {
	start := 0
	end := len(runes) - 1

	for start < end {
		if runes[start] != runes[end] {
			return false
		}
		start++
		end--
	}
	return true
}
