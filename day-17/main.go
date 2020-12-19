package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

// all possible coord changes
var chngs = [26][3]int{
	{-1, -1, -1}, {-1, -1, 0}, {-1, -1, 1},
	{-1, 0, -1}, {-1, 0, 0}, {-1, 0, 1},
	{-1, 1, -1}, {-1, 1, 0}, {-1, 1, 1},
	{0, -1, -1}, {0, -1, 0}, {0, -1, 1},
	{0, 0, -1} /*{0, 0, 0}, */, {0, 0, 1},
	{0, 1, -1}, {0, 1, 0}, {0, 1, 1},
	{1, -1, -1}, {1, -1, 0}, {1, -1, 1},
	{1, 0, -1}, {1, 0, 0}, {1, 0, 1},
	{1, 1, -1}, {1, 1, 0}, {1, 1, 1},
}

// printHumanReadable prints the data in human readable format (active
// cubes are displayed as #, inactive cubes as .), the third dimension (z)
// is represented by a [x,y] matrix for each z
func printHumanReadable(data map[[3]int]bool, depth int) {
	for z := -depth; z <= depth; z++ {
		for y := -depth; y <= depth; y++ {
			for x := -depth; x <= depth; x++ {
				v := data[[3]int{x, y, z}]
				if v {
					fmt.Print("#")
				} else {
					fmt.Print(".")
				}
			}
			fmt.Print("\n")
		}
		fmt.Print("\n")
	}
}

// activeNeighbours calculates the number of active cubes neighbouring
// to the specified coordinate in 3D space
func activeNeighbours(x int, y int, z int, data map[[3]int]bool) int {
	// calculate the number of active neighbours
	active := 0
	for _, chng := range chngs {
		nx := x + chng[0]
		ny := y + chng[1]
		nz := z + chng[2]
		if v, ok := data[[3]int{nx, ny, nz}]; ok && v {
			active++
		}
	}
	return active
}

func main() {
	// read the puzzle input
	b, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal("Cannot open the input file.")
	}

	// extract the puzzle input data
	rows := strings.Split(string(b), "\n")
	data := map[[3]int]bool{}
	for y, row := range rows {
		for x, c := range row {
			data[[3]int{x - 1, y - 1, -1}] = false
			data[[3]int{x - 1, y - 1, 0}] = c == '#'
			data[[3]int{x - 1, y - 1, 1}] = false
		}
	}

	// run the boot process (6 times, according to the assignment)
	for i := 0; i < 6; i++ {
		// wrap into even bigger cube (set cubes around to false)
		for c := range data {
			for _, chng := range chngs {
				nx := c[0] + chng[0]
				ny := c[1] + chng[1]
				nz := c[2] + chng[2]
				if _, ok := data[[3]int{nx, ny, nz}]; !ok {
					data[[3]int{nx, ny, nz}] = false
				}
			}
		}

		// test every stored cube to have active neighbours and set the new value
		newData := map[[3]int]bool{}
		for coord, active := range data {
			activeNeighbours := activeNeighbours(coord[0], coord[1], coord[2], data)
			if active {
				if activeNeighbours == 2 || activeNeighbours == 3 {
					newData[[3]int{coord[0], coord[1], coord[2]}] = true
				} else {
					newData[[3]int{coord[0], coord[1], coord[2]}] = false
				}
			} else {
				if activeNeighbours == 3 {
					newData[[3]int{coord[0], coord[1], coord[2]}] = true
				} else {
					newData[[3]int{coord[0], coord[1], coord[2]}] = false
				}
			}
		}
		data = newData
	}

	// sum and print the number of active cubes
	active := 0
	for _, v := range data {
		if v {
			active++
		}
	}
	fmt.Println("Part 1:", active)
}
