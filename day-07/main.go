package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// bagContent is a structure of bags which can be carried in another bag
type bagContent struct {
	name   string
	number int
}

// canCarry determines if the "search" bag can be carried in the "in" bag
func canCarry(search string, in string, bags map[string][]bagContent) bool {
	bag := bags[in]

	if len(bag) == 0 {
		return false
	}

	// test the current and the nested bags
	for _, n := range bag {
		if n.name == search || canCarry(search, n.name, bags) {
			return true
		}
	}

	return false
}

// haveToContain calculates how many bags the specified bag have to contain
func haveToContain(name string, bags map[string][]bagContent) int {
	if len(bags[name]) == 0 {
		return 0
	}

	// summarize the number of nested bags
	res := 0
	for _, v := range bags[name] {
		res += v.number * (1 + haveToContain(v.name, bags))
	}
	return res
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

	// process all lines
	bags := make(map[string][]bagContent)
	for scanner.Scan() {
		val := scanner.Text()

		// extract the bag data into better form
		bag := strings.Split(val[:len(val)-1], " bags contain ")
		color := bag[0]
		contains := regexp.MustCompile(" bags?(, )?").Split(bag[1], -1)
		contains = contains[:len(contains)-1]

		// store all bags in bags map
		bags[color] = []bagContent{}
		for _, v := range contains {
			c := strings.SplitN(v, " ", 2)

			// store only if the bag can content any other one
			if c[1] != "other" {
				num, _ := strconv.Atoi(c[0])
				bags[color] = append(bags[color], bagContent{
					name:   c[1],
					number: num,
				})
			}
		}
	}

	// calculate how many bags can carry my bag
	carryable := 0
	for k := range bags {
		if canCarry("shiny gold", k, bags) {
			carryable++
		}
	}

	// print the result
	fmt.Println("Part 1:", carryable)
	fmt.Println("Part 2:", haveToContain("shiny gold", bags))
}
