package main

import "fmt"

func main() {
	res := maxSubArray([]int{1, 2, 3, 4, -2, 5, -5, 9})
	fmt.Println(res)
}

func maxSubArray(nums []int) int {

	length := len(nums)

	maxSum, currSum := nums[0], nums[0]

	for i := 1; i < length; i++ {
		if currSum < 0 {
			currSum = 0
		}
		currSum += nums[i]
		maxSum = max(currSum, maxSum)
	}

	return maxSum

}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
