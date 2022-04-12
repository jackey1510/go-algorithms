package main

import (
	"fmt"
	"sort"
)

func main() {
	res := merge([][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}})
	fmt.Println(res)
	// res2 := merge([][]int{{1, 4}, {0, 0}})
	// fmt.Println(res2)
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

	sort.Sort(sorted)

	result := [][]int{intervals[0]}

	i := 1

	for i < length {

		curr := sorted[i]

		last := result[len(result)-1]

		// fmt.Println(i, result, curr, last)

		if isOverlap(curr, last) {
			// fmt.Println(curr, last, "overlaps")
			merged := merge2(last, curr)
			// fmt.Println(merged)
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
	// aStart, aEnd, bStart, bEnd := a[0], a[1], b[0], b[1]
	return a[0] <= b[1]

	// // a ends before b
	// // a0 - b0 - a1 - b1
	// if aStart <= bStart && bStart <= aEnd && aEnd < bEnd {
	// 	return true
	// }

	// // b ends before a
	// // b0 - a0 - b1 - a1
	// if bStart <= aStart && aStart <= bEnd && bEnd < aEnd {
	// 	return true
	// }

	// // a overlaps b
	// // a0 - b0 - b1 - a1
	// if aStart <= bStart && bEnd <= aEnd {
	// 	return true
	// }

	// // b overlaps a
	// // b0 - a0 - a1 - b1
	// if bStart <= aStart && aEnd <= bEnd {
	// 	return true
	// }

	// return false

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

func insert(result [][]int, el []int) [][]int {
	if len(result) == 0 {
		return append(result, el)
	}
	// fmt.Println(result, el)
	for i, v := range result {
		// fmt.Println(i, v)
		if el[0] < v[0] {
			left := result[0:i]
			right := append([][]int{}, result[i:]...)
			// fmt.Println(left, right, el, result)
			result = append(left, el)
			result = append(result, right...)
			return result
		}
	}
	return append(result, el)

}
