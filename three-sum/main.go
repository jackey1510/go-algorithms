package main

import "sort"

func main() {

}

// https://leetcode.com/problems/3sum/

type SortBy []int

func (a SortBy) Len() int           { return len(a) }
func (a SortBy) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a SortBy) Less(i, j int) bool { return a[i] < a[j] }
func threeSum(nums []int) [][]int {
	sorted := SortBy(nums)
	sort.Sort(sorted)

	result := [][]int{}

	for i, v := range sorted {
		if i > 0 && v == sorted[i-1] {
			continue
		}

	}
}

func twoSum(nums []int, v int) [][]int {
	
}
