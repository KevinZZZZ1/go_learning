package leetcode005

var low, maxLength = 0, 0

func LongestPalindrome(s string) string {
	low, maxLength = 0, 0
	sRunes := []rune(s)
	sLen := len(sRunes)

	if sLen < 2 {
		return s
	}

	// 尝试以字符串中某个字符开始，向左右扩展看是否能满足回文条件
	for i := 0; i < sLen-1; i++ {
		// 回文字符串为奇数时：例如：aba
		findPalindrome(sRunes, i, i)
		// 回文字符串为偶数，例如：abba
		findPalindrome(sRunes, i, i+1)
	}

	return string(sRunes[low : low+maxLength])
}

func findPalindrome(s []rune, startIndex int, endIndex int) {
	// 判断是否满足回文条件
	for startIndex > -1 && endIndex < len(s) && s[startIndex] == s[endIndex] {
		startIndex--
		endIndex++
	}

	if maxLength < endIndex-startIndex-1 {
		maxLength = endIndex - startIndex - 1
		low = startIndex + 1
	}

}
