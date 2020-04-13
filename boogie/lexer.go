package boogie

import (
	"strings"
)

type ContextState int

const (
	START ContextState = iota
	KEYWORD
)

type Lexeme struct {
	Id  string
	Val string
}

type Lexer struct {
	In       chan string
	parser   *Parser
	curChar  string
	curCtx   ContextState
	keywords []string
	buf      string
}

func NewLexer(in chan string) *Lexer {
	return &Lexer{
		In:       in,
		parser:   NewParser(),
		curCtx:   START,
		keywords: []string{"QUIT"},
	}
}

func (lexer *Lexer) Run() {
	go lexer.parser.Run()

	for {
		loc := <-lexer.In
		err := lexer.process(loc)

		if err != nil {
			panic(err)
		}
	}
}

func (lexer *Lexer) process(loc string) error {
	for _, r := range loc {
		lexer.curChar = string(r)

		switch lexer.curCtx {
		case START:
			lexer.buf += lexer.curChar

			if lexer.stringInSlice(lexer.buf, lexer.keywords) {
				lexer.curCtx = KEYWORD
			}
		case KEYWORD:
			lexer.parser.In <- Lexeme{Id: "key", Val: "QUIT"}
		}
	}

	return nil
}

func (lexer *Lexer) stringInSlice(str string, list []string) bool {
	for _, v := range list {
		if v == strings.ToUpper(str) {
			return true
		}
	}

	return false
}
