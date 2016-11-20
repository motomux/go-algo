package sort

// InsertionSort sorts slice of int by insertion sort
func InsertionSort(nums []int) {
	for i := 1; i < len(nums); i++ {
		for k := i; 0 < k; k-- {
			if nums[k] < nums[k-1] {
				nums[k], nums[k-1] = nums[k-1], nums[k]
			} else {
				continue
			}
		}
	}
}
