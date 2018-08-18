package lexer

import (
	"testing"

	"gitlab.com/frodsan/monkey/token"
)

func TestNextToken(t *testing.T) {
	input := "=+(){},;"

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	lex := New(input)

	for i, tt := range tests {
		tk := lex.NextToken()

		if tt.expectedType != tk.Type {
			t.Fatalf("tests[%d] - wrong token type. expected=%q, got=%q", i, tt.expectedType, tk.Type)
		}

		if tt.expectedLiteral != tk.Literal {
			t.Fatalf("tests[%d] - wrong literal. expected=%q, got=%q", i, tt.expectedLiteral, tk.Literal)
		}
	}
}
