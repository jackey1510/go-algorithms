package main

import (
	"fmt"
	"sort"
)

func main() {
	res := merge([][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}})
	fmt.Println(res)
	res3 := isOverlap([]int{8, 10}, []int{1, 6})
	fmt.Println(res3)
}

type SortBy [][]int

func (a SortBy) Len() int           { return len(a) }
func (a SortBy) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a SortBy) Less(i, j int) bool { return a[i][0] < a[j][0] }

func merge(intervals [][]int) [][]int {

	length := len(intervals)

	sorted := SortBy(intervals)

	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i][0] < sorted[j][0]
	})

	result := [][]int{intervals[0]}

	i := 1

	for i < length {
		curr := sorted[i]
		last := result[len(result)-1]
		if isOverlap(curr, last) {
			merged := merge2(last, curr)
			result[len(result)-1] = merged
			i++
			continue
		}
		result = append(result, curr)
		i++
	}

	return result
}

func isOverlap(a []int, b []int) bool {
	return a[0] <= b[1]
}

func merge2(a []int, b []int) []int {
	return []int{min(a[0], b[0]), max(a[1], b[1])}
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}
