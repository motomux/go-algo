package leetcode

// 57. Insert Interval
func insert(intervals []Interval, newInterval Interval) []Interval {
	if intervals == nil || len(intervals) == 0 {
		return []Interval{newInterval}
	}

	added := false
	res := make([]Interval, 0, len(intervals)+1)
	for _, interval := range intervals {
		if added || interval.End < newInterval.Start {
			res = append(res, interval)
		} else if newInterval.End < interval.Start {
			res = append(res, newInterval, interval)
			added = true
		} else {
			start := interval.Start
			end := interval.End
			if newInterval.Start < interval.Start {
				start = newInterval.Start
			}
			if newInterval.End > interval.End {
				end = newInterval.End
			}
			newInterval.Start = start
			newInterval.End = end
		}
	}

	if !added {
		res = append(res, newInterval)
	}

	return res
}
