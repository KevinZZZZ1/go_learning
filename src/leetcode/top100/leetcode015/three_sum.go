package leetcode015

import "sort"

func ThreeSum(nums []int) [][]int {
	numsLen := len(nums)
	if numsLen < 3 {
		return [][]int{}
	}

	// 先排序
	sort.Sort(sort.IntSlice(nums))

	ans := make([][]int, 0)

	for i := 0; i < len(nums)-2; i++ {
		// 如果当前处理的元素和上一个处理的元素相同，则直接跳过
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		j, k := i+1, len(nums)-1

		for j < k {
			if nums[i]+nums[j]+nums[k] == 0 {
				ans = append(ans, []int{nums[i], nums[j], nums[k]})

				// 处理前后两个指针对应的元素有相同值的情况
				for j < k && nums[k] == nums[k-1] {
					k--
				}

				for j < k && nums[j] == nums[j+1] {
					j++
				}

				// 由于要三个数相加等于0，而且上面去掉了重复的情况，所以在有序slice中不可能再存在三个值保持不变的情况，所以这里的双指针都要移动
				k--
				j++

			} else if nums[i]+nums[j]+nums[k] < 0 {
				j++
			} else {
				k--
			}
		}

	}

	return ans

}
