package scanner

import (
	"fmt"
	"strconv"
)

type Scanner struct {
	Source               string
	Tokens               []Token
	start, current, line int
}

func (s *Scanner) ScanTokens() ([]Token, error) {
	fmt.Println("Scanning tokens")

	for !s.isAtEnd() {
		s.start = s.current

		/* TODO: Handle error */
		_, err := s.scanToken()
		fmt.Println(err)
	}

	s.Tokens = append(s.Tokens, Token{EOF, "", 0, s.line})
	return s.Tokens, nil
}

/**
 * Core token scanning function
 *
 * @return Token
 *
 * - Check type of token depending on a rough classification ie single char, double char
 * - Handle each type according to a set of rules
 * - Append token in the function itself to scanner.Tokens slice
 * - Return appended token/err if any
 *
 * NOTE: 1.
 **/

func (s *Scanner) scanToken() (Token, error) {
	var token Token
	var err error

	c := s.advance()
	switch c {

	/* Single character tokens - except SLASH which is handled separately */
	case "(":
		token = Token{Type: LEFT_PAREN, Line: s.line}
	case ")":
		token = Token{Type: RIGHT_PAREN, Line: s.line}
	case "{":
		token = Token{Type: LEFT_BRACE, Line: s.line}
	case "}":
		token = Token{Type: RIGHT_BRACE, Line: s.line}
	case ",":
		token = Token{Type: COMMA, Line: s.line}
	case ".":
		token = Token{Type: DOT, Line: s.line}
	case "-":
		token = Token{Type: MINUS, Line: s.line}
	case "+":
		token = Token{Type: PLUS, Line: s.line}
	case ";":
		token = Token{Type: SEMICOLON, Line: s.line}
	case "*":
		token = Token{Type: STAR, Line: s.line}

	/* Double character tokens */
	case "!":
		if s.match("=") {
			token = Token{Type: BANG_EQUAL, Line: s.line}
		} else {
			token = Token{Type: BANG, Line: s.line}
		}
	case "=":
		if s.match("=") {
			token = Token{Type: EQUAL_EQUAL, Line: s.line}
		} else {
			token = Token{Type: EQUAL, Line: s.line}
		}
	case "<":
		if s.match("=") {
			token = Token{Type: LESS_EQUAL, Line: s.line}
		} else {
			token = Token{Type: LESS, Line: s.line}
		}

	case ">":
		if s.match("=") {
			token = Token{Type: GREATER_EQUAL, Line: s.line}
		} else {
			token = Token{Type: GREATER, Line: s.line}
		}

	/* Comment or single slash */
	case "/":
		if s.match("/") {
			for !s.isAtEnd() && s.peek() != "\n" {
				s.advance()
			}
			/* We ignore comments */
			return Token{}, nil
		} else {
			token = Token{Type: SLASH, Line: s.line}
		}

	/* Literals */

	case "\"":
		token, err = s.handleString()
		/* fmt.Printf("\nString token %s", token) */

	/* Newlines and whitespaces */
	case "\n", "\r":
		s.line += 1
		return Token{}, nil

	case " ", "\t":
		return Token{}, nil

	default:
		if isDigit(c) {
			token, err = s.handleNumber()
		} else if isAlpha(c) {
			token, err = s.handleIdentifier()
		} else {
			return Token{}, fmt.Errorf("invalid token")
		}
	}

	if err != nil {
		return Token{}, err
	}

	/* Lexemes for string literals handled with quotes removed in their own handler func */
	/* fmt.Printf("\nString token %s", token) */
	token.Lexeme = s.Source[s.start:s.current]
	s.Tokens = append(s.Tokens, token)

	return token, nil
}

/********* Functions to scan literals *********/

func (s *Scanner) handleString() (Token, error) {
	for !s.isAtEnd() && s.peek() != "\"" {

		if s.peek() == "\n" {
			s.line += 1
		}

		s.advance()

		/* Ended without terminating string */
		if s.isAtEnd() {
			return Token{}, fmt.Errorf("unterminated string")
		}
	}

	/* Consume closing quote */
	s.advance()

	token := Token{Type: STRING, Line: s.line, Literal: s.Source[s.start+1 : s.current-1]}
	/* fmt.Printf("\nString token %s", token) */
	return token, nil
}

func (s *Scanner) handleNumber() (Token, error) {
	/* Scan first part of a possible fraction */
	for isDigit(s.peek()) {
		s.advance()
	}

	if s.peek() == "." && isDigit(s.peekNext()) {
		s.advance()
		for isDigit(s.peek()) {
			s.advance()
		}
	}

	val, err := strconv.ParseFloat(s.Source[s.start:s.current], 64)
	if err != nil {
		return Token{}, fmt.Errorf("invalid number literal conversion")
	}

	token := Token{Type: NUMBER, Line: s.line, Literal: val}
	return token, nil
}

func (s *Scanner) handleIdentifier() (Token, error) {
	for isAlphaNumeric(s.peek()) {
		s.advance()
	}

	lexeme := s.Source[s.start:s.current]

	tokentype, exists := KeywordsAsTokenType[lexeme]
	if exists == true {
		token := Token{Type: tokentype, Line: s.line}
		return token, nil
	}

	token := Token{Type: IDENTIFIER, Literal: lexeme, Line: s.line}
	return token, nil
}

/********* Helper Methods *********/

func (s *Scanner) isAtEnd() bool {
	/* Ignore Newline */
	/* fmt.Printf("\n\n ISATEND! Current %d, len source %d", s.current, len(s.Source)) */
	return s.current >= len(s.Source)-1
}

func (s *Scanner) advance() string {
	c := s.Source[s.current]
	s.current += 1
	/* fmt.Printf("\n\n ADVANCE! Current %d, len source %d", s.current, len(s.Source)) */
	return string(c)
}

func (s *Scanner) match(tomatch string) bool {
	if s.isAtEnd() {
		return false
	}

	/* fmt.Printf("\n\n MATCH! Current %d, len source %d", s.current, len(s.Source)) */

	c := string(s.Source[s.current])
	if c != tomatch {
		return false
	}

	s.current += 1
	return true
}

func (s *Scanner) peek() string {
	if s.isAtEnd() {
		return "\\0"
	}

	return string(s.Source[s.current])
}

func (s *Scanner) peekNext() string {
	if s.isAtEnd() || s.current+1 >= len(s.Source) {
		return "\\0"
	}

	return string(s.Source[s.current+1])
}

/********* Helper Functions *********/

func isDigit(s string) bool {
	return s >= "0" && s <= "9"
}

func isAlpha(c string) bool {
	return c == "_" || (c >= "a" && c <= "z") || (c >= "A" && c <= "Z")
}

func isAlphaNumeric(c string) bool {
	return (isAlpha(c) || (c >= "0" && c <= "9"))
}
