package mydsl

import (
	"github.com/dezlitz/dsl"
)

const (
	TOKEN_LITERAL     dsl.TokenType = "LITERAL"
	TOKEN_PLUS        dsl.TokenType = "PLUS"
	TOKEN_MINUS       dsl.TokenType = "MINUS"
	TOKEN_MULTIPLY    dsl.TokenType = "MULTIPLY"
	TOKEN_DIVIDE      dsl.TokenType = "DIVIDE"
	TOKEN_OPEN_PAREN  dsl.TokenType = "OPEN_PAREN"
	TOKEN_CLOSE_PAREN dsl.TokenType = "CLOSE_PAREN"
	TOKEN_ASSIGN      dsl.TokenType = "ASSIGN"
	TOKEN_VARIABLE    dsl.TokenType = "VARIABLE"
	TOKEN_COMMENT     dsl.TokenType = "COMMENT"
	TOKEN_NL          dsl.TokenType = "NL"
	TOKEN_EOF         dsl.TokenType = "EOF"
)

func Scan(s *dsl.Scanner) dsl.Token {
	if recover {
		s.ExpectNot(dsl.ExpectNotRune{
			Runes: []rune{
				rune(0), '\n',
			},
			Fn:      nil,
			Options: dsl.ExpectRuneOptions{Multiple: true, Optional: true},
		})
		s.Expect(dsl.ExpectRune{
			Branches: []dsl.Branch{
				{Rn: '\n', Fn: nil},
			},
			Options: dsl.ExpectRuneOptions{Optional: true},
		})

		s.Match([]dsl.Match{{Literal: "", ID: dsl.TOKEN_UNKNOWN}})
		return s.Exit()
	}

	s.Call(skipWhitespace)
	s.Expect(dsl.ExpectRune{
		Branches: []dsl.Branch{
			{Rn: '-', Fn: nil},
			{Rn: '+', Fn: nil},
			{Rn: '*', Fn: nil},
			{Rn: '/', Fn: nil},
			{Rn: '(', Fn: nil},
			{Rn: ')', Fn: nil},
			{Rn: '\n', Fn: nil},
			{Rn: ':', Fn: assign},
			{Rn: '\'', Fn: comment},
			{Rn: '"', Fn: stringliteral},
			{Rn: rune(0), Fn: eof},
		},
		BranchRanges: []dsl.BranchRange{
			{StartRn: '0', EndRn: '9', Fn: literal},
			{StartRn: 'A', EndRn: 'Z', Fn: variable},
			{StartRn: 'a', EndRn: 'z', Fn: variable},
		},
	})
	s.Match([]dsl.Match{
		{Literal: "-", ID: TOKEN_MINUS},
		{Literal: "+", ID: TOKEN_PLUS},
		{Literal: "*", ID: TOKEN_MULTIPLY},
		{Literal: "/", ID: TOKEN_DIVIDE},
		{Literal: "(", ID: TOKEN_OPEN_PAREN},
		{Literal: ")", ID: TOKEN_CLOSE_PAREN},
		{Literal: "\n", ID: TOKEN_NL},
	})
	return s.Exit()
}

func eof(s *dsl.Scanner) {
	s.Match([]dsl.Match{{Literal: "", ID: TOKEN_EOF}})
}

func skipWhitespace(s *dsl.Scanner) {
	s.Expect(dsl.ExpectRune{
		Branches: []dsl.Branch{
			{Rn: ' ', Fn: nil},
			{Rn: '\t', Fn: nil},
		},
		Options: dsl.ExpectRuneOptions{Optional: true, Multiple: true, Skip: true},
	})
}

func variable(s *dsl.Scanner) {
	s.Expect(dsl.ExpectRune{
		Branches: []dsl.Branch{
			{Rn: '_', Fn: nil},
		},
		BranchRanges: []dsl.BranchRange{
			{StartRn: 'A', EndRn: 'Z', Fn: nil},
			{StartRn: 'a', EndRn: 'z', Fn: nil},
		},
		Options: dsl.ExpectRuneOptions{Multiple: true, Optional: true},
	})
	s.Match([]dsl.Match{{Literal: "", ID: TOKEN_VARIABLE}})
}

// ScanFn -> literal
func literal(s *dsl.Scanner) {
	s.Expect(dsl.ExpectRune{
		BranchRanges: []dsl.BranchRange{
			{StartRn: '0', EndRn: '9', Fn: nil},
		},
		Options: dsl.ExpectRuneOptions{Multiple: true, Optional: true},
	})
	s.Expect(dsl.ExpectRune{
		Branches: []dsl.Branch{
			{Rn: '.', Fn: fraction},
		},
		Options: dsl.ExpectRuneOptions{Optional: true},
	})
	s.Match([]dsl.Match{{Literal: "", ID: TOKEN_LITERAL}})
}

// ScanFn -> literal
func stringliteral(s *dsl.Scanner) {
	s.SkipRune()
	s.ExpectNot(dsl.ExpectNotRune{
		Runes: []rune{
			'"',
		},
		Fn:      nil,
		Options: dsl.ExpectRuneOptions{Multiple: true, Optional: true},
	})

	s.SkipRune()
	s.Match([]dsl.Match{{Literal: "", ID: TOKEN_LITERAL}})
}

// ScanFn -> number -> fraction
func fraction(s *dsl.Scanner) {
	s.Expect(dsl.ExpectRune{
		BranchRanges: []dsl.BranchRange{
			{StartRn: '0', EndRn: '9', Fn: nil},
		},
		Options: dsl.ExpectRuneOptions{Multiple: true},
	})
	s.Match([]dsl.Match{{Literal: "", ID: TOKEN_LITERAL}})
}

// ScanFn -> assign
func assign(s *dsl.Scanner) {
	s.Expect(dsl.ExpectRune{
		Branches: []dsl.Branch{
			{Rn: '=', Fn: nil},
		},
	})
	s.Match([]dsl.Match{{Literal: ":=", ID: TOKEN_ASSIGN}})
}

// ScanFn -> comment
func comment(s *dsl.Scanner) {
	s.SkipRune()
	s.ExpectNot(dsl.ExpectNotRune{
		Runes: []rune{
			rune(0), '\n',
		},
		Fn:      nil,
		Options: dsl.ExpectRuneOptions{Multiple: true, Optional: true},
	})
	s.Match([]dsl.Match{{Literal: "", ID: TOKEN_COMMENT}})
}
