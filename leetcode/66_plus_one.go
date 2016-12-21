package leetcode

// 66. Plus One
func plusOne(digits []int) []int {
	if len(digits) == 0 {
		return []int{}
	}

	res := make([]int, len(digits))
	carry := 1
	for i := len(digits) - 1; i >= 0; i-- {
		num := digits[i] + carry
		carry, res[i] = num/10, num%10
	}
	if carry > 0 {
		res = append([]int{carry}, res...)
	}
	return res
}
