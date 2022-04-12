package main

import "testing"

type testCase struct {
	desc   string
	input  []string
	result bool
}

func Test(t *testing.T) {
	testCases := []testCase{
		{
			desc:   "Expect true",
			input:  []string{"abc", "bca"},
			result: true,
		},
		{
			desc:   "Expect false",
			input:  []string{"abc", "bcd"},
			result: false,
		},
		{
			desc:   "Expect true",
			input:  []string{"abcde", "edcba"},
			result: true,
		},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			res := isAnagram(tC.input[0], tC.input[1])
			if res != tC.result {
				t.Errorf("Expect %v, Actual %v", tC.result, res)
			}

		})
	}
}
