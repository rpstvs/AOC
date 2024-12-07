package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Point struct {
	x int
	y int
}

var loops int

var dirs = [4]Point{{x: 0, y: -1}, {x: 1, y: 0}, {x: 0, y: 1}, {x: -1, y: 0}}

func main() {
	file, _ := os.Open("input.txt")

	defer file.Close()

	scanner := bufio.NewScanner(file)
	var line []string
	for scanner.Scan() {
		line2 := strings.Split(scanner.Text(), "\n")
		line = append(line, line2...)
	}

	sol := findGuard(line)
	steps := 0
	seen := make(map[Point]int)
	findCycles(line, sol, dirs[0], steps, seen)
	walk(line, sol, steps, dirs[0], seen)
	fmt.Println(loops)

}

func findGuard(labmap []string) Point {
	for i := 0; i < len(labmap); i++ {
		row := labmap[i]

		for j := 0; j < len(row); j++ {

			if row[j] == 94 {
				fmt.Println("got him")

				return Point{x: j, y: i}
			}
		}
	}
	return Point{}
}

func walk(labmap []string, curr Point, steps int, dir Point, seen map[Point]int) bool {
	steps++
	if steps > 10000 {
		loops += 1
		fmt.Println(labmap)
		return true
	}
	if curr.y >= len(labmap) || curr.x >= len(labmap[0]) || curr.x < 0 || curr.y < 0 {
		return true
	}
	if labmap[curr.y][curr.x] == '#' {
		return false
	}

	_, ok := seen[curr]
	if !ok {
		seen[curr] = 0
	}

	if walk(labmap, Point{x: curr.x + dir.x, y: curr.y + dir.y}, steps, dir, seen) {
		return true
	}
	newdir := changeDirection(dir)

	if walk(labmap, Point{x: curr.x + newdir.x, y: curr.y + newdir.y}, steps, newdir, seen) {
		return true
	}

	return false
}

func changeDirection(dir Point) Point {
	var newDirection Point
	for i, nextDir := range dirs {
		if dir.x == nextDir.x && dir.y == nextDir.y {
			if i == 3 {
				newDirection.x = dirs[0].x
				newDirection.y = dirs[0].y
				return newDirection
			}
			newDirection.x = dirs[i+1].x
			newDirection.y = dirs[i+1].y
			return newDirection
		}
	}
	return Point{}
}

func findCycles(labmap []string, curr Point, dir Point, steps int, seen map[Point]int) {

	for i := 0; i < len(labmap); i++ {
		row := []rune(labmap[i])
		for j := 0; j < len(row); j++ {
			tmp := labmap[i]
			if row[j] == 94 || row[j] == '#' {

				continue
			}
			row[j] = '#'
			labmap[i] = string(row)

			walkCycle(labmap, curr, steps, dir, seen)
			row[j] = '.'
			labmap[i] = tmp

		}
	}
	//fmt.Println(sum)
}

func walkCycle(labmap []string, curr Point, steps int, dir Point, seen map[Point]int) bool {

	steps++
	if steps > 100000 {
		loops++

		return true
	}
	if curr.y >= len(labmap) || curr.x >= len(labmap[0]) || curr.x < 0 || curr.y < 0 {
		return true
	}
	if labmap[curr.y][curr.x] == '#' {
		return false
	}

	_, ok := seen[curr]
	if !ok {
		seen[curr] = 0
	}

	if walkCycle(labmap, Point{x: curr.x + dir.x, y: curr.y + dir.y}, steps, dir, seen) {
		return true
	}
	dir = changeDirection(dir)

	if walkCycle(labmap, Point{x: curr.x + dir.x, y: curr.y + dir.y}, steps, dir, seen) {
		return true
	}

	dir = changeDirection(dir)

	if walkCycle(labmap, Point{x: curr.x + dir.x, y: curr.y + dir.y}, steps, dir, seen) {
		return true
	}
	dir = changeDirection(dir)

	if walkCycle(labmap, Point{x: curr.x + dir.x, y: curr.y + dir.y}, steps, dir, seen) {
		return true
	}

	return false

}
