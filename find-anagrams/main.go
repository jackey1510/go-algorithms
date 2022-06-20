package main

import "fmt"

func main() {
	fmt.Println(findAnagrams("abababc", "abc"))
}

func findAnagrams(s string, p string) []int {
	size := len(p)
	result := []int{}
	m := make(map[byte]int)
	pMap := make(map[byte]int)

	for i := range p {
		pMap[p[i]]++
	}

	for i := 0; i < size-1 && i < len(s); i++ {
		m[s[i]]++
	}

	for i := size - 1; i < len(s); i++ {
		m[s[i]]++
		isAnagram := true
		for j := range pMap {
			if m[j] != pMap[j] {
				isAnagram = false
				break
			}
		}

		if isAnagram {
			result = append(result, i-size+1)
		}

		m[s[i-size+1]]--
	}

	return result
}
