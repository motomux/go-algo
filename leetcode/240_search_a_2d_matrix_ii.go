package leetcode

func searchMatrix(matrix [][]int, target int) bool {
	for _, nums := range matrix {
		left, right := 0, len(nums)-1
		for left <= right {
			mid := (left + right) / 2
			if nums[mid] == target {
				return true
			} else if nums[mid] < target {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	return false
}
