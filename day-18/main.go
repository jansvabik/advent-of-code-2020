package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
	"strings"
)

// calculateSimpleExpression just sums or multiplies two values depending
// to the specified operator in op argument
func calculateSimpleExpression(as string, bs string, op string) (string, error) {
	a, _ := strconv.Atoi(as)
	b, _ := strconv.Atoi(bs)

	switch op {
	case "*":
		return strconv.Itoa(a * b), nil
	case "+":
		return strconv.Itoa(a + b), nil
	default:
		return strconv.Itoa(0), errors.New("calculate simple expr: not valid operator")
	}
}

// calculateSimpleExpressionQueue calculates the whole queue of simple
// operators (which are + and *) and returns the value as int and as string
// (the string return value is then used for further expression creating)
func calculateSimpleExpressionQueue(expr string) (int, string) {
	var fin int
	for {
		sexpr := regexp.MustCompile("([0-9]+)([\\*|\\+])([0-9]+)").FindStringSubmatch(expr)
		if len(sexpr) < 4 {
			break
		}

		res, _ := calculateSimpleExpression(sexpr[1], sexpr[3], sexpr[2])
		if len(sexpr[0]) != len(expr) {
			expr = res + expr[len(sexpr[0]):]
		} else {
			fin, _ = strconv.Atoi(res)
			break
		}
	}
	return fin, strconv.Itoa(fin)
}

// calculateExpression is a function which handles the whole mathematic expressions
// and tries to calculate their results (this result is then returned as an int)
func calculateExpression(expr string) int {
	var submatches []string
	for start := true; start || len(submatches) > 0; start = false {
		submatches = regexp.MustCompile("\\([0-9]+[\\*|\\+][0-9]+([\\*|\\+][0-9]+)*\\)").FindStringSubmatch(expr)
		if len(submatches) == 0 {
			break
		}

		parExpr := submatches[0][1 : len(submatches[0])-1]
		_, res := calculateSimpleExpressionQueue(parExpr)
		expr = strings.ReplaceAll(expr, submatches[0], res)
	}
	fin, _ := calculateSimpleExpressionQueue(expr)
	return fin
}

func main() {
	// read the puzzle input
	b, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal("Cannot open the input file.")
	}

	// extract the puzzle input data (remove spaces)
	rows := strings.Split(strings.ReplaceAll(string(b), " ", ""), "\n")

	// sum the results of all the math expressions in input
	sum := 0
	for _, row := range rows {
		sum += calculateExpression(row)
	}

	// nicely done!
	fmt.Println("Part 1:", sum)
}
