package main

import (
	"io/ioutil"
	"log"
	"sort"
	"strconv"
	"strings"
)

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

	return strings.Split(string(s), "\n")
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
