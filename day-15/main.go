package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	// read the puzzle input
	b, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal("Cannot open the input file.")
	}
	str := string(b)

	// extract the data
	startingNums := strings.Split(str, ",")
	numSaid := map[int][2]int{}
	for i, n := range startingNums {
		ni, _ := strconv.Atoi(n)
		numSaid[ni] = [2]int{0, i + 1}
	}

	// run the loop for the next turns
	lastNum, _ := strconv.Atoi(startingNums[len(startingNums)-1])
	for turn := len(startingNums) + 1; turn <= 2020; turn++ {
		// determine the current number
		var currentNum int
		if numSaid[lastNum][0] == 0 {
			currentNum = 0
		} else {
			currentNum = numSaid[lastNum][1] - numSaid[lastNum][0]
		}

		// store current number's turn
		if _, ok := numSaid[currentNum]; ok {
			numSaid[currentNum] = [2]int{numSaid[currentNum][1], turn}
		} else {
			numSaid[currentNum] = [2]int{0, turn}
		}

		// current num now becomes the latest one
		lastNum = currentNum
	}

	fmt.Println("Part 1:", lastNum)
}
