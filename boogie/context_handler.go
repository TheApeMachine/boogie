package boogie

type ContextHandler struct {
	handler interface {
		getNextState(lexer *Lexer) LexerContext
	}
}

func (context ContextHandler) getNextState(lexer *Lexer) LexerContext {
	return context.handler.getNextState(lexer)
}

type UndefinedContext struct{}

func (undefined UndefinedContext) getNextState(lexer *Lexer) LexerContext {
	if lexer.currentChar != utils.StringInList(lexer.currentChar, lexer.delimiters) {
		return UNDEFINED
	}

	return VARIABLE
}

type VariableContext struct{}

func (variable VariableContext) getNextState(lexer *Lexer) LexerContext {
	return VARIABLE
}
