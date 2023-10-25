package scanner

/**************** ENUM TYPES ****************/

type TokenType int
type Literal int

const (
	/* Single-character tokens */
	LEFT_PAREN TokenType = iota
	RIGHT_PAREN
	LEFT_BRACE
	RIGHT_BRACE
	COMMA
	DOT
	MINUS
	PLUS
	SEMICOLON
	SLASH
	STAR

	/* One or two character tokens */
	BANG
	BANG_EQUAL
	EQUAL
	EQUAL_EQUAL
	GREATER
	GREATER_EQUAL
	LESS
	LESS_EQUAL

	/* Keywords */
	AND
	CLASS
	ELSE
	FALSE
	FUN
	FOR
	IF
	NIL
	OR
	PRINT
	RETURN
	SUPER
	THIS
	TRUE
	VAR
	WHILE

	EOF
)

/* NOTE: 0 stands for nil */
const (
	IDENTIFIER Literal = iota + 1
	STRING
	NUMBER
)

/**************** Token Struct ****************/

type Token struct {
	Type    TokenType
	Lexeme  string
	Literal Literal
	Line    int
}

func (t Token) String() string {
	return "Type" + t.Type + "Lexeme" + t.Lexeme + "Literal" + t.Literal + "Line" + t.Line
}
