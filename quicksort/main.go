package main

import (
	"fmt"
)

func main() {

	arr := []int{38, 27, 43, 3, 3, 9, 25, 19, 2}

	fmt.Println(quicksort(arr))
	fmt.Println(mergeSort(arr))

}

type sortBy []int

func (a sortBy) Len() int           { return len(a) }
func (a sortBy) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a sortBy) Less(i, j int) bool { return a[i] < a[j] }

func quicksort(a sortBy) sortBy {
	length := a.Len()
	if length < 2 {
		return a
	}

	pivot := a[0]

	remain := a[1:]

	left, right := sortBy{}, sortBy{}

	for _, v := range remain {
		if v < pivot {
			left = append(left, v)
			continue
		}
		right = append(right, v)
	}

	return append(append(quicksort(left), pivot), quicksort(right)...)

}

func mergeSort(s sortBy) sortBy {

	length := s.Len()

	if length < 2 {
		return s
	}

	middle := length / 2

	left, right := s[0:middle], s[middle:]

	leftSorted, rightSorted := mergeSort(left), mergeSort(right)

	res := merge(leftSorted, rightSorted)

	return res
}

func merge(left sortBy, right sortBy) sortBy {
	merged := sortBy{}

	i, j := 0, 0

	for i < left.Len() && j < right.Len() {
		l, r := left[i], right[j]

		if l < r {
			merged = append(merged, l)
			i++
			continue
		}
		merged = append(merged, r)
		j++
	}

	merged = append(merged, left[i:]...)
	merged = append(merged, right[j:]...)

	return merged

}
