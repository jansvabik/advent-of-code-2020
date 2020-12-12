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

// fibonacci returns the n-th number of the fibonacci sequence
// thanks, Lukas Gurecky, for finding out that part 2 uses it!
func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

// calculatePossibilities calculates all the possible combinations
// of the adapters in bag, the number is returned from the function
func calculatePossibilities(joltage int, nums []int) int {
	nums = append([]int{0}, nums...)
	comb := 1

	// find groups of consecutive numbers
	for i := 0; i < len(nums); i++ {
		// calculate the number of consecutive numbers in this group
		lst := []int{nums[i]}
		for j := 1; j < 5; j++ {
			if i+j < len(nums) && nums[i+j] == lst[len(lst)-1]+1 {
				lst = append(lst, nums[i+j])
			}
		}

		// multiply the combinations by the ln-th number of fibonacci
		// if there are more than 2 numbers (means more than 1 combination)
		ln := len(lst)
		if ln > 2 {
			comb *= fibonacci(len(lst)+1) - 1
		}

		// skip the ln-1 next numbers
		i += ln - 1
	}

	return comb
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
