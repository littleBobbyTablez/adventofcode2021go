package main

func countIncreases(ints []int) int {
	acc := 0
	return countInceasesRec(ints, acc)
}

func countInceasesRec(ints []int, acc int) int {
	if len(ints) < 2 {
		return acc
	} else {
		if ints[0] < ints[1] {
			acc += 1
		}
		return countInceasesRec(ints[1:], acc)
	}
}

func sumOfThree(ints []int) []int {
	var result []int
	return sumOfThreeRec(ints, result)
}

func sumOfThreeRec(ints []int, result []int) []int {
	if len(ints) <= 2 {
		newResult := append(result, ints[0]+ints[1])
		return newResult
	} else {
		newResult := append(result, ints[0]+ints[1]+ints[2])
		return sumOfThreeRec(ints[1:], newResult)
	}
}
