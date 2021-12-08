package main

import (
	"fmt"
	"strconv"
	"strings"
)

type board struct {
	numbers [5][5]int
	visited [5][5]bool
}

type complete struct {
	isRow    bool
	index    int
	complete bool
}

type winningBoard struct {
	isRow            bool
	index            int
	rowOrColmunIndex int
}

func play(input []string) {
	numbers := getNumbersFromInput(input)
	boards := getBingoBoards(input[1:])

	playRec(numbers, boards, -1)
	playRecAndLoose(numbers, boards, -1)
}

func playRec(numbers []int, boards []board, last int) {
	winningIndex := checkForWinningBoard(boards)
	if winningIndex.index != -1 {
		fmt.Printf("Winning Board Index: %d\n", winningIndex.index)
		fmt.Printf("Winning Board:\n%d\n%t\n", boards[winningIndex.index].numbers, boards[winningIndex.index].visited)
		fmt.Printf("Winning line is Row: %t\n", winningIndex.isRow)
		fmt.Printf("Winning row Index is: %d\n", winningIndex.rowOrColmunIndex)
		fmt.Printf("Last Number: %d\n", last)

		score := evalScore(boards[winningIndex.index]) * last

		fmt.Printf("The score is: %d\n", score)
		return
	}

	number := numbers[0]
	for i, b := range boards {
		for j, r := range b.numbers {
			for k, c := range r {
				if c == number {
					boards[i].visited[j][k] = true
				}
			}
		}
	}
	playRec(numbers[1:], boards, number)
}

func playRecAndLoose(numbers []int, boards []board, last int) {
	winningIndex := checkForWinningBoard(boards)
	if len(boards) == 0 {
		return
	}
	if len(boards) == 1 && winningIndex.index != -1 {
		fmt.Printf("Loosing Board Index: %d\n", winningIndex.index)
		fmt.Printf("Loosing Board:\n%d\n%t\n", boards[winningIndex.index].numbers, boards[winningIndex.index].visited)
		fmt.Printf("Loosing line is Row: %t\n", winningIndex.isRow)
		fmt.Printf("Loosing row Index is: %d\n", winningIndex.rowOrColmunIndex)
		fmt.Printf("Last Number: %d\n", last)

		score := evalScore(boards[winningIndex.index]) * last

		fmt.Printf("The score is: %d\n", score)
		return
	}

	if winningIndex.index != -1 {
		var boardsLeft []board
		for i, b := range boards {
			if i != winningIndex.index {
				boardsLeft = append(boardsLeft, b)
			}
		}
		playRecAndLoose(numbers, boardsLeft, last)
	} else {
		number := numbers[0]
		for i, b := range boards {
			for j, r := range b.numbers {
				for k, c := range r {
					if c == number {
						boards[i].visited[j][k] = true
					}
				}
			}
		}
		playRecAndLoose(numbers[1:], boards, number)
	}
}

func evalScore(b board) int {
	sum := 0
	for i, r := range b.visited {
		for j, c := range r {
			if !c {
				sum += b.numbers[i][j]
			}
		}
	}
	return sum
}

func checkForWinningBoard(boards []board) winningBoard {
	for i, b := range boards {
		c := checkBoard(b.visited)
		if c.complete {
			return winningBoard{c.isRow, i, c.index}
		}
	}
	return winningBoard{false, -1, -1}
}

func checkBoard(b [5][5]bool) complete {
	for i, r := range b {
		rowComplete := check(r, 0)
		if rowComplete {
			return complete{true, i, true}
		}
	}
	for i := range b {
		column := boardColumn(b, i)
		columnComplete := check(column, 0)
		if columnComplete {
			return complete{false, i, true}
		}
	}

	return complete{false, -1, false}
}

func check(r [5]bool, index int) bool {
	b := r[index]
	if index == 4 {
		return b
	}
	if !b {
		return false
	}
	return check(r, index+1)
}

func boardColumn(b [5][5]bool, columnIndex int) [5]bool {
	column := [5]bool{false}
	for i, row := range b {
		column[i] = row[columnIndex]
	}
	return column
}

func getNumbersFromInput(input []string) []int {
	str := strings.Split(input[0], ",")
	var numbers []int
	for _, s := range str {
		i, _ := strconv.Atoi(s)
		numbers = append(numbers, i)
	}
	return numbers
}

func getBingoBoards(input []string) []board {
	var boards []board
	for _, b := range input {
		boards = append(boards, parseBoard(b))
	}
	return boards
}
func parseBoard(input string) board {
	rows := strings.Split(input, "\n")
	var splitRows [5][5]int

	for i, row := range rows {
		trimmed := strings.Trim(row, " ")
		split := strings.Split(trimmed, " ")
		filtered := filterEmpty(split)
		ints := [5]int{0}
		for j, s := range filtered {
			n, _ := strconv.Atoi(s)
			ints[j] = n
		}
		splitRows[i] = ints
	}
	notVisitedRow := [5]bool{false, false, false, false, false}
	return board{splitRows, [5][5]bool{notVisitedRow}}
}

func filterEmpty(split []string) []string {
	var output []string
	for _, s := range split {
		if s != "" {
			output = append(output, s)
		}
	}
	return output
}
