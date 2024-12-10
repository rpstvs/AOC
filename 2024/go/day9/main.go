package main

import (
	"fmt"
	"os"
	"slices"
)

func main() {
	dat, _ := os.ReadFile("input.txt")
	memMap := convertStringtomap(string(dat))
	arrange(memMap)
	sol1 := calculateChecksum(memMap)
	sol2 := solve(string(dat))
	fmt.Println(sol2)
	fmt.Println(sol1)
}

func solve(input string) int {
	var line [][2]int
	for i, char := range input {
		n := -1
		if i%2 == 0 {
			n = i / 2
		}
		line = append(line, [2]int{n, int(char - '0')})
	}

	for i := len(line) - 1; i >= 0; i-- {
		if line[i][0] == -1 {
			continue
		}
		for j := 0; j < i; j++ {
			if line[j][0] != -1 {
				continue
			}
			if line[i][1] > line[j][1] {
				continue
			}

			line[i], line[j] = line[j], line[i]
			line = slices.Insert(line, j+1, [2]int{-1, line[i][1] - line[j][1]})
			line[i+1][1] = line[j][1]
			i++
			break
		}
	}

	total := 0
	c := -1
	for i := range line {
		for range line[i][1] {
			c++
			if line[i][0] == -1 {
				continue
			}
			total += line[i][0] * c
		}
	}
	return total
}

func convertStringtomap(s string) []int {

	var mapMemory []int
	var id int
	var freelen int

	for i := 0; i < len(s); i++ {

		for j := 0; j < int(s[i]-'0'); j++ {
			if i%2 == 0 {
				mapMemory = append(mapMemory, id)

			} else {
				mapMemory = append(mapMemory, -1)
				freelen++
			}
		}
		if i%2 == 0 {

			id++
		}
	}

	return mapMemory
}

func calculateChecksum(arr []int) int {
	var sum int
	for i := 0; i < len(arr); i++ {

		if arr[i] < 0 {
			continue

		}

		sum += arr[i] * i

	}

	return sum
}

func arrange(arr []int) {
	l := 0
	r := len(arr) - 1

	for {

		for arr[l] >= 0 {
			l++
		}

		for arr[r] < 0 {
			r--
		}

		if l >= r {
			break
		}

		arr[l] = arr[r]
		arr[r] = -1

	}
}
