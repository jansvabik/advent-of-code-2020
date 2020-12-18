package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"math"
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

// applyMaskV1 applies the mask according to the AOC14/1 rules
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

// applyMaskV2 applies the mask according to the AOC14/2 rules
func applyMaskV2(val string, mask string) []string {
	combs := []string{}                            // list of addr combinations (to return)
	xcount := strings.Count(mask, "X")             // the number of x chars in mask
	combcount := int(math.Pow(2, float64(xcount))) // the number of possible combinations

	// find every combination and transform the value accordingly
	for i := 0; i < combcount; i++ {
		// represent this exact combination in binary
		bin := strconv.FormatInt(int64(i), 2)
		blz := ensureBinaryValStringLength(bin, xcount)

		// apply the mask changes specified in the task rules
		str := []rune(val)
		xindex := 0
		for i, r := range mask {
			switch r {
			case 'X':
				str[i] = rune(blz[xindex])
				xindex++
			case '1':
				str[i] = '1'
			}
		}

		// store the new address in array to return
		combs = append(combs, string(str))
	}
	return combs
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
	}

	// run the docking program!
	var mask string
	mem1 := map[string]int64{}
	mem2 := map[string]int64{}
	for _, l := range lines {
		if l[0] == "mask" {
			mask = l[1]
		} else {
			// extracting and calculating values
			addr := l[0][4 : len(l[0])-1]                      // get mem addr
			origdec, _ := strconv.Atoi(l[1])                   // convert to dec int
			bin := strconv.FormatInt(int64(origdec), 2)        // convert to bin str
			bin = ensureBinaryValStringLength(bin, 36)         // add leading zeros
			fin, _ := applyMaskV1(bin, mask)                   // apply mask for pt1
			findec, _ := strconv.ParseInt(fin, 2, 64)          // convert to dec int
			addrdec, _ := strconv.Atoi(addr)                   // decimal addr
			addrbin := strconv.FormatInt(int64(addrdec), 2)    // binary addr
			addrbin = ensureBinaryValStringLength(addrbin, 36) // add lead zeros to addr

			// part 1
			mem1[addr] = findec // write to memory

			// part 2
			addrs := applyMaskV2(addrbin, mask)
			for _, a := range addrs {
				finaddrdec, _ := strconv.ParseInt(a, 2, 64) // final addr decimal
				finaddrstr := strconv.Itoa(int(finaddrdec)) // final addr as a str
				mem2[finaddrstr] = int64(origdec)           // write to memory
			}
		}
	}

	// sum all values and print the results
	sum1 := int64(0)
	sum2 := int64(0)
	for _, v := range mem1 {
		sum1 += v
	}
	for _, v := range mem2 {
		sum2 += v
	}

	fmt.Println("Part 1:", sum1)
	fmt.Println("Part 2:", sum2)
}
