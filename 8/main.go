package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"marshallformula.codes/utils"
)

func main() {

	is, err := utils.InputScanner("input.txt")
	defer is.Close()

	if err != nil {
		log.Fatalln(err)
	}

	grid := make([][]int, 0)

	is.Scan(func(s string) {
		trees := make([]int, 0, len(s))
		for _, tree := range strings.Split(s, "") {
			num, err := strconv.Atoi(tree)
			if err != nil {
				log.Fatalln("Couldn't convert to number", tree)
			}
			trees = append(trees, num)
		}
		grid = append(grid, trees)
	})

	highScore := 0
	columns := len(grid[0])
	rows := len(grid)

	// all edges will have at least 1 viewing distance of 0
	// this should make the score of the tree always 0
	for i := 1; i < rows-1; i++ {
		for j := 1; j < columns-1; j++ {

			currentTree := grid[i][j]

			// the highest score this could get is 4
			// making an assumption that at least 12 tree will have a higher score
			if currentTree == 0 {
				continue
			}

			upCount := 0
			for up := i - 1; up >= 0; up-- {
				if grid[up][j] < currentTree {
					upCount++
				} else {
					upCount++
					break
				}
			}

			rightCount := 0
			for right := j + 1; right < columns; right++ {
				if grid[i][right] < currentTree {
					rightCount++
				} else {
					rightCount++
					break
				}
			}

			downCount := 0
			for down := i + 1; down < len(grid); down++ {
				if grid[down][j] < currentTree {
					downCount++
				} else {
					downCount++
					break
				}
			}

			leftCount := 0
			for left := j - 1; left >= 0; left-- {
				if grid[i][left] < currentTree {
					leftCount++
				} else {
					leftCount++
					break
				}
			}

			score := upCount * rightCount * downCount * leftCount

			if score > highScore {
				highScore = score
			}

		}
	}

	fmt.Println(highScore)

}
