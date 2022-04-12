package main

import "fmt"

func main() {
	for i := 1; i < 15; i++ {
		fmt.Println(canWinNim(i), i)
	}
}

func canWinNim(n int) bool {

	if n < 4 {
		return true
	}
	// dp := make([]bool, n+1)

	// dp[1], dp[2], dp[3] = true, true, true

	// for i := 4; i <= n; i++ {
	// 	dp[i] = !(dp[i-1] && dp[i-2] && dp[i-3])
	// }

	return n%4 != 0

	// return dp[n]

}
