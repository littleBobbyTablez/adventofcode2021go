package main

import (
	"fmt"
)

type point struct {
	x int
	y int
}

type line struct {
	start point
	end   point
}

func findHorizontalAndVeticalIntersections(input []line) {
	var filtered []line
	for _, l := range input {
		if l.start.x == l.end.x || l.start.y == l.end.y {
			filtered = append(filtered, l)
		}
	}

	findIntersections(filtered)
}

func findIntersections(input []line) {

	var normalized []line
	for _, l := range input {
		if l.start.x > l.end.x || l.start.y > l.end.y {
			normalized = append(normalized, line{l.end, l.start})
		} else {
			normalized = append(normalized, line{l.start, l.end})
		}
	}

	var visited = make(map[point]bool)
	var overlapping = make(map[point]bool)
	var points = getAllPoints(normalized)
	fmt.Println(findOverlapping(points, visited, overlapping))
}

func findOverlapping(input []point, visited map[point]bool, overlapping map[point]bool) int {
	if len(input) == 0 {
		return len(overlapping)
	}
	subject := input[0]

	if visited[subject] {
		overlapping[subject] = true
		visited[subject] = true
		return findOverlapping(input[1:], visited, overlapping)
	}
	visited[subject] = true
	return findOverlapping(input[1:], visited, overlapping)
}

func getAllPoints(lines []line) []point {
	var output []point
	for _, l := range lines {
		points := getPointsOnLine(l)
		for _, p := range points {
			output = append(output, p)
		}
	}
	return output
}

func getPointsOnLine(subject line) []point {
	var points []point
	if isHorizontal(subject) {
		for i := subject.start.x; i <= subject.end.x; i++ {
			points = append(points, point{i, subject.start.y})
		}
	} else if isVertical(subject) {
		for i := subject.start.y; i <= subject.end.y; i++ {
			points = append(points, point{subject.start.x, i})
		}
	} else if isProportional(subject) {
		for i := 0; i <= subject.end.y-subject.start.y; i++ {
			points = append(points, point{subject.start.x + i, subject.start.y + i})
		}
	} else {
		if subject.start.x > subject.end.x {
			for i := 0; i <= subject.end.y-subject.start.y; i++ {
				points = append(points, point{subject.start.x - i, subject.start.y + i})
			}
		} else {
			for i := 0; i <= subject.start.y-subject.end.y; i++ {
				points = append(points, point{subject.start.x + i, subject.start.y - i})
			}
		}
	}
	return points
}

func isProportional(l line) bool {
	return l.start.x <= l.end.x && l.start.y <= l.end.y
}

func isHorizontal(l line) bool {
	return l.start.y == l.end.y
}

func isVertical(l line) bool {
	return l.start.x == l.end.x
}
