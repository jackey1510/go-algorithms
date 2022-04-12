package main

import (
	"fmt"
	"sort"
)

func main() {
	res := groupAnagrams([]string{"abc", "bac", "aaa", "bbb", "bbb", "abcwd", "cwdab"})
	fmt.Println(res)
}

func groupAnagrams(strs []string) [][]string {
	res := [][]string{}

	ms := make(map[string][]string)

	for _, v := range strs {
		sorted := sortString(v)
		ms[sorted] = append(ms[sorted], v)

	}

	for _, v := range ms {
		res = append(res, v)
	}

	return res

}

func sortString(s string) string {
	sortBy := SortBy(s)
	sort.Sort(sortBy)
	return string(sortBy)

}

type SortBy []rune

func (a SortBy) Len() int           { return len(a) }
func (a SortBy) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a SortBy) Less(i, j int) bool { return a[i] < a[j] }
