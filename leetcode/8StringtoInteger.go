package leetcode

import (
	"math"
)

func myAtoi(str string) int {
	myByte := []byte(str)
	var sign int32 = 0
	var result int32
	hasCount := false
	for _, v := range myByte {
		if v >= '0' && v <= '9' {
			hasCount = true
			diff := int32(v - '0')
			tmp := result*10 + diff
			if tmp/10 != result {
				if sign == -1 {
					result = math.MinInt32
				} else {
					result = math.MaxInt32
				}
				return int(result)
			}
			result = tmp
			continue
		}
		if v == '-' {
			if sign != 0 || hasCount {
				break
			}
			sign = -1
			hasCount = true
			continue
		}
		if v == '+' {
			if sign != 0 || hasCount {
				break
			}
			sign = 1
			hasCount = true
			continue
		}
		if v != ' ' {
			break
		} else if hasCount {
			break
		}
	}
	if sign == 0 {
		sign = 1
	}
	return int(sign * result)
}
