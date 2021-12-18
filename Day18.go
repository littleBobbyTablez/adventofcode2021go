package main

import (
	"regexp"
	"strconv"
	"strings"
)

func addAll(pairs []string) string {
	start := pairs[0]
	for _, pair := range pairs[1:] {
		start = addNumbers(start, pair)
	}
	return start
}

func addNumbers(s1 string, s2 string) string {
	sub := combine(s1, s2)
	sub = handle(sub)
	return sub
}

func handle(before string) string {
	sub := before
	first := findFirstPairOfNumbers(sub)
	sub = explodeAll("", sub, first)
	sub = splitOneValue(sub)

	if sub == before {
		return sub
	}

	return handle(sub)
}

func combine(s1 string, s2 string) string {
	return "[" + s1 + "," + s2 + "]"
}

func splitOneValue(s string) string {
	r, _ := regexp.Compile("\\d\\d")
	toSplit := r.FindString(s)
	if toSplit == "" {
		return s
	}
	split := strings.Split(s, toSplit)
	res := split[0]
	res += splitValue(toSplit) + split[1]

	for _, v := range split[2:] {
		res += toSplit + v
	}

	return res
}

func splitValue(v string) string {
	i, _ := strconv.Atoi(v)
	l := i / 2
	r := i/2 + i%2

	return combine(strconv.Itoa(l), strconv.Itoa(r))
}

func explodeAll(pre string, s string, pair string) string {
	if pair == "" {
		return pre + s
	}

	newPre, post := splitAtPair(s, pair)
	pre += newPre

	if isReadyToExplode(pre) {
		l, r := findNumbersInPair(pair)
		toReplaceL := findFirstNumberToTheLeft(pre)
		toReplaceR := findFirstNumberToTheRight(post)

		if toReplaceL != -1 {
			pre = replaceLeft(toReplaceL, l+toReplaceL, pre)
		}
		if toReplaceR != -1 {
			post = replaceRight(toReplaceR, r+toReplaceR, post)
		}
		next := findFirstPairOfNumbers(post)
		return explodeAll(pre+"0", post, next)
	} else {
		next := findFirstPairOfNumbers(post)
		return explodeAll(pre+pair, post, next)
	}
}

func replaceRight(r int, nr int, s string) string {
	sr := strconv.Itoa(r)
	split := strings.Split(s, sr)

	res := ""
	for i, p := range split[:len(split)-1] {
		if i == 0 {
			res += p + strconv.Itoa(nr)
		} else {
			res += p + sr
		}
	}

	return res + split[len(split)-1]
}

func replaceLeft(l int, nl int, s string) string {
	sl := strconv.Itoa(l)
	split := strings.Split(s, sl)

	res := split[0]
	for i, p := range split[1:] {
		if i == len(split)-2 {
			res += strconv.Itoa(nl) + p
		} else {
			res += sl + p
		}
	}

	return res
}

func isReadyToExplode(s string) bool {
	r, _ := regexp.Compile("\\[")
	r2, _ := regexp.Compile("\\]")
	open := len(r.FindAllString(s, -1))
	closed := len(r2.FindAllString(s, -1))

	return open-closed > 3
}

func findFirstPairOfNumbers(s string) string {
	r, _ := regexp.Compile("\\[\\d+,\\d+\\]")

	return r.FindString(s)
}

func findFirstNumberToTheLeft(s string) int {
	r, _ := regexp.Compile("\\d+")
	n := r.FindAllString(s, -1)
	if len(n) == 0 {
		return -1
	}
	left, _ := strconv.Atoi(n[len(n)-1])

	return left
}

func findFirstNumberToTheRight(s string) int {
	r, _ := regexp.Compile("\\d+")
	n := r.FindString(s)
	if n == "" {
		return -1
	}
	right, _ := strconv.Atoi(n)

	return right
}

func findNumbersInPair(s string) (int, int) {
	r, _ := regexp.Compile("\\d+")
	allString := r.FindAllString(s, -1)
	left, _ := strconv.Atoi(allString[0])
	right, _ := strconv.Atoi(allString[1])

	return left, right
}

func splitAtPair(s string, pair string) (string, string) {
	split := strings.Split(s, pair)
	pre := split[0]
	post := split[1]

	for _, v := range split[2:] {
		post += pair + v
	}
	return pre, post
}

func calculateMagnitude(s string) int {
	pair := findFirstPairOfNumbers(s)
	if pair == "" {
		res, _ := strconv.Atoi(s)
		return res
	}
	pre, post := splitAtPair(s, pair)
	a, b := findNumbersInPair(pair)
	magnitude := strconv.Itoa((3 * a) + (2 * b))

	return calculateMagnitude(pre + magnitude + post)
}

func findLargestSum(numbers []string) int {
	max := 0
	for i := 0; i < len(numbers); i++ {
		for j := 0; j < len(numbers); j++ {
			if i != j {
				b := calculateMagnitude(addNumbers(numbers[j], numbers[i]))
				if b > max {
					max = b
				}
			}
		}
	}
	return max
}
