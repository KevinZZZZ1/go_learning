package leetcode003

import "testing"

func TestLengthOfLongestSubstring(t *testing.T) {
	var testString string = "abcabcbb"
	var expected int = 3

	ans := LengthOfLongestSubstring(testString)

	if ans != expected {
		t.Error("ans is not expected")
	}
}
