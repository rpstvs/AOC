package main

import (
	"bufio"
	"fmt"
	"os"
)

type Point struct {
	x, y int
}

var dirs = [4]Point{{x: 0, y: -1}, {x: 1, y: 0}, {x: 0, y: 1}, {x: -1, y: 0}}

func main() {
	grid := getinput()
	ans1, ans2 := findTrailHeads(grid)
	fmt.Println(ans1)
	fmt.Println(ans2)
}

func getinput() [][]int {
	file, _ := os.Open("input.txt")

	defer file.Close()

	scanner := bufio.NewScanner(file)
	var grid [][]int

	for scanner.Scan() {
		var line = scanner.Text()
		var row []int

		for _, char := range line {
			var num = int(char - '0')
			row = append(row, num)
		}
		grid = append(grid, row)

	}
	return grid
}

type point struct {
	x, y int
}

func findNext(input [][]int, current point) []point {
	validNextSteps := []point{}
	// can check left
	if current.x > 0 && input[current.y][current.x-1] == input[current.y][current.x]+1 {
		validNextSteps = append(validNextSteps, point{current.x - 1, current.y})
	}

	// can check right
	if current.x < len(input[0])-1 && input[current.y][current.x+1] == input[current.y][current.x]+1 {
		validNextSteps = append(validNextSteps, point{current.x + 1, current.y})
	}

	// can check up
	if current.y > 0 && input[current.y-1][current.x] == input[current.y][current.x]+1 {
		validNextSteps = append(validNextSteps, point{current.x, current.y - 1})
	}

	// can check down
	if current.y < len(input)-1 && input[current.y+1][current.x] == input[current.y][current.x]+1 {
		validNextSteps = append(validNextSteps, point{current.x, current.y + 1})
	}
	return validNextSteps
}

func findScore(input [][]int, start point, trailHeads map[point]struct{}, count int) (map[point]struct{}, int) {
	if input[start.y][start.x] == 9 {
		if _, ok := trailHeads[start]; !ok {
			trailHeads[start] = struct{}{}
		}
		return trailHeads, count + 1
	}
	nextSteps := findNext(input, start)
	if len(nextSteps) == 0 {
		return trailHeads, count
	}

	for _, step := range nextSteps {
		trailHeads, count = findScore(input, step, trailHeads, count)
	}
	return trailHeads, count
}

func findTrailHeads(input [][]int) (int, int) {
	countScore := 0
	countRating := 0
	for j, row := range input {
		for i, char := range row {
			if char == 0 {
				score, rating := findScore(input, point{i, j}, make(map[point]struct{}), 0)
				countScore += len(score)
				countRating += rating
			}
		}
	}
	return countScore, countRating
}
