package main

import (
	"strconv"
	"strings"
)

type display struct {
	digits [10]string
	number [4]string
}

func countEasyNumbers(input []display) int {
	count := 0
	for _, d := range input {
		for _, n := range d.number {
			length := len(n)
			if length == 2 || length == 3 || length == 4 || length == 7 {
				count++
			}
		}
	}
	return count
}

func calculateSumOfNumbers(input []display) int {
	sum := 0
	for _, d := range input {
		dict := parseDisplay(d)
		var l []int
		for _, s := range d.number {
			l = append(l, dict[s])
		}
		var sl []string
		for _, i := range l {
			sl = append(sl, strconv.Itoa(i))
		}

		s := strings.Join(sl, "")

		atoi, _ := strconv.Atoi(s)
		sum += atoi
	}
	return sum
}

func parseDisplay(d display) map[string]int {
	dict, unique, lengthSix, lengthFife := addUniqueAndCollectNonUnique(d)

	dict = addSixer(lengthSix, unique, dict)
	dict = addFifer(lengthFife, unique, dict)
	return dict
}

func addUniqueAndCollectNonUnique(d display) (map[string]int, map[int]string, []string, []string) {
	dict := make(map[string]int)
	unique := make(map[int]string)

	var lengthSix []string
	var lengthFife []string

	for _, digit := range d.digits {
		l := len(digit)
		switch l {
		case 2:
			dict[digit] = 1
			unique[1] = digit
		case 3:
			dict[digit] = 7
			unique[7] = digit
		case 4:
			dict[digit] = 4
			unique[4] = digit
		case 7:
			dict[digit] = 8
			unique[8] = digit
		case 6:
			lengthSix = append(lengthSix, digit)
		case 5:
			lengthFife = append(lengthFife, digit)
		default:
		}
	}
	return dict, unique, lengthSix, lengthFife
}

func addFifer(lengthFife []string, unique map[int]string, dict map[string]int) map[string]int {
	for _, fife := range lengthFife {
		diff4 := len(diff(fife, unique[4]))
		diff3 := len(diff(fife, unique[1]))

		switch {
		case diff4 == 2:
			dict[fife] = 2
		case diff4 == 1 && diff3 == 0:
			dict[fife] = 3
		case diff4 == 1 && diff3 == 1:
			dict[fife] = 5
		default:
		}
	}
	return dict
}

func addSixer(lengthSix []string, unique map[int]string, dict map[string]int) map[string]int {
	for _, six := range lengthSix {
		diff4 := len(diff(six, unique[4]))
		diff1 := len(diff(six, unique[1]))
		diff7 := len(diff(six, unique[7]))

		switch {
		case diff4 == 0:
			dict[six] = 9
		case diff1 == 0 && diff4 == 1:
			dict[six] = 0
		case diff7 == 1:
			dict[six] = 6
		default:
		}
	}
	return dict
}

func diff(a string, b string) []string {
	var diff []string

	for _, s := range strings.Split(b, "") {
		if !stringContains(a, s) {
			diff = append(diff, s)
		}
	}
	return diff
}

func stringContains(a string, b string) bool {
	for _, e := range strings.Split(a, "") {
		if e == b {
			return true
		}
	}
	return false
}
