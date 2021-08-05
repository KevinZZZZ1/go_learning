package leetcode004

import "testing"

func TestFindMedianSortedArrays(t *testing.T) {
	var testCases = []struct {
		nums1    []int
		nums2    []int
		expected float64
	}{
		{[]int{1, 3}, []int{2}, 2.0},
		{[]int{1, 2}, []int{3, 4}, 2.5},
		{[]int{0, 0}, []int{0, 0}, 0.0},
		{[]int{}, []int{1}, 1.0},
		{[]int{2}, []int{}, 2.0},
	}

	for _, testCase := range testCases {
		actual := FindMedianSortedArrays(testCase.nums1, testCase.nums2)

		if actual != testCase.expected {
			t.Error("not match")
		}
	}

}
