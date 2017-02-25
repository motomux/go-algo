package solution

import (
	"reflect"
	"testing"

	"github.com/motomux/pretty"
)

func TestAdd(t *testing.T) {
	tests := map[string]struct {
		new  *Intervals
		in   Interval
		post *Intervals
	}{
		"should add non-overlapped interval": {
			new: &Intervals{
				list: []Interval{
					Interval{1, 5},
				},
			},
			in: Interval{10, 15},
			post: &Intervals{
				list: []Interval{
					Interval{1, 5},
					Interval{10, 15},
				},
			},
		},

		"should add when list is empty": {
			new: &Intervals{
				list: []Interval{},
			},
			in: Interval{10, 15},
			post: &Intervals{
				list: []Interval{
					Interval{10, 15},
				},
			},
		},

		"should not add when there is the same interval": {
			new: &Intervals{
				list: []Interval{
					Interval{10, 15},
				},
			},
			in: Interval{10, 15},
			post: &Intervals{
				list: []Interval{
					Interval{10, 15},
				},
			},
		},

		"should not add when there is a interval which covers given interval": {
			new: &Intervals{
				list: []Interval{
					Interval{10, 15},
				},
			},
			in: Interval{12, 13},
			post: &Intervals{
				list: []Interval{
					Interval{10, 15},
				},
			},
		},

		"should add and merge given interval with existing intervals when there are overlapped intervals": {
			new: &Intervals{
				list: []Interval{
					Interval{10, 15},
					Interval{25, 30},
					Interval{35, 45},
					Interval{55, 65},
					Interval{65, 75},
				},
			},
			in: Interval{20, 50},
			post: &Intervals{
				list: []Interval{
					Interval{10, 15},
					Interval{55, 65},
					Interval{65, 75},
					Interval{20, 50},
				},
			},
		},
	}

	for k, test := range tests {
		t.Run(k, func(t *testing.T) {
			intervals := test.new
			intervals.Add(test.in)
			if !reflect.DeepEqual(intervals, test.post) {
				t.Errorf("diff=%v", pretty.Diff(intervals, test.post))
			}
		})
	}
}
