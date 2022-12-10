package main

import (
	"fmt"
	"log"
	"marshallformula.codes/utils"
)

func allUnique(chunk *string) bool {
	set := make(map[rune]struct{})

	for _, v := range *chunk {
		set[v] = struct{}{}
	}

	return len(set) == 4
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

	for i := 4; i <= len(input); i++ {
		chunk := input[i-4 : i]

		if allUnique(&chunk) {
			fmt.Println("Found first unique chunk:", chunk)
			fmt.Println(i)
			return
		}

	}
}
