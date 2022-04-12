package main

import "testing"

type testC struct {
	desc   string
	input  []int
	result int
}

func Test(t *testing.T) {
	testCases := []testC{
		{
			desc:   "[10,9,2,5,3,7,101,18], 4",
			input:  []int{10, 9, 2, 5, 3, 7, 101, 18},
			result: 4,
		},
		{
			desc:   "[0,1,0,3,2,3], 4",
			input:  []int{0, 1, 0, 3, 2, 3},
			result: 4,
		},
		{
			desc:   " [7,7,7,7,7,7,7], 1",
			input:  []int{7, 7, 7, 7, 7, 7, 7},
			result: 1,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			res := lengthOfLIS(tC.input)
			if res != tC.result {
				t.Errorf("%v,%v", res, tC.result)
			}
		})
	}
}
