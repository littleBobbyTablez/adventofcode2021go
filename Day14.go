package main

func insert(input map[string]int, rules map[string]string, count int) int {
	for i := 0; i < count; i++ {
		input = insertOnce(input, rules)
	}
	return countAndSubtract(input)
}

func insertOnce(input map[string]int, rules map[string]string) map[string]int {
	output := make(map[string]int)
	for k, v := range input {
		output[k] = v
	}
	for k, v := range input {
		elem := rules[k]
		first := string(k[0]) + elem
		output[first] += v
		second := elem + string(k[1])
		output[second] += v
		output[k] = output[k] - v
	}

	return output
}

func countAndSubtract(input map[string]int) int {
	most := 0
	fewest := -1

	m := make(map[string]int)
	for k, v := range input {
		first := string(k[0])
		second := string(k[1])

		m[first] += v
		m[second] += v
	}
	for _, v := range m {
		if fewest == -1 {
			fewest = v
		}
		if v > most {
			most = v
		}
		if v < fewest {
			fewest = v
		}
	}

	// % are necessary to care for edges
	return (most - fewest + ((most % 2) - (fewest % 2))) / 2
}
