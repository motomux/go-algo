package main

import "testing"

func TestBinarySearch(t *testing.T) {
	type (
		in struct {
			nums []int
			num  int
		}
		out struct {
			index int
			err   error
		}
	)

	tests := map[string]struct {
		in
		out
	}{
		"case1": {
			in{[]int{1, 2, 3, 4, 5}, 2},
			out{1, nil},
		},
		"case2": {
			in{[]int{1, 2, 3, 4, 5}, 1},
			out{0, nil},
		},
		"case3": {
			in{[]int{1, 2, 3, 4, 5}, 5},
			out{4, nil},
		},
		"case4": {
			in{[]int{1, 2, 3, 4, 5}, 6},
			out{-1, ErrNotFound},
		},
		"case5": {
			in{[]int{1, 2, 3, 4, 5}, -1},
			out{-1, ErrNotFound},
		},
	}

	for k, test := range tests {
		t.Run(k, func(t *testing.T) {
			in, out := test.in, test.out
			index, err := BinarySearch(in.nums, in.num)
			if err != out.err {
				t.Errorf("actual: %v, expected: %v", err, out.err)
			}
			if index != out.index {
				t.Errorf("actual: %v, expected: %v", index, out.index)
			}
		})
	}
}
