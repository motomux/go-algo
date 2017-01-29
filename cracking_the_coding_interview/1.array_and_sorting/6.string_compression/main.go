package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("vim-go")
}

// StringCompression compresses input and return it if compressed string is shorter than input
// i.e aabbbbbccd -> a2b5c2d1
func StringCompression(s string) string {
	var count int
	compressed := make([]byte, 0)
	for i := 0; i < len(s); i++ {
		count++
		if len(s)-1 == i || s[i] != s[i+1] {
			sCount := strconv.Itoa(count)
			compressed = append(compressed, s[i])
			compressed = append(compressed, []byte(sCount)...)
			count = 0
		}
	}

	if sCompressed := string(compressed); len(sCompressed) < len(s) {
		return sCompressed
	}
	return s
}
