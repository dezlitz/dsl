package mydsl

import (
	"github.com/madlitz/go-dsl"
)

func Scan(s *dsl.Scanner) dsl.Token {
	if recover {
		s.Expect(dsl.ExpectRune{
			Branches: []dsl.Branch{
				{rune(0), nil},
				{'\n', nil}},
			Options: dsl.ScanOptions{Multiple: true, Invert: true, Optional: true}})
		s.Match([]dsl.Match{{"", "UNKNOWN"}})
		return s.Exit()
	}

	s.Expect(dsl.ExpectRune{
		Branches: []dsl.Branch{
			{' ', whitespace},
			{'\t', whitespace}},
		Options: dsl.ScanOptions{Optional: true}})
	s.Expect(dsl.ExpectRune{
		Branches: []dsl.Branch{
			{'-', nil},
			{'+', nil},
			{'*', nil},
			{'/', nil},
			{'(', nil},
			{')', nil},
			{'\n', nil},
			{':', assign},
			{'\'', comment},
			{'"', stringliteral},
			{rune(0), eof}},
		BranchRanges: []dsl.BranchRange{
			{'0', '9', literal},
			{'A', 'Z', variable},
			{'a', 'z', variable}}})
	s.Match(([]dsl.Match{
		{"-", "MINUS"},
		{"+", "PLUS"},
		{"*", "MULTIPLY"},
		{"/", "DIVIDE"},
		{"(", "OPEN_PAREN"},
		{")", "CLOSE_PAREN"},
		{"\n", "NL"}}))
	s.Expect(dsl.ExpectRune{
		Branches: []dsl.Branch{
			{' ', nil},
			{'\t', nil}},
		Options: dsl.ScanOptions{Multiple: true, Optional: true}})
	return s.Exit()
}

func eof(s *dsl.Scanner) {
	s.Match([]dsl.Match{{"", "EOF"}})
}

func whitespace(s *dsl.Scanner) {
	s.Expect(dsl.ExpectRune{
		Branches: []dsl.Branch{
			{' ', nil},
			{'\t', nil}},
		Options: dsl.ScanOptions{Optional: true, Multiple: true}})
	s.Match([]dsl.Match{{"", "WS"}})
}

func variable(s *dsl.Scanner) {
	s.Expect(dsl.ExpectRune{
		Branches: []dsl.Branch{
			{'_', nil}},
		BranchRanges: []dsl.BranchRange{
			{'A', 'Z', nil},
			{'a', 'z', nil}},
		Options: dsl.ScanOptions{Multiple: true, Optional: true}})
	s.Match([]dsl.Match{{"", "VARIABLE"}})
}

// ScanFn -> literal
func literal(s *dsl.Scanner) {
	s.Expect(dsl.ExpectRune{
		BranchRanges: []dsl.BranchRange{
			{'0', '9', nil}},
		Options: dsl.ScanOptions{Multiple: true, Optional: true}})
	s.Expect(dsl.ExpectRune{
		Branches: []dsl.Branch{
			{'.', fraction}},
		Options: dsl.ScanOptions{Optional: true}})
	s.Match([]dsl.Match{{"", "LITERAL"}})
}

// ScanFn -> literal
func stringliteral(s *dsl.Scanner) {
	s.SkipRune()
	s.Expect(dsl.ExpectRune{
		Branches: []dsl.Branch{
			{'"', nil}},
		Options: dsl.ScanOptions{Multiple: true, Invert: true, Optional: true}})
	s.SkipRune()
	s.Match([]dsl.Match{{"", "LITERAL"}})
}

// ScanFn -> number -> fraction
func fraction(s *dsl.Scanner) {
	s.Expect(dsl.ExpectRune{
		BranchRanges: []dsl.BranchRange{
			{'0', '9', nil}},
		Options: dsl.ScanOptions{Multiple: true}})
	s.Match([]dsl.Match{{"", "LITERAL"}})
}

// ScanFn -> assign
func assign(s *dsl.Scanner) {
	s.Expect(dsl.ExpectRune{
		Branches: []dsl.Branch{
			{'=', nil}}})
	s.Match([]dsl.Match{{":=", "ASSIGN"}})
}

// ScanFn -> comment
func comment(s *dsl.Scanner) {
	s.SkipRune()
	s.Expect(dsl.ExpectRune{
		Branches: []dsl.Branch{
			{rune(0), nil},
			{'\n', nil}},
		Options: dsl.ScanOptions{Multiple: true, Invert: true, Optional: true}})
	s.Match([]dsl.Match{{"", "COMMENT"}})
}
