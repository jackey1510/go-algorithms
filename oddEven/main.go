package main

import "fmt"

func main() {
	slice := getSlice()
	for _, num := range slice {
		print(num, isEven(num))
	}
}

func isEven(number int) bool {
	return number%2 == 0
}

func getSlice() []int {
	slice := []int{}
	for i := 0; i <= 10; i++ {
		slice = append(slice, i)
	}
	return slice
}

func print(number int, isEven bool) {
	if isEven {
		fmt.Printf("%d is even\n", number)
		return
	}
	fmt.Printf("%d is odd\n", number)
}
