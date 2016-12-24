package leetcode

// 246. Strobogrammatic Number
func isStrobogrammatic(num string) bool {
	for i, j := 0, len(num)-1; i <= j; i, j = i+1, j-1 {
		n := string([]byte{num[i], num[j]})
		switch n {
		case "00":
		case "11":
		case "88":
		case "69":
		case "96":
		default:
			return false
		}
	}

	return true
}
