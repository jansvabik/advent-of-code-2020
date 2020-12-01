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
	printed2, printed3 := false, false
	for _, a := range numbers {
		for _, b := range numbers {
			if a+b == 2020 && !printed2 {
				fmt.Println("2 values: ", a*b)
				printed2 = true
			}

			for _, c := range numbers {
				if a+b+c == 2020 && !printed3 {
					fmt.Println("3 values: ", a*b*c)
					printed3 = true
				}
			}
		}
	}
}
