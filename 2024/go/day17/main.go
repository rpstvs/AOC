package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	a, b, c, program := parseinput()

	output := solve(a, b, c, program)
	fmt.Println(strings.Join(toStr(output), ","))
}

func parseinput() (int, int, int, []int) {
	var a, b, c int
	var program []int
	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
		if strings.HasPrefix(line, "Register A: ") {

			val := strings.Fields(line)
			a, _ = strconv.Atoi(val[2])
		}
		if strings.HasPrefix(line, "Register B: ") {
			val := strings.Fields(line)
			b, _ = strconv.Atoi(val[2])
		}
		if strings.HasPrefix(line, "Register C: ") {
			val := strings.Fields(line)
			c, _ = strconv.Atoi(val[2])
		}
		if strings.HasPrefix(line, "Program:") {
			val := strings.Fields(line)

			for _, ch := range val[1] {
				if ch == ',' {
					continue
				}
				val, _ := strconv.Atoi(string(ch))

				program = append(program, val)
			}
		}

	}
	return a, b, c, program
}

func solve(a, b, c int, code []int) []int {
	output := make([]int, 0)
	for ip := 0; ip < len(code); {
		var (
			op    = code[ip]
			arg   = code[ip+1]
			combo = 0
		)
		switch arg {
		case 1, 2, 3:
			combo = arg
		case 4:
			combo = a
		case 5:
			combo = b
		case 6:
			combo = c
		}
		switch op {
		case 0:
			a >>= combo
		case 1:
			b ^= arg
		case 2:
			b = combo % 8
		case 3:
			if a != 0 {
				ip = arg
				continue
			}
		case 4:
			b ^= c
		case 5:
			output = append(output, combo%8)
		case 6:
			b = a >> combo
		case 7:
			c = a >> combo
		}

		ip += 2
	}
	return output
}

func toStr(x []int) []string {
	s := make([]string, len(x))
	for i := 0; i < len(x); i++ {
		s[i] = strconv.Itoa(x[i])
	}
	return s
}
