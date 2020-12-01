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

	// create an array of values
	var numbers []int
	for scanner.Scan() {
		val, _ := strconv.Atoi(scanner.Text())
		numbers = append(numbers, val)
	}

	// test every combination of values
	stop := false
	for _, a := range numbers {
		for _, b := range numbers {
			for _, c := range numbers {
				if a+b+c == 2020 {
					fmt.Println(a * b * c)
					stop = true
				}

				// break the loop if the values were found
				if stop {
					break
				}
			}

			// break the loop if the values were found
			if stop {
				break
			}
		}

		// break the loop if the values were found
		if stop {
			break
		}
	}
}
