package leetcode

func maxArea(height []int) int {
	left := 0
	right := len(height) - 1
	result := 0
	for left < right {
		result = max(result, (right-left)*min(height[left], height[right]))
		if height[left] > height[right] {
			right--
		} else {
			left++
		}
	}
	return result
}

func min(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}

/*
func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
*/
