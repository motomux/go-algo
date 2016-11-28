package leetcode

func containsNearbyDuplicate(nums []int, k int) bool {
	m := make(map[int]int)
	for i, num := range nums {
		index, ok := m[num]
		if ok && i-index <= k {
			return true
		}
		m[num] = i
	}
	return false
}
