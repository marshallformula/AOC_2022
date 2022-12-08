package main

import (
	"errors"
	"fmt"
	"log"

	"marshallformula.codes/utils"
)

func assignPriorities() map[int32]int16 {
	priorities := make(map[int32]int16)
	var priorityValue int16 = 1

	for i := 'a'; i <= 'z'; i++ {
		priorities[i] = priorityValue
		priorityValue++
	}

	for i := 'A'; i <= 'Z'; i++ {
		priorities[i] = priorityValue
		priorityValue++
	}

	return priorities

}

/* func sortTheSack(sack *[]int32) {
	sort.SliceStable(*sack, func(i, j int) bool {
		return (*sack)[i] < (*sack)[j]
	})
} */

func toSetIsh[T comparable](values *[]T) map[T]struct{} {

	hash := make(map[T]struct{})

	for _, v := range *values {
		hash[v] = struct{}{}
	}

	return hash
}

func uniqueify[T comparable](values *[]T) []T {
	set := toSetIsh(values)
	keys := make([]T, 0, len(set))
	for k := range set {
		keys = append(keys, k)
	}
	return keys
}

func findTheDupe(sack1 *[]int32, sack2 *[]int32) (int32, error) {

	sack1Set := toSetIsh(sack1)
	uniqueSack2 := uniqueify(sack2)

	for _, v := range uniqueSack2 {
		if _, ok := sack1Set[v]; ok {
			return v, nil
		}
	}

	return 0, errors.New("Could not find the duplicate item")
}

func main() {
	is, err := utils.InputScanner("input.txt")
	defer is.Close()

	if err != nil {
		log.Fatalln(err)
	}

	priorityValues := assignPriorities()

	sum := 0

	is.Scan(func(val string) {
		letters := []rune(val)
		length := len(letters)

		sack1 := letters[:(length / 2)]
		sack2 := letters[(length / 2):]

		duplicateItem, err := findTheDupe(&sack1, &sack2)

		if err != nil {
			log.Fatalln(err)
		}

		priority, ok := priorityValues[duplicateItem]

		if !ok {
			log.Fatalf("Couldn't find priority for %c\n", duplicateItem)
		}

		sum += int(priority)

	})

	fmt.Println("Sum of Priorities: ", sum)
}
