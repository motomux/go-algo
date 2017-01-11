package main

import (
	"reflect"
	"testing"
)

func TestSpiral(t *testing.T) {
	tests := map[string]struct {
		in  [][]int
		out []int
	}{
		"1x1": {
			in: [][]int{
				[]int{1},
			},
			out: []int{1},
		},
		"2x2": {
			in: [][]int{
				[]int{1, 2},
				[]int{3, 4},
			},
			out: []int{1, 2, 4, 3},
		},
		"3x3": {
			in: [][]int{
				[]int{1, 2, 3},
				[]int{4, 5, 6},
				[]int{7, 8, 9},
			},
			out: []int{1, 2, 3, 6, 9, 8, 7, 4, 5},
		},
		"4x4": {
			in: [][]int{
				[]int{1, 2, 3, 4},
				[]int{5, 6, 7, 8},
				[]int{9, 10, 11, 12},
				[]int{13, 14, 15, 16},
			},
			out: []int{1, 2, 3, 4, 8, 12, 16, 15, 14, 13, 9, 5, 6, 7, 11, 10},
		},
		"5x5": {
			in: [][]int{
				[]int{1, 2, 3, 4, 5},
				[]int{6, 7, 8, 9, 10},
				[]int{11, 12, 13, 14, 15},
				[]int{16, 17, 18, 19, 20},
				[]int{21, 22, 23, 24, 25},
			},
			out: []int{1, 2, 3, 4, 5, 10, 15, 20, 25, 24, 23, 22, 21, 16, 11, 6, 7, 8, 9, 14, 19, 18, 17, 12, 13},
		},
		"6x6": {
			in: [][]int{
				[]int{1, 2, 3, 4, 5, 6},
				[]int{7, 8, 9, 10, 11, 12},
				[]int{13, 14, 15, 16, 17, 18},
				[]int{19, 20, 21, 22, 23, 24},
				[]int{25, 26, 27, 28, 29, 30},
				[]int{31, 32, 33, 34, 35, 36},
			},
			out: []int{1, 2, 3, 4, 5, 6, 12, 18, 24, 30, 36, 35, 34, 33, 32, 31, 25, 19, 13, 7, 8, 9, 10, 11, 17, 23, 29, 28, 27, 26, 20, 14, 15, 16, 22, 21},
		},
	}

	for k, test := range tests {
		t.Run(k, func(t *testing.T) {
			out := spiral(test.in)
			if !reflect.DeepEqual(out, test.out) {
				t.Errorf("actual=%v, expected=%v", out, test.out)
			}
		})
	}
}
