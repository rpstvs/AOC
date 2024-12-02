package main

import (
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {

	simMap := map[int]int{}
	var similarity int
	dat, err := os.ReadFile("input.txt")

	if err != nil {
		fmt.Println(err)
	}

	lines := strings.Split(string(dat), "\n")

	var left []int
	var right []int
	for _, line := range lines {

		Numbers := strings.Split(line, "   ")
		leftNumber, _ := strconv.Atoi(Numbers[0])
		rightNumber, _ := strconv.Atoi(Numbers[1])
		left = append(left, leftNumber)
		right = append(right, rightNumber)
		simMap[leftNumber] = 0
	}

	slices.Sort(left)
	slices.Sort(right)
	var sum int
	for i, _ := range left {
		sum += int(math.Abs(float64(left[i]) - float64(right[i])))
		_, exists := simMap[right[i]]
		if exists {
			simMap[right[i]] += 1
		}
	}

	for _, val := range left {
		similarity += val * simMap[val]
	}

	fmt.Println(similarity)
}
