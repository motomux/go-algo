package main

import "fmt"

func main() {
	fmt.Println("vim-go")
}

// CheckOneAway checks inputs are one away (delete, insert, or replace) each other
func CheckOneAway(s1, s2 string) bool {
	diff := len(s1) - len(s2)
	if diff == 1 {
		return CheckAdd(s1, s2)
	} else if diff == -1 {
		return CheckAdd(s2, s1)
	} else if len(s1) == len(s2) {
		return CheckOneReplace(s1, s2)
	}

	return false
}

// CheckAdd checks if s1 has one additional character
func CheckAdd(s1, s2 string) bool {
	var diff int
	for i, k := 0, 0; k < len(s2); {
		if s1[i] != s2[k] {
			if diff == 0 {
				diff++
				i++
			} else {
				return false
			}
		} else {
			i++
			k++
		}
	}

	return true
}

// CheckOneReplace checks if s1 has one character which is replaced from s2
func CheckOneReplace(s1, s2 string) bool {
	var diff int
	for i, k := 0, 0; i < len(s1); i, k = i+1, k+1 {
		if s1[i] != s2[k] {
			if diff == 0 {
				diff++
			} else {
				return false
			}
		}
	}

	return true
}
