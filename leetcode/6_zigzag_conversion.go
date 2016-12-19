package leetcode

// 6. ZigZag Conversion
func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}

	bs := make([][]byte, numRows)

	i := 0
	for i < len(s) {
		for j := 0; i < len(s) && j < numRows; j++ {
			bs[j] = append(bs[j], s[i])
			i++
		}
		for j := numRows - 2; i < len(s) && j > 0; j-- {
			bs[j] = append(bs[j], s[i])
			i++
		}
	}

	var res []byte
	for _, b := range bs {
		res = append(res, b...)
	}
	return string(res)
}
