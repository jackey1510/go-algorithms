package main

func main() {

}

//https://leetcode.com/problems/ransom-note/

func canConstruct(ransomNote string, magazine string) bool {
	runeMap := getRuneMap(magazine)

	for _, r := range ransomNote {
		if runeMap[r] == 0 {
			return false
		}
		runeMap[r]--

	}
	return true
}

func getRuneMap(s string) map[rune]int {
	m := make(map[rune]int)
	for _, v := range []rune(s) {
		m[v]++
	}
	return m
}
