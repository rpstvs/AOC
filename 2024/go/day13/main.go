package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Point struct {
	x, y int
}

type Eq struct {
	A, B, R Point
}

func main() {
	Equations := parseInput()

	tokens := 0
	for _, equation := range Equations {
		tokens += solve(equation)
	}

	fmt.Println(tokens)
}

func parseInput() []Eq {
	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	eqs := []Eq{}
	equation := Eq{}

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			equation = Eq{}
		}
		var x, y int
		switch {
		case strings.HasPrefix(line, "Button A"):
			fmt.Sscanf(line, "Button A: X+%d, Y+%d", &x, &y)
			equation.A = Point{x, y}
		case strings.HasPrefix(line, "Button B"):
			fmt.Sscanf(line, "Button B: X+%d, Y+%d", &x, &y)
			equation.B = Point{x, y}
		case strings.HasPrefix(line, "Prize"):
			fmt.Sscanf(line, "Prize: X=%d, Y=%d", &x, &y)
			equation.R = Point{x + 10000000000000, y + 10000000000000}
			eqs = append(eqs, equation)
		}

	}

	return eqs

}
func solve(Eq Eq) int {

	b := (Eq.A.x*Eq.R.y - Eq.A.y*Eq.R.x) / (Eq.A.x*Eq.B.y - Eq.A.y*Eq.B.x)
	a := (Eq.R.x - b*Eq.B.x) / Eq.A.x

	if a*Eq.A.x+b*Eq.B.x == Eq.R.x && a*Eq.A.y+b*Eq.B.y == Eq.R.y {
		return 3*a + b
	}
	return 0
}
