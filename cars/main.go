package main

func main() {

}

func Solution(P []int, S []int) int {
	// write your code in Go 1.4

	cars := countSort(S)

	sumOfPeople, noOfCars, peopleTaken := 0, 0, 0

	for _, p := range P {
		sumOfPeople += p
	}

	for _, seats := range cars {
		peopleTaken += seats
		noOfCars++
		if peopleTaken >= sumOfPeople {
			return noOfCars
		}
	}

	return noOfCars
}

func countSort(s []int) []int {
	m := make(map[int]int)
	max := s[0]
	for _, v := range s {
		m[v]++
		if v > max {
			max = v
		}
	}
	result := []int{}

	for i := 0; i <= max; i++ {
		count := m[i]
		for count > 0 {
			result = append(result, i)
			count--
		}
	}
	return result
}
