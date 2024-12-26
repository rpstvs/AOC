package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	x, y int
}

var numKeypad = map[byte]Point{'7': {0, 0}, '8': {0, 1}, '9': {0, 2},
	'4': {1, 0}, '5': {1, 1}, '6': {1, 2},
	'1': {2, 0}, '2': {2, 1}, '3': {2, 2},
	'_': {3, 0}, '0': {3, 1}, 'A': {3, 2}}

var dirKeypad = map[byte]Point{'_': {0, 0}, '^': {0, 1}, 'A': {0, 2},
	'<': {1, 0}, 'V': {1, 1}, '>': {1, 2},
}

type key struct {
	seq   string
	depth int
}

var cachePress = map[key]int{}

func main() {
	sequences := parseInput()

	solvept1(sequences)
	solvept2(sequences)
}

func solvept1(sequences []string) {
	var sum int
	for _, code := range sequences {
		directions := pressNumbers(code)
		seqLength := press(directions, 2)

		numericPart, _ := strconv.Atoi(code[:len(code)-1])

		sum += seqLength * numericPart
	}
	fmt.Println(sum)
}

func solvept2(sequences []string) {
	var sum int
	for _, code := range sequences {
		directions := pressNumbers(code)
		seqLength := press(directions, 25)

		numericPart, _ := strconv.Atoi(code[:len(code)-1])

		sum += seqLength * numericPart
	}
	fmt.Println(sum)
}

func press(code string, depth int) int {
	if depth == 0 {
		return len(code)
	}

	if size, ok := cachePress[key{code, depth}]; ok {
		return size
	}

	totalSize := 0

	for _, chunk := range strings.SplitAfter(code, "A") {
		totalSize += press(pressDirections(chunk), depth-1)

	}

	cachePress[key{code, depth}] = totalSize
	return totalSize
}

func pressDirections(dir string) string {
	hole := dirKeypad['_']
	curr := dirKeypad['A']

	seq := ""

	for _, c := range dir {
		goal := dirKeypad[byte(c)]

		deltaH, deltaV := goal.y-curr.y, goal.x-curr.x

		hs, vs := getMovements(deltaH, deltaV)

		seq += prioritize(hs, vs, curr, goal, hole)
		curr.x += deltaV
		curr.y += deltaH

		seq += "A"
	}
	return seq
}

func pressNumbers(dir string) string {
	hole := numKeypad['_']
	curr := numKeypad['A']

	seq := ""

	for _, c := range dir {
		goal := numKeypad[byte(c)]

		deltaH, deltaV := goal.y-curr.y, goal.x-curr.x

		hs, vs := getMovements(deltaH, deltaV)

		seq += prioritize(hs, vs, curr, goal, hole)
		curr.x += deltaV
		curr.y += deltaH

		seq += "A"
	}
	return seq
}

func getMovements(deltaH, deltaV int) (string, string) {
	var hs, vs string

	for deltaH > 0 {
		hs += ">"
		deltaH--
	}

	for deltaH < 0 {
		hs += "<"
		deltaH++
	}

	for deltaV > 0 {
		vs += "V"
		deltaV--
	}

	for deltaV < 0 {
		vs += "^"
		deltaV++
	}
	return hs, vs

}

func prioritize(hs, vs string, from, to, avoid Point) string {
	switch {
	case from.x == avoid.x && to.y == avoid.y:
		return vs + hs
	case from.y == avoid.y && to.x == avoid.x:
		return hs + vs
	case strings.Contains(hs, "<"):
		return hs + vs
	default:
		return vs + hs
	}
}

func parseInput() []string {
	file, _ := os.Open("input.txt")

	defer file.Close()
	sequences := []string{}
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		sequences = append(sequences, line)
	}
	return sequences
}
