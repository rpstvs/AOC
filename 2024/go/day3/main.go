package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	//readFile, err := os.ReadFile("input.txt")

	//regex := regexp.MustCompile("mul\\((\\d{1,3},\\d{1,3})\\)")

	//matches := regex.FindAllStringSubmatch(string(readFile), -1)

	//sum = getSum(matches)
	DoSum()

}

func getSum(matches [][]string) int {
	multList := make([][]int, len(matches))
	var sum int
	for i, _ := range matches {
		mult := strings.Split(matches[i][1], ",")
		v1, _ := strconv.Atoi(mult[0])
		v2, _ := strconv.Atoi(mult[1])
		sum += v1 * v2
		multList[i] = []int{v1, v2}
	}
	return sum
}

func DoSum() int {
	readFile, err := os.ReadFile("input.txt")

	if err != nil {
		fmt.Println(err)

	}

	dontMult := strings.Split(string(readFile), "don't()")
	doSplits := make([]string, 0)
	for i, val := range dontMult {
		if i == 0 {
			doSplits = append(doSplits, val)
			continue
		}
		doSplit := strings.Split(val, "do()")
		doSplits = append(doSplits, doSplit[1:]...)
	}
	sum := 0

	doString := strings.Join(doSplits, "")
	regex := regexp.MustCompile("mul\\((\\d{1,3},\\d{1,3})\\)")

	matches := regex.FindAllStringSubmatch(doString, -1)

	sum = getSum(matches)

	println(sum)

	return sum
}
