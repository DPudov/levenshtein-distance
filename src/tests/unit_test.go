package tests

import (
	"levenshtein-distance/levenshtein"
	"testing"
)

func typicalTestLevenshteinRecursive(t *testing.T, s1, s2 string, expected int) {
	got := levenshtein.LevenshteinRecursive(s1, s2)
	if got != expected {
		t.Errorf("Result from %s, %s = %d; want %d", s1, s2, got, expected)
	}
}


func typicalTestLevenshteinIterative(t *testing.T, s1, s2 string, expected int) {
	got, _ := levenshtein.LevenshteinIterative(s1, s2)
	if got != expected {
		t.Errorf("Result from %s, %s = %d; want %d", s1, s2, got, expected)
	}
}

func typicalTestLevenshteinDamerau(t *testing.T, s1, s2 string, expected int) {
	got, _ := levenshtein.LevenshteinDamerau(s1, s2)
	if got != expected {
		t.Errorf("Result from %s, %s = %d; want %d", s1, s2, got, expected)
	}
}


func TestLevenshteinRecursiveEmpty(t *testing.T) {
	s1 := ""
	s2 := ""
	expected := 0
	typicalTestLevenshteinRecursive(t, s1, s2, expected)
}

func TestLevenshteinRecursiveEqual(t *testing.T) {
	s1 := "abc"
	s2 := "abc"
	expected := 0
	typicalTestLevenshteinRecursive(t, s1, s2, expected)
}

func TestLevenshteinRecursiveEqualLen(t *testing.T) {
	s1 := "abc"
	s2 := "def"
	expected := 3
	typicalTestLevenshteinRecursive(t, s1, s2, expected)
}

func TestLevenshteinRecursiveInsert(t *testing.T) {
	s1 := "abc"
	s2 := "abcdef"
	expected := 3
	typicalTestLevenshteinRecursive(t, s1, s2, expected)
}

func TestLevenshteinRecursiveDelete(t *testing.T) {
	s1 := "abcdef"
	s2 := "abc"
	expected := 3
	typicalTestLevenshteinRecursive(t, s1, s2, expected)
}

func TestLevenshteinRecursiveInsertDelete(t *testing.T) {
	s1 := "xyzabc"
	s2 := "abcdef"
	expected := 6
	typicalTestLevenshteinRecursive(t, s1, s2, expected)
}

func TestLevenshteinRecursiveReplace(t *testing.T) {
	s1 := "abkdef"
	s2 := "abcdefg"
	expected := 2
	typicalTestLevenshteinRecursive(t, s1, s2, expected)
}

func TestLevenshteinRecursiveTransposition(t *testing.T) {
	s1 := "abc"
	s2 := "acb"
	expected := 2
	typicalTestLevenshteinRecursive(t, s1, s2, expected)
}

// Levenshtein-iterative
func TestLevenshteinIterativeEmpty(t *testing.T) {
	s1 := ""
	s2 := ""
	expected := 0
	typicalTestLevenshteinIterative(t, s1, s2, expected)
}

func TestLevenshteinIterativeEqual(t *testing.T) {
	s1 := "abc"
	s2 := "abc"
	expected := 0
	typicalTestLevenshteinIterative(t, s1, s2, expected)
}

func TestLevenshteinIterativeEqualLen(t *testing.T) {
	s1 := "abc"
	s2 := "def"
	expected := 3
	typicalTestLevenshteinIterative(t, s1, s2, expected)
}

func TestLevenshteinIterativeInsert(t *testing.T) {
	s1 := "abc"
	s2 := "abcdef"
	expected := 3
	typicalTestLevenshteinIterative(t, s1, s2, expected)
}

func TestLevenshteinIterativeDelete(t *testing.T) {
	s1 := "abcdef"
	s2 := "abc"
	expected := 3
	typicalTestLevenshteinIterative(t, s1, s2, expected)
}

func TestLevenshteinIterativeInsertDelete(t *testing.T) {
	s1 := "xyzabc"
	s2 := "abcdef"
	expected := 6
	typicalTestLevenshteinIterative(t, s1, s2, expected)
}

func TestLevenshteinIterativeReplace(t *testing.T) {
	s1 := "abkdef"
	s2 := "abcdefg"
	expected := 2
	typicalTestLevenshteinIterative(t, s1, s2, expected)
}

func TestLevenshteinIterativeTransposition(t *testing.T) {
	s1 := "abc"
	s2 := "acb"
	expected := 2
	typicalTestLevenshteinIterative(t, s1, s2, expected)
}

// Levenshtein-Damerau
func TestLevenshteinDamerauEmpty(t *testing.T) {
	s1 := ""
	s2 := ""
	expected := 0
	typicalTestLevenshteinDamerau(t, s1, s2, expected)
}

func TestLevenshteinDamerauEqual(t *testing.T) {
	s1 := "abc"
	s2 := "abc"
	expected := 0
	typicalTestLevenshteinDamerau(t, s1, s2, expected)
}

func TestLevenshteinDamerauEqualLen(t *testing.T) {
	s1 := "abc"
	s2 := "def"
	expected := 3
	typicalTestLevenshteinDamerau(t, s1, s2, expected)
}

func TestLevenshteinDamerauInsert(t *testing.T) {
	s1 := "abc"
	s2 := "abcdef"
	expected := 3
	typicalTestLevenshteinDamerau(t, s1, s2, expected)
}

func TestLevenshteinDamerauDelete(t *testing.T) {
	s1 := "abcdef"
	s2 := "abc"
	expected := 3
	typicalTestLevenshteinDamerau(t, s1, s2, expected)
}

func TestLevenshteinDamerauInsertDelete(t *testing.T) {
	s1 := "xyzabc"
	s2 := "abcdef"
	expected := 6
	typicalTestLevenshteinDamerau(t, s1, s2, expected)
}

func TestLevenshteinDamerauReplace(t *testing.T) {
	s1 := "abkdef"
	s2 := "abcdefg"
	expected := 2
	typicalTestLevenshteinDamerau(t, s1, s2, expected)
}

func TestLevenshteinDamerauTransposition(t *testing.T) {
	s1 := "abc"
	s2 := "acb"
	expected := 1
	typicalTestLevenshteinDamerau(t, s1, s2, expected)
}