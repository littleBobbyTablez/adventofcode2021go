package main

import "fmt"

type flashPoint struct {
	x int
	y int
}

func countFalshes(input [10][10]int, step int, flashes int) int {

	if step == 400 {
		return flashes
	}

	fpMap := make(map[flashPoint]bool)

	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			input[i][j] = input[i][j] + 1
		}
	}
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {

			i2 := input[i][j]
			if i2 >= 10 {
				fpMap[flashPoint{i, j}] = true
				input, flashes = flash(input, i, j, flashes, fpMap)
			}
		}
	}

	for k, _ := range fpMap {
		input[k.x][k.y] = 0
	}

	if len(fpMap) == 100 {
		fmt.Println(step + 1)
	}
	return countFalshes(input, step+1, flashes)
}

func flash(input [10][10]int, i int, j int, flashes int, fpMap map[flashPoint]bool) ([10][10]int, int) {
	for k := i - 1; k <= i+1; k++ {
		for l := j - 1; l <= j+1; l++ {
			if (k != i || l != j) && k >= 0 && k < 10 && l >= 0 && l < 10 {
				input[k][l] = input[k][l] + 1
				if input[k][l] == 10 {
					fpMap[flashPoint{k, l}] = true
					input, flashes = flash(input, k, l, flashes, fpMap)
				}
			}

		}
	}
	input[i][j] = 0

	return input, flashes + 1
}
