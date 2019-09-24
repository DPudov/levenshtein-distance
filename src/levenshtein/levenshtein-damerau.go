package levenshtein

func LevenshteinDamerau(first string, second string) (int, [][] int) {
	lenFirst := len(first)
	lenSecond := len(second)
	rows := lenFirst + 1
	cols := lenSecond + 1
	matrix := allocateMatrix(rows, cols)

	for i := 0; i < rows; i++ {
		matrix[i][0] = i
	}
	for i := 1; i < cols; i++ {
		matrix[0][i] = i
	}

	add := 0
	for j := 1; j < cols; j++ {
		for i := 1; i < rows; i++ {
			if first[i-1] == second[j-1] {
				add = 0
			} else {
				add = 1
			}
			matrix[i][j] = MinThree(matrix[i-1][j]+1,
				matrix[i][j-1]+1,
				matrix[i-1][j-1]+add)

			if i > 1 && j > 1 && first[i-1] == second[j-2] && first[i-2] == second[j-1] {
				matrix[i][j] = Min(matrix[i][j], matrix[i-2][j-2]+add)
			}
		}
	}

	return matrix[lenFirst][lenSecond], matrix
}
