package main

import (
	"sort"
)

func findFirstCorruptElement(input []string) int {
	errorValues := initErrorValueMap()

	sum := 0
	for _, line := range input {
		firstError := findFirstError(line)
		sum += errorValues[firstError]
	}

	return sum
}

func calculateUnfinishedScore(input []string) int {
	var scores []int

	for _, s := range input {
		scores = append(scores, findUnfinished(s))
	}

	zeroCounter := 0
	for _, s := range scores {
		if s == 0 {
			zeroCounter++
		}
	}

	sort.Ints(scores)
	noZeros := scores[zeroCounter:]

	return noZeros[len(noZeros)/2]
}

func findUnfinished(line string) int {
	var stack []string
	completionMap := initCompletionValueMap()
	noErrors := true

	for _, s := range line {
		x := string(s)
		if isOpening(x) {
			stack = append(stack, x)
		} else {
			l := len(stack) - 1
			n := stack[l]
			if closingFitsOpening(x, n) {
				stack = stack[:l]
			} else {
				noErrors = false
			}
		}
	}

	var reversed []string
	if len(stack) > 0 && noErrors {
		reversed = reverseSlice(stack)
	}

	return calculateScore(reversed, completionMap)
}

func calculateScore(reversed []string, completionMap map[string]int) int {
	total := 0

	for _, s := range reversed {
		total = (total * 5) + completionMap[s]
	}

	return total
}

func reverseSlice(s []string) []string {
	var output []string

	for i := len(s) - 1; i >= 0; i-- {
		output = append(output, s[i])
	}

	return output
}

func findFirstError(line string) string {
	var stack []string

	for _, s := range line {
		x := string(s)
		if isOpening(x) {
			stack = append(stack, x)
		} else {
			l := len(stack) - 1
			n := stack[l]
			if closingFitsOpening(x, n) {
				stack = stack[:l]
			} else {
				return x
			}
		}
	}
	return ""
}

func closingFitsOpening(x string, n string) bool {
	switch {
	case x == ")" && n == "(":
		return true
	case x == "]" && n == "[":
		return true
	case x == "}" && n == "{":
		return true
	case x == ">" && n == "<":
		return true
	default:
		return false
	}
}

func isOpening(x string) bool {
	return x == "(" || x == "[" || x == "{" || x == "<"
}

func initErrorValueMap() map[string]int {
	errorValues := make(map[string]int)

	errorValues[""] = 0
	errorValues[")"] = 3
	errorValues["]"] = 57
	errorValues["}"] = 1197
	errorValues[">"] = 25137

	return errorValues
}

func initCompletionValueMap() map[string]int {
	completionValues := make(map[string]int)

	completionValues["("] = 1
	completionValues["["] = 2
	completionValues["{"] = 3
	completionValues["<"] = 4

	return completionValues
}
