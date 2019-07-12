package scanner

import (
	"unicode/utf8"

	"github.com/tsukanov-as/radar/bsl/tokens"
)

// Scanner ...
type Scanner struct {
	src      string
	offset   int
	rdOffset int
	ch       rune
	Line     int
	eot      bool
	err      int

	lit      string
	pos      int
	begpos   int
	endpos   int
	endline  int
	comments map[int]string
}

func (s *Scanner) readRune() {
	s.offset = s.rdOffset
	r, w := rune(s.src[s.offset]), 1
	switch {
	case r == 0:
		s.error("illegal character NUL")
	case r >= 0x80:
		r, w = utf8.DecodeRuneInString(s.src[s.offset:])
		if r == utf8.RuneError && w == 1 {
			s.error("illegal UTF-8 encoding")
		}
	}
	s.rdOffset += w
	s.pos++
	s.ch = r
}

// Init ...
func (s *Scanner) Init(src string) {
	s.src = src
	s.Line = 1
	s.comments = make(map[int]string)
	s.offset = 0
	s.rdOffset = 0
	if len(s.src) > 0 {
		s.readRune()
		s.eot = false
	} else {
		s.eot = true
		s.ch = -1 // eof
	}
}

func (s *Scanner) next() {
	if s.rdOffset < len(s.src) {
		if s.ch == '\n' {
			s.Line++
		}
		s.readRune()
	} else {
		s.offset = len(s.src)
		s.eot = true
		s.ch = -1 // eof
	}
}

func isLetter(ch rune) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || 0x410 <= ch && ch <= 0x44F || ch == '_' || ch == 0x451 // ch >= 0x80 && unicode.IsLetter(ch)
}

func isDigit(ch rune) bool {
	return '0' <= ch && ch <= '9' // || ch >= 0x80 && unicode.IsDigit(ch)
}

func isSpace(ch rune) bool {
	return ch == ' ' || ch == '\t' || ch == '\n' || ch == '\r' || ch == 0xA0 // || ch >= 0x80 && unicode.IsSpace(ch)
}

func (s *Scanner) error(msg string) {
	println(msg)
}

func (s *Scanner) skipWhitespace() {
	for isSpace(s.ch) {
		s.next()
	}
}

func (s *Scanner) scanComment() string {
	s.next()
	start := s.offset
	for s.ch != '\n' && !s.eot {
		s.next()
	}
	return s.src[start:s.offset]
}

func (s *Scanner) scanIdentifier() string {
	start := s.offset
	s.next()
	for isLetter(s.ch) || isDigit(s.ch) {
		s.next()
	}
	return s.src[start:s.offset]
}

func (s *Scanner) scanNumber() string {
	start := s.offset
	s.next()
	for isDigit(s.ch) {
		s.next()
	}
	if s.ch == '.' {
		s.next()
		for isDigit(s.ch) {
			s.next()
		}
	}
	return s.src[start:s.offset]
}

func (s *Scanner) scanString() string {
	start := s.offset
	s.ch = '"'
	for s.ch == '"' {
		s.next()
		for s.ch != '"' && s.ch != '\n' && !s.eot {
			s.next()
		}
		if s.ch == '"' {
			s.next()
		}
	}
	return s.src[start:s.offset]
}

func (s *Scanner) scanDateTime() string {
	s.next()
	start := s.offset
	for s.ch != '\'' && !s.eot {
		s.next()
	}
	return s.src[start:s.offset]
}

// Scan ...
func (s *Scanner) Scan() (tok tokens.Token, lit string) {

	s.endpos = s.pos
	s.endline = s.Line

	if s.lit != "" && s.lit[len(s.lit)-1] == '\n' {
		s.Line++
	}

scan:

	s.skipWhitespace()

	switch ch := s.ch; {
	case isLetter(ch):
		lit = s.scanIdentifier()
		tok = tokens.Lookup(lit)
	case isDigit(ch):
		lit = s.scanNumber()
		tok = tokens.NUMBER
	default:
		switch ch {
		case '"':
			lit = s.scanString()
			if lit[len(lit)-1] == '"' {
				tok = tokens.STRING
			} else {
				tok = tokens.STRINGBEG
			}
		case '|':
			lit = s.scanString()
			if lit[len(lit)-1] == '"' {
				tok = tokens.STRINGEND
			} else {
				tok = tokens.STRINGMID
			}
		case '\'':
			lit = s.scanDateTime()
			if s.ch == '\'' {
				s.next()
			} else {
				s.error("error: date not closed")
			}
			tok = tokens.DATETIME
		case '=':
			tok = tokens.EQL
			s.next()
		case '<':
			s.next()
			if s.ch == '=' {
				tok = tokens.LEQ
				s.next()
			} else {
				tok = tokens.LSS
			}
		case '>':
			s.next()
			if s.ch == '=' {
				tok = tokens.GEQ
				s.next()
			} else {
				tok = tokens.GTR
			}
		case '+':
			tok = tokens.ADD
			s.next()
		case '-':
			tok = tokens.SUB
			s.next()
		case '*':
			tok = tokens.MUL
			s.next()
		case '/':
			s.next()
			if s.ch == '/' {
				s.comments[s.Line] = s.scanComment()
				tok = tokens.COMMENT
				s.next()
				goto scan
			} else {
				tok = tokens.DIV
			}
		case '%':
			tok = tokens.MOD
			s.next()
		case '(':
			tok = tokens.LPAREN
			s.next()
		case ')':
			tok = tokens.RPAREN
			s.next()
		case '[':
			tok = tokens.LBRACK
			s.next()
		case ']':
			tok = tokens.RBRACK
			s.next()
		case '?':
			tok = tokens.TERNARY
			s.next()
		case ',':
			tok = tokens.COMMA
			s.next()
		case '.':
			tok = tokens.PERIOD
			s.next()
		case ':':
			tok = tokens.COLON
			s.next()
		case ';':
			tok = tokens.SEMICOLON
			s.next()
		case '&':
			s.next()
			if !isLetter(s.ch) {
				s.error("error: expected directive")
			}
			lit = s.scanIdentifier()
			tok = tokens.LookupDirective(lit)
			if tok == tokens.BAD {
				s.error("error: unknown directive")
			}
		case '#':
			s.next()
			s.skipWhitespace()
			if !isLetter(s.ch) {
				s.error("error: expected directive")
			}
			lit = s.scanIdentifier()
			tok = tokens.LookupPrepInst(lit)
			if tok == tokens.BAD {
				s.error("error: unknown preprocessor")
			}
		case '~':
			s.next()
			s.skipWhitespace()
			if isLetter(s.ch) || isDigit(s.ch) {
				lit = s.scanIdentifier()
			}
			tok = tokens.LABEL
		case -1:
			tok = tokens.EOF
		}
	}

	return tok, lit
}
