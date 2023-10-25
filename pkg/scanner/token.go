package scanner

import "fmt"

/********* Enum Types *********/

type TokenType int

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

	/* Literals */
	IDENTIFIER
	STRING
	NUMBER

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

/********* Token Struct *********/

type Token struct {
	Type    TokenType
	Lexeme  string
	Literal any
	Line    int
}

/******** Token Helpers  ********/

var TokenTypeAsLexeme map[TokenType]string = map[TokenType]string{
	LEFT_PAREN:  "(",
	RIGHT_PAREN: ")",
	LEFT_BRACE:  "{",
	RIGHT_BRACE: "}",
	COMMA:       ",",
	DOT:         ".",
	MINUS:       "-",
	PLUS:        "+",
	SEMICOLON:   ";",
	SLASH:       "/",
	STAR:        "*",

	/* One or two character tokens */
	BANG:          "!",
	BANG_EQUAL:    "!=",
	EQUAL:         "=",
	EQUAL_EQUAL:   "==",
	GREATER:       ">",
	GREATER_EQUAL: ">=",
	LESS:          "<",
	LESS_EQUAL:    "<=",

	/* Literals */
	IDENTIFIER: "IDENTIFIER",
	STRING:     "STRING",
	NUMBER:     "NUMBER",

	/* Keywords */
	AND:    "and",
	CLASS:  "class",
	ELSE:   "else",
	FALSE:  "false",
	FUN:    "fun",
	FOR:    "for",
	IF:     "if",
	NIL:    "nil",
	OR:     "or",
	PRINT:  "print",
	RETURN: "return",
	SUPER:  "super",
	THIS:   "this",
	TRUE:   "true",
	VAR:    "var",
	WHILE:  "while",

	/* EOF */
}

var KeywordsAsTokenType map[string]TokenType = map[string]TokenType{
	"and":    AND,
	"class":  CLASS,
	"else":   ELSE,
	"false":  FALSE,
	"fun":    FUN,
	"for":    FOR,
	"if":     IF,
	"nil":    NIL,
	"or":     OR,
	"print":  PRINT,
	"return": RETURN,
	"super":  SUPER,
	"this":   THIS,
	"true":   TRUE,
	"var":    VAR,
	"while":  WHILE,
}

func (t Token) String() string {
	tokenType, _ := TokenTypeAsLexeme[t.Type]

	return fmt.Sprintf("\nType %s Lexeme %s Literal %v Line %d", tokenType, t.Lexeme, t.Literal, t.Line)
}
