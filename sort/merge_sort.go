package sort

import "math"

// MargeSort sorts slice of int by merge sort
func MargeSort(nums []int, left, right int) {
	if left < right {
		mid := (left + right) / 2
		MargeSort(nums, left, mid)
		MargeSort(nums, mid+1, right)
		Marge(nums, left, mid, right)
	}
}

// Marge marge slice of int by 2 finger algorithm
func Marge(nums []int, left, mid, right int) {
	leftNums := make([]int, mid-left+1)
	copy(leftNums, nums[left:])
	leftNums = append(leftNums, math.MaxInt64)
	rightNums := make([]int, right-mid)
	copy(rightNums, nums[mid+1:])
	rightNums = append(rightNums, math.MaxInt64)
	l, r := 0, 0
	for i := left; i <= right; i++ {
		if leftNums[l] < rightNums[r] {
			nums[i] = leftNums[l]
			l++
		} else {
			nums[i] = rightNums[r]
			r++
		}
	}
}
