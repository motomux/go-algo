package main

import "testing"

func Test(t *testing.T) {
	tests := map[string]struct {
		s1, s2 string
		out    bool
	}{
		"case1": {"abcdef", "abcdef", true},
		"case2": {"abcdef", "efabcd", true},
		"case3": {"abcdef", "abcde", false},
		"case4": {"abcdef", "abcdeff", false},
		"case5": {"abcdef", "fbcdea", false},
	}

	for k, test := range tests {
		t.Run(k, func(t *testing.T) {
			out := StringRotation(test.s1, test.s2)
			if out != test.out {
				t.Errorf("actual=%t expected=%t", out, test.out)
			}
		})
	}
}
