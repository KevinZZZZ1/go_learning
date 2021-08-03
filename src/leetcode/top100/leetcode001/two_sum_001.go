package leetcode001

func TwoSum(nums []int, target int) []int {
	var numLen int = len(nums)
	var numToIndex map[int]int = make(map[int]int, numLen)

	for i := 0; i < numLen; i++ {
		numToIndex[nums[i]] = i
	}

	for i := 0; i < numLen; i++ {
		left := target - nums[i]

		index, ok := numToIndex[left]

		// 在map中找到，且不是自己
		if ok && index != i {
			return []int{i, index}
		}

	}

	return []int{}
}
