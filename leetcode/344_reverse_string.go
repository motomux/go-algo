package leetcode

// 344. Reverse String
func reverseString(s string) string {
	l := len(s)
	if l == 0 {
		return ""
	}
	res := make([]byte, l)
	for i := 0; i <= l/2; i++ {
		res[i], res[l-1-i] = s[l-1-i], s[i]
	}
	return string(res)
}
