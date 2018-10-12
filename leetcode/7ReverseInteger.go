package leetcode

func reverse(x int) int {
	myX := int32(x)
	var result int32
	for myX != 0 {
		tail := myX % 10
		newResult := result*10 + tail
		if newResult/10 != result {
			return 0
		}
		result = newResult
		myX = myX / 10
	}
	return int(result)
}
