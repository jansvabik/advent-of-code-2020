package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
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

	// create a slice of all numbers
	var numbers []int
	for scanner.Scan() {
		val, _ := strconv.Atoi(scanner.Text())
		numbers = append(numbers, val)
	}

	preamble := numbers[:25] // preamble list
	npri := 0                // next preamble replacement index
	invalidNumber := -1      // the invalid number

	// find the invalid number
	for _, val := range numbers[25:] {
		valid := false
		for i, a := range preamble {
			for j, b := range preamble {
				if i == j {
					continue
				}

				if a+b == val {
					valid = true
				}
			}
		}

		// store the invalid number if found
		if !valid {
			invalidNumber = val
			fmt.Println("Part 1:", val)
			break
		}

		// refresh preamble array
		preamble[npri] = val
		npri++
		if npri >= len(preamble) {
			npri = 0
		}
	}

	// find contiguous set - the weakness
	for i, v := range numbers {
		acc, min, max := v, v, v
		for _, w := range numbers[i+1:] {
			acc += w

			// store the min and max values
			if w < min {
				min = w
			}
			if w > max {
				max = w
			}

			// print the result on match
			if acc == invalidNumber {
				fmt.Println("Part 2:", min+max)
			}
		}
	}
}
