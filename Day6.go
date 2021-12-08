package main

import "fmt"

func growEfficiently(population map[int]int, day int, target int) {
	if day == target {
		fmt.Println(sumValues(population))
		return
	}

	nextPopulation := make(map[int]int)
	nextPopulation[8] = population[0]
	for i := 0; i < 8; i++ {
		nextPopulation[i] = population[i+1]
	}
	nextPopulation[6] += population[0]

	growEfficiently(nextPopulation, day+1, target)
}

func sumValues(someMap map[int]int) int {
	sum := 0
	for _, v := range someMap {
		sum += v
	}
	return sum
}
