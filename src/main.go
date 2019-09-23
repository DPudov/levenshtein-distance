package main

import "levenshtein-distance/levenshtein"

func main() {
	f := "12354"
	s := "13245"
	println(levenshtein.LevenshteinRecursive(f, s))
	println(levenshtein.LevenshteinIterative(f, s))
	println(levenshtein.LevenshteinDamerau(f, s))
}
