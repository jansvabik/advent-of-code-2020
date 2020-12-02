package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
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

	// create a data structure for every line
	valid1 := 0
	valid2 := 0
	for scanner.Scan() {
		val := scanner.Text()

		// store data for each line
		parts := strings.Split(val, " ")
		len := strings.Split(parts[0], "-")
		a, _ := strconv.Atoi(len[0])
		b, _ := strconv.Atoi(len[1])
		letter := string(parts[1][0])
		password := parts[2]

		// count the letter occurences
		occured := 0
		for _, l := range password {
			if string(l) == letter {
				occured++
			}
		}

		// test part 1 validity
		if occured >= a && occured <= b {
			valid1++
		}

		// test part 2 validity
		cond1 := string(password[a-1]) == letter && string(password[b-1]) != letter
		cond2 := string(password[a-1]) != letter && string(password[b-1]) == letter
		if cond1 || cond2 {
			valid2++
		}
	}

	// print the number of valid passwords
	fmt.Println("Part 1:", valid1)
	fmt.Println("Part 2:", valid2)
}
