package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	Pins   = 5
	height = 7
)

type schematics [Pins]int

func main() {
	keys, locks := parseInput()

	count := 0

	for _, key := range keys {
		for _, lock := range locks {
			if !overlap(key, lock) {
				count++
			}
		}
	}
	fmt.Println(count)
}

func overlap(key, lock schematics) bool {
	for i := 0; i < Pins; i++ {
		if key[i]+lock[i] > height-2 {
			return true
		}
	}
	return false
}

func parseInput() ([]schematics, []schematics) {
	var keys, locks []schematics

	file, _ := os.Open("input.txt")

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for {
		m := [height]string{}
		for i := 0; i < height; i++ {
			if !scanner.Scan() {
				return keys, locks
			}
			m[i] = scanner.Text()
		}
		val, isLock := parseSchema(m)

		if isLock {
			locks = append(locks, val)

		} else {
			keys = append(keys, val)
		}
		scanner.Scan()
	}
}

func parseSchema(m [height]string) (schematics, bool) {
	res := schematics{}

	for j := 0; j < Pins; j++ {
		count := 0

		for i := 1; i < height-1; i++ {
			if m[i][j] == '#' {
				count++
			}
		}
		res[j] = count
	}
	return res, m[0] == "#####" && m[height-1] == "....."
}
