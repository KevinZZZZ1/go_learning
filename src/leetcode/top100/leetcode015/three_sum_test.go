package leetcode015

import "testing"

func TestThreeSum(t *testing.T) {
	var testCases = []struct {
		in       []int
		expected [][]int
	}{
		{[]int{-1, 0, 1, 2, -1, -4}, [][]int{{-1, -1, 2}, {-1, 0, 1}}},
		{[]int{}, [][]int{}},
		{[]int{0}, [][]int{}},
	}

	for i, testCase := range testCases {
		actual := ThreeSum(testCase.in)
		t.Logf("test case %v, got answer %v, expected %v", i, actual, testCase.in)
	}

}
