package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

type Rule struct {
	left, op, right string
}

func main() {
	solvept1()
	solvept2()
}

func solvept1() {
	facts, rules := parseInput()
	z := GetNumber("z", facts, rules)
	fmt.Println(z)
}

func parseInput() (map[string]bool, map[string]Rule) {
	facts := map[string]bool{}
	file, _ := os.Open("input.txt")

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			break
		}

		ff := strings.Split(line, ": ")
		facts[ff[0]] = ff[1] == "1"

	}

	var Rules = map[string]Rule{}

	for scanner.Scan() {
		ff := strings.Fields(scanner.Text())
		Rules[ff[4]] = Rule{left: ff[0], op: ff[1], right: ff[2]}
	}
	return facts, Rules

}

func GetNumber(letter string, facts map[string]bool, rules map[string]Rule) int {
	var z int

	for i := 0; i < 64; i++ {

		if solve(fmt.Sprintf("%s%02d", letter, i), facts, rules) {
			z |= 1 << i
		}
	}
	return z
}

func solve(target string, facts map[string]bool, rules map[string]Rule) bool {

	if v, ok := facts[target]; ok {
		return v
	}

	rule, ok := rules[target]

	if !ok {
		return false
	}

	left := solve(rule.left, facts, rules)
	right := solve(rule.right, facts, rules)

	var result bool

	switch rule.op {
	case "AND":
		result = left && right
	case "OR":
		result = left || right
	case "XOR":
		result = left != right
	}

	facts[target] = result
	return result
}

func solvept2() {
	_, rules := parseInput()

	GenerateDot(rules)

	list := []string{"z14", "z27", "z39", "msq", "qwf", "mps", "cnk", "vhm"}
	sort.Strings(list)
	fmt.Println(strings.Join(list, ","))

}

func GenerateDot(rules map[string]Rule) {
	f, err := os.Create("graph.dot")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	color := map[string]string{
		"XOR": "red",
		"AND": "blue",
		"OR":  "green",
	}

	fmt.Fprint(f, "digraph {\n")
	for rName, r := range rules {
		fmt.Fprintf(f, "%s -> %s [color=\"%s\"];\n", r.left, rName, color[r.op])
		fmt.Fprintf(f, "%s -> %s [color=\"%s\"];\n", r.right, rName, color[r.op])
	}

	var xys, zs []string
	for i := 0; i <= 45; i++ {
		xys = append(xys, fmt.Sprintf("x%02d", i))
		xys = append(xys, fmt.Sprintf("y%02d", i))
		zs = append(zs, fmt.Sprintf("z%02d", i))
	}
	fmt.Fprintf(f, "{rank = min;\n %s ; \n};\n", strings.Join(xys, " -> "))
	fmt.Fprintf(f, "{rank = max;\n %s ; \n};\n", strings.Join(zs, " -> "))
	fmt.Fprint(f, "}\n")

}
