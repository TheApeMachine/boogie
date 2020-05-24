package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/theapemachine/boogie/apeterm"
	"github.com/theapemachine/boogie/boogie"
	"github.com/theapemachine/boogie/eddie"
)

func main() {
	editFile := flag.String("e", "test.boo", "input file to edit")
	flag.Parse()

	prog := os.Args[1]

	if editFile != nil {
		ui := apeterm.NewUI()
		buffer := eddie.NewBuffer(ui, editFile)
		//buffer.Load()
		buffer.SetFocus()
		prog = buffer.GetData()
	}

	program := boogie.NewProgram(prog)
	lexer := boogie.NewLexer(program)
	parser := boogie.NewParser(lexer)
	evaluator := boogie.NewEvaluator(parser)

	for out := range evaluator.GenerateOutput() {
		fmt.Print(*out)
	}
}
