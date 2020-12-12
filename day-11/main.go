package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// occupiedAround calculates the number of occupied seats around the specified coordinate
func occupiedNearest(x int, y int, seats [][]rune) int {
	occ := 0

	// all seats around
	chngs := [8][2]int{{-1, -1}, {0, -1}, {1, -1}, {-1, 0}, {1, 0}, {-1, 1}, {0, 1}, {1, 1}}

	// calculate the number of occupied seats around
	for _, mv := range chngs {
		nx := x + mv[0]
		ny := y + mv[1]
		if nx >= 0 && ny >= 0 && nx < len(seats[0]) && ny < len(seats) {
			if seats[ny][nx] == '#' {
				occ++
			}
		}
	}

	return occ
}

// occupiedNearest calculates how many seats are occupied in all 8 sides from the coordinate
func occupiedAround(x int, y int, seats [][]rune) int {
	occ := 0

	// all seats around
	chngs := [8][2]int{{-1, -1}, {0, -1}, {1, -1}, {-1, 0}, {1, 0}, {-1, 1}, {0, 1}, {1, 1}}

	// calculate the number of occupied seats around
	for _, mv := range chngs {
		nx := x
		ny := y
		for {
			nx += mv[0]
			ny += mv[1]
			if nx >= 0 && ny >= 0 && nx < len(seats[0]) && ny < len(seats) {
				if seats[ny][nx] == '.' {
					continue
				}
				if seats[ny][nx] == '#' {
					occ++
				}
			}
			break
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
	var rows1 [][]rune
	var rows2 [][]rune
	for scanner.Scan() {
		rows1 = append(rows1, []rune(scanner.Text()))
		rows2 = append(rows2, []rune(scanner.Text()))
	}

	// run the simulation loop for part 1
	for {
		chngs := 0
		newMatrix := make([][]rune, len(rows1))
		for y, row := range rows1 {
			newMatrix[y] = make([]rune, len(rows1[0]))
			for x, v := range row {
				occa := occupiedNearest(x, y, rows1)
				if v == 'L' && occa == 0 {
					newMatrix[y][x] = '#'
					chngs++
				} else if v == '#' && occa >= 4 {
					newMatrix[y][x] = 'L'
					chngs++
				} else {
					newMatrix[y][x] = rows1[y][x]
				}
			}
		}

		// replace the old state by the new one
		copy(rows1, newMatrix)

		// if there was no change, we are done
		if chngs == 0 {
			break
		}
	}

	// run the simulation loop for part 2
	for {
		chngs := 0
		newMatrix := make([][]rune, len(rows2))
		for y, row := range rows2 {
			newMatrix[y] = make([]rune, len(rows2[0]))
			for x, v := range row {
				occa := occupiedAround(x, y, rows2)
				if v == 'L' && occa == 0 {
					newMatrix[y][x] = '#'
					chngs++
				} else if v == '#' && occa >= 5 {
					newMatrix[y][x] = 'L'
					chngs++
				} else {
					newMatrix[y][x] = rows2[y][x]
				}
			}
		}

		// replace the old state by the new one
		copy(rows2, newMatrix)

		// if there was no change, we are done
		if chngs == 0 {
			break
		}
	}

	// calculate the number of occupied seats for both parts
	occ1 := 0
	occ2 := 0
	for i := range rows1 {
		for _, v := range rows1[i] {
			if v == '#' {
				occ1++
			}
		}
		for _, v := range rows2[i] {
			if v == '#' {
				occ2++
			}
		}
	}
	fmt.Println("Part 1:", occ1)
	fmt.Println("Part 2:", occ2)
}
