package main

import "fmt"

func main() {
	fmt.Println("vim-go")
}

// RotateMatrix rotates matrix
func RotateMatrix(m [][]int) {
	if len(m) == 0 {
		return
	}

	for i := 0; i < (len(m)+1)*2; i++ {
		size := len(m) - i*2
		for k := 0; k < size-1; k++ {
			m[i][i+k], m[i+k][i+size-1], m[i+size-1][i+size-1-k], m[i+size-1-k][i] = m[i+size-1-k][i], m[i][i+k], m[i+k][i+size-1], m[i+size-1][i+size-1-k]
		}
	}
}
