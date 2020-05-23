package boogie

import (
	"bufio"
	"log"
	"os"
)

// LineOfCode is a custom type that refers to string.
// It is used to keep track of the intention of the value.
type LineOfCode string

// Program is a wrapper around the program file data.
// It is responsible for generating lines of code that
// can be consumed by a lexer.
type Program struct {
	fileHandler *os.File
}

func NewProgram(filePath string) *Program {
	fh, err := os.Open(filePath)

	if err != nil {
		log.Fatal(err)
	}

	return &Program{
		fileHandler: fh,
	}
}

func (program *Program) GenerateLines() chan *LineOfCode {
	out := make(chan *LineOfCode)
	scanner := bufio.NewScanner(program.fileHandler)

	go func() {
		defer program.fileHandler.Close()
		defer close(out)

		for scanner.Scan() {
			loc := LineOfCode(scanner.Text())
			out <- &loc
		}

		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
	}()

	return out
}
