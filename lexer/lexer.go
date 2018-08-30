package lexer

import "github.com/khanhtc1202/k/token"

type Lexer struct {
}

func New(input string) *Lexer {
	return &Lexer{}
}

func (l *Lexer) NextToken() token.Token {

}
