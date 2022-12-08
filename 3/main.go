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

func findTheCommon(groupItems *[][]int32) (int32, error) {
	if len(*groupItems) != 3 {
		return 0, errors.New(fmt.Sprintf("Invalid group size: %d", len(*groupItems)))
	}

	sack0 := toSetIsh(&(*groupItems)[0])
	sack1 := toSetIsh(&(*groupItems)[1])
	sack2 := toSetIsh(&(*groupItems)[2])

	for val := range sack0 {
		_, ok1 := sack1[val]
		_, ok2 := sack2[val]
		if ok1 && ok2 {
			return val, nil
		}
	}

	return 0, errors.New("Could not find the common item")
}

func main() {
	is, err := utils.InputScanner("input.txt")
	defer is.Close()

	if err != nil {
		log.Fatalln(err)
	}

	priorityValues := assignPriorities()

	sum := 0

	elfGroup := make([][]int32, 3, 3)
	groupdIdx := 0

	is.Scan(func(val string) {
		letters := []rune(val)

		elfGroup[groupdIdx] = uniqueify(&letters)

		if groupdIdx < 2 {

			groupdIdx++

		} else {

			commonItem, err := findTheCommon(&elfGroup)

			if err != nil {
				log.Fatalln(err)
			}

			priority, ok := priorityValues[commonItem]

			if !ok {
				log.Fatalf("Couldn't find priority for %c\n", commonItem)
			}

			sum += int(priority)
			groupdIdx = 0
			elfGroup = make([][]int32, 3, 3)
		}
	})

	fmt.Println("Sum of Priorities: ", sum)
}
