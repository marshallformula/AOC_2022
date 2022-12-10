package main

import (
	"fmt"
	"log"
	"marshallformula.codes/utils"
)

func allUnique(chunk *string, markerLength int) bool {
	set := make(map[rune]struct{})

	for _, v := range *chunk {
		set[v] = struct{}{}
	}

	return len(set) == markerLength
}

func main() {

	is, err := utils.InputScanner("input.txt")
	defer is.Close()

	if err != nil {
		log.Fatalln(err)
	}

	var input string

	is.Scan(func(val string) {
		// we know it's only 1 line
		input = val
	})

	markerLength := 14
	for i := markerLength; i <= len(input); i++ {
		chunk := input[i-markerLength : i]

		if allUnique(&chunk, markerLength) {
			fmt.Println("Found first unique chunk:", chunk)
			fmt.Println(i)
			return
		}

	}
}
