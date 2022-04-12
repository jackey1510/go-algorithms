package main

func main() {

}

func productExceptSelf(nums []int) []int {

	prefix, suffix, length := 1, 1, len(nums)
	result := []int{}

	for i := 0; i < length; i++ {
		result = append(result, 1)
	}

	for i := range nums {
		result[i] *= prefix
		prefix *= nums[i]
		result[length-i-1] *= suffix
		suffix *= nums[length-i-1]
	}
	return result
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}
