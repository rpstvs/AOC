package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const size = 70

type Point struct {
	x, y int
}

var (
	start = Point{0, 0}
	end   = Point{size, size}
	Dirs  = []Point{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
)

func main() {
	mem := parseinput()
	corrupted := map[Point]struct{}{}
	for _, p := range mem[:1024] {
		corrupted[p] = struct{}{}
	}
	fmt.Println(len(BFS(corrupted)))
}

func parseinput() []Point {
	var points []Point
	file, _ := os.Open("example.txt")

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		vals := strings.Split(line, ",")
		val1, _ := strconv.Atoi(vals[0])
		val2, _ := strconv.Atoi(vals[1])

		p := Point{val1, val2}
		points = append(points, p)
	}
	return points
}

func BFS(corrupted map[Point]struct{}) []Point {
	queue := []Point{start}
	visited := map[Point]struct{}{start: {}}
	parent := map[Point]Point{}

	for len(queue) > 0 {
		var curr Point
		curr, queue = queue[0], queue[1:]

		if curr == end {
			var path []Point
			for p := end; p != start; p = parent[p] {
				path = append(path, p)

			}
			return path
		}
		for _, d := range Dirs {

			n := Point{curr.x + d.x, curr.y + d.y}

			if n.x < 0 || n.x > size || n.y < 0 || n.y > size {
				continue

			}

			if _, ok := corrupted[n]; ok {
				continue
			}

			if _, ok := visited[n]; !ok {
				visited[n] = struct{}{}
				parent[n] = curr
				queue = append(queue, n)
			}

		}
	}
	return nil
}
