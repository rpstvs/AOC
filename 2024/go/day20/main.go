package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

type Point struct {
	x, y int
}

var Delta = []Point{{0, 1}, {1, 0}, {-1, 0}, {0, -1}}

func main() {
	solvePt1(2, 100)
	solvePt2(20, 100)
}

func solvePt1(CheatTime, minSavedTime int) {

	walls, start, end := parseInput()

	path := bfs(walls, start, end)
	sum := 0
	for i := 0; i < len(path)-1; i++ {
		for j := i + 1; j < len(path); j++ {
			dist := distance(path[i], path[j])

			if dist > 0 && dist <= CheatTime {
				saved := j - i - dist
				if saved >= minSavedTime {
					sum++
				}
			}
		}
	}

	fmt.Println(sum)
}

func solvePt2(CheatTime, minSavedTime int) {

	walls, start, end := parseInput()

	path := bfs(walls, start, end)
	sum := 0
	for i := 0; i < len(path)-1; i++ {
		for j := i + 1; j < len(path); j++ {
			dist := distance(path[i], path[j])

			if dist > 0 && dist <= CheatTime {
				saved := j - i - dist
				if saved >= minSavedTime {
					sum++
				}
			}
		}
	}

	fmt.Println(sum)
}

func parseInput() (map[Point]struct{}, Point, Point) {
	file, _ := os.Open("example.txt")
	defer file.Close()
	walls := map[Point]struct{}{}
	scanner := bufio.NewScanner(file)
	var start, end Point
	for i := 0; scanner.Scan(); i++ {
		for j, ch := range scanner.Text() {
			switch ch {
			case '#':
				walls[Point{i, j}] = struct{}{}
			case 'S':
				start = Point{i, j}
			case 'E':
				end = Point{i, j}
			}

		}
	}
	return walls, start, end

}

func bfs(walls map[Point]struct{}, start, end Point) []Point {
	queue := []Point{start}
	visited := map[Point]struct{}{start: {}}
	parent := map[Point]Point{}

	for len(queue) > 0 {
		var curr Point

		curr, queue = queue[0], queue[1:]

		if curr == end {
			return getPath(parent, start, end)
		}

		for _, d := range Delta {

			next := Point{curr.x + d.x, curr.y + d.y}

			if _, ok := walls[next]; ok {
				continue

			}

			if _, ok := visited[next]; !ok {
				visited[next] = struct{}{}
				parent[next] = curr
				queue = append(queue, next)
			}
		}
	}
	return nil
}

func getPath(parent map[Point]Point, start, end Point) []Point {
	l := make([]Point, 0)

	for p := end; p != start; p = parent[p] {
		l = append(l, p)
	}
	l = append(l, start)
	slices.Reverse(l)
	return l
}

func distance(p1, p2 Point) int {
	return abs(p2.x-p1.x) + abs(p2.y-p1.y)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
