package linkedlist

import (
	"reflect"
	"testing"

	"github.com/kr/pretty"
)

func Test(t *testing.T) {
	tests := map[string]struct {
		node *Node
		k    int
		out  *Node
	}{
		"last 1st": {
			node: &Node{1, &Node{2, &Node{3, &Node{4, &Node{5, nil}}}}},
			k:    1,
			out:  &Node{5, nil},
		},
		"last 3rd": {
			node: &Node{1, &Node{2, &Node{3, &Node{4, &Node{5, nil}}}}},
			k:    3,
			out:  &Node{3, &Node{4, &Node{5, nil}}},
		},
		"last 5th": {
			node: &Node{1, &Node{2, &Node{3, &Node{4, &Node{5, nil}}}}},
			k:    5,
			out:  &Node{1, &Node{2, &Node{3, &Node{4, &Node{5, nil}}}}},
		},
		"last 0": {
			node: &Node{1, &Node{2, &Node{3, &Node{4, &Node{5, nil}}}}},
			k:    0,
			out:  nil,
		},
		"k is greater than length of list": {
			node: &Node{1, &Node{2, &Node{3, &Node{4, &Node{5, nil}}}}},
			k:    6,
			out:  nil,
		},
	}

	for k, test := range tests {
		t.Run(k, func(t *testing.T) {
			out := LastKthNode(test.node, test.k)
			if !reflect.DeepEqual(out, test.out) {
				t.Errorf("diff=%v", pretty.Diff(out, test.out))
			}
		})
	}
}
