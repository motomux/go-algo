package leetcode

// 136. Single Number
func singleNumber(nums []int) int {
	m := make(map[int]int)
	for _, num := range nums {
		if count, ok := m[num]; !ok {
			m[num] = 1
		} else {
			m[num] = count + 1
		}
	}

	for k, count := range m {
		if count == 1 {
			return k
		}
	}

	return -1
}
