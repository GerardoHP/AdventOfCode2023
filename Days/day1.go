package Days

func Day1(codes []string) int {
	result := 0
	for _, v := range codes {
		result += decode(v)
	}

	return result
}

func decode(code string) int {
	a := make([]int, 0)
	for _, v := range code {
		if isInt, n := runeToInt(v); isInt {
			if len(a) < 2 {
				a = append(a, n)
			} else {
				a[1] = n
			}
		}
	}

	if len(a) == 0 {
		return 0
	}

	if len(a) == 1 {
		return a[0]*10 + a[0]
	}

	return a[0]*10 + a[1]
}

func runeToInt(r rune) (bool, int) {
	str := string(r)
	switch str {
	case "0":
		return true, 0
	case "1":
		return true, 1
	case "2":
		return true, 2
	case "3":
		return true, 3
	case "4":
		return true, 4
	case "5":
		return true, 5
	case "6":
		return true, 6
	case "7":
		return true, 7
	case "8":
		return true, 8
	case "9":
		return true, 9
	default:
		return false, -1
	}
}
