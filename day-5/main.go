package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

func main() {
	// open file with the input values
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal("Cannot open the input file.")
	}
	defer file.Close()

	// open a scanner to read from the file
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	// process all lines
	max := 0
	var ids []int
	for scanner.Scan() {
		val := scanner.Text()
		id := 0

		// plane rows
		addend := 128
		for _, l := range val[:7] {
			addend /= 2
			if l == 'B' {
				id += addend
			}
		}
		id *= 8

		// plane cols
		addend = 8
		for _, l := range val[7:] {
			addend /= 2
			if l == 'R' {
				id += addend
			}
		}

		// test the highest value
		if id > max {
			max = id
		}

		// store in id array
		ids = append(ids, id)
	}

	// print the result
	fmt.Println("Part 1:", max)

	// sort the ids and find the missing number
	sort.Ints(ids)
	for i, id := range ids {
		next := id + 1
		if ids[i+1] != next {
			fmt.Println("Part 2:", next)
			break
		}
	}
}
