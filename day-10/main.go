package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

// calculateAdapterDifferences calculates the number of diffs
// in adapter combinations trying to reach the maximum possible joltage
func calculateAdapterDifferences(nums []int) (int, int, int) {
	joltage := 0
	diff := [3]int{0, 0, 0}
	for {
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

// calculatePossibilities calculates all the possible combinations
// of the adapters in bag, the number is returned from the function
func calculatePossibilities(joltage int, nums []int) int {
	poss := 1

	totalComb := 1
	subtract := 0
	for i, n := range nums[:len(nums)-1] {
		comb := 0
		for j := 1; j < 4; j++ {
			if i+j < len(nums) {
				if nums[i+j] < n+4 {
					if nums[i+j] > n+1 {
						fmt.Println(n, nums[i+1])
						subtract++
					}
					comb++
				}
			}
		}
		totalComb *= comb
	}

	totalComb -= subtract

	fmt.Println("totalcomb", totalComb)

	return poss
}

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

	// create a slice of all numbers
	var nums []int
	for scanner.Scan() {
		val, _ := strconv.Atoi(scanner.Text())
		nums = append(nums, val)
	}
	sort.Ints(nums)

	// calculate all the differences to the max
	d1, _, d3 := calculateAdapterDifferences(nums)
	fmt.Println("Part 1:", d1*d3)

	// calculate all the possible combinations
	fmt.Println("Part 2:", calculatePossibilities(0, nums))
}
