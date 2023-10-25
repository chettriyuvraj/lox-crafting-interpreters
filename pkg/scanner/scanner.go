package scanner

import "fmt"

type Scanner struct {
	Source               string
	Tokens               []Token
	start, current, line int
}

func (s Scanner) ScanTokens() ([]Token, error) {
	fmt.Println("Scanning tokens")

	for !s.isAtEnd() {
		s.start = s.current
		s.scanToken()
	}

	s.Tokens = append(s.Tokens, Token{EOF, "", 0, s.line})
	return s.Tokens, nil
}

func (s Scanner) isAtEnd() bool {
	return true
}

func (s Scanner) scanToken() (Token, error) {
	return Token{}, nil
}
