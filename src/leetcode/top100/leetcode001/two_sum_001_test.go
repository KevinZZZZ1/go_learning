package leetcode001

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {

	var twoSumTests = []struct {
		in       []int
		target   int
		expected []int
	}{
		{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
		{[]int{3, 2, 4}, 6, []int{1, 2}},
		{[]int{3, 3}, 6, []int{0, 1}},
	}

	for _, test := range twoSumTests {
		actual := TwoSum(test.in, test.target)

		isMatched := reflect.DeepEqual(actual, test.expected)

		if !isMatched {
			t.Errorf("TwoSum(%v, %d); expected %v; actual %v", test.in, test.target, test.expected, actual)
		}

	}

}
