package main

import "fmt"

func fibor(n int) int {
	if n <= 1 {
		return n
	}
	return fibor(n-1) + fibor(n-2)
}

func fiboi() func() int {
	x, y := 0, 1
	return func() int {
		r := x
		x, y = y, x+y
		return r
	}
}

func fiboc(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func main() {
	// for i := 0; i < 100; i++ {
	// 	fmt.Printf("%d ", fibor(i))
	// }
	fmt.Println()
	next := fiboi()
	for i := 0; i < 200; i++ {
		fmt.Printf("%d ", next())
	}
	fmt.Println()
	c := make(chan int, 200)
	go fiboc(cap(c), c)
	for i := range c {
		fmt.Printf("%d ", i)
	}
}
