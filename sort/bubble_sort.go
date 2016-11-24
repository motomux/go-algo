package sort

// BubbleSort sorts slice of int with bubble sort
func BubbleSort(nums []int) {
	for i := len(nums) - 1; i > 0; i-- {
		for k := 0; k < i; k++ {
			if nums[k] > nums[k+1] {
				nums[k], nums[k+1] = nums[k+1], nums[k]
			}
		}
	}
}
