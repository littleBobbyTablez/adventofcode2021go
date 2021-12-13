package main

import "fmt"

type dot struct {
	x int
	y int
}

type fold struct {
	x    bool
	line int
}

func foldAll(dots map[dot]bool, ins []fold) map[dot]bool {
	dotMap := make(map[dot]bool)
	dotMap = dots
	for _, in := range ins {
		dotMap = foldOnce(dotMap, in)
	}

	return dotMap
}

func foldOnce(dots map[dot]bool, ins fold) map[dot]bool {
	dotMap := make(map[dot]bool)

	for d, _ := range dots {
		if ins.x {
			if d.x < ins.line {
				dotMap[d] = true
			} else {
				newX := ins.line - (d.x - ins.line)
				dotMap[dot{newX, d.y}] = true
			}
		} else {
			if d.y < ins.line {
				dotMap[d] = true
			} else {
				newY := ins.line - (d.y - ins.line)
				dotMap[dot{d.x, newY}] = true
			}
		}
	}

	return dotMap
}

func printDots(dotMap map[dot]bool) {
	xMax := 0
	yMax := 0

	for k, _ := range dotMap {
		if k.x > xMax {
			xMax = k.x
		}
		if k.y > yMax {
			yMax = k.y
		}
	}

	for y := 0; y <= yMax; y++ {
		for x := 0; x <= xMax; x++ {
			if dotMap[dot{x, y}] {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}

		fmt.Println("")
	}
}
