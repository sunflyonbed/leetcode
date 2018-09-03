package leetcode

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	size1 := len(nums1)
	size2 := len(nums2)
	mid1 := (size1 + size2 + 1) / 2
	mid2 := (size1 + size2 + 2) / 2
	return float64(mySort(nums1, nums2, mid1)+mySort(nums1, nums2, mid2)) / 2
}

func mySort(min, max []int, index int) int {
	size1 := len(min)
	size2 := len(max)
	if size1 == 0 && size2 == 0 {
		return 0
	}
	if size1 > size2 {
		return mySort(max, min, index)
	}
	if size1 == 0 {
		return max[index-1]
	}
	if index == 1 {
		return myMin(min[0], max[0])
	}
	index1 := myMin(size1, index/2)
	index2 := myMin(size2, index/2)
	if min[index1-1] > max[index2-1] {
		return mySort(min, max[index2:], index-index2)
	} else {
		return mySort(min[index1:], max, index-index1)
	}
}

func myMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}
