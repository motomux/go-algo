package main

import (
	"reflect"
	"testing"

	"github.com/kr/pretty"
)

func TestURLify(t *testing.T) {
	tests := map[string]struct {
		s      []byte
		length int
		post   []byte
	}{
		"case01": {[]byte("hello world  "), 11, []byte("hello%20world")},
		"case02": {[]byte("hello world from Golang      "), 23, []byte("hello%20world%20from%20Golang")},
		"case03": {[]byte(""), 0, []byte("")},
	}

	for k, test := range tests {
		t.Run(k, func(t *testing.T) {
			URLify(test.s, test.length)
			if !reflect.DeepEqual(test.s, test.post) {
				t.Errorf("diff=%v", pretty.Diff(test.s, test.post))
			}
		})
	}
}
