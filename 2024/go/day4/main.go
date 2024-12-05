package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

const charIgnore = ' '

var masks = [][]string{
	{
		"XMAS",
	},
	{
		"SAMX",
	},
	{
		"X",
		"M",
		"A",
		"S",
	},
	{
		"S",
		"A",
		"M",
		"X",
	},
	{
		"X   ",
		" M  ",
		"  A ",
		"   S",
	},
	{
		"   X",
		"  M ",
		" A  ",
		"S   ",
	},
	{
		"S   ",
		" A  ",
		"  M ",
		"   X",
	},
	{
		"   S",
		"  A ",
		" M  ",
		"X   ",
	},
}

var masks2 = [][]string{
	{
		"M S",
		" A ",
		"M S",
	},
	{
		"S S",
		" A ",
		"M M",
	},
	{
		"S M",
		" A ",
		"S M",
	},
	{
		"M M",
		" A ",
		"S S",
	},
}

func main() {

	err := error(nil)
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("Error: %s", err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)
	lines := make([]string, 0)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	fmt.Println(lines)
	sum := 0
	sum2 := 0

	for _, mask := range masks {
		sum += getResults(mask, lines)
	}

	for _, mask2 := range masks2 {
		sum2 += Part2(mask2, lines)
	}

	fmt.Println(sum)
	fmt.Println(sum2)

}

func getResults(mask []string, lines []string) int {

	count := 0
	for y := 0; y <= len(lines)-len(mask); y++ {
		for x := 0; x <= len(lines[0])-len(mask[0]); x++ {
			found := 0
			total := 0

			for dy := 0; dy < len(mask); dy++ {
				for dx := 0; dx < len(mask[0]); dx++ {
					if mask[dy][dx] != charIgnore {
						total++
					}
					if mask[dy][dx] == lines[y+dy][x+dx] {
						found++
					}
				}
			}
			if found == total {
				count++
			}

		}
	}
	return count
}

func Part2(mask []string, lines []string) int {
	count := 0
	for y := 0; y <= len(lines)-len(mask); y++ {
		for x := 0; x <= len(lines[0])-len(mask[0]); x++ {
			found := 0
			total := 0

			for dy := 0; dy < len(mask); dy++ {
				for dx := 0; dx < len(mask[0]); dx++ {
					if mask[dy][dx] != charIgnore {
						total++
					}
					if mask[dy][dx] == lines[y+dy][x+dx] {
						found++
					}
				}
			}
			if found == total {
				count++
			}

		}
	}
	return count
}
