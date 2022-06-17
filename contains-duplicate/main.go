package main

func main() {

}

func containsDuplicate(nums []int) bool {
	m := make(map[int]bool)

	for _, v := range nums {
		if m[v] {
			return false
		}
		m[v] = true
	}
	return true
}

func lengthOfLIS(nums []int) int {

	dp := []int{}

	for range nums {
		dp = append(dp, 1)
	}

	max := 1

	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				dp[i] = maxInt(dp[j]+1, dp[i])
				max = maxInt(dp[i], max)
			}
		}

	}

	return max
}

func maxInt(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
