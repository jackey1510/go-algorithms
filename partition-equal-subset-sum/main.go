package main

import "fmt"

func main() {
	fmt.Println(canPartition([]int{1, 5, 11, 5}))
}

func canPartition(nums []int) bool {

	sum := 0

	for _, v := range nums {
		sum += v
	}

	if sum%2 != 0 {
		return false
	}

	target := sum / 2

	return subsetSum(nums, target, len(nums))
}

func subsetSum(numbers []int, target int, n int) bool {
	dp := make([][]bool, n)

	for i := 0; i < n; i++ {
		if dp[i] == nil {
			dp[i] = make([]bool, target+1)
		}
		dp[i][0] = true
	}

	for i := 1; i < n; i++ {
		for j := 1; j < target+1; j++ {
			if numbers[i] <= j {
				dp[i][j] = dp[i-1][j-numbers[i]] || dp[i-1][j]
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}

	return dp[n-1][target]

}
