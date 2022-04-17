package scanner

import (
	"glox/glerrors"
)

type Scanner struct {
	source string
	Tokens []Token

	start   int
	current int
	line    int
}

func NewScanner(source string) *Scanner {
	return &Scanner{source: source, line: 1, start: 0, current: 0}
}

func (s *Scanner) isAtEnd() bool {
	return s.current >= len(s.source)
}

func (s *Scanner) ScanTokens() []Token {

	for !s.isAtEnd() {
		s.start = s.current
		s.scanToken()
	}

	s.Tokens = append(s.Tokens, NewToken(Eof, "", nil, s.line))
	return s.Tokens
}

func iif(expr bool, then TokenType, el TokenType) TokenType {
	if expr {
		return then
	}

	return el
}

func (s *Scanner) scanToken() {
	c := s.advance()
	switch c {
	case '(':
		s.addTokenType(LeftParen)
		break
	case ')':
		s.addTokenType(RightParen)
		break
	case '{':
		s.addTokenType(LeftBrace)
		break
	case '}':
		s.addTokenType(RightBrace)
		break
	case ',':
		s.addTokenType(Comma)
		break
	case '.':
		s.addTokenType(Dot)
		break
	case '-':
		s.addTokenType(Minus)
		break
	case '+':
		s.addTokenType(Plus)
		break
	case ';':
		s.addTokenType(Semicolon)
		break
	case '*':
		s.addTokenType(Star)
		break
	case '!':
		s.addTokenType(iif(s.match('='), BangEqual, Bang))
		break
	case '=':
		s.addTokenType(iif(s.match('='), EqualEqual, Equal))
		break
	case '<':
		s.addTokenType(iif(s.match('='), LessEqual, Less))
		break
	case '>':
		s.addTokenType(iif(s.match('='), GreaterEqual, Greater))
		break
	case '/':
		if s.match('/') {
			// A comment goes until the end of the line.
			for s.peek() != '\n' && !s.isAtEnd() {
				s.advance()
			}
		} else {
			s.addTokenType(Slash)
		}
		break
	case ' ', '\r', '\t':
		// Ignore whitespace.
		break
	case '\n':
		s.line++
	case '"':
		s.string()
		break
	default:
		glerrors.Error(s.line, "Unexpected character")
	}
}

func (s *Scanner) string() {

	for s.peek() != '"' && !s.isAtEnd() {
		if s.peek() == '\n' {
			s.line++
		}
		s.advance()
	}

	if s.isAtEnd() {
		glerrors.Error(s.line, "Unterminated string")
		return
	}

	// The closing "
	s.advance()

	val := s.source[s.start+1 : s.current-1]
	s.addToken(String, val)

}

func (s *Scanner) peek() byte {
	if s.isAtEnd() {
		return '\000'
	}

	return s.source[s.current]
}

func (s *Scanner) match(expected byte) bool {
	if s.isAtEnd() {
		return false
	}

	if s.source[s.current] != expected {
		return false
	}

	s.current++
	return true
}

func (s *Scanner) advance() byte {
	s.current++
	return s.source[s.current-1]
}

func (s *Scanner) addTokenType(t TokenType) {
	s.addToken(t, nil)
}

func (s *Scanner) addToken(t TokenType, literal any) {
	text := s.source[s.start:s.current]
	s.Tokens = append(s.Tokens, NewToken(t, text, literal, s.line))
}
