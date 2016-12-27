package leetcode

// 409. Longest Palindrome
func longestPalindrome(s string) int {
	bytes := make(map[byte]int, 0)
	for i := 0; i < len(s); i++ {
		b := s[i]
		count := bytes[b]
		bytes[b] = count + 1
	}

	res := 0
	hasPrime := false
	for _, count := range bytes {
		x, y := count/2, count%2
		res += x * 2
		if y == 1 {
			hasPrime = true
		}
	}

	if hasPrime {
		res++
	}
	return res
}
