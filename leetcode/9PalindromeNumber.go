package leetcode

/*
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	var list []int
	for x > 0 {
		list = append(list, x%10)
		x = x / 10
	}
	size := len(list)
	for i := 0; i < size/2; i++ {
		if list[i] != list[size-1-i] {
			return false
		}
	}
	return true
}
*/

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	tmp := x
	var newX int
	for tmp > 0 {
		newX = newX*10 + tmp%10
		tmp = tmp / 10
	}
	return newX == x
}
