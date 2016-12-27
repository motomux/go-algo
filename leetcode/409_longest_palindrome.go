package leetcode

// 409. Longest Palindrome
func longestPalindrome(s string) int {
	bytes := make(map[byte]int, 0)
	count := 0
	for i := 0; i < len(s); i++ {
		b := s[i]
		_, ok := bytes[b]
		if ok {
			count++
			delete(bytes, b)
		} else {
			bytes[b] = 1
		}
	}
	if len(bytes) > 0 {
		return count*2 + 1
	}
	return count * 2
}
