package main

import "testing"

func Test(t *testing.T) {
	tests := map[string]struct {
		s1, s2 string
		out    bool
	}{
		"same-case1":          {s1: "abcdef", s2: "abcdef", out: true},
		"add-case1":           {s1: "abcde", s2: "abcdef", out: true},
		"add-case2":           {s1: "abcdef", s2: "abcde", out: true},
		"add-case3":           {s1: "bcdef", s2: "abcdef", out: true},
		"add-case4":           {s1: "abcdef", s2: "bcdef", out: true},
		"add-case5":           {s1: "abdef", s2: "abcdef", out: true},
		"add-case6":           {s1: "abcdef", s2: "abdef", out: true},
		"add-case7":           {s1: "abcdef", s2: "abdeg", out: false},
		"replace-case1":       {s1: "abcdff", s2: "abcdef", out: true},
		"replace-case2":       {s1: "abcdef", s2: "abcdff", out: true},
		"replace-case3":       {s1: "abcdef", s2: "abcdfe", out: false},
		"more-than-two-case1": {s1: "abcd", s2: "abcdef", out: false},
	}

	for k, test := range tests {
		t.Run(k, func(t *testing.T) {
			out := CheckOneAway(test.s1, test.s2)
			if out != test.out {
				t.Errorf("actual=%t, expected=%t", out, test.out)
			}
		})
	}
}
