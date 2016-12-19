package leetcode

// 6. ZigZag Conversion
func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}

	strs := make([][]rune, numRows)
	diff := numRows - 2

	for i, c := range s {
		p := i % (numRows + diff)
		if p < numRows {
			strs[p] = append(strs[p], c)
		} else {
			p = (p - numRows - diff) * -1
			strs[p] = append(strs[p], c)
		}
	}

	res := ""
	for _, r := range strs {
		res += string(r)
	}
	return res
}
