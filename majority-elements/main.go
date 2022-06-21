package main

import "fmt"

//https://leetcode.com/problems/majority-element/
func main() {
	fmt.Println(majorityElement([]int{2, 2, 1, 1, 1, 2, 2, 1, 1, 2, 2}))
}

func majorityElement(nums []int) int {
	m := make(map[int]int)

	half := len(nums) / 2

	for _, v := range nums {
		m[v]++
		if m[v] > half {
			return v
		}
	}

	return -1
}
