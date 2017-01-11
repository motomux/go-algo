package main

import "fmt"

func main() {
	m0 := [][]int{
		[]int{1},
	}
	m1 := [][]int{
		[]int{1, 2},
		[]int{3, 4},
	}
	m2 := [][]int{
		[]int{1, 2, 3},
		[]int{4, 5, 6},
		[]int{7, 8, 9},
	}
	m3 := [][]int{
		[]int{1, 2, 3, 4},
		[]int{5, 6, 7, 8},
		[]int{9, 10, 11, 12},
		[]int{13, 14, 15, 16},
	}
	m4 := [][]int{
		[]int{1, 2, 3, 4, 5},
		[]int{6, 7, 8, 9, 10},
		[]int{11, 12, 13, 14, 15},
		[]int{16, 17, 18, 19, 20},
		[]int{21, 22, 23, 24, 25},
	}

	fmt.Println(spiral(m0))
	fmt.Println(spiral(m1))
	fmt.Println(spiral(m2))
	fmt.Println(spiral(m3))
	fmt.Println(spiral(m4))

}

func spiral(m [][]int) []int {
	res := make([]int, 0, len(m)*len(m)*2)

	for k := 0; k < (len(m)+1)/2; k++ {
		x := k
		length := len(m) - k
		// top
		for i := x; i < length; i++ {
			res = append(res, m[x][i])
		}

		// right
		for i := x + 1; i < length; i++ {
			res = append(res, m[i][length-1])
		}

		// bottom
		for i := length - 2; i >= x; i-- {
			res = append(res, m[length-1][i])
		}

		for i := length - 2; i > x; i-- {
			res = append(res, m[i][x])
		}
	}

	return res
}
