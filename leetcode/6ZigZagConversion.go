package leetcode

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	myByte := []byte(s)
	size := len(myByte)
	newResult := make([]byte, 0, size)
	set := (numRows - 1) * 2
	for i := 0; i < numRows; i++ {
		for index := 0; ; {
			if newIndex := index*set + i; newIndex < size {
				newResult = append(newResult, myByte[newIndex])
			} else {
				break
			}
			if i != 0 && i != numRows-1 {
				if newIndex := (index+1)*set - i; newIndex < size {
					newResult = append(newResult, myByte[newIndex])
				} else {
					break
				}
			}
			index++
		}
	}
	return string(newResult)
}
