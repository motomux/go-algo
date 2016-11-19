package main

import (
	"reflect"
	"testing"
)

func TestInsertionSort(t *testing.T) {
	tests := map[string]struct {
		in  []int
		out []int
	}{
		"case1": {
			in:  []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			out: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		"case2": {
			in:  []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0},
			out: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		"case3": {
			in:  []int{4, 6, 2, 1, 8, 9, 0, 7, 5, 3},
			out: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
	}

	for k, test := range tests {
		t.Run(k, func(t *testing.T) {
			InsertionSort(test.in)
			if !reflect.DeepEqual(test.in, test.out) {
				t.Errorf("actual: %v, expected: %v", test.in, test.out)
			}
		})
	}
}
