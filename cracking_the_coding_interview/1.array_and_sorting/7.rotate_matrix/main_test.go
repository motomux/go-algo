package main

import (
	"reflect"
	"testing"

	"github.com/kr/pretty"
)

func Test(t *testing.T) {
	tests := map[string]struct {
		in, post [][]int
	}{

		"0x0": {
			in: [][]int{
				[]int{},
			},
			post: [][]int{
				[]int{},
			},
		},

		"1x1": {
			in: [][]int{
				[]int{1},
			},
			post: [][]int{
				[]int{1},
			},
		},

		"2x2": {
			in: [][]int{
				[]int{1, 2},
				[]int{3, 4},
			},
			post: [][]int{
				[]int{3, 1},
				[]int{4, 2},
			},
		},

		"3x3": {
			in: [][]int{
				[]int{1, 2, 3},
				[]int{4, 5, 6},
				[]int{7, 8, 9},
			},
			post: [][]int{
				[]int{7, 4, 1},
				[]int{8, 5, 2},
				[]int{9, 6, 3},
			},
		},

		"4x4": {
			in: [][]int{
				[]int{1, 2, 3, 4},
				[]int{5, 6, 7, 8},
				[]int{9, 10, 11, 12},
				[]int{13, 14, 15, 16},
			},
			post: [][]int{
				[]int{13, 9, 5, 1},
				[]int{14, 10, 6, 2},
				[]int{15, 11, 7, 3},
				[]int{16, 12, 8, 4},
			},
		},

		"5x5": {
			in: [][]int{
				[]int{1, 2, 3, 4, 5},
				[]int{6, 7, 8, 9, 10},
				[]int{11, 12, 13, 14, 15},
				[]int{16, 17, 18, 19, 20},
				[]int{21, 22, 23, 24, 25},
			},
			post: [][]int{
				[]int{21, 16, 11, 6, 1},
				[]int{22, 17, 12, 7, 2},
				[]int{23, 18, 13, 8, 3},
				[]int{24, 19, 14, 9, 4},
				[]int{25, 20, 15, 10, 5},
			},
		},
	}

	for k, test := range tests {
		t.Run(k, func(t *testing.T) {
			RotateMatrix(test.in)
			if !reflect.DeepEqual(test.in, test.post) {
				t.Errorf("diff=%v", pretty.Diff(test.in, test.post))
			}
		})
	}
}
