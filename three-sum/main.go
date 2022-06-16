package main

import (
	"fmt"
	"sort"
)

// https://leetcode.com/problems/3sum/

func main() {
	// fmt.Println(twoSum([]int{0, 1, 3, 4, 5}, 4))
	// fmt.Println(twoSum([]int{0, 1, 3, 4, 5}, 3))
	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4}))
	fmt.Println(threeSum([]int{0}))
}

type SortBy []int

func (a SortBy) Len() int           { return len(a) }
func (a SortBy) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a SortBy) Less(i, j int) bool { return a[i] < a[j] }

func threeSum(nums []int) [][]int {
	sorted := SortBy(nums)
	sort.Sort(sorted)

	res := [][]int{}

	for i := range nums {
		fmt.Println(i, res)

		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		twoSum := twoSum(nums[i+1:], nums[i]*-1)
		for _, v := range twoSum {
			v = append(v, nums[i])
			res = append(res, v)
		}
	}

	return res

}

func twoSum(nums []int, target int) [][]int {

	start, end := 0, len(nums)-1

	res := [][]int{}

	for start < end {

		if start > 0 && nums[start] == nums[start-1] {
			start++
			continue
		}

		sum := nums[start] + nums[end]

		fmt.Println(start, end, res)

		if sum == target {
			res = append(res, []int{nums[start], nums[end]})
			start++
			continue
		}

		if sum > target {
			end--
			continue
		}

		start++
	}

	return res
}
