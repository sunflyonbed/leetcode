package leetcode

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	list := make([][]byte, len(strs))
	var minSize, minIndex int
	for index, v := range strs {
		list[index] = []byte(v)
		if len(list[index]) == 0 {
			return ""
		}
		if len(list[index]) < minSize || minSize == 0 {
			minSize = len(list[index])
			minIndex = index
		}
	}
	matchSize := len(list[minIndex])
	for index, v := range list {
		if index == minIndex {
			continue
		}
		for i := 0; i < matchSize; i++ {
			if list[minIndex][i] != v[i] {
				matchSize = i
				break
			}
		}
		if matchSize == 0 {
			return ""
		}
	}
	return string(list[minIndex][:matchSize])
}
