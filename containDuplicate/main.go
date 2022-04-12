package main

import "fmt"

func main() {

}

// ContainsDuplicate checks whether there is duplicate in array
func ContainsDuplicate(nums []int) bool {

	numMap := make(map[int]int)

	for _, v := range nums {
		if numMap[v] > 0 {
			fmt.Println(numMap)
			return true
		}
		numMap[v]++
	}
	return false
}
