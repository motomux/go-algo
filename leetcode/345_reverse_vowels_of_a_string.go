package leetcode

// 345. Reverse Vowels of a String
func reverseVowels(s string) string {
	res := make([]byte, 0)
	vowels := make([]byte, 0)
	for i := 0; i < len(s); i++ {
		switch b := s[i]; b {
		case 'a', 'i', 'u', 'e', 'o', 'A', 'I', 'U', 'E', 'O':
			vowels = append(vowels, b)
		}
	}

	for i, k := 0, len(vowels)-1; i < len(s); i++ {
		switch b := s[i]; b {
		case 'a', 'i', 'u', 'e', 'o', 'A', 'I', 'U', 'E', 'O':
			res = append(res, vowels[k])
			k--
		default:
			res = append(res, b)
		}
	}
	return string(res)
}
