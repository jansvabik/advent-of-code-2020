package advmath

import (
	"regexp"
	"strconv"
	"strings"
)

// calculateSimpleExpression calculates, as it says, simple expressions
// these are the expressions which contains only + and * operators, like
// e.g. 1+5+7*24+4*6 etc., there must not be any bracket in the simple expr
func calculateSimpleExpression(expr string) (int, string) {
	intsToMultiply := []int{}

	// sum the splitted values
	exprsToMultiply := strings.Split(expr, "*")
	for _, etm := range exprsToMultiply {
		sum := 0
		espl := strings.Split(etm, "+")
		for _, e := range espl {
			v, _ := strconv.Atoi(e)
			sum += v
		}
		intsToMultiply = append(intsToMultiply, sum)
	}

	// multiply all the values
	multiplied := 1
	for _, v := range intsToMultiply {
		multiplied *= v
	}

	return multiplied, strconv.Itoa(multiplied)
}

// CalculateExpression is a function which handles the whole mathematic expressions
// and tries to calculate their results (this result is then returned as an int)
func CalculateExpression(expr string) int {
	var submatches []string
	for start := true; start || len(submatches) > 0; start = false {
		submatches = regexp.MustCompile("\\([0-9]+[\\*|\\+][0-9]+([\\*|\\+][0-9]+)*\\)").FindStringSubmatch(expr)
		if len(submatches) == 0 {
			break
		}

		parExpr := submatches[0][1 : len(submatches[0])-1]
		_, res := calculateSimpleExpression(parExpr)
		expr = strings.ReplaceAll(expr, submatches[0], res)
	}
	fin, _ := calculateSimpleExpression(expr)
	return fin
}
