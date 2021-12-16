package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	//executeDay1()
	//executeDay2()
	//executeDay3()
	//executeDay4()
	//executeDay5()
	//executeDay6()
	//executeDay7()
	//executeDay8()
	//executeDay9()
	//executeDay10()
	//executeDay11()
	//executeDay12()
	//executeDay13()

	//executeDay14()
	//ececuteDay15()

	executeDay16()
}

func executeDay16() {
	pathToHex := "resources/hexVals.txt"
	hexVals := parseHexVals(pathToHex)
	pathToFile := "resources/Day16Input.txt"
	s, _ := ioutil.ReadFile(pathToFile)

	binary := translateToBinary(string(s), hexVals)
	//binary := translateToBinary("880086C3E88112", hexVals)

	p := packageParser{bs(binary), 0, 0}
	fmt.Println(p.parse())
}

func ececuteDay15() {
	//pathToFile := "resources/Day15Example.txt"
	pathToFile := "resources/Day15Input.txt"

	input := readFileToRiskLevelMap(pathToFile)

	findShortestPath(point{0, 0}, input, point{99, 99})
	findShortestPathWithRealMap(point{0, 0}, input, point{499, 499})
}

func executeDay14() {
	//pathToFile := "resources/Day14Example.txt"
	pathToFile := "resources/Day14Input.txt"

	startingMap, rules := readFileInsertRules(pathToFile)

	fmt.Println(insert(startingMap, rules, 40))
}

func executeDay13() {
	//pathToFile := "resources/Day13Example.txt"
	pathToFile := "resources/Day13Input.txt"
	inputDots, inputFolds := readFileDotsAndFolds(pathToFile)

	once := foldOnce(inputDots, inputFolds[0])
	fmt.Println(len(once))
	all := foldAll(inputDots, inputFolds)
	fmt.Println(len(all))
	printDots(all)
}

func executeDay12() {
	//pathToFile := "resources/Day12Example.txt"
	pathToFile := "resources/Day12Input.txt"
	input := readFileToPathMap(pathToFile)

	//fmt.Println(findPaths(input))
	findPaths2(input)
}

func executeDay11() {
	pathToFile := "resources/Day11Input.txt"
	input := readFileOctupusMap(pathToFile)

	fmt.Println(countFalshes(input, 0, 0))
}

func executeDay10() {
	pathToFile := "resources/Day10Input.txt"
	input := readFileToStringSlice(pathToFile)

	fmt.Println(findFirstCorruptElement(input))
	fmt.Println(calculateUnfinishedScore(input))
}

func executeDay9() {
	pathToFile := "resources/Day9InputMark.txt"
	input := readFileHeightMap(pathToFile)

	fmt.Println(findLowPoints(input))
}

func executeDay8() {
	pathToFile := "resources/Day8Input.txt"
	input := readFileToDisplays(pathToFile)
	fmt.Println(countEasyNumbers(input))
	fmt.Println(calculateSumOfNumbers(input))
}

func executeDay7() {
	pathToFile := "resources/Day7Input.txt"
	input := readFileRowToInt(pathToFile)

	fmt.Println(findBestPosition(input))
}

func executeDay6() {
	pathToFile := "resources/Day6Input.txt"
	input := readFileRowToFishMap(pathToFile)
	growEfficiently(input, 0, 256)
}

func executeDay5() {
	pathToFile := "resources/Day5Input.txt"
	input := readFileToLines(pathToFile)

	findHorizontalAndVeticalIntersections(input)
	findIntersections(input)
}

func executeDay4() {
	pathToFile := "/Users/jonas.stendel/Projekte/AdventOfCode2021/src/main/resources/Day4_1.txt"

	input := readFileToBingoBoards(pathToFile)
	//fmt.Println(len(readFileToBingoBoards(pathToFile)))
	play(input)
}

func executeDay3() {
	pathToFile := "/Users/jonas.stendel/Projekte/AdventOfCode2021/src/main/resources/Day3_1.txt"
	input := readFileToStringSlice(pathToFile)
	fineGammaAndEpsilonRate(input)
	findOxyAndCo2Rate(input)
}

func executeDay1() {
	pathToFile := "/Users/jonas.stendel/Projekte/AdventOfCode2021/src/main/resources/Day1_1.txt"
	input := readFileToIntSlice(pathToFile)
	fmt.Println(countIncreases(input))

	summed := sumOfThree(input)
	fmt.Println(countIncreases(summed))
}

func executeDay2() {
	pathToFile := "/Users/jonas.stendel/Projekte/AdventOfCode2021/src/main/resources/Day2_1.txt"
	input := readFileToCommandList(pathToFile)

	target := moveToTarget(input, position{0, 0})
	fmt.Println(target)
	fmt.Println(target.y * target.x)
	withAim := moveToTargetWithAim(input, aimPosition{0, 0, 0})
	fmt.Println(withAim)
	fmt.Println(withAim.x * withAim.y)
}
