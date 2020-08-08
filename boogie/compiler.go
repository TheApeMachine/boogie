package boogie

type Compiler struct{}

func NewCompiler() *Compiler {
	return &Compiler{}
}

func (compiler *Compiler) CompileFromAST(parser *Parser) []uint8 {
	_ = parser.BuildAST()
	return []uint8{}
}
