package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"marshallformula.codes/utils"
)

type assignment struct {
	lowerBound int
	upperBound int
}

func splitIt(str *string, separator string) []string {
	return strings.Split(*str, separator)
}

func parseAssignment(data *string) (*assignment, error) {
	bounds := splitIt(data, "-")

	lower, err := strconv.Atoi(bounds[0])

	if err != nil {
		return nil, err
	}

	upper, err := strconv.Atoi(bounds[1])

	if err != nil {
		return nil, err
	}

	return &assignment{lower, upper}, nil

}

func checkContains(first *assignment, second *assignment) bool {
	if (first.lowerBound <= second.lowerBound) && (first.upperBound >= second.upperBound) {
		return true
	}
	return false
}

func checkOverlap(first *assignment, second *assignment) bool {
	return (first.lowerBound >= second.lowerBound && first.lowerBound <= second.upperBound) ||
		(first.upperBound >= second.lowerBound && first.upperBound <= second.upperBound)
}

func main() {

	is, err := utils.InputScanner("input.txt")
	defer is.Close()

	if err != nil {
		log.Fatalln(err)
	}

	overlapping := 0

	is.Scan(func(val string) {

		pairs := splitIt(&val, ",")

		if len(pairs) != 2 {
			log.Fatal("Invalid pair of assignment", pairs)
		}

		first, err := parseAssignment(&pairs[0])

		if err != nil {
			log.Fatalln(err)
		}

		second, err := parseAssignment(&pairs[1])

		if err != nil {
			log.Fatalln(err)
		}

		if checkContains(first, second) || checkContains(second, first) {
			overlapping++
		} else if checkOverlap(first, second) {
			overlapping++
		}

	})

	fmt.Println(overlapping)

}
