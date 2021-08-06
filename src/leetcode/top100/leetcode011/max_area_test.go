package leetcode011

import "testing"

func TestMaxArea(t *testing.T) {
	var testCases = []struct {
		heights  []int
		expected int
	}{
		{[]int{1, 8, 6, 2, 5, 4, 8, 3, 7}, 49},
		{[]int{1, 1}, 1},
		{[]int{4, 3, 2, 1, 4}, 16},
		{[]int{1, 2, 1}, 2},
	}

	for i, testCase := range testCases {
		actual := MaxArea(testCase.heights)
		if actual != testCase.expected {
			t.Errorf("case %v, get answer %v, but expected %v", i, actual, testCase.expected)
		}

	}

}
