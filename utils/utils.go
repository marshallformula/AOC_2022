package utils

import (
	"bufio"
	"os"
)

type inputScanner struct {
	input *os.File
}

func (is *inputScanner) Close() {
	is.input.Close()
}

func (is *inputScanner) Scan(fn func(string)) {
	scanner := bufio.NewScanner(is.input)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
    fn(scanner.Text())
	}

}

func InputScanner(inputFile string) (*inputScanner, error) {
	input, err := os.Open(inputFile)
	
	if err != nil {
	return nil, err
	}

	return &inputScanner{input}, nil
}
