package main

// ZeroMatrix pads 0s when any element of the same row or column has 0
func ZeroMatrix(m [][]int) {
	r := make([]bool, len(m))
	c := make([]bool, len(m))

	for i := 0; i < len(m); i++ {
		for k := 0; k < len(m[i]); k++ {
			if m[i][k] == 0 {
				r[i] = true
				c[k] = true
			}
		}
	}

	for i, v := range r {
		if v {
			padRow(m, i)
		}
	}

	for i, v := range c {
		if v {
			padCol(m, i)
		}
	}
}

func padRow(m [][]int, i int) {
	for k := 0; k < len(m[i]); k++ {
		m[i][k] = 0
	}
}

func padCol(m [][]int, i int) {
	for k := 0; k < len(m); k++ {
		m[k][i] = 0
	}
}
