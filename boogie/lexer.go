package boogie

type Lexer struct {
	program *Program
}

func NewLexer(program *Program) *Lexer {
	return &Lexer{
		program: program,
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
	return &Lexeme{}
}
