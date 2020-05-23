package boogie

type TokenType int

const (
	DELIM TokenType = iota
)

type TokenValue string

type Lexeme struct {
	tokenType  TokenType
	tokenValue TokenValue
}
