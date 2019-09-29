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

func levenshteinRecursiveModule(first, second string, result, minval int) int {
	lenFirst := len(first)
	lenSecond := len(second)
	if result >= minval {
		return minval
	} else if lenFirst == 0 {
		return result +lenSecond
	} else if lenSecond == 0 {
		return result + lenFirst
	} else {
		add := 1
		if first[lenFirst-1] == second[lenSecond-1] {
			add = 0
		}
		r1 := levenshteinRecursiveModule(first[:lenFirst-1], second[:lenSecond-1], result+add, minval)
		r2 := levenshteinRecursiveModule(first, second[:lenSecond-1], result+1, Min(r1, minval))
		r3 := levenshteinRecursiveModule(first[:lenFirst-1], second, result+1, MinThree(r1, r2, minval))
		return MinThree(r1, r2, r3)
	}
}

func LevenshteinRecursiveOptimized(first string, second string) int {
	minval := Max(len(first), len(second))
	result := 0
	return levenshteinRecursiveModule(first, second, result, minval)
}
