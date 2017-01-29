package main

import "testing"

func Test(t *testing.T) {
	tests := map[string]struct {
		in, out string
	}{
		"compress-case1":     {"aaabbbcccddd", "a3b3c3d3"},
		"compress-case2":     {"aaabbbcccd", "a3b3c3d1"},
		"compress-case3":     {"abbbcccddd", "a1b3c3d3"},
		"compress-case4":     {"aaaaaaaaaaabb", "a11b2"},
		"not_compress-case1": {"abcd", "abcd"},
		"not_compress-case2": {"aabbccdd", "aabbccdd"},
	}

	for k, test := range tests {
		t.Run(k, func(t *testing.T) {
			out := StringCompression(test.in)
			if out != test.out {
				t.Errorf("actual=%s expected=%s", out, test.out)
			}
		})
	}
}
