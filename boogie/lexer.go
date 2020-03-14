package boogie

import "fmt"

type Lexer struct {
	In chan string
}

func NewLexer(in chan string) *Lexer {
	return &Lexer{
		In: in,
	}
}

func (lexer Lexer) Run() {
	for {
		msg := <-lexer.In
		fmt.Println(msg)
	}
}
