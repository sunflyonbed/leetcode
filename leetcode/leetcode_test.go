package leetcode

import (
	"testing"
)

const (
	P4_TEST1_MAX = 3
	P4_TEST2_MAX = 3
)

func Test_Problem4(t *testing.T) {
	var test1 []int
	for i := 0; i < P4_TEST1_MAX; i++ {
		var test2 []int
		for j := 0; j < P4_TEST1_MAX; j++ {
			result := findMedianSortedArrays(test1, test2)
			t.Logf("test: i:%v j:%v result:%v", test1, test2, result)
			test2 = append(test2, j)
		}
		test1 = append(test1, i)
	}
	l1 := []int{1, 3}
	l2 := []int{2}
	result := findMedianSortedArrays(l1, l2)
	t.Logf("test: i:%v j:%v result:%v", l1, l2, result)

}
