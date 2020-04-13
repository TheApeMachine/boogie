package boogie

type Parser struct {
	In  chan Lexeme
	ast map[string][]string
}

func NewParser() *Parser {
	ast := make(map[string][]string)
	ast["start"] = make([]string, 0)

	return &Parser{
		In:  make(chan Lexeme),
		ast: ast,
	}
}

func (parser *Parser) Run() {
	for {
		lex := <-parser.In
		_ = parser.process(lex)
	}
}

func (parser *Parser) process(lex Lexeme) error {
	parser.ast["start"] = append(parser.ast["start"], lex.Val)
	return nil
}
