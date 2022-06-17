package main

func main() {

}

func sortColors(nums []int) {
	m := make(map[int]int)
	for _, v := range nums {
		m[v]++
	}
	color := 0
	for index := 0; index < len(nums); index++ {
		for m[color] == 0 {
			color++
		}
		nums[index] = color
	}
}
