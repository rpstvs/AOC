package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Changes [4]int

func main() {
	solvept1()
	solvept2()
}

func solvept1() {
	nums := parseInput()
	sum := 0
	for _, num := range nums {
		for i := 0; i < 2000; i++ {
			num = Compute(num)
		}
		sum += num
	}
	fmt.Println(sum)

}

func solvept2() {
	nums := parseInput()

	priceAfter := map[int]map[Changes]int{}
	discovered := map[Changes]map[int]struct{}{}

	for buyerid, secret := range nums {
		if _, ok := priceAfter[buyerid]; !ok {
			priceAfter[buyerid] = map[Changes]int{}
		}
		ring, i := Changes{}, 0

		oldPrice := secret % 10

		for t := 1; t <= 2000; t++ {
			secret = Compute(secret)

			newPrice := secret % 10

			delta := newPrice - oldPrice

			oldPrice = newPrice

			ring[i], i = delta, (i+1)%4
			if t >= 4 {
				change := Changes{ring[i], ring[(i+1)%4], ring[(i+2)%4], ring[(i+3)%4]}
				if _, ok := discovered[change]; !ok {
					discovered[change] = map[int]struct{}{}
				}
				discovered[change][buyerid] = struct{}{}

				if _, ok := priceAfter[buyerid][change]; !ok {
					priceAfter[buyerid][change] = newPrice
				}
			}
		}
	}

	maxBananas := 0

	for change, affectedIds := range discovered {
		bananas := 0
		for buyerid := range affectedIds {
			bananas += priceAfter[buyerid][change]
		}
		if bananas > maxBananas {
			maxBananas = bananas
		}
	}
	fmt.Println(maxBananas)

}

func Compute(secret int) int {
	secret = ((secret * 64) ^ secret) % 16777216
	secret = ((secret / 32) ^ secret) % 16777216
	secret = ((secret * 2048) ^ secret) % 16777216
	return secret
}

func parseInput() []int {
	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	nums := []int{}

	for scanner.Scan() {
		val, _ := strconv.Atoi(scanner.Text())

		nums = append(nums, val)
	}
	return nums
}
