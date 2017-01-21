package main

import "testing"

func TestCheckPlindromePermutation(t *testing.T) {
	tests := map[string]struct {
		s   string
		out bool
	}{
		"case01": {"helleh", true},
		"case02": {"hellieh", true},
		"case03": {"helehl", true},
		"case04": {"hele", false},
		"case05": {"hello world", false},
		"case06": {"", true},
	}

	for k, test := range tests {
		t.Run(k, func(t *testing.T) {
			out := CheckPlindromePermutation(test.s)
			if out != test.out {
				t.Errorf("actual=%t expected=%t", out, test.out)
			}
		})
	}
}

func TestCheckPlindromePermutation2(t *testing.T) {
	tests := map[string]struct {
		s   string
		out bool
	}{
		"case01": {"helleh", true},
		"case02": {"hellieh", true},
		"case03": {"helehl", true},
		"case04": {"hele", false},
		"case05": {"hello world", false},
		"case06": {"", true},
	}

	for k, test := range tests {
		t.Run(k, func(t *testing.T) {
			out := CheckPlindromePermutation2(test.s)
			if out != test.out {
				t.Errorf("actual=%t expected=%t", out, test.out)
			}
		})
	}
}
