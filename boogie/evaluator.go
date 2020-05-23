package boogie

type Evaluator struct {
	parser *Parser
}

func NewEvaluator(parser *Parser) *Evaluator {
	return &Evaluator{
		parser: parser,
	}
}

func (evaluator *Evaluator) GenerateOutput() chan *string {
	out := make(chan *string)

	go func() {
		defer close(out)
		str := "test"
		out <- &str
	}()

	return out
}
