package main

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"

	"marshallformula.codes/utils"
)

type instruction struct {
	quantity int
	from     int
	to       int
}

// index of where each stack starts
var stackIndexes = []int16{0, 4, 8, 12, 16, 20, 24, 28, 32}
var stacks = make([][]string, 9, 9)
var instructions = make([]instruction, 0)
var instructionRegex *regexp.Regexp

func cleanRow(line *string) []string {
	row := make([]string, 0, 9)
	for _, start := range stackIndexes {
		item := (*line)[start : start+3]
		row = append(row, strings.TrimSpace(strings.Trim(item, "[]")))
	}

	return row
}

// initial stack poulation
func populateStacks(rows *[][]string) {
	for i := len(*rows) - 1; i >= 0; i-- {
		for j, crate := range (*rows)[i] {
			if len(crate) > 0 {
				stack := stacks[j]
				if stack == nil {
					stack = make([]string, 0)
				}
				stacks[j] = append(stack, crate)
			}
		}
	}
}

func parseInstruction(instructionLine *string) {
	match := instructionRegex.FindStringSubmatch(*instructionLine)
	if (len(match)) != 4 {
		log.Fatal("Didn't process the instruction regex", *instructionLine)
	}

	quantity, qErr := strconv.Atoi(match[1])
	from, fErr := strconv.Atoi(match[2])
	to, tErr := strconv.Atoi(match[3])

	if qErr != nil || fErr != nil || tErr != nil {
		log.Fatal("Couldn't cast string to number", qErr, fErr, tErr)
	}

	instructions = append(instructions, instruction{quantity, from, to})
}

func performInstruction(inst *instruction) {
	// since arrays are 0 index - just to make things more readable
	fromStackIdx := inst.from - 1
	toStackIdx := inst.to - 1

	// pull out the stacks involved
	fromStack := stacks[fromStackIdx]
	toStack := stacks[toStackIdx]

	// repeat as many times as instructed
	for i := 0; i < inst.quantity; i++ {

		// find the top crate
		lastIdx := len(fromStack) - 1
		crate := fromStack[lastIdx]

		// add the crate on the new stack
		toStack = append(toStack, crate)
		// pop the create off the old stack
		fromStack = fromStack[:lastIdx]

	}

	// update the stacks in the data structure
	stacks[fromStackIdx] = fromStack
	stacks[toStackIdx] = toStack
}

func process(rows *[][]string) func(string) {

	startedInstructions := false

	return func(val string) {

		if startedInstructions {
			if len(val) < 1 {
				// empty line between stacks and instructions
				return
			}

			a := strings.TrimSpace(val)
			parseInstruction(&a)

			return
		}

		if len(val) == 0 {
			fmt.Println("finished stacks")
		}

		row := cleanRow(&val)

		if row[0] == "1" {
			fmt.Println("Finished Reading Stack Rows")
			startedInstructions = true
			return
		}

		*rows = append(*rows, row)
	}
}

func main() {

	is, err := utils.InputScanner("input.txt")
	defer is.Close()

	if err != nil {
		log.Fatalln(err)
	}

	instructionRegex = regexp.MustCompile("move ([0-9]+) from ([1-9]) to ([1-9])")

	crateRows := make([][]string, 0)
	processor := process(&crateRows)

	is.Scan(processor)

	populateStacks(&crateRows)

	for _, i := range instructions {
		performInstruction(&i)
	}

	topStacks := ""

	for i, x := range stacks {
		fmt.Println(i, x)
		topStacks += x[len(x)-1]
	}

	fmt.Println(topStacks)
}
