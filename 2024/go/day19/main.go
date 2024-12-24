package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	sum := 0
	patterns, designs := parseinput()
	ways := 0
	cache := map[string]int{}
	for _, design := range designs {
		if part1(patterns, design) {
			sum += 1
		}

	}

	for _, design := range designs {
		ways += part2(patterns, design, cache)
	}
	fmt.Println(ways)
}

func parseinput() ([]string, []string) {
	file, _ := os.Open("input.txt")

	defer file.Close()

	scanner := bufio.NewScanner(file)
	patterns := []string{}
	designs := []string{}
	for scanner.Scan() {
		if scanner.Text() == "" {
			break
		}
		line := strings.Split(scanner.Text(), ", ")

		for _, pattern := range line {

			patterns = append(patterns, pattern)

		}

	}

	for scanner.Scan() {
		line := scanner.Text()

		designs = append(designs, line)
	}

	return patterns, designs
}

func part1(patterns []string, d string) bool {
	if len(d) == 0 {
		return true
	}

	for _, p := range patterns {
		if strings.HasPrefix(d, p) {
			found := part1(patterns, d[len(p):])

			if found {
				return true
			}
		}
	}
	return false
}

func part2(patterns []string, d string, cache map[string]int) int {
	if len(d) == 0 {
		return 1
	}

	if w, ok := cache[d]; ok {
		return w
	}
	w := 0
	for _, p := range patterns {
		if strings.HasPrefix(d, p) {
			w += part2(patterns, d[len(p):], cache)
		}
	}
	cache[d] = w
	return w
}
