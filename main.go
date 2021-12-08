package main

import (
	"fmt"
)

func main() {
	//executeDay1()
	//executeDay2()
	//executeDay3()
	//executeDay4()
	//executeDay5()
	//executeDay6()
	//executeDay7()

	executeDay8()
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
