package main

import (
	"fmt"
	"strconv"
	"strings"
)

var iterations = 25

/*
if the stone is 0, mark it as 1

if the stone has even number of digits, separate it in 2 stones, MSB and LSB

Other stones , multiply by 2024
*/
func main() {
	result := convertString("7568 155731 0 972 1 6919238 80646 22")
	var sum int

	res := make(map[int]int, 0)

	for _, val := range result {
		res[val]++
	}

	for idx := 0; idx < 75; idx++ {
		res = iterate(res)
	}

	for _, count := range res {
		sum += count
	}
	fmt.Println(sum)
}

func evenDigits(num int) int {
	var slc []int

	for num > 0 {
		slc = append(slc, num%10)
		num /= 10
	}

	if len(slc)%2 != 0 {
		return 0
	}

	return 1
}

func iterate(nums map[int]int) map[int]int {

	res := make(map[int]int, len(nums))

	for val, count := range nums {
		isEvent := evenDigits(val)
		if val == 0 {
			res[1] += count

		} else if isEvent == 0 {
			val *= 2024
			res[val] += count
		} else {
			stone1, stone2 := splitStone(val)
			res[stone1] += count
			res[stone2] += count
		}

	}
	return res
}

func convertString(s string) []int {
	var result []int

	s1 := strings.Split(s, " ")
	for _, num := range s1 {
		v, _ := strconv.Atoi(num)
		result = append(result, v)
	}

	return result
}

func splitStone(num int) (int, int) {
	stoneString := strconv.Itoa(num)
	stone1, stone2 := stoneString[:len(stoneString)/2], stoneString[len(stoneString)/2:]
	v1, _ := strconv.Atoi(stone1)
	v2, _ := strconv.Atoi(stone2)

	return v1, v2
}
