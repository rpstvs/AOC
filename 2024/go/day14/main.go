package main

import (
	"bufio"
	"fmt"
	"os"
)

type Point struct {
	x, y int
}

type robot struct {
	p   Point
	vel Point
}

func main() {
	robots := parseInput()

	//steps(robots)
	sol2 := 0
	for sol2 = 1; sol2 < 1000000; sol2++ {
		updatePos(robots)
		q := checkcenter(robots)

		fmt.Println("number robots: ", q)
		if q > len(robots)/2 {

			fmt.Println(sol2)
			return
		}

	}
	qs := checkQuarter(robots)

	for _, q := range qs {
		if q > len(robots)/2 {
			fmt.Println(sol2)
		}

	}

}

func parseInput() []robot {

	file, _ := os.Open("example.txt")

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var x, y, vx, vy int
	robotMap := []robot{}
	for scanner.Scan() {
		line := scanner.Text()

		fmt.Sscanf(line, "p=%d,%d v=%d,%d", &x, &y, &vx, &vy)
		robotMap = append(robotMap, robot{Point{x, y}, Point{vx, vy}})

	}
	return robotMap
}

func updatePos(robots []robot) {

	for i, _ := range robots {
		robots[i].p.x = (robots[i].p.x + robots[i].vel.x + 101) % 101
		robots[i].p.y = (robots[i].p.y + robots[i].vel.y + 103) % 103
	}

}

func checkQuarter(robots []robot) []int {
	q1, q2, q3, q4 := 0, 0, 0, 0
	qs := []int{}
	for i := 0; i < len(robots); i++ {
		if robots[i].p.x < 101/2 && robots[i].p.y < 103/2 {
			q1++
		}
		if robots[i].p.x > 101/2 && robots[i].p.y < 103/2 {
			q2++
		}
		if robots[i].p.x < 101/2 && robots[i].p.y > 103/2 {
			q3++
		}
		if robots[i].p.x > 101/2 && robots[i].p.y > 103/2 {
			q4++
		}
	}

	qs = append(qs, q1, q2, q3, q4)
	return qs
}

func checkcenter(robots []robot) int {
	q := 0

	for i := 0; i < len(robots); i++ {
		if robots[i].p.x < 75 && robots[i].p.x > 25 && robots[i].p.y < 75 && robots[i].p.y > 25 {
			q++
		}

	}

	return q
}

func steps(robots []robot) {

	total := len(robots)

	for steps := 1; steps < 1000000; steps++ {
		seen := map[Point]struct{}{}
		for _, robot := range robots {
			nextX := (robot.p.x + steps*robot.vel.x) % 101
			nextY := (robot.p.y + steps*robot.vel.y) % 103

			if _, ok := seen[Point{nextX, nextY}]; !ok {
				seen[Point{nextX, nextY}] = struct{}{}
			}
		}
		if len(seen) == total {
			fmt.Println(steps)
			break
		}
	}
}
