package main

import "fmt"

func main() {
	fmt.Println("vim-go")
}

// CheckPermutation checks if s2 is permutation of s1
// Only checks ascii code
func CheckPermutation(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	table := [128]int{}
	for _, c := range s1 {
		table[c]++
	}

	for _, c := range s2 {
		if table[c] < 1 {
			return false
		}
		table[c]--
	}

	return true
}

// CheckPermutation2 checks if s2 is permutation of s1
// Only checks alphabet
// Case sensitive
func CheckPermutation2(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	var table int64
	for _, c := range s1 {
		mask := int64(1 << uint(c-'A'))
		if table&mask == 0 {
			table |= mask
		} else {
			table ^= mask
		}
	}

	for _, c := range s2 {
		mask := int64(1 << uint(c-'A'))
		if table&mask == 0 {
			table |= mask
		} else {
			table ^= mask
		}
	}

	if table == 0 {
		return true
	}
	return false
}
