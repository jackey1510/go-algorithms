package main

import "fmt"

func main() {
	// res1 := characterReplacement("ABAB", 2)
	// res2 := characterReplacement("AABABBA", 1)
	// res3 := characterReplacement("AABABBA", 2)
	// res4 := characterReplacement("AABABBA", 3)
	// res5 := characterReplacement("AAAA", 2)
	// res6 := characterReplacement("BAAA", 1)
	res7 := characterReplacement("PNDQPERNFSSSRDEQLFOF", 4)
	// fmt.Println(res1, res2, res3, res4, res5, res6)
	fmt.Println(res7)
}

func characterReplacement(s string, k int) int {

	runes := []rune(s)
	maxLength, currentLength := 1, 1

	m := make(map[rune]int)

	maxChar := 1
	m[runes[0]]++

	left := 0

	for right := 1; right < len(s); right++ {
		currentLength = right - left + 1
		char := runes[right]
		m[char]++
		maxChar = max(maxChar, m[char])
		fmt.Println(right, left, char, maxChar, m)
		if currentLength-maxChar <= k {
			maxLength = max(currentLength, maxLength)
			continue
		}
		m[runes[left]]--
		left++
		maxChar = 1
	}

	return maxLength
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
