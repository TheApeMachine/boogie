package boogie

import (
	"fmt"
	"log"
)

type LexerContext int

const (
	UNDEFINED LexerContext = iota
	VARIABLE
	ASSIGNMENT
	OPERATOR
	NUMBER
)

type Lexer struct {
	program     *Program
	delimiters  []string
	context     LexerContext
	currentChar string
	buf         string
}

func NewLexer(program *Program) *Lexer {
	return &Lexer{
		program:    program,
		delimiters: []string{" ", "\n"},
	}
}

func (lexer *Lexer) GenerateLexemes() chan *Lexeme {
	out := make(chan *Lexeme)

	go func() {
		defer close(out)

		for loc := range lexer.program.GenerateLines() {
			out <- lexer.parseLexemes(loc)
		}
	}()

	return out
}

func (lexer *Lexer) parseLexemes(loc *LineOfCode) *Lexeme {
	dref := *loc

	log.Println(fmt.Sprintf("boogie.parseLexemes(%v)", dref))

	for _, r := range dref {
		lexer.currentChar = string(r)
		fmt.Println(lexer.currentChar)

		switch lexer.context {
		case UNDEFINED:
			lexer.context = ContextHandler{UndefinedContext{}}.getNextState()
		}

		fmt.Println(fmt.Sprintf("context: %d", lexer.context))

		lexer.buf += lexer.currentChar
	}

	return &Lexeme{}
}
