package main

import (
	"bufio"
	"fmt"
	"os"
)

type Point struct {
	x, y int
}

func main() {
	grid, size := parseInput()
	price := 0
	price2 := 0

	visited := map[Point]struct{}{}

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if _, ok := visited[Point{i, j}]; ok {
				continue

			}
			//area, fence := explore(grid, Point{i, j}, visited)
			area2, fence2 := explore2(grid, Point{i, j}, visited)
			//price += area * fence
			price2 += area2 * fence2
		}

	}
	fmt.Println(price, price2)
}

func parseInput() (map[Point]uint8, int) {
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	grid := map[Point]uint8{}
	var i int
	for i = 0; scanner.Scan(); i++ {
		line := scanner.Text()
		for j, c := range line {
			grid[Point{x: j, y: i}] = uint8(c)
		}

	}

	return grid, i
}

func explore(grid map[Point]uint8, p Point, visited map[Point]struct{}) (int, int) {
	visited[p] = struct{}{}

	perimeter, area := 0, 0
	dirs := []Point{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	var c Point
	toDo := []Point{p}

	for len(toDo) > 0 {
		c, toDo = toDo[0], toDo[1:]
		area++

		for _, dir := range dirs {
			fmt.Println("im here")
			next := Point{c.x + dir.x, c.y + dir.y}
			val, ok := grid[next]
			if ok && val == grid[p] {

				if _, ok := visited[next]; !ok {
					toDo = append(toDo, next)
					visited[next] = struct{}{}
				}
			} else {
				perimeter++
			}
		}

	}

	return area, perimeter
}

func explore2(m map[Point]uint8, p Point, visited map[Point]struct{}) (int, int) {
	visited[p] = struct{}{}

	id := m[p]

	area, corners := 0, 0

	var c Point
	toDo := []Point{p}
	for len(toDo) > 0 {
		c, toDo = toDo[0], toDo[1:]

		area++

		for _, d := range []Point{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} {
			next := Point{c.x + d.x, c.y + d.y}
			val, ok := m[next]
			if ok && val == id {
				if _, ok := visited[next]; !ok {
					toDo = append(toDo, next)
					visited[next] = struct{}{}
				}
			}
		}

		for _, d := range []Point{{-1, -1}, {1, 1}, {1, -1}, {-1, 1}} {
			// convex corner
			if m[Point{c.x + d.x, c.y}] != id &&
				m[Point{c.x, c.y + d.y}] != id {
				corners++
			}
			// concave corner
			if m[Point{c.x + d.x, c.y}] == id &&
				m[Point{c.x, c.y + d.y}] == id &&
				m[Point{c.x + d.x, c.y + d.y}] != id {
				corners++
			}
		}

	}

	return area, corners
}
