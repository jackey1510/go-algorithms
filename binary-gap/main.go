package main

import "fmt"

func main() {
	testCases := make(map[int]int)
	testCases[1041] = 5
	testCases[15] = 0
	testCases[32] = 0

	for k, v := range testCases {
		fmt.Println(Solution(k) == v)
	}
}

func Solution(N int) int {

	gap, lastBit, i := 0, -1, 0

	for N > 0 {
		bit := N & 1
		if bit == 1 {
			if lastBit != -1 {
				gap = max(gap, i-lastBit-1)
			}
			lastBit = i
		}

		N = N >> 1
		i++
	}
	return gap
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
