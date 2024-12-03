package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	dat, _ := os.ReadFile("input.txt")

	//var safe bool

	var count int

	lines := strings.Split(string(dat), "\n")

	for _, line := range lines {
		numbers := strings.Split(line, " ")
		if !isReportSafe(numbers) {
			numbers2 := make([]string, len(numbers)-1)

			for i := 0; i < len(numbers); i++ {
				copy(numbers2, numbers[:i])
				copy(numbers2[i:], numbers[i+1:])
				fmt.Println(numbers2)
				if isReportSafe(numbers2) {
					count++
					break
				}
			}
		} else {
			count++
		}

	}
	fmt.Println(count)
}

func isReportSafe(numbers []string) bool {
	var direction []bool

	for i := 0; i < len(numbers)-1; i++ {
		a, _ := strconv.Atoi(numbers[i])
		b, _ := strconv.Atoi(numbers[i+1])
		sum := a - b

		if sum < 0 {
			direction = append(direction, false)
		} else {
			direction = append(direction, true)
		}

		level := int(math.Abs(float64(a) - float64(b)))

		if level < 1 || level > 3 || level == 0 {
			return false
		}

		if i >= 1 {
			if direction[i] != direction[i-1] {
				return false
			}
		}

	}
	return true
}
