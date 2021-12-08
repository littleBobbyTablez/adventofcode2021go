package main

func findBestPosition(positions []int) int {

	startValue := 0
	for _, v := range positions {
		startValue += calculateFuel(v)
	}

	return findBestPositionRec(positions, startValue, 1)
}

func findBestPositionRec(positions []int, previousBest int, subject int) int {
	if subject == 2001 {
		return previousBest
	}

	fuel := 0
	for _, v := range positions {
		fuel += calculateFuel(abs(v - subject))
	}

	if previousBest > fuel {
		return findBestPositionRec(positions, fuel, subject+1)
	} else {
		return findBestPositionRec(positions, previousBest, subject+1)
	}
}

func calculateFuel(distance int) int {
	fuel := 0
	for i := 1; i <= distance; i++ {
		fuel += i
	}
	return fuel
}

func sum(positions []int) int {
	sum := 0
	for _, v := range positions {
		sum += v
	}
	return sum
}

func abs(i int) int {
	if i >= 0 {
		return i
	}
	return i * -1
}
