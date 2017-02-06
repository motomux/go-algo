package main

import "strings"

// StringRotation checks if s2 is rotation of s1
func StringRotation(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	return strings.Contains(s1+s1, s2)
}
