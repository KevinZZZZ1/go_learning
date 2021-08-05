package leetcode004

func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	len1, len2 := len(nums1), len(nums2)

	if len1 == 0 && len2 == 0 {
		return 0.0
	}

	if len1 != 0 && len2 == 0 {
		return findMedianSortedArray(nums1)
	}

	if len2 != 0 && len1 == 0 {
		return findMedianSortedArray(nums2)
	}

	mergedArray := make([]int, len1+len2)
	curIndex1, curIndex2, curMergeIndex := 0, 0, 0

	for curMergeIndex < len1+len2 {
		if curIndex1 < len1 && curIndex2 < len2 {
			val1, val2 := nums1[curIndex1], nums2[curIndex2]
			if val1 < val2 {
				mergedArray[curMergeIndex] = val1
				curIndex1++
				curMergeIndex++
			} else {
				mergedArray[curMergeIndex] = val2
				curIndex2++
				curMergeIndex++
			}
		} else if curIndex1 < len1 && curIndex2 >= len2 {
			mergedArray[curMergeIndex] = nums1[curIndex1]
			curMergeIndex++
			curIndex1++
		} else if curIndex1 >= len1 && curIndex2 < len2 {
			mergedArray[curMergeIndex] = nums2[curIndex2]
			curMergeIndex++
			curIndex2++
		}

	}

	return findMedianSortedArray(mergedArray)
}

func findMedianSortedArray(nums []int) float64 {
	lenNums := len(nums)

	if lenNums <= 0 {
		return 0.0
	}

	if lenNums%2 == 0 {
		return (float64(nums[lenNums/2]) + float64(nums[lenNums/2-1])) / 2
	} else {
		return float64(nums[lenNums/2])
	}
}
