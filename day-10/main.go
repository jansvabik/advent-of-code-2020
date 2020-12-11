package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

// calculateAdapterDifferences calculates the number of diffs
// in adapter combinations trying to reach the maximum possible joltage
func calculateAdapterDifferences(nums []int) (int, int, int) {
	joltage := 0
	diff := [3]int{0, 0, 0}
	for true {
		startJoltage := joltage
		found := false
		for i := 1; !found && i <= 3; i++ {
			for _, n := range nums {
				if joltage+i == n {
					joltage = n
					diff[i-1]++
					found = true
					break
				}
			}
		}

		// if the adapter was found, try to find the next one
		if found {
			continue
		}

		// if there was no change in joltage, there is no next adapter
		if startJoltage == joltage {
			break
		}
	}

	return diff[0], diff[1], diff[2] + 1
}

func main() {
	// open file with the input values
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal("Cannot open the input file.")
		os.Exit(1)
	}
	defer file.Close()

	// open a scanner to read from the file
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	// create a slice of all numbers
	var nums []int
	for scanner.Scan() {
		val, _ := strconv.Atoi(scanner.Text())
		nums = append(nums, val)
	}

	// calculate all the differences to the max
	d1, _, d3 := calculateAdapterDifferences(nums)
	fmt.Println("Part 1:", d1*d3)
}
