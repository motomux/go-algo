package main

import "errors"

// ErrNotFound is given when specified num was not found
var ErrNotFound = errors.New("Not found")

// BinarySearch searches specified number in the sorted slice
func BinarySearch(nums []int, num int) (index int, err error) {
	l, r := 0, len(nums)-1
	for l <= r {
		m := (l + r) / 2
		if nums[m] == num {
			return m, nil
		} else if nums[m] < num {
			l = m + 1
		} else {
			r = m - 1
		}
	}

	return -1, ErrNotFound
}
