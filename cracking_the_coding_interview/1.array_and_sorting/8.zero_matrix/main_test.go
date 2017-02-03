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

		"case1": {
			in: [][]int{
				[]int{1, 3, 5, 8, 0},
				[]int{1, 3, 5, 8, 1},
				[]int{1, 3, 5, 8, 1},
				[]int{1, 3, 5, 8, 1},
				[]int{1, 3, 5, 8, 1},
			},
			post: [][]int{
				[]int{0, 0, 0, 0, 0},
				[]int{1, 3, 5, 8, 0},
				[]int{1, 3, 5, 8, 0},
				[]int{1, 3, 5, 8, 0},
				[]int{1, 3, 5, 8, 0},
			},
		},

		"case2": {
			in: [][]int{
				[]int{0, 3, 5, 8, 1},
				[]int{1, 3, 5, 8, 1},
				[]int{1, 3, 5, 8, 1},
				[]int{1, 3, 5, 8, 1},
				[]int{1, 3, 5, 8, 1},
			},
			post: [][]int{
				[]int{0, 0, 0, 0, 0},
				[]int{0, 3, 5, 8, 1},
				[]int{0, 3, 5, 8, 1},
				[]int{0, 3, 5, 8, 1},
				[]int{0, 3, 5, 8, 1},
			},
		},

		"case3": {
			in: [][]int{
				[]int{1, 3, 5, 8, 1},
				[]int{1, 0, 5, 8, 1},
				[]int{1, 3, 0, 8, 1},
				[]int{1, 3, 5, 8, 1},
				[]int{1, 3, 5, 8, 1},
			},
			post: [][]int{
				[]int{1, 0, 0, 8, 1},
				[]int{0, 0, 0, 0, 0},
				[]int{0, 0, 0, 0, 0},
				[]int{1, 0, 0, 8, 1},
				[]int{1, 0, 0, 8, 1},
			},
		},
	}

	for k, test := range tests {
		t.Run(k, func(t *testing.T) {
			ZeroMatrix(test.in)
			if !reflect.DeepEqual(test.in, test.post) {
				t.Errorf("diff=%v", pretty.Diff(test.in, test.post))
			}
		})
	}
}
