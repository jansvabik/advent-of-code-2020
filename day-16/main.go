package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
	"strings"
)

// validityCondition represents the conditions used
// for validating tickets, there are two arrays of
// two integers, these two integers are used as
// lowest possible and largest possible numbers
// (values) of the validated number (value in ticket)
type validityCondition struct {
	Name      string
	Intervals [2][2]int
}

func main() {
	// read the puzzle input
	b, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal("Cannot open the input file.")
	}

	// extract the data
	d := strings.Split(string(b), "\n\n")
	notes := strings.Split(d[0], "\n")
	myTicket := strings.Split(d[1], "\n")[1]
	nearbyTickets := strings.Split(d[2], "\n")[1:]

	// validity conditions
	valconds := []validityCondition{}
	for _, n := range notes {
		// extract the conditions data
		s := regexp.MustCompile(": | or ").Split(n, 3)
		i1s := strings.Split(s[1], "-")
		i2s := strings.Split(s[2], "-")

		// convert to integers
		i10, _ := strconv.Atoi(i1s[0])
		i11, _ := strconv.Atoi(i1s[1])
		i20, _ := strconv.Atoi(i2s[0])
		i21, _ := strconv.Atoi(i2s[1])

		// add to the list of conditions
		valconds = append(valconds, validityCondition{
			Name: s[0],
			Intervals: [2][2]int{
				{i10, i11},
				{i20, i21},
			},
		})
	}

	// part 1 - get the sum of absolutely invalid values
	// also create a list of valid tickets to further processing
	invalidValuesSum := 0
	validTickets := [][]int{}
	for _, nt := range nearbyTickets {
		// create a slice of integers (ticket values)
		svs := strings.Split(nt, ",")
		vals := []int{}
		for _, sv := range svs {
			v, _ := strconv.Atoi(sv)
			vals = append(vals, v)
		}

		// test every value to be valid or invalid
		allValuesValid := true
		for _, v := range vals {
			isValid := false
			for _, conds := range valconds {
				if (conds.Intervals[0][0] <= v && conds.Intervals[0][1] >= v) || (conds.Intervals[1][0] <= v && conds.Intervals[1][1] >= v) {
					isValid = true
					break
				}
			}
			if !isValid {
				invalidValuesSum += v
				allValuesValid = false
			}
		}

		// if there were all values valid, the whole ticket is also valid
		if allValuesValid {
			validTickets = append(validTickets, vals)
		}
	}

	// part 2 - connect all values to matching name (valconds)
	nameToI := map[string]int{}
	for {
		chngs := 0
		for i := 0; i < len(validTickets[0]); i++ {
			// extract i-th values from the tickets to be tested
			validFor := []string{}
			for _, conds := range valconds {
				allValid := true
				for _, t := range validTickets {
					v := t[i]
					if ok := (conds.Intervals[0][0] <= v && conds.Intervals[0][1] >= v) || (conds.Intervals[1][0] <= v && conds.Intervals[1][1] >= v); !ok {
						allValid = false
						break
					}
				}
				if _, ok := nameToI[conds.Name]; allValid && !ok {
					validFor = append(validFor, conds.Name)
				}
			}

			// if this is valid for just one field, remember it
			if len(validFor) == 1 {
				nameToI[validFor[0]] = i
				chngs++
			}
		}
		if chngs == 0 {
			break
		}
	}

	// get indexes of the data starting with "departure" word
	depIndexes := []int{}
	for n, v := range nameToI {
		if len(n) >= 9 && n[0:9] == "departure" {
			depIndexes = append(depIndexes, v)
		}
	}

	// multiply the departure data values in my ticket
	mtvalss := strings.Split(myTicket, ",")
	pt2res := 1
	for _, di := range depIndexes {
		v, _ := strconv.Atoi(mtvalss[di])
		pt2res *= v
	}

	// print the results
	fmt.Println("Part 1:", invalidValuesSum)
	fmt.Println("Part 2:", pt2res)
}
