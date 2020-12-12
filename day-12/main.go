package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

// getNewFacing determines new ship facing depending to the rotation and angle
func getNewFacing(curr rune, rot rune, angle int) rune {
	// possible faces and its indexes
	faces := [4]rune{'N', 'E', 'S', 'W'}
	startIndex := map[rune]int{'N': 0, 'E': 1, 'S': 2, 'W': 3}

	// rotate until the angle is reached
	index := startIndex[curr]
	for i := 0; i < angle/90; i++ {
		// change operation depending to the rotation
		if rot == 'R' {
			index++
		} else {
			index--
		}

		// if out of range, correct the difference
		if index >= len(faces) {
			index -= 4
		} else if index < 0 {
			index += 4
		}
	}
	return faces[index]
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

	// create a slice of all commands
	moves := [4]float64{0, 0, 0, 0}
	mi := map[rune]int{'N': 0, 'E': 1, 'S': 2, 'W': 3}
	currentFacing := 'E'
	for scanner.Scan() {
		val := scanner.Text()
		ins := []rune(val[0:1])[0]
		arg, _ := strconv.Atoi(val[1:])

		// process the instruction
		switch ins {
		case 'N', 'E', 'S', 'W':
			moves[mi[ins]] += float64(arg)
		case 'L', 'R':
			currentFacing = getNewFacing(currentFacing, ins, arg)
		case 'F':
			moves[mi[currentFacing]] += float64(arg)
		}
	}

	// calculate the manhattan distance
	manhattanDist := math.Abs(moves[mi['S']]-moves[mi['N']]) + math.Abs(moves[mi['E']]-moves[mi['W']])
	fmt.Println(manhattanDist)
}
