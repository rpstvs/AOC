package main

import (
	"bufio"
	"fmt"
	"os"
)

const Size = 50

type Point struct {
	x, y int
}

func main() {
	freqs := parseInput()
	solvePart1(freqs)
	solvePart2(freqs)
}

func (p Point) isValid() bool {
	return p.x >= 0 && p.x < Size && p.y >= 0 && p.y < Size
}

func parseInput() map[byte][]Point {
	file, _ := os.Open("input.txt")

	defer file.Close()
	scanner := bufio.NewScanner(file)
	freqs := map[byte][]Point{}
	for i := 0; scanner.Scan(); i++ {
		line := scanner.Text()

		for j, c := range line {
			if c == '.' {
				continue
			}
			ch := byte(c)
			freqs[ch] = append(freqs[ch], Point{j, i})
		}
	}
	fmt.Println(len(freqs))
	return freqs
}

func solvePart1(freqs map[byte][]Point) {

	antiNode := map[Point]struct{}{}

	for _, locs := range freqs {
		for a := 0; a < len(locs)-1; a++ {

			for b := a + 1; b < len(locs); b++ {
				delta := Point{locs[b].x - locs[a].x, locs[b].y - locs[a].y}

				anti1 := Point{locs[a].x - delta.x, locs[a].y - delta.y}

				if anti1.x >= 0 && anti1.x < Size && anti1.y >= 0 && anti1.y < Size {
					antiNode[anti1] = struct{}{}
				}

				anti2 := Point{locs[b].x + delta.x, locs[b].y + delta.y}

				if anti2.x >= 0 && anti2.x < Size && anti2.y >= 0 && anti2.y < Size {
					antiNode[anti2] = struct{}{}
				}

			}
		}
	}
	fmt.Println(len(antiNode))
}

func solvePart2(freqs map[byte][]Point) {

	antiNode := map[Point]struct{}{}

	for _, locs := range freqs {
		for a := 0; a < len(locs)-1; a++ {

			for b := a + 1; b < len(locs); b++ {
				delta := Point{locs[b].x - locs[a].x, locs[b].y - locs[a].y}

				outofbound := 0

				for period := 0; outofbound < 2; period++ {
					outofbound = 0
					anti1 := Point{locs[a].x - period*delta.x, locs[a].y - period*delta.y}

					if anti1.x >= 0 && anti1.x < Size && anti1.y >= 0 && anti1.y < Size {
						antiNode[anti1] = struct{}{}
					} else {
						outofbound++
					}

					anti2 := Point{locs[b].x + period*delta.x, locs[b].y + period*delta.y}

					if anti2.x >= 0 && anti2.x < Size && anti2.y >= 0 && anti2.y < Size {
						antiNode[anti2] = struct{}{}
					} else {
						outofbound++
					}
				}

			}
		}
	}
	fmt.Println(len(antiNode))
}
