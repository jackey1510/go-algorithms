package main

func main() {

}

type arr []int

func lengthOfLIS(nums []int) int {
	dp := []int{}
	for range nums {
		dp = append(dp, 1)
	}
	max := 1

	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = maxNumber(dp[i], dp[j]+1)
				max = maxNumber(dp[i], max)
			}
		}
	}
	return max

}

func (a arr) last() int {
	return a[len(a)-1]
}

func maxNumber(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func minNumber(a int, b int) int {
	if a < b {
		return a
	}
	return b
}
