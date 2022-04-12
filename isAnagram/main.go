package main

func main() {
	isAnagram("abc", "bca")
}

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	ms := make(map[string]int)
	mt := make(map[string]int)
	for i := range s {
		ms[s[i:i+1]]++
		mt[t[i:i+1]]++
	}

	for i := range ms {
		if ms[i] != mt[i] {
			return false
		}
	}
	return true
}
