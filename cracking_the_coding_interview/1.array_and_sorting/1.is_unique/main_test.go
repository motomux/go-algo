package main

import "testing"

func TestIsUnique(t *testing.T) {
	tests := map[string]struct {
		input  string
		output bool
	}{
		"case01": {"abcdefg", true},
		"case02": {"Aa", true},
		"case03": {"aacdefg", false},
		"case04": {"abcdefgg", false},
		"case05": {"abcdefga", false},
	}

	for k, test := range tests {
		t.Run(k, func(t *testing.T) {
			output := IsUnique(test.input)
			if output != test.output {
				t.Errorf("actual=%t expected=%t", output, test.output)
			}
		})
	}
}

func TestIsUnique2(t *testing.T) {
	tests := map[string]struct {
		input  string
		output bool
	}{
		"case01": {"abcdefg", true},
		"case02": {"Aa", true},
		"case03": {"aacdefg", false},
		"case04": {"abcdefgg", false},
		"case05": {"abcdefga", false},
	}

	for k, test := range tests {
		t.Run(k, func(t *testing.T) {
			output := IsUnique2(test.input)
			if output != test.output {
				t.Errorf("actual=%t expected=%t", output, test.output)
			}
		})
	}
}
