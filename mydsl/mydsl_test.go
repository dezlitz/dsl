// Copyright (c) 2015 Des Little <deslittle@gmail.com>
// All rights reserved. Use of this source code is governed by a LGPL v3
// license that can be found in the LICENSE file.

package mydsl

import (
	"bufio"
	"bytes"
	"testing"
	"os"
	"encoding/json"
	"fmt"
    "github.com/Autoblocks/go-dsl"
)

func TestPrintAST(t *testing.T) {
	reader := bytes.NewBufferString(
		`a := 1 * 5 + 7
		b := 3.45 * 44.21 / (4 + a) 'A Simple Expression
		double(a + b)`)
	bufreader := bufio.NewReader(reader)
	ts := NewTokenSet()
    ns := NewNodeSet()
    logfilename := "TestPrintAST.log"
    logfile, err := os.Create(logfilename)
    if err != nil {
		t.Fatal("Error: Could not create log file " + logfilename + ": " + err.Error())
	}
	ast, _ := dsl.ParseAndLog(Parse, Scan, ts, ns, bufreader, logfile)
	logfile.Close()
	
	astjson, _ := json.Marshal(ast)
	fmt.Print(string(astjson))
	if(string(astjson) != `{"root":{"type":"ROOT","tokens":null,"children":[{"type":"ASSIGNMENT","tokens":[{"ID":"VARIABLE","Literal":"a","Line":1,"Position":1}],"children":[{"type":"TERMINAL","tokens":[{"ID":"LITERAL","Literal":"1","Line":1,"Position":6}],"children":null},{"type":"EXPRESSION","tokens":[{"ID":"MULTIPLY","Literal":"*","Line":1,"Position":8}],"children":[{"type":"TERMINAL","tokens":[{"ID":"LITERAL","Literal":"5","Line":1,"Position":10}],"children":null},{"type":"EXPRESSION","tokens":[{"ID":"PLUS","Literal":"+","Line":1,"Position":12}],"children":[{"type":"TERMINAL","tokens":[{"ID":"LITERAL","Literal":"7","Line":1,"Position":14}],"children":null}]}]}]},{"type":"ASSIGNMENT","tokens":[{"ID":"VARIABLE","Literal":"b","Line":2,"Position":3}],"children":[{"type":"TERMINAL","tokens":[{"ID":"LITERAL","Literal":"3.45","Line":2,"Position":8}],"children":null},{"type":"EXPRESSION","tokens":[{"ID":"MULTIPLY","Literal":"*","Line":2,"Position":13}],"children":[{"type":"TERMINAL","tokens":[{"ID":"LITERAL","Literal":"44.21","Line":2,"Position":15}],"children":null},{"type":"EXPRESSION","tokens":[{"ID":"DIVIDE","Literal":"/","Line":2,"Position":21}],"children":[{"type":"EXPRESSION","tokens":[{"ID":"OPEN_PAREN","Literal":"(","Line":2,"Position":23}],"children":[{"type":"TERMINAL","tokens":[{"ID":"LITERAL","Literal":"4","Line":2,"Position":24}],"children":null},{"type":"EXPRESSION","tokens":[{"ID":"PLUS","Literal":"+","Line":2,"Position":26}],"children":[{"type":"TERMINAL","tokens":[{"ID":"VARIABLE","Literal":"a","Line":2,"Position":28}],"children":null}]}]},{"type":"TERMINAL","tokens":[{"ID":"CLOSE_PAREN","Literal":")","Line":2,"Position":29}],"children":null}]}]}]},{"type":"COMMENT","tokens":[{"ID":"COMMENT","Literal":"A Simple Expression","Line":2,"Position":32}],"children":null},{"type":"CALL","tokens":[{"ID":"VARIABLE","Literal":"double","Line":3,"Position":3}],"children":[{"type":"TERMINAL","tokens":[{"ID":"VARIABLE","Literal":"a","Line":3,"Position":10}],"children":null},{"type":"EXPRESSION","tokens":[{"ID":"PLUS","Literal":"+","Line":3,"Position":12}],"children":[{"type":"TERMINAL","tokens":[{"ID":"VARIABLE","Literal":"b","Line":3,"Position":14}],"children":null}]}]}]}}`){
		t.Errorf("JSON malformed.")
	}
}

func TestDSL(t *testing.T) {
	
	reader := bytes.NewBufferString(
`a := 1 * 5 + 7
b := 3.45 * 44.21 / (4 + a) 'A Simple Expression
double(a + b)`)
	bufreader := bufio.NewReader(reader)
    ts := NewTokenSet()
    ns := NewNodeSet()
    logfilename := "TestDSL.log"
    logfile, fileErr := os.Create(logfilename)
    if fileErr != nil {
		t.Fatal("Error: Could not create log file " + logfilename + ": " + fileErr.Error())
	}
	ast, errs := dsl.ParseAndLog(Parse, Scan, ts, ns, bufreader, logfile)
    logfile.Close()
	cases := []dsl.Node {
		{Type: "TERMINAL", Tokens: []dsl.Token{{"LITERAL", "1", 1, 6}}},
		{Type: "TERMINAL", Tokens: []dsl.Token{{"LITERAL", "5", 1, 10}}},
		{Type: "TERMINAL", Tokens: []dsl.Token{{"LITERAL", "7", 1, 14}}},
		{Type: "EXPRESSION", Tokens: []dsl.Token{{"PLUS", "+", 1, 12}}},
		{Type: "EXPRESSION", Tokens: []dsl.Token{{"MULTIPLY", "*", 1, 8}}},
		{Type: "ASSIGNMENT", Tokens: []dsl.Token{{"VARIABLE", "a", 1, 1}}},
		{Type: "TERMINAL", Tokens: []dsl.Token{{"LITERAL", "3.45", 2, 6}}},
		{Type: "TERMINAL", Tokens: []dsl.Token{{"LITERAL", "44.21", 2, 13}}},
		{Type: "TERMINAL", Tokens: []dsl.Token{{"LITERAL", "4", 2, 22}}},
		{Type: "TERMINAL", Tokens: []dsl.Token{{"VARIABLE", "a", 2, 26}}},
		{Type: "EXPRESSION", Tokens: []dsl.Token{{"PLUS", "+", 2, 24}}},
		{Type: "EXPRESSION", Tokens: []dsl.Token{{"OPEN_PAREN", "(", 2, 21}}},
		{Type: "TERMINAL", Tokens: []dsl.Token{{"CLOSE_PAREN", ")", 2, 27}}},
		{Type: "EXPRESSION", Tokens: []dsl.Token{{"DIVIDE", "/", 2, 19}}},
		{Type: "EXPRESSION", Tokens: []dsl.Token{{"MULTIPLY", "*", 2, 11}}},
		{Type: "ASSIGNMENT", Tokens: []dsl.Token{{"VARIABLE", "b", 2, 1}}},
		{Type: "COMMENT", Tokens: []dsl.Token{{"COMMENT", "A Simple Expression", 2, 30}}},
		{Type: "TERMINAL", Tokens: []dsl.Token{{"VARIABLE", "a", 3, 8}}},
		{Type: "TERMINAL", Tokens: []dsl.Token{{"VARIABLE", "b", 3, 12}}},
		{Type: "EXPRESSION", Tokens: []dsl.Token{{"PLUS", "+", 3, 10}}},
		{Type: "CALL", Tokens: []dsl.Token{{"VARIABLE", "double", 3, 1}}},
		{Type: "ROOT", Tokens: []dsl.Token{{"", "", 0, 0}}},
	}
	count := 0
	ast.Print()
	fmt.Println()
	ast.Inspect(func(node *dsl.Node)(){
		if count > len(cases) - 1{
			t.Fatalf("Too many nodes.")
		}
		if cases[count].Type != node.Type{
			t.Errorf("Line: %v:%v Node: \"%v\" Wanted node type %v, found %v", cases[count].Tokens[0].Line, cases[count].Tokens[0].Position, 
                node.Type, cases[count].Type, node.Type)
		}
        for i, token := range node.Tokens{
           if cases[count].Tokens[i].ID != token.ID{
			 t.Errorf("Line: %v:%v Token: \"%v\" Wanted token ID %v, found %v", cases[count].Tokens[i].Line, cases[count].Tokens[i].Position, 
                token.Literal, cases[count].Tokens[i].ID, token.ID)
		   }
           if cases[count].Tokens[i].Literal != token.Literal{
			 t.Errorf("Line: %v:%v ID: \"%v\" Wanted token literal \"%v\", found \"%v\"", cases[count].Tokens[i].Line, cases[count].Tokens[i].Position, 
                token.ID, cases[count].Tokens[i].Literal, token.Literal)
		   }
           if cases[count].Tokens[i].Line != token.Line{
			 t.Errorf("Line: %v:%v Token: \"%v\" Wanted token line %v, found %v", cases[count].Tokens[i].Line, cases[count].Tokens[i].Position, 
                token.Literal, cases[count].Tokens[i].Line, token.Line)
		   }
           if cases[count].Tokens[i].Position != token.Position{
			 t.Errorf("Line: %v:%v Token: \"%v\" Wanted token position %v, found %v", cases[count].Tokens[i].Line, cases[count].Tokens[i].Position, 
                token.Literal, cases[count].Tokens[i].Position, token.Position)
		   } 
        }
		count++
	})
	if count != len(cases){
		t.Errorf("Not enough nodes.")
	}

	if errs != nil {
		t.Fail()
		for _, err := range errs {
			t.Error(err.String())
		}
	}

}

func TestTokenExpectedButNotFoundError(t *testing.T) {
	reader := bytes.NewBufferString(
		`a error := 1 * 5 + 7
		b := 3.45 * 44.21 / (4 + a) 'A Simple Expression
		double(a + b)`)
	bufreader := bufio.NewReader(reader)
	ts := NewTokenSet()
	ns := NewNodeSet()
	logfilename := "TestTokenExpectedButNotFoundError.log"
    logfile, fileErr := os.Create(logfilename)
    if fileErr != nil {
		t.Fatal("Error: Could not create log file " + logfilename + ": " + fileErr.Error())
	}
	_, errs := dsl.ParseAndLog(Parse, Scan, ts, ns, bufreader, logfile)

	if len(errs) != 1 {
		t.Fail()
		t.Error("Should only report exactly 1 error")
	}
	err := errs[0];
	if err.Code != dsl.TOKEN_EXPECTED_NOT_FOUND {
		t.Fail()
		t.Errorf("Expected error code 'Token expected but not found'. Found error: '%v", err.Error)
	}
	if err.Line != 1 {
		t.Fail()
		t.Errorf("Expected error line 1. Found line: %v", err.Line)
	}
	if err.StartPosition != 2 {
		t.Fail()
		t.Errorf("Expected error start position 2. Found position: %v", err.StartPosition)
	}
	if err.EndPosition != 7 {
		t.Fail()
		t.Errorf("Expected error end position 7. Found position: %v", err.EndPosition)
	}

}

func TestRuneExpectedButNotFoundError(t *testing.T) {
	reader := bytes.NewBufferString(
		`_ := 1 * 5 + 7
		b := 3.45 * 44.21 / (4 + a) 'A Simple Expression
		double(a + b)`)
	bufreader := bufio.NewReader(reader)
	ts := NewTokenSet()
	ns := NewNodeSet()
	logfilename := "TestRuneExpectedButNotFoundError.log"
    logfile, fileErr := os.Create(logfilename)
    if fileErr != nil {
		t.Fatal("Error: Could not create log file " + logfilename + ": " + fileErr.Error())
	}
	_, errs := dsl.ParseAndLog(Parse, Scan, ts, ns, bufreader, logfile)

	if len(errs) != 1 {
		t.Fail()
		t.Error("Should only exactly 1 error")
	}
	err := errs[0];
	if err.Code != dsl.RUNE_EXPECTED_NOT_FOUND {
		t.Fail()
		t.Errorf("Expected error code 'Rune expected but not found'. Found error: '%v", err.Error)
	}
	if err.Line != 1 {
		t.Fail()
		t.Errorf("Expected error line 1. Found line: %v", err.Line)
	}
	if err.StartPosition != 0 {
		t.Fail()
		t.Errorf("Expected error start position 0. Found position: %v", err.StartPosition)
	}
	if err.EndPosition != 1 {
		t.Fail()
		t.Errorf("Expected error end position 1. Found position: %v", err.EndPosition)
	}

}



