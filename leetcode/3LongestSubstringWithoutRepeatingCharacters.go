package leetcode

func lengthOfLongestSubstring(s string) int {
	list := []byte(s)
	indexMap := make(map[byte]int)
	for _, v := range list {
		indexMap[v] = -1
	}
	var length int
	var pos int = -1
	for index, v := range list {
		pos = max(pos, indexMap[v])
		length = max(length, index-pos)
		indexMap[v] = index
	}
	return length
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
