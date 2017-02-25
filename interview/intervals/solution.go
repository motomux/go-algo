package solution

// Interval represents interval
type Interval struct {
	Start, End int
}

// Intervals represents slice of interval
type Intervals struct {
	list []Interval
}

// Add adds interval to intervals
func (intervals *Intervals) Add(newIntvl Interval) {
	var newList []Interval
	for _, intvl := range intervals.list {
		if intvl.End < newIntvl.Start || intvl.Start > newIntvl.End {
			newList = append(newList, intvl)
		} else {
			newIntvl.Start = min(newIntvl.Start, intvl.Start)
			newIntvl.End = max(newIntvl.End, intvl.End)
		}
	}
	newList = append(newList, newIntvl)
	intervals.list = newList
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
