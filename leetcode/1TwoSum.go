package leetcode

func twoSum(nums []int, target int) []int {
	list := make(map[int]int)
	for index, v := range nums {
		needed := target - v
		if needIndex, exist := list[v]; exist {
			if needIndex == index {
				continue
			}
			return []int{needIndex, index}
		}
		list[needed] = index
	}
	return nil
}
