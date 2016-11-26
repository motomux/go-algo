package sort

// SelectionSort sorts slice of nums by selection sort
func SelectionSort(nums []int) {
	for i := 0; i < len(nums); i++ {
		min := nums[i]
		minIndex := i
		for k := i + 1; k < len(nums); k++ {
			if min > nums[k] {
				min = nums[k]
				minIndex = k
			}
		}
		nums[i], nums[minIndex] = nums[minIndex], nums[i]
	}
}
