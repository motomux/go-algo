package main

import "fmt"

func main() {
	fmt.Println("vim-go")
}

// IsUnique checks if an input contains only unique characters
// Only checks ASCII code
func IsUnique(s string) bool {
	table := [128]int{}

	for _, c := range s {
		if table[c] == 1 {
			return false
		}
		table[c]++
	}

	return true
}

// IsUnique2 checks if an input contains only unique characters with bit vector
// Only checks alphabet
// Case sensitive
func IsUnique2(s string) bool {
	var table int64
	for _, c := range s {
		mask := int64(1 << uint(c-'A'))
		if table&mask > 0 {
			return false
		}
		table |= mask
	}

	return true
}
