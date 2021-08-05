package leetcode003

type EMPTY struct{}

func LengthOfLongestSubstring(s string) int {
	sRunes := []rune(s)
	lenRunes := len(sRunes)
	set := make(map[rune]EMPTY, lenRunes)

	ans, startIndex, endIndex := 0, 0, 0

	for startIndex < lenRunes && endIndex < lenRunes {
		sRune := sRunes[endIndex]
		_, contains := set[sRune]
		if !contains {
			set[sRune] = EMPTY{}
			ans = maxInt(ans, endIndex-startIndex+1)
			endIndex++
		} else {
			set = make(map[rune]EMPTY, lenRunes)
			startIndex++
			endIndex = startIndex
		}
	}

	return ans
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}

	return b
}
