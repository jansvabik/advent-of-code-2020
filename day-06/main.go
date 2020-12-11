package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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
	anyone := 0
	everyone := 0
	firstPerson := true
	var anyoneLetters = make(map[rune]struct{})
	var everyoneLetters = make(map[rune]struct{})
	for scanner.Scan() {
		line := scanner.Text()

		// process the data if empty line
		if len(line) == 0 {
			anyone += len(anyoneLetters)
			everyone += len(everyoneLetters)
			anyoneLetters = make(map[rune]struct{})
			everyoneLetters = make(map[rune]struct{})
			firstPerson = true
			continue
		}

		for _, l := range line {
			// add all letters uniquely to anyoneLetters
			if _, ok := anyoneLetters[l]; !ok {
				anyoneLetters[l] = struct{}{}
			}

			// if this is the first person, add its letters to everyoneLetters
			if firstPerson {
				everyoneLetters[l] = struct{}{}
			}
		}

		// if this is not the first person, compare letters in everyoneLetters and line
		if !firstPerson {
			for s := range everyoneLetters {
				found := false
				for _, l := range line {
					if s == l {
						found = true
					}
				}

				// if the tested everyoneLetter is not in line, remove it
				if !found {
					delete(everyoneLetters, s)
				}
			}
		}

		// in the next round, it won't be the first person
		firstPerson = false
	}

	// add last line
	anyone += len(anyoneLetters)
	everyone += len(everyoneLetters)

	// print the gold star results
	fmt.Println("Part 1:", anyone)
	fmt.Println("Part 2:", everyone)
}
