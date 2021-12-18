package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"sort"
	"strconv"
	"strings"
)

func readFileOctupusMap(pathToFile string) [10][10]int {
	s, err := ioutil.ReadFile(pathToFile)
	if err != nil {
		log.Fatalf("Error readinf File: %s\n", err)
	}

	input := strings.Split(string(s), "\n")
	var values [10][10]int

	for i := range input {
		str := input[i]
		split := strings.Split(str, "")

		for j, v := range split {
			n, _ := strconv.Atoi(v)
			values[i][j] = n
		}
	}
	return values
}

func readFileToRiskLevelMap(pathToFile string) map[point]int {
	s, err := ioutil.ReadFile(pathToFile)
	if err != nil {
		log.Fatalf("Error readinf File: %s\n", err)
	}

	input := strings.Split(string(s), "\n")

	output := make(map[point]int)

	for i, v := range input {
		split := strings.Split(v, "")
		for j, n := range split {
			lvl, _ := strconv.Atoi(n)
			output[point{i, j}] = lvl
		}
	}

	return output
}

func readFileInsertRules(pathToFile string) (map[string]int, map[string]string) {
	s, err := ioutil.ReadFile(pathToFile)
	if err != nil {
		log.Fatalf("Error readinf File: %s\n", err)
	}

	input := strings.Split(string(s), "\n\n")
	rules := make(map[string]string)

	ruleStrings := strings.Split(input[1], "\n")

	for _, rs := range ruleStrings {
		split := strings.Split(rs, " -> ")
		rules[split[0]] = split[1]
	}

	startingMap := make(map[string]int)
	start := input[0]
	for i := 0; i < len(start)-1; i++ {
		startingMap[string(start[i])+string(start[i+1])] += 1
	}

	return startingMap, rules
}

func readFileToIntSlice(pathToFile string) []int {
	s, err := ioutil.ReadFile(pathToFile)
	if err != nil {
		log.Fatalf("Error readinf File: %s\n", err)
	}

	input := strings.Split(string(s), "\n")
	var values []int

	for i := range input {
		str := input[i]
		number, err := strconv.Atoi(str)
		if err != nil {
			log.Fatalf("Error converting string to int: %s\n", err)
		}
		values = append(values, number)
	}
	return values
}

func readFileRowToInt(pathToFile string) []int {
	s, err := ioutil.ReadFile(pathToFile)
	if err != nil {
		log.Fatalf("Error readinf File: %s\n", err)
	}

	input := strings.Split(string(s), ",")
	var values []int

	for i := range input {
		str := input[i]
		number, err := strconv.Atoi(str)
		if err != nil {
			log.Fatalf("Error converting string to int: %s\n", err)
		}
		values = append(values, number)
	}
	return values
}

func readFileHeightMap(pathToFile string) [100][100]int {
	s, err := ioutil.ReadFile(pathToFile)
	if err != nil {
		log.Fatalf("Error readinf File: %s\n", err)
	}

	input := strings.Split(string(s), "\n")
	var values [100][100]int

	for i := range input {
		str := input[i]
		split := strings.Split(str, "")

		for j, v := range split {
			n, _ := strconv.Atoi(v)
			values[i][j] = n
		}
	}
	return values
}

func readFileDotsAndFolds(pathToFile string) (map[dot]bool, []fold) {
	s, err := ioutil.ReadFile(pathToFile)
	if err != nil {
		log.Fatalf("Error readinf File: %s\n", err)
	}

	input := strings.Split(string(s), "\n\n")
	dots := make(map[dot]bool)
	var folds []fold
	dotInput := strings.Split(input[0], "\n")

	for _, v := range dotInput {
		split := strings.Split(v, ",")
		x, _ := strconv.Atoi(split[0])
		y, _ := strconv.Atoi(split[1])
		dots[dot{x, y}] = true
	}

	foldInput := strings.Split(input[1], "\n")

	for _, v := range foldInput {
		split := strings.Split(v, " ")
		elem := strings.Split(split[2], "=")
		l, _ := strconv.Atoi(elem[1])
		foldInstruction := fold{elem[0] == "x", l}
		folds = append(folds, foldInstruction)
	}

	return dots, folds
}

func readFileRowToFishMap(pathToFile string) map[int]int {
	s, err := ioutil.ReadFile(pathToFile)
	if err != nil {
		log.Fatalf("Error readinf File: %s\n", err)
	}

	input := strings.Split(string(s), ",")
	var values []int

	for i := range input {
		str := input[i]
		number, err := strconv.Atoi(str)
		if err != nil {
			log.Fatalf("Error converting string to int: %s\n", err)
		}
		values = append(values, number)
	}
	output := make(map[int]int)
	for _, value := range values {
		output[value] += 1
	}

	return output
}

func readFileToStringSlice(pathToFile string) []string {
	s, err := ioutil.ReadFile(pathToFile)
	if err != nil {
		log.Fatalf("Error readinf File: %s\n", err)
	}

	split := strings.Split(string(s), "\n")

	var trimmed []string
	for _, s2 := range split {

		trimmed = append(trimmed, strings.Trim(s2, " \n"))
	}
	return trimmed
}

func parseHexVals(pathToFile string) map[string]string {
	s, err := ioutil.ReadFile(pathToFile)
	if err != nil {
		log.Fatalf("Error readinf File: %s\n", err)
	}

	values := strings.Split(string(s), "\n")

	m := make(map[string]string)

	for _, v := range values {
		split := strings.Split(v, " = ")
		m[split[0]] = split[1]
	}
	return m
}

func readFileToPathMap(pathToFile string) map[string][]string {
	s, err := ioutil.ReadFile(pathToFile)
	if err != nil {
		log.Fatalf("Error readinf File: %s\n", err)
	}

	paths := strings.Split(string(s), "\n")
	output := make(map[string][]string)

	for _, path := range paths {
		split := strings.Split(path, "-")
		key := split[0]
		value := split[1]
		output[key] = append(output[key], value)
		if value != "end" {
			output[value] = append(output[value], key)
		}
	}

	return output
}

func readFileToBingoBoards(pathToFile string) []string {
	s, err := ioutil.ReadFile(pathToFile)
	if err != nil {
		log.Fatalf("Error readinf File: %s\n", err)
	}

	return strings.Split(string(s), "\n\n")
}

func readFileToCommandList(pathToFile string) []command {
	s, err := ioutil.ReadFile(pathToFile)
	if err != nil {
		log.Fatalf("Error readinf File: %s\n", err)
	}

	input := strings.Split(string(s), "\n")
	var commands []command

	for i := range input {
		str := input[i]
		split := strings.Split(str, " ")
		number, err := strconv.Atoi(split[1])
		if err != nil {
			log.Fatalf("Error converting string to int: %s\n", err)
		}
		commands = append(commands, command{split[0], number})
	}
	return commands
}

func readFileToDisplays(pathToFile string) []display {
	s, err := ioutil.ReadFile(pathToFile)
	if err != nil {
		log.Fatalf("Error readinf File: %s\n", err)
	}

	input := strings.Split(string(s), "\n")
	var displays []display

	for i := range input {
		str := input[i]
		split := strings.Split(str, " | ")
		d := display{[10]string{""}, [4]string{""}}
		digits := strings.Split(split[0], " ")
		numbers := strings.Split(split[1], " ")

		var sortedDigits []string
		var sortedNumbers []string

		for _, digit := range digits {
			dig := strings.Split(digit, "")
			sort.Strings(dig)
			sortedDigits = append(sortedDigits, strings.Join(dig, ""))
		}

		for _, n := range numbers {
			num := strings.Split(n, "")
			sort.Strings(num)
			sortedNumbers = append(sortedNumbers, strings.Join(num, ""))
		}

		for j, digit := range sortedDigits {
			d.digits[j] = digit
		}
		for k, number := range sortedNumbers {
			d.number[k] = number
		}

		displays = append(displays, d)
	}
	return displays
}

func readFileListOfPairs(pathToFile string) []string {
	s, err := ioutil.ReadFile(pathToFile)
	if err != nil {
		log.Fatalf("Error readinf File: %s\n", err)
	}

	split := strings.Split(string(s), "\n")

	return split
}

func readFileToLines(pathToFile string) []line {
	s, err := ioutil.ReadFile(pathToFile)
	if err != nil {
		log.Fatalf("Error readinf File: %s\n", err)
	}

	split := strings.Split(string(s), "\n")
	var output []line
	for _, e := range split {
		stringPoints := strings.Split(e, " -> ")
		l := line{parsePoint(stringPoints[0]), parsePoint(stringPoints[1])}
		output = append(output, l)
	}
	return output
}

func parsePoint(s string) point {
	strings := strings.Split(s, ",")
	x, _ := strconv.Atoi(strings[0])
	y, _ := strconv.Atoi(strings[1])
	return point{x, y}
}

func diffFiles(p1 string, p2 string) {
	one := readFileToStringSlice(p1)
	two := readFileToStringSlice(p2)

	for _, s := range one {
		isInBoth := false
		for _, i := range two {
			if i == s {
				isInBoth = true
			}
		}
		if !isInBoth {
			fmt.Println(s)
		}
	}
}
