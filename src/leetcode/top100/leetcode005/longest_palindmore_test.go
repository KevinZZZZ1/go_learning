package leetcode005

import "testing"

func TestLongestPalindrome(t *testing.T) {
	var testCases = []struct {
		s        string
		expected string
	}{
		{"babad", "bab"},
		{"cbbd", "bb"},
		{"a", "a"},
		{"ac", "a"},
	}

	for _, testCase := range testCases {
		actual := LongestPalindrome(testCase.s)
		if actual != testCase.expected {
			t.Errorf("input: %v, get answer: %v, but expected: %v", testCase.s, actual, testCase.expected)
		}

	}
}
