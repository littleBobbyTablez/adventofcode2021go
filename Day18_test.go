package main

import (
	"fmt"
	"strings"
	"testing"
)

func Test_explosions(t *testing.T) {
	tests := []string{
		"[[6,[5,[4,[3,2]]]],1]",
		"[7,[6,[5,[4,[3,2]]]]]",
		"[[[[[9,8],1],2],3],4]",
		"[[3,[2,[1,[7,3]]]],[6,[5,[4,[3,2]]]]]",
		"[[3,[2,[8,0]]],[9,[5,[4,[3,2]]]]]",
		"[[[0,[4,5]],[0,0]],[[[4,5],[2,6]],[9,5]]]",
	}
	want := []string{
		"[[6,[5,[7,0]]],3]",
		"[7,[6,[5,[7,0]]]]",
		"[[[[0,9],2],3],4]",
		"[[3,[2,[8,0]]],[9,[5,[7,0]]]]",
		"[[3,[2,[8,0]]],[9,[5,[7,0]]]]",
		"[[[0,[4,5]],[0,0]],[[[4,5],[2,6]],[9,5]]]",
	}
	for i, tt := range tests {
		t.Run(tt, func(t *testing.T) {
			first := findFirstPairOfNumbers(tt)
			actual := explodeAll("", tt, first, make(map[string]int))
			if want[i] != actual {
				fmt.Println(actual)
				fmt.Println(want[i])
				t.Fail()
			}
		})
	}
}

func Test_numbers(t *testing.T) {
	tests := []struct {
		name string
		sut  [2]string
	}{
		{"1", [2]string{"[[[[4,3],4],4],[7,[[8,4],9]]]", "[1,1]"}},
		{"2", [2]string{"[[[0,[4,5]],[0,0]],[[[4,5],[2,6]],[9,5]]]", "[7,[[[3,7],[4,3]],[[6,3],[8,8]]]]"}},
		{"2", [2]string{"[[[[4,0],[5,4]],[[7,7],[6,0]]],[[8,[7,7]],[[7,9],[5,0]]]]", "[[2,[[0,8],[3,4]]],[[[6,7],1],[7,[1,6]]]]"}},
	}
	want := []string{
		"[[[[0,7],4],[[7,8],[6,0]]],[8,1]]",
		"[[[[4,0],[5,4]],[[7,7],[6,0]]],[[8,[7,7]],[[7,9],[5,0]]]]",
		"[[[[6,7],[6,7]],[[7,7],[0,7]]],[[[8,7],[7,7]],[[8,8],[8,0]]]]",
	}
	for i, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := addNumbers(test.sut[0], test.sut[1])

			if actual != want[i] {
				fmt.Println(actual)
				fmt.Println(want)
				t.Fail()
			}
		})
	}
}

func Test_addAll(t *testing.T) {
	tests := []string{
		"[1,1]\n[2,2]\n[3,3]\n[4,4]",
		"[1,1]\n[2,2]\n[3,3]\n[4,4]\n[5,5]",
		"[1,1]\n[2,2]\n[3,3]\n[4,4]\n[5,5]\n[6,6]",
		"[[[0,[4,5]],[0,0]],[[[4,5],[2,6]],[9,5]]]\n[7,[[[3,7],[4,3]],[[6,3],[8,8]]]]\n[[2,[[0,8],[3,4]]],[[[6,7],1],[7,[1,6]]]]\n[[[[2,4],7],[6,[0,5]]],[[[6,8],[2,8]],[[2,1],[4,5]]]]\n[7,[5,[[3,8],[1,4]]]]\n[[2,[2,2]],[8,[8,1]]]\n[2,9]\n[1,[[[9,3],9],[[9,0],[0,7]]]]\n[[[5,[7,4]],7],1]\n[[[[4,2],2],6],[8,7]]",
	}
	want := []string{
		"[[[[1,1],[2,2]],[3,3]],[4,4]]",
		"[[[[3,0],[5,3]],[4,4]],[5,5]]",
		"[[[[5,0],[7,4]],[5,5]],[6,6]]",
		"[[[[8,7],[7,7]],[[8,6],[7,7]]],[[[0,7],[6,6]],[8,7]]]",
	}
	for i, test := range tests {
		t.Run(test, func(t *testing.T) {
			sut := strings.Split(test, "\n")
			actual := addAll(sut)

			if actual != want[i] {
				fmt.Println(actual)
				fmt.Println(want)
				t.Fail()
			}
		})
	}
}

func Test_Magnitude(t *testing.T) {
	tests := []struct {
		name string
		sut  string
	}{
		{"1", "[[1,2],[[3,4],5]]"},
		{"2", "[[[[0,7],4],[[7,8],[6,0]]],[8,1]]"},
		{"3", "[[[[8,7],[7,7]],[[8,6],[7,7]]],[[[0,7],[6,6]],[8,7]]]"},
		//{"4", ""},
	}

	want := []int{
		143,
		1384,
		3488,
	}
	for i, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := calculateMagniture(test.sut)

			//fmt.Println(actual)
			if actual != want[i] {
				fmt.Println(actual, want[i])
				t.Fail()
			}
		})
	}
}
