package main

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"

	"marshallformula.codes/utils"
)

type instruction interface {
	hacky_go_garbage()
}

type registerAdd int
type noop struct{}

func (ra registerAdd) hacky_go_garbage() {}
func (n noop) hacky_go_garbage()         {}

func newInstruction(fields []string) (instruction, error) {
	if len(fields) == 2 {
		reg, err := strconv.Atoi(fields[1])
		if err != nil {
			return noop{}, err
		}
		return registerAdd(reg), nil
	}

	if len(fields) == 1 {
		return noop{}, nil
	}

	return noop{}, errors.New(fmt.Sprint("Recieved invalid instruction: ", fields))
}

type cycleList []int

func (cycles *cycleList) signalAt(cycleNum int) int {
	sum := 1

	for i := 0; i < cycleNum-1; i++ {
		sum += (*cycles)[i]
	}

	return sum * cycleNum

}

func main() {

	is, err := utils.InputScanner("input.txt")
	defer is.Close()

	if err != nil {
		log.Fatalln(err)
	}

	instructions := make([]*instruction, 0)

	is.Scan(func(s string) {
		f := strings.Fields(s)

		inst, err := newInstruction(f)

		if err != nil {
			log.Fatal(err)
		}

		instructions = append(instructions, &inst)

	})

	cycles := make(cycleList, 0)

	for _, inst := range instructions {
		switch i := (*inst).(type) {

		case noop:
			cycles = append(cycles, 0)

		case registerAdd:
			cycles = append(cycles, 0)
			cycles = append(cycles, int(i))

		default:
			log.Fatalln("This shouldn't be possible")
		}
	}

	/* sum := 0

	for _, s := range []int{20, 60, 100, 140, 180, 220} {

		sum += cycles.signalAt(s)

	}

	fmt.Println(sum) */

	registerX := 1
	pixel := 0

	for _, c := range cycles {

		if utils.Abs(registerX-pixel) < 2 {
			fmt.Print("#")
		} else {
			fmt.Print(".")
		}

		registerX += c

		if (pixel+1)%40 == 0 {
			fmt.Println()
			pixel = 0
			continue
		}

		pixel++

	}

}
