package levenshtein

func LevenshteinRecursive(first string, second string) int {
	lenFirst := len(first)
	lenSecond := len(second)
	if Min(lenFirst, lenSecond) == 0 {
		return Max(lenFirst, lenSecond)
	} else {
		add := 1
		if first[lenFirst-1] == second[lenSecond-1] {
			add = 0
		}

		return MinThree(
			LevenshteinRecursive(first[:lenFirst-1], second)+1,
			LevenshteinRecursive(first, second[:lenSecond-1])+1,
			LevenshteinRecursive(first[:lenFirst-1], second[:lenSecond-1])+add)
	}
}
