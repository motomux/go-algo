package main

import (
	"reflect"
	"testing"

	"github.com/kr/pretty"
)

func Test(t *testing.T) {
	tests := map[string]struct {
		in, post *Node
	}{
		"no_dup-1": {
			in:   &Node{1, &Node{2, &Node{3, &Node{4, &Node{5, nil}}}}},
			post: &Node{1, &Node{2, &Node{3, &Node{4, &Node{5, nil}}}}},
		},
		"dup-1": {
			in:   &Node{1, &Node{1, &Node{3, &Node{4, &Node{5, nil}}}}},
			post: &Node{1, &Node{3, &Node{4, &Node{5, nil}}}},
		},
		"last_dup-1": {
			in:   &Node{1, &Node{2, &Node{3, &Node{4, &Node{1, nil}}}}},
			post: &Node{1, &Node{2, &Node{3, &Node{4, nil}}}},
		},
		"all_dup-1": {
			in:   &Node{1, &Node{1, &Node{1, &Node{1, &Node{1, nil}}}}},
			post: &Node{1, nil},
		},
		"empty-1": {
			in:   nil,
			post: nil,
		},
	}

	for k, test := range tests {
		t.Run(k, func(t *testing.T) {
			RemoveDup(test.in)
			if !reflect.DeepEqual(test.in, test.post) {
				t.Errorf("diff=%v", pretty.Diff(test.in, test.post))
			}
		})
	}
}
