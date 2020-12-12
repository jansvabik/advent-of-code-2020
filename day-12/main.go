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

// rotatePoint rotates the point around the ship
// the point coordinates are relatives from the ship
func rotatePoint(p map[rune]int, rot rune, angle int) map[rune]int {
	nextSide := map[rune]rune{'N': 'E', 'E': 'S', 'S': 'W', 'W': 'N'}
	prevSide := map[rune]rune{'N': 'W', 'E': 'N', 'S': 'E', 'W': 'S'}
	for i := 0; i < angle/90; i++ {
		newPoint := make(map[rune]int, 4)
		for k, v := range p {
			if rot == 'R' {
				newPoint[nextSide[k]] = v
			} else {
				newPoint[prevSide[k]] = v
			}
		}
		p = newPoint
	}
	return p
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

	// create a slice of all instructions
	nums := []string{}
	for scanner.Scan() {
		nums = append(nums, scanner.Text())
	}

	// part 1 variables
	currentFacing := 'E'                               // current ship facing (starting to east)
	moves1 := [4]float64{0, 0, 0, 0}                   // ship moves during the navigating
	mi := map[rune]int{'N': 0, 'E': 1, 'S': 2, 'W': 3} // index to the moves array for each side

	// part 2 variables
	negSide := map[rune]rune{'N': 'S', 'E': 'W', 'S': 'N', 'W': 'E'} // negatives of each side
	wp := map[rune]int{'N': 1, 'E': 10}                              // current waypoint direction
	moves2 := [4]float64{0, 0, 0, 0}                                 // ship moves during the navigation

	// process the instructions
	for _, val := range nums {
		ins := []rune(val[0:1])[0]
		arg, _ := strconv.Atoi(val[1:])

		// for part 1
		switch ins {
		case 'N', 'E', 'S', 'W':
			moves1[mi[ins]] += float64(arg)
		case 'L', 'R':
			currentFacing = getNewFacing(currentFacing, ins, arg)
		case 'F':
			moves1[mi[currentFacing]] += float64(arg)
		}

		// for part 2
		switch ins {
		case 'N', 'E':
			wp[ins] += arg
		case 'S', 'W':
			wp[negSide[ins]] -= arg
		case 'L', 'R':
			wp = rotatePoint(wp, ins, arg)
		case 'F':
			for k, v := range wp {
				moves2[mi[k]] += float64(arg * v)
			}
		}
	}

	// calculate the manhattan distance for each puzzle part
	md1 := math.Abs(moves1[mi['S']]-moves1[mi['N']]) + math.Abs(moves1[mi['E']]-moves1[mi['W']])
	md2 := math.Abs(moves2[mi['S']]-moves2[mi['N']]) + math.Abs(moves2[mi['E']]-moves2[mi['W']])
	fmt.Println("Part 1:", md1)
	fmt.Println("Part 2:", md2)
}
