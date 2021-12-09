package main

type lowPoint struct {
	i int
	j int
}

func findLowPoints(input [100][100]int) (int, int) {
	var heights []int
	basins := make(map[lowPoint]int)

	for i, row := range input {
		for j, c := range row {
			u := i - 1
			d := i + 1
			l := j - 1
			r := j + 1
			u, d, l, r = filterEdgeCasesForLowPoints(u, l, d, r)

			if c < input[u][j] && c < input[i][r] && c < input[d][j] && c < input[i][l] {
				heights = append(heights, c)
				basins[lowPoint{i, j}] = 1
			}
		}
	}

	for i, row := range input {
		for j, c := range row {
			if c != 9 && basins[lowPoint{i, j}] == 0 {
				k, l := findBasin(i, j, input, basins)
				basins[lowPoint{k, l}] += 1
			}
		}
	}

	var largestBasins = [3]int{0}

	for _, v := range basins {
		largestBasins = replaceIfLarger(v, largestBasins)
	}

	basinSize := 1
	for _, s := range largestBasins {
		if s != 0 {
			basinSize = basinSize * s
		}
	}
	output := 0
	for _, h := range heights {
		output += h + 1
	}

	return output, basinSize
}

func replaceIfLarger(i int, a [3]int) [3]int {
	index := 0
	for j, v := range a {
		if v < a[index] {
			index = j
		}
	}

	if i > a[index] {
		a[index] = i
	}

	return a
}

func findBasin(i int, j int, input [100][100]int, basins map[lowPoint]int) (int, int) {
	lp := lowPoint{i, j}
	if basins[lp] != 0 {
		return i, j
	}

	u := i - 1
	d := i + 1
	l := j - 1
	r := j + 1
	up, down, left, right := filterEdgeCases(u, l, d, r, i, j, input)

	switch {
	case up <= right && up <= down && up <= left:
		return findBasin(u, j, input, basins)
	case down <= right && down <= up && down <= left:
		return findBasin(d, j, input, basins)
	case right <= up && right <= down && right <= left:
		return findBasin(i, r, input, basins)
	case left <= up && left <= down && left <= right:
		return findBasin(i, l, input, basins)
	}
	return i, j
}

func filterEdgeCases(u int, l int, d int, r int, i int, j int, input [100][100]int) (int, int, int, int) {
	var up int
	var down int
	var left int
	var right int

	if u >= 0 {
		up = input[u][j]
	} else {
		up = 9
	}
	if l < 0 {
		left = 9
	} else {
		left = input[i][l]
	}
	if d > 99 {
		down = 9
	} else {
		down = input[d][j]
	}
	if r > 99 {
		right = 9
	} else {
		right = input[i][r]
	}
	return up, down, left, right
}

func filterEdgeCasesForLowPoints(u int, l int, d int, r int) (int, int, int, int) {
	if u < 0 {
		u = u + 2
	}
	if l < 0 {
		l = l + 2
	}
	if d > 99 {
		d = d - 2
	}
	if r > 99 {
		r = r - 2
	}
	return u, d, l, r
}
