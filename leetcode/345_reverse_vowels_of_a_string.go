package leetcode

// 345. Reverse Vowels of a String

func reverseVowels(s string) string {
	res := make([]byte, len(s))
	copy(res, s)
	vowels := map[byte]bool{'a': true, 'i': true, 'u': true, 'e': true, 'o': true, 'A': true, 'I': true, 'U': true, 'E': true, 'O': true}
	for i, k := 0, len(s)-1; i < k; {
		for i < k {
			if _, ok := vowels[s[i]]; ok {
				break
			}
			i++
		}
		for i < k {
			if _, ok := vowels[s[k]]; ok {
				break
			}
			k--
		}
		res[i], res[k] = res[k], res[i]
		i++
		k--
	}
	return string(res)
}
