package main

import "fmt"

func main() {
	fmt.Println(firstBadVersion(10000000))
}

// https://leetcode.com/problems/first-bad-version/
/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

func isBadVersion(version int) bool { return version >= 7822 }

func firstBadVersion(n int) int {
	return recur(n, n)
}

func recur(n int, sub int) int {
	sub = sub / 2
	if isBadVersion(n) {
		if !isBadVersion(n - 1) {
			return n
		}
		return recur(n-max(sub, 1), sub)
	}

	return recur(n+max(sub, 1), sub)
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
