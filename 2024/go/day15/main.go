package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

var Width, Height = 50, 50
var NotFound = errors.New("not found")

type Point struct {
	x, y int
}

func main() {

	grid, movements, robot := parseInput()
	fmt.Println(robot)
	for _, moves := range movements {
		robot = part1sol(grid, moves, robot)
	}

	fmt.Println(gps(grid, 'O'))

	solvePart2()

}

func solvePart2() {
	m, movements, robot := parseInput()
	m, robot = grow(m, robot)
	for _, moves := range movements {
		robot = apply2(m, moves, robot)
	}
	fmt.Println(gps(m, '['))
}

func parseInput() (map[Point]byte, []byte, Point) {
	var robot Point
	m := map[Point]byte{}

	file, _ := os.Open("input.txt")

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for i := 0; scanner.Scan(); i++ {
		line := scanner.Text()
		if line == "" {
			break

		}
		for j, c := range line {
			if c == '@' {
				robot = Point{i, j}
				m[Point{i, j}] = '.'

			} else {
				m[Point{i, j}] = byte(c)
			}
		}

	}

	moves := []byte{}
	for scanner.Scan() {
		moves = append(moves, []byte(scanner.Text())...)
	}
	return m, moves, robot

}

func getDelta(move byte) Point {

	switch move {
	case '<':
		return Point{0, -1}
	case '>':
		return Point{0, 1}
	case '^':
		return Point{-1, 0}
	case 'v':
		return Point{1, 0}

	}
	panic("invalid move")

}

func findeNextEmpty(m map[Point]byte, p, delta Point) (Point, error) {
	for {
		p = Point{p.x + delta.x, p.y + delta.y}

		switch m[p] {
		case '.':
			return p, nil
		case '#':
			return Point{}, NotFound

		}
	}
}

func gps(m map[Point]byte, char byte) int {
	sum := 0
	for i := 0; i < Height; i++ {
		for j := 0; j < Width; j++ {
			if m[Point{i, j}] == char {
				sum += 100*i + j
			}
		}
	}
	return sum
}

func part1sol(m map[Point]byte, move byte, robot Point) Point {
	delta := getDelta(move)

	nextEmpty, err := findeNextEmpty(m, robot, delta)

	if err != nil {
		return robot
	}

	closest := Point{robot.x + delta.x, robot.y + delta.y}

	if m[closest] == 'O' {
		m[closest] = '.'
		m[nextEmpty] = 'O'
	}
	return closest
}

func grow(m map[Point]byte, robot Point) (map[Point]byte, Point) {
	newM := make(map[Point]byte, len(m)*2)
	for i := 0; i < Height; i++ {
		for j := 0; j < Width; j++ {
			p := Point{i, j}
			switch m[p] {
			case '#':
				newM[Point{p.x, 2 * p.y}] = '#'
				newM[Point{p.x, 2*p.y + 1}] = '#'
			case 'O':
				newM[Point{p.x, 2 * p.y}] = '['
				newM[Point{p.x, 2*p.y + 1}] = ']'
			case '.':
				newM[Point{p.x, 2 * p.y}] = '.'
				newM[Point{p.x, 2*p.y + 1}] = '.'
			case '@':
				newM[Point{p.x, 2 * p.y}] = '@'
				newM[Point{p.x, 2*p.y + 1}] = '.'
			}
		}
	}
	Width *= 2
	return newM, Point{robot.x, 2 * robot.y}
}

func apply2(m map[Point]byte, move byte, robot Point) Point {
	delta := getDelta(move)

	nextEmpty, err := findeNextEmpty(m, robot, delta)
	if err != nil {
		return robot
	}

	if move == '<' || move == '>' {
		for curr := nextEmpty; curr != robot; {
			closest := Point{curr.x, curr.y - delta.y}
			m[curr], m[closest] = m[closest], m[curr]
			curr = closest
		}
		return Point{robot.x, robot.y + delta.y}
	}

	if move == '^' || move == 'v' {
		affected, maxLevel, err := affectedVertically(m, robot, delta.x)
		if err != nil {
			return robot
		}
		for x := maxLevel; x != robot.x; x -= delta.x {
			for col := range affected[x] {
				m[Point{x + delta.x, col}], m[Point{x, col}] = m[Point{x, col}], m[Point{x + delta.x, col}]
			}
		}

		return Point{robot.x + delta.x, robot.y}
	}

	panic("unreachable")
}

func affectedVertically(m map[Point]byte, robot Point, deltaX int) (map[int]map[int]struct{}, int, error) {
	affected := map[int]map[int]struct{}{
		robot.x: {robot.y: {}},
	}

	for currX := robot.x; ; currX += deltaX {
		newCols, err := newColumns(m, currX+deltaX, affected[currX])
		if err != nil {
			return nil, 0, err
		}

		if len(newCols) == 0 {
			return affected, currX, nil
		}

		affected[currX+deltaX] = newCols
	}
}

func newColumns(m map[Point]byte, nextX int, columns map[int]struct{}) (map[int]struct{}, error) {
	newCols := map[int]struct{}{}
	for col := range columns {
		switch m[Point{nextX, col}] {
		case '#':
			return nil, NotFound
		case '[':
			newCols[col] = struct{}{}
			newCols[col+1] = struct{}{}
		case ']':
			newCols[col] = struct{}{}
			newCols[col-1] = struct{}{}
		}
	}
	return newCols, nil
}
