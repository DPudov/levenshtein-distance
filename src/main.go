package main

import (
	"fmt"
	"levenshtein-distance/levenshtein"
	"os"
)

const (
	once      = "once"
	many      = "many"
	visualize = "visualize"
)

func takeInputString() string {
	fmt.Println("Please input string:")
	var result string
	_, err := fmt.Scanln(&result)
	for err != nil {
		fmt.Println("Incorrect input. Try again...")
		_, err = fmt.Scanln(&result)
	}
	return result
}

func handleOnce(visualize bool) {
	first := takeInputString()
	second := takeInputString()
	fmt.Println("Result is...")
	recursive := levenshtein.LevenshteinRecursive(first, second)
	iterative, matrixIt := levenshtein.LevenshteinIterative(first, second)
	damerau, matrixDa := levenshtein.LevenshteinDamerau(first, second)
	fmt.Println("For Levenshtein recursive: ", recursive)
	fmt.Println("For Levenshtein iterative: ", iterative)
	if visualize {
		levenshtein.PrintMatrixWithStrings(first, second, matrixIt)
	}
	fmt.Println("For Levenshtein-Damerau: ", damerau)
	if visualize {
		levenshtein.PrintMatrixWithStrings(first, second, matrixDa)
	}
}

func handleMany(visualize bool) {
	for {
		handleOnce(visualize)
	}
}

func printUsage() {
	fmt.Println("You may run program with those arguments:\n" +
		"once Input two strings, get result. May be combined with 'visualize'\n" +
		"visualize Prints result matrices\n" +
		"many Takes several inputs until kill signal")
}

func main() {
	argsWithoutProgram := os.Args[1:]
	l := len(argsWithoutProgram)
	if l > 0 {
		z := argsWithoutProgram[0]
		if l > 1 {
			f := argsWithoutProgram[1]
			if z == once && f == visualize {
				handleOnce(true)
			} else if z == many && f == visualize {
				handleMany(true)
			} else {
				printUsage()
			}
		} else {
			if z == once {
				handleOnce(false)
			} else if z == many {
				handleMany(false)
			} else {
				printUsage()
			}
		}
	} else {
		printUsage()
	}
}
