package main

import (
	"bufio"
	"fmt"
	"levenshtein-distance/levenshtein"
	"os"
)

const (
	once      = "once"
	many      = "many"
	visualize = "visualize"
)

func takeInputString(in *bufio.Reader) string {
	fmt.Println("Please input string:")
	result, err := in.ReadString('\n')
	for err != nil {
		fmt.Println("Incorrect input. Try again...")
		result, err = in.ReadString('\n')
	}
	result = result[:len(result)-1]
	return result
}

func handleOnce(visualize bool, in *bufio.Reader) {
	first := takeInputString(in)
	second := takeInputString(in)
	fmt.Println("Result is...")
	fmt.Println("Strings:", first, second)
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

func handleMany(visualize bool, in *bufio.Reader) {
	for {
		handleOnce(visualize, in)
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
	in := bufio.NewReader(os.Stdin)

	l := len(argsWithoutProgram)
	if l > 0 {
		z := argsWithoutProgram[0]
		if l > 1 {
			f := argsWithoutProgram[1]
			if z == once && f == visualize {
				handleOnce(true, in)
			} else if z == many && f == visualize {
				handleMany(true, in)
			} else {
				printUsage()
			}
		} else {
			if z == once {
				handleOnce(false, in)
			} else if z == many {
				handleMany(false, in)
			} else {
				printUsage()
			}
		}
	} else {
		printUsage()
	}

}
