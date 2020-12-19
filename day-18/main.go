package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"./advmath"
	"./simplemath"
)

func main() {
	// read the puzzle input
	b, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal("Cannot open the input file.")
	}

	// extract the puzzle input data (remove spaces)
	rows := strings.Split(strings.ReplaceAll(string(b), " ", ""), "\n")

	// sum the results of all the math expressions in input
	sum1, sum2 := 0, 0
	for _, row := range rows {
		sum1 += simplemath.CalculateExpression(row)
		sum2 += advmath.CalculateExpression(row)
	}

	// nicely done!
	fmt.Println("Part 1:", sum1)
	fmt.Println("Part 2:", sum2)
}
