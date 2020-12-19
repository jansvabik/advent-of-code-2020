package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

// cond describes exactly one rule with all its subrules
// or subconditions a message can be validated according to
type cond struct {
	Type    int // 0 = Char, 1 = Chain, 2 = OrChain
	Char    string
	Chain   []int
	OrChain [2][]int
}

// rlist is a list of rules
var rlist = map[int]cond{}

// isValid determines if the specified message is valid according to the
// rule specified in ri argument, the rules have to be passed to the fnc,
// function returns the boolean result and the length of the validated str
func isValid(msg string, ri int) (bool, int) {
	rule := rlist[ri]
	// fmt.Println("Testing", msg, "for rule", ri, rule)

	// char condition
	if rule.Type == 0 && string(msg[0]) == rule.Char {
		return true, 1
	}

	// standard chain condition
	if rule.Type == 1 {
		validatedLength := 0
		for _, cond := range rule.Chain {
			v, l := isValid(msg[validatedLength:], cond)
			if !v {
				return false, 0
			}
			validatedLength += l
		}
		return true, validatedLength
	}

	// or chain condition
	if rule.Type == 2 {
		// first block
		validAtLeastOneBlock := true
		validatedLength := 0
		for _, cond := range rule.OrChain[0] {
			v, l := isValid(msg[validatedLength:], cond)
			if !v {
				validAtLeastOneBlock = false
				break
			}
			validatedLength += l
		}
		if validAtLeastOneBlock {
			return true, validatedLength
		}

		// second block
		validAtLeastOneBlock = true
		validatedLength = 0
		for _, cond := range rule.OrChain[1] {
			v, l := isValid(msg[validatedLength:], cond)
			if !v {
				validAtLeastOneBlock = false
				break
			}
			validatedLength += l
		}

		return validAtLeastOneBlock, validatedLength
	}

	return false, 0
}

func main() {
	// read the puzzle input
	b, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal("Cannot open the input file.")
	}

	// extract the puzzle input data
	data := strings.Split(string(b), "\n\n")
	strConditions := strings.Split(data[0], "\n")
	messages := strings.Split(data[1], "\n")

	// create a map of all rules (well structured rules map)
	for _, sc := range strConditions {
		// extract the rule data
		spl := strings.Split(sc, ": ")
		ruleNo, _ := strconv.Atoi(spl[0])

		// create a condition structure for further use
		c := cond{}
		if spl[1][0] == '"' {
			c.Type = 0
			c.Char = spl[1][1:2]
		} else if strings.Contains(spl[1], "|") {
			c.Type = 2

			// create a list of OR conditions
			c.OrChain = [2][]int{}
			orConds := strings.Split(spl[1], " | ")
			for i, oc := range orConds {
				strInts := strings.Split(oc, " ")
				for _, si := range strInts {
					in, _ := strconv.Atoi(si)
					c.OrChain[i] = append(c.OrChain[i], in)
				}
			}
		} else {
			c.Type = 1

			// create a list of chained conditions
			chain := strings.Split(spl[1], " ")
			for _, el := range chain {
				i, _ := strconv.Atoi(el)
				c.Chain = append(c.Chain, i)
			}
		}

		// register the rule
		rlist[ruleNo] = c
	}

	// count the valid rules and print the results
	valid := 0
	for _, m := range messages {
		v, l := isValid(m, 0)
		if v && l == len(m) {
			valid++
		}
	}
	fmt.Println("Part 1:", valid)
}
