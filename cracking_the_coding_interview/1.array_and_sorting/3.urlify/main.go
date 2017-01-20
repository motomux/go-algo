package main

import "fmt"

func main() {
	fmt.Println("vim-go")
}

// URLify mutates input by replacing white space with %20
func URLify(s []byte, length int) {
	var spaceCount int
	for i := 0; i < length; i++ {
		if s[i] == ' ' {
			spaceCount++
		}
	}

	newLength := length + spaceCount*2
	for i, k := length-1, newLength-1; 0 <= i; i-- {
		if s[i] == ' ' {
			s[k] = '0'
			s[k-1] = '2'
			s[k-2] = '%'
			k -= 3
		} else {
			s[k] = s[i]
			k--
		}
	}
}
