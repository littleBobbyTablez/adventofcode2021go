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

func foldAll(dots []dot, ins []fold) map[dot]bool {
	output := dots
	dotMap := make(map[dot]bool)
	for _, in := range ins {
		output, dotMap = foldOnce(output, in)
	}

	return dotMap
}

func foldOnce(dots []dot, ins fold) ([]dot, map[dot]bool) {
	dotMap := make(map[dot]bool)

	for _, d := range dots {
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

	var output []dot

	for k, _ := range dotMap {
		output = append(output, k)
	}

	return output, dotMap

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
