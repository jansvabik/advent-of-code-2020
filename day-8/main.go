package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

// cmd is a structure for one command, which consists of (ins)truction and (arg)ument
type cmd struct {
	ins string
	arg int
}

// isInSlice returns a boolean true if "n" is "in" sl or false if "n" is not in "sl"
func isInSlice(n int, sl []int) bool {
	for _, v := range sl {
		if n == v {
			return true
		}
	}
	return false
}

// isInfiniteLoop tests the command list to be an inf loop of not to be an inf loop and returns boolean value and acc value
func isInfiniteLoop(cmds []cmd) (bool, int) {
	var calledInstructions []int
	acc := 0
	for i := 0; i < len(cmds); i++ {
		// test if this instruction has been already called
		if isInSlice(i, calledInstructions) {
			return true, acc
		}
		calledInstructions = append(calledInstructions, i)

		// "run" the command
		switch cmds[i].ins {
		case "nop":
			break
		case "acc":
			acc += cmds[i].arg
			break
		case "jmp":
			i += cmds[i].arg - 1
			break
		}
	}
	return false, acc
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

	// process all lines
	var cmds []cmd
	for scanner.Scan() {
		val := scanner.Text()
		ins := val[:3]
		arg, _ := strconv.Atoi(val[4:])
		cmds = append(cmds, cmd{
			ins: ins,
			arg: arg,
		})
	}

	// part 1 - what is the acc value before starting inf loop
	_, acc := isInfiniteLoop(cmds)
	fmt.Println("Part 1:", acc)

	// part 2 - what has to be changed to break the inf loop
	for i, c := range cmds {
		// duplicate the command slice
		cmdsDebug := make([]cmd, len(cmds))
		copy(cmdsDebug, cmds)

		// change exactly one jmp/nop command
		if c.ins == "nop" {
			cmdsDebug[i].ins = "jmp"
		} else if c.ins == "jmp" {
			cmdsDebug[i].ins = "nop"
		}

		// if this is not an inf loop, print acc
		inf, acc := isInfiniteLoop(cmdsDebug)
		if !inf {
			fmt.Println("Part 2:", acc)
		}
	}
}
