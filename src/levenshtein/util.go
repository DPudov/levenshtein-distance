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

func PrintMatrix(matrix [][]int) {
	for i := range matrix {
		for j := range matrix[i] {
			print(matrix[i][j], " ")
		}
		println()
	}
}

func PrintMatrixWithStrings(first string, second string, matrix [][]int) {
	print("s1/s2 ")
	f := []rune(first)
	s := []rune(second)
	for i := range s {
		print(string(s[i]), " ")
	}
	println()
	for i := range matrix {
		if i > 0 {
			print(string(f[i-1]), "   ")
		} else {
			print("    ")
		}
		for j := range matrix[i] {
			print(matrix[i][j], " ")
		}
		println()
	}
}
