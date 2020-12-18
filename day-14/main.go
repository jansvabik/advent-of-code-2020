package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

// ensureBinaryValStringLength ensures that the val string is length
// exactly the ln characters, this function adds leading zeros if
// the wanted ln is higher then the val length, if the val length
// is lower than the specified ln, the string will be truncated
// from the left by ln-len(val) characters, up to the specified ln
func ensureBinaryValStringLength(val string, ln int) string {
	nval := make([]rune, ln)
	for i := range nval {
		if i < ln-len(val) {
			nval[i] = '0'
		} else {
			nval[i] = rune(val[i-(ln-len(val))])
		}
	}
	return string(nval)
}

// applyMaskV1 applies the mask depending to the AOC14/1 rules
func applyMaskV1(val string, mask string) (string, error) {
	// ensure the same length of val and mask
	if len(val) != len(mask) {
		return val, errors.New("apply mask: val and string have to be same length")
	}

	// create new value with mask applied
	nval := make([]rune, len(mask))
	for i, m := range mask {
		if m != 'X' {
			nval[i] = m
		} else {
			nval[i] = []rune(val)[i]
		}
	}
	return string(nval), nil
}

func main() {
	// read the puzzle input
	b, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal("Cannot open the input file.")
	}
	str := string(b)

	// extract the data
	data := strings.Split(str, "\n")
	var lines [][]string
	for _, l := range data {
		spl := strings.Split(l, " = ")
		lines = append(lines, spl)
		fmt.Println(spl)
	}

	// part 1
	var mask string
	mem1 := map[string]int64{}
	for _, l := range lines {
		if l[0] == "mask" {
			mask = l[1]
		} else {
			addr := l[0][4 : len(l[0])-1]              // get mem addr
			dec, _ := strconv.Atoi(l[1])               // convert to dec int
			bin := strconv.FormatInt(int64(dec), 2)    // convert to bin str
			bin = ensureBinaryValStringLength(bin, 36) // add leading zeros
			fin, _ := applyMaskV1(bin, mask)           // apply mask for pt1
			findec, _ := strconv.ParseInt(fin, 2, 64)  // convert to dec int
			mem1[addr] = findec                        // write to memory
		}
	}

	// sum all values and print the results
	sum1 := int64(0)
	for _, v := range mem1 {
		sum1 += v
	}

	fmt.Println("Part 1:", sum1)
}
