package leetcode

func romanToInt(s string) int {
	roman := []byte(s)
	var result int
	size := len(s)
	for i := 0; i < size; i++ {
		switch roman[i] {
		case 'M':
			result += 1000
		case 'D':
			result += 500
		case 'C':
			if i == size-1 {
				result += 100
				break
			}
			if roman[i+1] == 'M' {
				result += 900
				i++
			} else if roman[i+1] == 'D' {
				result += 400
				i++
			} else {
				result += 100
			}
		case 'L':
			result += 50
		case 'X':
			if i == size-1 {
				result += 10
				break
			}
			if roman[i+1] == 'C' {
				result += 90
				i++
			} else if roman[i+1] == 'L' {
				result += 40
				i++
			} else {
				result += 10
			}
		case 'V':
			result += 5
		case 'I':
			if i == size-1 {
				result += 1
				break
			}
			if roman[i+1] == 'X' {
				result += 9
				i++
			} else if roman[i+1] == 'V' {
				result += 4
				i++
			} else {
				result += 1
			}
		}
	}
	return result
}
