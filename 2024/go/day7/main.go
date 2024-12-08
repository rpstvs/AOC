package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.Open("input.txt")

	defer file.Close()

	scanner := bufio.NewScanner(file)
	var line []string
	for scanner.Scan() {
		line2 := (scanner.Text())
		line = append(line, line2)
	}
	sum := 0
	sum2 := 0
	for _, result := range line {
		target, inputs := parseEquation(result)

		if solveEquation(1, inputs[0], target, inputs, false) {

			sum += target

		}

		if solveEquation(1, inputs[0], target, inputs, true) {
			sum2 += target
		}
	}
	//Solve(line)
	answer := strconv.Itoa(sum)
	answer2 := strconv.Itoa(sum2)
	fmt.Println(answer)
	fmt.Println(answer2)
}

func parseEquation(eq string) (int, []int) {
	var equation []int
	result := strings.Split(eq, ": ")
	target, _ := strconv.Atoi(result[0])

	operands := strings.Split(result[1], " ")
	for _, numberString := range operands {
		number, _ := strconv.Atoi(numberString)

		equation = append(equation, number)
	}

	return target, equation
}

func solveEquation(currIdx, solution, target int, inputs []int, concatenation bool) bool {

	if currIdx == len(inputs) {
		return solution == target
	}

	curr := inputs[currIdx]

	if solveEquation(currIdx+1, solution+curr, target, inputs, concatenation) {
		return true
	}
	if solveEquation(currIdx+1, solution*curr, target, inputs, concatenation) {
		return true
	}

	if concatenation {
		concantNumber, _ := strconv.Atoi(strconv.Itoa(solution) + strconv.Itoa(inputs[currIdx]))

		concatValues := []int{concantNumber}
		concatValues = append(concatValues, inputs[currIdx+1:]...)

		if solveEquation(1, concatValues[0], target, concatValues, concatenation) {
			return true
		}
	}

	return false

}
