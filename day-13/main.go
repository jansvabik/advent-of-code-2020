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
	data := strings.Split(str, "\n")
	first, _ := strconv.Atoi(data[0])
	buses := []int{}
	for _, v := range strings.Split(data[1], ",") {
		if v != "x" {
			cv, _ := strconv.Atoi(v)
			buses = append(buses, cv)
		}
	}

	// find the airport bus with minimum waiting time
	minWaitTime := buses[0]
	minBusID := buses[0]
	for _, b := range buses {
		wtime := b - first%b
		if wtime < minWaitTime {
			minWaitTime = wtime
			minBusID = b
		}
	}

	// print the results
	fmt.Println("Part 1:", minWaitTime*minBusID)
}
