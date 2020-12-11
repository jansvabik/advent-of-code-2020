package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// validateSimple validates the provided keys existency
func validateSimple(p map[string]string) bool {
	keys := []string{"ecl", "pid", "eyr", "hcl", "byr", "iyr", "hgt"}
	for _, key := range keys {
		if _, ok := p[key]; !ok {
			return false
		}
	}
	return true
}

// validateAdvanced validates the conditions for each passport data
func validateAdvanced(p map[string]string) bool {
	// fields existency check
	if !validateSimple(p) {
		return false
	}

	// byr
	byr, _ := strconv.Atoi(p["byr"])
	if byr < 1920 || byr > 2002 {
		return false
	}

	// iyr
	iyr, _ := strconv.Atoi(p["iyr"])
	if iyr < 2010 || iyr > 2020 {
		return false
	}

	// eyr
	eyr, _ := strconv.Atoi(p["eyr"])
	if eyr < 2020 || eyr > 2030 {
		return false
	}

	// hgt
	hgtUnit := p["hgt"][len(p["hgt"])-2:]
	hgtValue, _ := strconv.Atoi(p["hgt"][:len(p["hgt"])-2])
	unitOk := hgtUnit != "cm" && hgtUnit != "in"
	cmOk := hgtUnit == "cm" && (hgtValue < 150 || hgtValue > 193)
	inOk := hgtUnit == "in" && (hgtValue < 59 || hgtValue > 76)
	if unitOk || cmOk || inOk {
		return false
	}

	// hcl
	if match, _ := regexp.MatchString(`^#[a-z0-9]{6}$`, p["hcl"]); !match {
		return false
	}

	// ecl
	if match, _ := regexp.MatchString(`^amb|blu|brn|gry|grn|hzl|oth$`, p["ecl"]); !match {
		return false
	}

	// test pid
	if match, _ := regexp.MatchString(`^[0-9]{9}$`, p["pid"]); !match {
		return false
	}

	return true
}

// incrementIfValid increments the counter if the passport is valid
func incrementIfValid(p map[string]string, cs *int, ca *int) {
	if validateSimple(p) {
		*cs++
	}
	if validateAdvanced(p) {
		*ca++
	}
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

	// test every passport
	passport := make(map[string]string)
	validSimple, validAdvanced := 0, 0
	for scanner.Scan() {
		val := scanner.Text()

		// if this is an empty line, make test
		if len(val) < 1 {
			incrementIfValid(passport, &validSimple, &validAdvanced)
			passport = make(map[string]string)
			continue
		}

		// extract the data and add them to the passport map
		settings := strings.Split(val, " ")
		for i := 0; i < len(settings); i++ {
			s := strings.Split(settings[i], ":")
			passport[s[0]] = s[1]
		}
	}

	// test the last passport
	incrementIfValid(passport, &validSimple, &validAdvanced)

	// print the results
	fmt.Println("Part 1:", validSimple)
	fmt.Println("Part 2:", validAdvanced)
}
