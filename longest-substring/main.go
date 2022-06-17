package main

import "fmt"

func main() {
	// fmt.Println(lengthOfLongestSubstring("abcaad"))
	// fmt.Println(lengthOfLongestSubstring("aaaa"))
	// fmt.Println(lengthOfLongestSubstring("abcadbe"))
	fmt.Println(lengthOfLongestSubstring("tmmzuxt"))

}

// https://leetcode.com/problems/longest-substring-without-repeating-characters/

func lengthOfLongestSubstring(s string) int {
	l := len(s)
	if l < 2 {
		return l
	}

	start, end := 0, 1

	ru := []rune(s)

	result := 1

	m := make(map[rune]int)

	for end <= l {
		char := ru[end-1]
		m[char]++
		if hasRepeated(char, m) {
			start += findFirstOccurrenceIndex(ru[start:end], char, m) + 1
			end++
			continue
		}
		result = max(end-start, result)
		end++
	}

	return result
}

func hasRepeated(s rune, m map[rune]int) bool {
	if m[s] > 1 {
		return true
	}
	return false
}

func findFirstOccurrenceIndex(rs []rune, r rune, m map[rune]int) int {
	for i := range rs {
		m[rs[i]]--
		if rs[i] == r {
			return i
		}
	}
	return -1
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
