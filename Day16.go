package main

import (
	"strconv"
)

type bs string

type packageParser struct {
	input      bs
	pos        int
	versionSum int
}

func (p *packageParser) parse() int {
	vs := p.readInt(3)
	id := p.readInt(3)

	p.versionSum += vs

	if id == 4 {
		return p.parseLiteral("")
	} else {
		values := p.parseOperational()
		return calculateValues(values, id)
	}
}

func (p *packageParser) parseOperational() []int {
	var values []int
	if p.readInt(1) == 1 {
		count := p.readInt(11)
		for i := 0; i < count; i++ {
			values = append(values, p.parse())
		}
	} else {
		length := p.readInt(15)
		end := p.pos + length
		for p.pos < end {
			values = append(values, p.parse())
		}
	}
	return values
}

func (p *packageParser) parseLiteral(acc bs) int {
	if p.readInt(1) == 1 {
		i := p.input[p.pos : p.pos+4]
		p.pos += 4
		return p.parseLiteral(acc + i)
	} else {
		i := p.input[p.pos : p.pos+4]
		p.pos += 4
		res := (acc + i).asInt()
		return res
	}
}

func (p *packageParser) readInt(l int) int {
	i := p.input[p.pos : p.pos+l].asInt()
	p.pos += l
	return i
}

func (s bs) asInt() int {
	y, _ := strconv.ParseInt(string(s), 2, 64)
	return int(y)
}

func calculateValues(values []int, id int) int {
	switch id {
	case 0:
		sum := 0
		for _, v := range values {
			sum += v
		}
		return sum
	case 1:
		sum := 1
		for _, v := range values {
			sum *= v
		}
		return sum
	case 2:
		min := values[0]
		for _, v := range values {
			if min > v {
				min = v
			}
		}
		return min
	case 3:
		max := 0
		for _, v := range values {
			if max < v {
				max = v
			}
		}
		return max
	case 5:
		if values[0] > values[1] {
			return 1
		}
		return 0
	case 6:
		if values[0] < values[1] {
			return 1
		}
		return 0
	case 7:
		if values[0] == values[1] {
			return 1
		}
		return 0
	}
	return values[0]
}

func translateToBinary(input string, hexVals map[string]string) []byte {

	binaryString := ""

	for _, i := range input {
		binaryString = binaryString + hexVals[string(i)]
	}

	return []byte(binaryString)
}
