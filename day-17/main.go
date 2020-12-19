package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

// all possible coordination changes in 3D space
var chngs3D = [26][3]int{
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

// all possible coordination changes in 4D space
var chngs4D = [80][4]int{
	{-1, -1, -1, -1}, {-1, -1, -1, 0}, {-1, -1, -1, 1},
	{-1, -1, 0, -1}, {-1, -1, 0, 0}, {-1, -1, 0, 1},
	{-1, -1, 1, -1}, {-1, -1, 1, 0}, {-1, -1, 1, 1},
	{-1, 0, -1, -1}, {-1, 0, -1, 0}, {-1, 0, -1, 1},
	{-1, 0, 0, -1}, {-1, 0, 0, 0}, {-1, 0, 0, 1},
	{-1, 0, 1, -1}, {-1, 0, 1, 0}, {-1, 0, 1, 1},
	{-1, 1, -1, -1}, {-1, 1, -1, 0}, {-1, 1, -1, 1},
	{-1, 1, 0, -1}, {-1, 1, 0, 0}, {-1, 1, 0, 1},
	{-1, 1, 1, -1}, {-1, 1, 1, 0}, {-1, 1, 1, 1},
	{0, -1, -1, -1}, {0, -1, -1, 0}, {0, -1, -1, 1},
	{0, -1, 0, -1}, {0, -1, 0, 0}, {0, -1, 0, 1},
	{0, -1, 1, -1}, {0, -1, 1, 0}, {0, -1, 1, 1},
	{0, 0, -1, -1}, {0, 0, -1, 0}, {0, 0, -1, 1},
	{0, 0, 0, -1} /*{0, 0, 0, 0},*/, {0, 0, 0, 1},
	{0, 0, 1, -1}, {0, 0, 1, 0}, {0, 0, 1, 1},
	{0, 1, -1, -1}, {0, 1, -1, 0}, {0, 1, -1, 1},
	{0, 1, 0, -1}, {0, 1, 0, 0}, {0, 1, 0, 1},
	{0, 1, 1, -1}, {0, 1, 1, 0}, {0, 1, 1, 1},
	{1, -1, -1, -1}, {1, -1, -1, 0}, {1, -1, -1, 1},
	{1, -1, 0, -1}, {1, -1, 0, 0}, {1, -1, 0, 1},
	{1, -1, 1, -1}, {1, -1, 1, 0}, {1, -1, 1, 1},
	{1, 0, -1, -1}, {1, 0, -1, 0}, {1, 0, -1, 1},
	{1, 0, 0, -1}, {1, 0, 0, 0}, {1, 0, 0, 1},
	{1, 0, 1, -1}, {1, 0, 1, 0}, {1, 0, 1, 1},
	{1, 1, -1, -1}, {1, 1, -1, 0}, {1, 1, -1, 1},
	{1, 1, 0, -1}, {1, 1, 0, 0}, {1, 1, 0, 1},
	{1, 1, 1, -1}, {1, 1, 1, 0}, {1, 1, 1, 1},
}

// printHumanReadable3D prints the 3D space data in human readable format
// (active cubes are displayed as #, inactive cubes as .), the third
// dimension (z) is represented by a [x,y] matrix for each z
func printHumanReadable3D(data map[[3]int]bool, depth int) {
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

// activeNeighbours3D calculates the number of active cubes neighbouring
// to the specified coordinate in 3D space
func activeNeighbours3D(x int, y int, z int, data map[[3]int]bool) int {
	// calculate the number of active neighbours
	active := 0
	for _, chng := range chngs3D {
		nx := x + chng[0]
		ny := y + chng[1]
		nz := z + chng[2]
		if v, ok := data[[3]int{nx, ny, nz}]; ok && v {
			active++
		}
	}
	return active
}

// activeNeighbours4D calculates the number of active cubes neighbouring
// to the specified coordinate in 4D space
func activeNeighbours4D(w int, x int, y int, z int, data map[[4]int]bool) int {
	// calculate the number of active neighbours
	active := 0
	for _, chng := range chngs4D {
		nw := w + chng[0]
		nx := x + chng[1]
		ny := y + chng[2]
		nz := z + chng[3]
		if v, ok := data[[4]int{nw, nx, ny, nz}]; ok && v {
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
	data3D := map[[3]int]bool{}
	data4D := map[[4]int]bool{}
	for y, row := range rows {
		for x, c := range row {
			data3D[[3]int{x - 1, y - 1, 0}] = c == '#'
			data4D[[4]int{x - 1, y - 1, 0, 0}] = c == '#'
		}
	}

	// run the boot process (6 times, according to the assignment)
	for i := 0; i < 6; i++ {
		// wrap into even bigger cube (set cubes around to false)
		for c := range data3D {
			for _, chng := range chngs3D {
				nx := c[0] + chng[0]
				ny := c[1] + chng[1]
				nz := c[2] + chng[2]
				if _, ok := data3D[[3]int{nx, ny, nz}]; !ok {
					data3D[[3]int{nx, ny, nz}] = false
				}
			}
		}

		// test every stored cube to have active neighbours and set the new value
		newData3D := map[[3]int]bool{}
		for coord, active := range data3D {
			activeNeighbours := activeNeighbours3D(coord[0], coord[1], coord[2], data3D)
			if active {
				if activeNeighbours == 2 || activeNeighbours == 3 {
					newData3D[[3]int{coord[0], coord[1], coord[2]}] = true
				} else {
					newData3D[[3]int{coord[0], coord[1], coord[2]}] = false
				}
			} else {
				if activeNeighbours == 3 {
					newData3D[[3]int{coord[0], coord[1], coord[2]}] = true
				} else {
					newData3D[[3]int{coord[0], coord[1], coord[2]}] = false
				}
			}
		}
		data3D = newData3D
	}

	// run the boot process (6 times, according to the assignment)
	for i := 0; i < 6; i++ {
		// wrap into even bigger cube (set cubes around to false)
		for c := range data4D {
			for _, chng := range chngs4D {
				nw := c[0] + chng[0]
				nx := c[1] + chng[1]
				ny := c[2] + chng[2]
				nz := c[3] + chng[3]
				if _, ok := data4D[[4]int{nw, nx, ny, nz}]; !ok {
					data4D[[4]int{nw, nx, ny, nz}] = false
				}
			}
		}

		// test every stored cube to have active neighbours and set the new value
		newData4D := map[[4]int]bool{}
		for coord, active := range data4D {
			activeNeighbours := activeNeighbours4D(coord[0], coord[1], coord[2], coord[3], data4D)
			if active {
				if activeNeighbours == 2 || activeNeighbours == 3 {
					newData4D[[4]int{coord[0], coord[1], coord[2], coord[3]}] = true
				} else {
					newData4D[[4]int{coord[0], coord[1], coord[2], coord[3]}] = false
				}
			} else {
				if activeNeighbours == 3 {
					newData4D[[4]int{coord[0], coord[1], coord[2], coord[3]}] = true
				} else {
					newData4D[[4]int{coord[0], coord[1], coord[2], coord[3]}] = false
				}
			}
		}
		data4D = newData4D
	}

	// sum and print the number of active cubes
	active3D := 0
	active4D := 0
	for _, v := range data3D {
		if v {
			active3D++
		}
	}
	for _, v := range data4D {
		if v {
			active4D++
		}
	}
	fmt.Println("Part 1:", active3D)
	fmt.Println("Part 2:", active4D)
}
