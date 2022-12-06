package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

type ElfSnacks []int

func (snackers ElfSnacks) topThreeSum() int {
	sum := 0
	for _, cals := range snackers[len(snackers)-3:] {
		sum += cals
	}
	return sum
}

func main() {

	input, err := os.Open("input.txt")
	defer input.Close()

	if err != nil {
		log.Fatalln(err)
	}

	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanLines)

	currentElf := 0
	elves := ElfSnacks{}

	for scanner.Scan() {
		item := scanner.Text()
		newElf := len(item) == 0

		if newElf {

			elves = append(elves, currentElf)
			currentElf = 0
			continue
		}

		calories, err := strconv.Atoi(item)

		if err != nil {
			log.Fatalln(err)
		}

		currentElf += calories

	}

	sort.Ints(elves)
	fmt.Printf("Sum of top 3: %d\n", elves.topThreeSum())
}
