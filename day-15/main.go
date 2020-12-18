package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

// getNumberAtPosition gets the number which is said at the given
// turn (position) in the Christmas Elves game and returns it
func getNumberAtPosition(lastStartNumber int, startNumbers int, said map[int][2]int, pos int) int {
	// run the loop for the next turns
	lastNum := lastStartNumber
	for turn := startNumbers + 1; turn <= pos; turn++ {
		// determine the current number
		var currentNum int
		if said[lastNum][0] == 0 {
			currentNum = 0
		} else {
			currentNum = said[lastNum][1] - said[lastNum][0]
		}

		// store current number's turn
		if _, ok := said[currentNum]; ok {
			said[currentNum] = [2]int{said[currentNum][1], turn}
		} else {
			said[currentNum] = [2]int{0, turn}
		}

		// current num now becomes the latest one
		lastNum = currentNum
	}
	return lastNum
}

func main() {
	// read the puzzle input
	b, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal("Cannot open the input file.")
	}
	str := string(b)

	// extract the data
	startingNums := strings.Split(str, ",")
	said1 := map[int][2]int{}
	said2 := map[int][2]int{}
	for i, n := range startingNums {
		ni, _ := strconv.Atoi(n)
		said1[ni] = [2]int{0, i + 1}
		said2[ni] = [2]int{0, i + 1}
	}

	// calculate the numbers at the given position and print the results
	last, _ := strconv.Atoi(startingNums[len(startingNums)-1])
	p2020 := getNumberAtPosition(last, len(startingNums), said1, 2020)
	p30000000 := getNumberAtPosition(last, len(startingNums), said2, 30000000)

	fmt.Println("Part 1:", p2020)
	fmt.Println("Part 2:", p30000000)
}
