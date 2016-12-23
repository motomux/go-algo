package leetcode

// 246. Strobogrammatic Number
func isStrobogrammatic(num string) bool {
	for i := 0; i < len(num)/2; i++ {
		n1, n2 := num[i], num[len(num)-1-i]
		switch {
		case n1 == '9' && n2 == '6':
		case n1 == '6' && n2 == '9':
		case n1 != '8' && n1 != '1' && n1 != '0':
			return false
		case n1 != n2:
			return false
		}
	}

	if len(num)%2 == 0 {
		return true
	}
	m := num[len(num)/2]
	switch {
	case m == '0' || m == '1' || m == '8':
		return true
	default:
		return false
	}
}
