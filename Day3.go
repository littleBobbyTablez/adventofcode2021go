package main

import (
	"fmt"
	"strconv"
)

func fineGammaAndEpsilonRate(binaryStrings []string) {
	asInts := parseInput(binaryStrings)

	result := countOnes(asInts)

	breakPoint := len(asInts) / 2
	gammaList := [12]int{0}
	for j, k := range result {
		if k > breakPoint {
			gammaList[j] = 1
		}
	}

	epsilonList := [12]int{0}
	for h, g := range result {
		if g < breakPoint {
			epsilonList[h] = 1
		}
	}

	gammaRate := intSliceToNumber(gammaList)
	fmt.Println(gammaRate)
	epsilonRate := intSliceToNumber(epsilonList)
	fmt.Println(epsilonRate)
	fmt.Println(gammaRate * epsilonRate)
}

func countOnes(asInts [][12]int) [12]int {
	result := [12]int{0}

	for _, l := range asInts {
		for i, v := range l {
			result[i] += v
		}
	}
	return result
}

func findOxyAndCo2Rate(binaryStrings []string) {
	asInts := parseInput(binaryStrings)

	oxyList := filterRecOxy(asInts, 0)
	co2List := filterRecCo2(asInts, 0)
	oxyRange := intSliceToNumber(oxyList)
	fmt.Println(oxyRange)
	co2Range := intSliceToNumber(co2List)
	fmt.Println(co2Range)
	fmt.Println(oxyRange * co2Range)

}

func filterRecOxy(ints [][12]int, index int) [12]int {
	size := len(ints)
	if size == 1 {
		return ints[0]
	}

	ones := countOnes(ints)
	breakPoint := (size / 2) + (size % 2)
	toFilter := findNumberToFilterOxy(ones, index, breakPoint)

	var filtered [][12]int
	for _, v := range ints {
		if v[index] == toFilter {
			filtered = append(filtered, v)
		}
	}
	return filterRecOxy(filtered, index+1)
}

func filterRecCo2(ints [][12]int, index int) [12]int {
	size := len(ints)
	if size == 1 {
		return ints[0]
	}

	ones := countOnes(ints)
	breakPoint := (size / 2) + (size % 2)
	toFilter := findNumberToFilterCo2(ones, index, breakPoint)

	var filtered [][12]int
	for _, v := range ints {
		if v[index] == toFilter {
			filtered = append(filtered, v)
		}
	}
	return filterRecCo2(filtered, index+1)
}

func findNumberToFilterOxy(ones [12]int, index int, breakPoint int) int {
	if ones[index] >= breakPoint {
		return 1
	}
	return 0
}

func findNumberToFilterCo2(ones [12]int, index int, breakPoint int) int {
	if ones[index] < breakPoint {
		return 1
	}
	return 0
}

func intSliceToNumber(ints [12]int) int {
	var s string
	for _, v := range ints {
		s += strconv.Itoa(v)
	}

	parseInt, err := strconv.ParseInt(s, 2, 32)
	if err != nil {
		fmt.Println("Error converting String!")
	}
	return int(parseInt)
}

func parseInput(binaryStrings []string) [][12]int {
	var asRunes [][]rune
	for _, s := range binaryStrings {
		asRunes = append(asRunes, []rune(s))
	}

	var asInts [][12]int
	for _, l := range asRunes {
		ints := [12]int{0}
		for i, r := range l {
			ints[i] = int(r - '0')
		}
		asInts = append(asInts, ints)
	}
	return asInts
}
