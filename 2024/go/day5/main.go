package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.Open("input.txt")

	defer file.Close()

	scanner := bufio.NewScanner(file)
	var updates [][]int
	var Pairs [][]int
	for scanner.Scan() {
		if len(scanner.Text()) == 0 {
			continue
		}
		if len(scanner.Text()) > 6 {
			pagesUpd := strings.Split(scanner.Text(), ",")
			var tmp []int
			for i := 0; i < len(pagesUpd); i++ {
				val1, _ := strconv.Atoi(pagesUpd[i])
				tmp = append(tmp, val1)
			}
			updates = append(updates, tmp)
		} else {
			pagesPairs := strings.Split(scanner.Text(), "|")
			val1, _ := strconv.Atoi(pagesPairs[0])
			val2, _ := strconv.Atoi(pagesPairs[1])
			intArr := []int{val1, val2}
			Pairs = append(Pairs, intArr)
		}

	}
	var newUpdate []int
	var sumMid int
	var sumNotOrdered int
	for _, update := range updates {
		newUpdate = getPrio(Pairs, update)
		if reflect.DeepEqual(update, newUpdate) {
			sumMid += getMiddleElem(newUpdate)
		} else {
			sumNotOrdered += getMiddleElem(newUpdate)
		}

	}

	fmt.Println(sumMid)
	fmt.Println(sumNotOrdered)

}

func getPrio(pairs [][]int, update []int) []int {
	newUpdate := make([]int, len(update))

	for i, _ := range update {
		prio := 0
		for j := 0; j < len(update); j++ {

			if foundinList([]int{update[i], update[j]}, pairs) {
				prio++
			}
		}
		newUpdate[(len(update)-1)-prio] = update[i]
	}
	return newUpdate
	// (len(newarrya)-1) - prio
}

func foundinList(pair []int, pairs [][]int) bool {
	var found bool
	for _, pair2 := range pairs {
		found = reflect.DeepEqual(pair, pair2)
		if found {
			return true
		} else {
			continue

		}
	}
	return found
}

func getMiddleElem(intArr []int) int {
	return intArr[len(intArr)/2]
}
