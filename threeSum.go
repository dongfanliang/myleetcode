package leetcode

import (
	"sort"
)

// # 15
func threeSum(nums []int) [][]int {
	res := make([][]int, 0, 2)
	count := len(nums)
	if count < 3 {
		return res
	}

	sort.Ints(nums)
	for i := 0; i < count-2; i++ {
		if nums[i] > 0 {
			break
		}

		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		left := i + 1
		right := count - 1
		for left < right {
			s := nums[i] + nums[left] + nums[right]
			if s == 0 {
				res = append(res, []int{nums[i], nums[left], nums[right]})
				for left < right && nums[left] == nums[left+1] {
					left += 1
				}
				for left < right && nums[right] == nums[right-1] {
					right -= 1
				}
				left += 1
				right -= 1
			} else if s < 0 {
				left += 1
			} else {
				right -= 1
			}
		}
	}

	return res
}
