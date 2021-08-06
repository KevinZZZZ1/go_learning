package leetcode011

func MaxArea(height []int) int {
	// 双指针法，leftIndex和rightIndex先分别指向切片的首尾，然后分别向中间移动
	leftIndex, rightIndex := 0, len(height)-1
	maxArea := 0

	for leftIndex < rightIndex {
		// 计算围成的面积
		tmpArea := (rightIndex - leftIndex) * minInt(height[leftIndex], height[rightIndex])
		maxArea = maxInt(tmpArea, maxArea)

		// 移动左，右指针。移动切片中对应值较小的，如果是左指针就向后移，右指针向前移
		if height[leftIndex] > height[rightIndex] {
			rightIndex--
		} else {
			leftIndex++
		}

	}

	return maxArea
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func minInt(a, b int) int {
	if a > b {
		return b
	}

	return a
}
