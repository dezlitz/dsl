// Copyright (c) 2024 Dez Little <deslittle@gmail.com>
// All rights reserved. Use of this source code is governed by a LGPL v3
// license that can be found in the LICENSE file.

package dsl

import (
	"bytes"
	"fmt"
	"strings"
)

type ErrorCode int

const (
	FILE_NOT_FOUND ErrorCode = iota // 0
	COULD_NOT_CREATE_FILE
	TOKEN_EXPECTED_NOT_FOUND
	EXPECTED_TOKEN_NOT_IN_TOKENSET
	SCANNED_TOKEN_NOT_IN_TOKENSET
	RUNE_EXPECTED_NOT_FOUND
	NODE_NOT_IN_NODESET
)

// Errors contain the error text, the line and positions the error occurred on, and
// a string containing the input text from that line
type Error struct {
	Code          ErrorCode
	Error         error
	LineString    string
	StartLine     int
	StartPosition int
	EndLine       int
	EndPosition   int
}

func (e *Error) String() string {
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("\nError Line:%v %v\n", e.StartLine, e.Error))
	buf.WriteString(e.LineString + "\n")
	for i := 1; i < e.StartPosition-1; i++ {
		if e.LineString[i] == '\t' {
			buf.WriteString("\t")
		} else {
			buf.WriteString(" ")
		}
	}
	buf.WriteString("^")
	if e.StartLine != e.EndLine {
		for i := e.StartPosition; i < strings.Count(e.LineString, ""); i++ {
			buf.WriteString("-")
		}
	} else {
		for i := e.StartPosition; i < e.EndPosition-1; i++ {
			buf.WriteString("-")
		}
	}

	buf.WriteString("^\n")
	return buf.String()
}
