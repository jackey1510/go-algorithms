package main

import "testing"

func Test(t *testing.T) {
	res := ContainsDuplicate([]int{1, 2, 3, 1})
	if !res {
		t.Errorf("res: %v", res)
	}
}
