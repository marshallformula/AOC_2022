package main

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"

	"marshallformula.codes/utils"
)

var motions []*motion = make([]*motion, 0)

type position struct {
	x int
	y int
}

func (p *position) move(direction *string) {
	switch *direction {
	case "U":
		p.y++
	case "D":
		p.y--
	case "R":
		p.x++
	case "L":
		p.x--
	default:
		log.Fatal("Should not get here")
	}
}

func (p *position) asKey() string {
	return fmt.Sprint(p.x, ".", p.y)
}

type motion struct {
	direction string
	steps     int
}

func fromFields(fields []string) (*motion, error) {
	if len(fields) != 2 {
		return nil, errors.New(fmt.Sprintf("Invalid fields for motion. %s", fields))
	}

	direction := fields[0]
	steps, err := strconv.Atoi(fields[1])

	if err != nil {
		return nil, err
	}

	return &motion{direction, steps}, nil
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func moveTail(head *position, tail *position) {
	xDistance := head.x - tail.x
	yDistance := head.y - tail.y

	absX := Abs(xDistance)
	absY := Abs(yDistance)

	moveRight := xDistance > 0
	moveUp := yDistance > 0

	// they are close enough
	if (absX + absY) < 2 {
		return
	}

	if (absX + absY) == 2 {
		// positions are diagonal
		if absX == absY {
			return
		}

		// head is on the same y, just need to move it horizontally
		if absX > absY {
			// if the distane is positive move right otherwise move left
			if moveRight {
				tail.x++
			} else {
				tail.x--
			}
		} else {
			// if the distane is positive move up otherwise move down
			if moveUp {
				tail.y++
			} else {
				tail.y--
			}
		}
		return
	}

	// we need to move on both axis
	if (absX + absY) <= 4 {

		// need to get x to be 1 distance and y to be 0
		// positive move right ...
		if moveRight {
			tail.x++
		} else {
			tail.x--
		}

		if moveUp {
			tail.y++
		} else {
			tail.y--
		}
		// }

		return
	}

	log.Fatalln("Inexplicable position (head,tail)", absX, absY, head, tail)

}

func walk() map[string]struct{} {
	visited := make(map[string]struct{}, 0)

	knots := make([]*position, 0)

	for i := 0; i < 10; i++ {
		knots = append(knots, &position{0, 0})
	}

	// record first position
	visited[knots[0].asKey()] = struct{}{}

	for _, motion := range motions {
		for i := 0; i < motion.steps; i++ {

			head := knots[0]
			head.move(&motion.direction)

			for k := 1; k < len(knots); k++ {
				intHead := knots[k-1]
				tail := knots[k]

				moveTail(intHead, tail)

			}

			visited[knots[9].asKey()] = struct{}{}
		}
	}

	return visited
}

func main() {

	is, err := utils.InputScanner("input.txt")
	defer is.Close()

	if err != nil {
		log.Fatalln(err)
	}

	is.Scan(func(s string) {
		fields := strings.Fields(s)
		motion, err := fromFields(fields)

		if err != nil {
			log.Fatalln(err)
		}

		motions = append(motions, motion)
	})

	visited := walk()

	fmt.Println(len(visited))

}
