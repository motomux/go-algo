package main

import "fmt"

func main() {
	fmt.Println("vim-go")
}

// CheckPlindromePermutation checks if input is plindrome permutation (can be generated as the same forwards and backwards)
// Only checks ascii
func CheckPlindromePermutation(s string) bool {
	table := [128]int{} // Assumption
	oddCount := 0

	for _, c := range s {
		if c == ' ' {
			continue
		}

		table[c]++
		if table[c]%2 == 0 {
			oddCount--
		} else {
			oddCount++
		}
	}

	if oddCount > 1 {
		return false
	}
	return true
}

// CheckPlindromePermutation2 checks if input is plindrome permutation (can be generated as the same forwards and backwards)
// Only checks ascii
// Case sensitive
func CheckPlindromePermutation2(s string) bool {
	var table int64
	for _, c := range s {
		if c == ' ' {
			continue
		}

		mask := int64(1 << uint(c-'A'))
		if table&mask == 0 {
			table |= mask
		} else {
			table ^= mask
		}
	}

	if table == 0 || (table-1)&table == 0 {
		return true
	}

	return false
}
