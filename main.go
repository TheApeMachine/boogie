package main

import (
	"fmt"
	"os"

	"github.com/theapemachine/boogie/boogie"
)

func main() {
	program := boogie.NewProgram(os.Args[1])
	lexer := boogie.NewLexer(program)
	parser := boogie.NewParser(lexer)
	evaluator := boogie.NewEvaluator(parser)

	for out := range evaluator.GenerateOutput() {
		fmt.Print(*out)
	}
}
