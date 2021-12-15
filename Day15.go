package main

import (
	"fmt"
)

func findShortestPath(start point, values map[point]int, end point) {

	distances := make(map[point]int)
	distances[point{0, 0}] = 0

	findPathUntilItDoesNotChange(start, values, distances, end, -1)
}

func findShortestPathWithRealMap(start point, values map[point]int, end point) {
	realMap := calculateRealMap(values, end)

	distances := make(map[point]int)
	distances[point{0, 0}] = 0

	findPathUntilItDoesNotChange(start, realMap, distances, end, -1)
}

func findPathUntilItDoesNotChange(p point, values map[point]int, distances map[point]int, end point, last int) {
	resultMap := findPathRec(p, values, distances, end)

	result := resultMap[end]
	if last == result {
		fmt.Println(result)
		return
	}
	findPathUntilItDoesNotChange(p, values, distances, end, result)
}

func calculateRealMap(values map[point]int, end point) map[point]int {

	output := make(map[point]int)
	startSize := (end.x + 1) / 5
	realSize := startSize * 5
	for i := 0; i < realSize; i++ {
		for j := 0; j < realSize; j++ {

			p := point{i % startSize, j % startSize}
			original := values[p]
			newValue := (((original - 1) + j/startSize + i/startSize) % 9) + 1
			output[point{i, j}] = newValue
		}
	}
	return output
}

func findPathRec(p point, values map[point]int, distances map[point]int, end point) map[point]int {
	neighbours := getNeighbours(p)

	for _, n := range neighbours {

		if n.x >= 0 && n.y >= 0 && n.x <= end.x && n.y <= end.y {
			newPath := distances[p] + values[n]

			if oldPath, exists := distances[n]; (exists && newPath < oldPath) || !exists {
				distances[n] = newPath
			}
		}
		if n.x == 0 && n.y == 0 {
			distances[n] = 0
		}
	}

	if p == end {
		return distances
	}

	if p.y == end.y {
		return findPathRec(point{p.x + 1, 0}, values, distances, end)
	} else {
		return findPathRec(point{p.x, p.y + 1}, values, distances, end)
	}
}

func getNeighbours(p point) []point {
	var neighbours []point
	u := p.x - 1
	d := p.x + 1
	l := p.y - 1
	r := p.y + 1

	neighbours = append(neighbours, point{u, p.y})
	neighbours = append(neighbours, point{d, p.y})
	neighbours = append(neighbours, point{p.x, r})
	neighbours = append(neighbours, point{p.x, l})

	return neighbours
}
