package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// occupiedAround calculates the number of occupied seats around the specified coordinate
func occupiedAround(x int, y int, seats [][]rune) int {
	occ := 0

	// all seats around
	chngs := [8][2]int{{-1, -1}, {0, -1}, {1, -1}, {-1, 0}, {1, 0}, {-1, 1}, {0, 1}, {1, 1}}

	// calculate the number of occupied seats around
	for _, mv := range chngs {
		nx := mv[0] + x
		ny := mv[1] + y
		if nx >= 0 && ny >= 0 && nx < len(seats[0]) && ny < len(seats) {
			if seats[ny][nx] == '#' {
				occ++
			}
		}
	}

	return occ
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

	// create a slice of all rows
	var rows [][]rune
	for scanner.Scan() {
		rows = append(rows, []rune(scanner.Text()))
	}

	// run the simulation loop
	for true {
		chngs := 0
		newMatrix := make([][]rune, len(rows))
		for y, row := range rows {
			newMatrix[y] = make([]rune, len(rows[0]))
			for x, v := range row {
				occa := occupiedAround(x, y, rows)
				if v == 'L' && occa == 0 {
					newMatrix[y][x] = '#'
					chngs++
				} else if v == '#' && occa >= 4 {
					newMatrix[y][x] = 'L'
					chngs++
				} else if v == '.' {
					newMatrix[y][x] = '.'
				} else {
					newMatrix[y][x] = rows[y][x]
				}
			}
		}

		// replace the old state by the new one
		copy(rows, newMatrix)

		// if there was no change, we are done
		if chngs == 0 {
			break
		}
	}

	// calculate the number of occupied seats
	occ := 0
	for _, row := range rows {
		for _, v := range row {
			if v == '#' {
				occ++
			}
		}
	}
	fmt.Println("Part 1:", occ)
}
