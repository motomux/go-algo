package main

import "testing"

func TestCheckPermutation(t *testing.T) {
	tests := map[string]struct {
		s1, s2 string
		out    bool
	}{
		"case01": {"hello", "ehllo", true},
		"case02": {"hello", "olleh", true},
		"case03": {"hello", "hllo", false},
		"case04": {"hello", "heloo", false},
	}

	for k, test := range tests {
		t.Run(k, func(t *testing.T) {
			out := CheckPermutation(test.s1, test.s2)
			if out != test.out {
				t.Errorf("actual=%t expected=%t", out, test.out)
			}
		})
	}
}

func TestCheckPermutation2(t *testing.T) {
	tests := map[string]struct {
		s1, s2 string
		out    bool
	}{
		"case01": {"hello", "ehllo", true},
		"case02": {"hello", "olleh", true},
		"case03": {"hello", "hllo", false},
		"case04": {"hello", "heloo", false},
	}

	for k, test := range tests {
		t.Run(k, func(t *testing.T) {
			out := CheckPermutation2(test.s1, test.s2)
			if out != test.out {
				t.Errorf("actual=%t expected=%t", out, test.out)
			}
		})
	}
}
