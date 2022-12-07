package main

import (
	"fmt"
	"log"
	"strings"

	"marshallformula.codes/utils"
)

type Result int

const (
	WIN Result = iota
	LOSE
	DRAW
)

const (
	ROUND_WINNER = 6
	ROUND_DRAW   = 3
	ROUND_LOSER  = 0
	ROCK         = 1
	PAPER        = 2
	SCISSORS     = 3
)

type Shape int

func (a *Shape) play(b *Shape) Result {
	switch *a {

	case ROCK:
		switch *b {
		case PAPER:
			return LOSE
		case SCISSORS:
			return WIN
		default:
			return DRAW
		}

	case PAPER:
		switch *b {
		case ROCK:
			return WIN
		case SCISSORS:
			return LOSE
		default:
			return DRAW
		}

	default:
		switch *b {
		case ROCK:
			return LOSE
		case PAPER:
			return WIN

		default:
			return DRAW
		}
	}
}

type ShapeDecoder map[string]Shape

var shapeDecoder = ShapeDecoder{
	"A": ROCK,
	"B": PAPER,
	"C": SCISSORS,
}

type StrategyDecoder map[string]Result

var strategyDecoder = StrategyDecoder{
	"X": LOSE,
	"Y": DRAW,
	"Z": WIN,
}

func (strategy *Result) choseShape(opponentShape *Shape) Shape {
	switch *strategy {
	case WIN:
		switch *opponentShape {
		case ROCK:
			return PAPER
		case PAPER:
			return SCISSORS
		default:
			return ROCK
		}

	case LOSE:
		switch *opponentShape {
		case ROCK:
			return SCISSORS

		case PAPER:
			return ROCK

		default:
			return PAPER
		}

	default:
		return *opponentShape
	}
}

func roundProcessor(elfScore *int, myScore *int) func(string) {
	return func(round string) {
		plays := strings.Fields(round)

		elfShape, elfOk := shapeDecoder[plays[0]]
		myStrategy, myOk := strategyDecoder[plays[1]]

		if !(elfOk && myOk) {
			log.Fatalln("Couldn't decode the points for the round", plays)
		}

		myShape := myStrategy.choseShape(&elfShape)

		switch elfShape.play(&myShape) {

		case WIN:
			// elf wins
			*elfScore += (ROUND_WINNER + int(elfShape))
			*myScore += (ROUND_LOSER + int(myShape))

		case LOSE:
			// elf loses
			*elfScore += (ROUND_LOSER + int(elfShape))
			*myScore += (ROUND_WINNER + int(myShape))

		default:
			*elfScore += (ROUND_DRAW + int(elfShape))
			*myScore += (ROUND_DRAW + int(myShape))

		}

	}

}

func main() {

	is, err := utils.InputScanner("input.txt")
	defer is.Close()

	if err != nil {
		log.Fatalln(err)
	}

	elfScore := 0
	myScore := 0

	is.Scan(roundProcessor(&elfScore, &myScore))

	fmt.Printf("Elf: %d, Me: %d\n", elfScore, myScore)
}
