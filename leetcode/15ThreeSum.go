package leetcode

import "sort"

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	size := len(nums)
	var result [][]int
	for i := 0; i < size-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		left := i + 1
		right := size - 1
		toCal := 0 - nums[i]
		for left < right {
			if total := nums[left] + nums[right]; total == toCal {
				result = append(result, []int{nums[i], nums[left], nums[right]})
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				for left < right && nums[right] == nums[right-1] {
					right--
				}
				left++
				right--
			} else if total > toCal {
				right--
			} else {
				left++
			}
		}
	}
	return result
}
