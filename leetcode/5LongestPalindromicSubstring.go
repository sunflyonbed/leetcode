package leetcode

func longestPalindrome(s string) string {
	list := []byte(s)
	size := len(list)
	if size == 0 {
		return ""
	}
	if size == 1 {
		return s
	}
	var longIndex, maxLen int
	for i := 0; i < size; {
		if size-i < maxLen/2 {
			break
		}
		frontIndex := i
		tailIndex := i
		for tailIndex < size-1 && list[tailIndex] == list[tailIndex+1] {
			tailIndex++
		}
		i = tailIndex + 1
		for tailIndex < size-1 && frontIndex > 0 && list[frontIndex-1] == list[tailIndex+1] {
			tailIndex++
			frontIndex--
		}
		thisLen := tailIndex - frontIndex + 1
		if thisLen > maxLen {
			maxLen = thisLen
			longIndex = frontIndex
		}
	}
	return string(list[longIndex : longIndex+maxLen])
}
