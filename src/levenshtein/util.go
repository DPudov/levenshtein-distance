package levenshtein

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func MinThree(a, b, c int) int {
	if a < b {
		if a < c {
			return a
		}
		return c
	} else {
		if b < c {
			return b
		}
		return c
	}
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
