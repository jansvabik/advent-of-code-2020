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

	// create an array of lines
	var rows []string
	for scanner.Scan() {
		rows = append(rows, scanner.Text())
	}

	// all slopes to test
	slopes := [][]int{
		{1, 1},
		{3, 1},
		{5, 1},
		{7, 1},
		{1, 2},
	}

	// multiplied result (for part 2)
	multiplied := 1

	// calculate every defined slope
	for _, slope := range slopes {
		x, y := 0, 0
		right, down := slope[0], slope[1]
		trees := 0
		for y < len(rows)-1 {
			// move right
			x += right

			// if the x coord is higher than the line length, go back to line start
			rowLength := len(rows[y])
			if x >= rowLength {
				x -= rowLength
			}

			// move down
			y += down

			// tree test
			if rows[y][x] == '#' {
				trees++
			}
		}

		// display part 1 result
		if right == 3 && down == 1 {
			fmt.Println("Part 1:", trees)
		}

		// multiply the number of trees in this slope
		multiplied *= trees
	}

	fmt.Println("Part 2:", multiplied)
}
