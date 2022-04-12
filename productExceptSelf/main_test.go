package main

import "testing"

func Test(t *testing.T) {
	res := productExceptSelf([]int{1, 2, 3})
	if !equal([]int{6, 3, 2}, res) {
		t.Errorf("%v", res)
	}
}

func Test2(t *testing.T) {
	res := productExceptSelf([]int{1, 2, 3, 4})
	if !equal([]int{24, 12, 8, 6}, res) {
		t.Errorf("%v", res)
	}
}

func Test3(t *testing.T) {
	res := productExceptSelf([]int{-1, 1, 0, -3, 3})
	if !equal([]int{0, 0, 9, 0, 0}, res) {
		t.Errorf("%v", res)
	}
}

func equal(a []int, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
