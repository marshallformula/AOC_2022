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

	visibleTrees := make([]int, 0)
	columns := len(grid[0])
	rows := len(grid)

	for i := 1; i < rows-1; i++ {
		for j := 1; j < columns-1; j++ {

			currentTree := grid[i][j]

			if currentTree == 0 {
				continue
			}

			visible := true

			for up := i - 1; up >= 0; up-- {
				if grid[up][j] >= currentTree {
					visible = false
					break
				}
			}

			if visible {
				visibleTrees = append(visibleTrees, currentTree)
				continue
			}
			visible = true

			for right := j + 1; right < columns; right++ {
				if grid[i][right] >= currentTree {
					visible = false
					break
				}
			}

			if visible {
				visibleTrees = append(visibleTrees, currentTree)
				continue
			}
			visible = true

			for down := i + 1; down < len(grid); down++ {
				if grid[down][j] >= currentTree {
					visible = false
					break
				}
			}

			if visible {
				visibleTrees = append(visibleTrees, currentTree)
				continue
			}
			visible = true

			for left := j - 1; left >= 0; left-- {
				if grid[i][left] >= currentTree {
					visible = false
					break
				}
			}

			if visible {
				visibleTrees = append(visibleTrees, currentTree)
			}

		}
	}

	edgeTrees := ((rows) * 2) + (((columns) - 2) * 2)

	// fmt.Println(visibleTrees)
	fmt.Println(len(visibleTrees) + edgeTrees)

}
