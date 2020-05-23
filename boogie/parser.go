package boogie

type ASTNode struct{}

type AST struct{}

type Parser struct {
	lexer *Lexer
}

func NewParser(lexer *Lexer) *Parser {
	return &Parser{
		lexer: lexer,
	}
}

func (parser *Parser) BuildAST() *AST {
	var ast AST

	for lexeme := range parser.lexer.GenerateLexemes() {
		_ = lexeme
	}

	return &ast
}
