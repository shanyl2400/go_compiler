package object

import (
	"bytes"
	"fmt"
	"shanyl2400/go_compiler/ast"
	"strings"
)

const (
	INTEGER_OBJ = "INTEGER"
	BOOLEAN_OBJ = "BOOLEAN"
	NULL_OBJ    = "NULL"

	FUNCTION_OBJ     = "FUNCTION"
	RETURN_VALUE_OBJ = "RETURN_VALUE"
	ERROR_OBJ        = "ERROR"
)

type ObjectType string
type Object interface {
	Type() ObjectType
	Inspect() string
}

type Integer struct {
	Value int64
}

func (i *Integer) Inspect() string {
	return fmt.Sprintf("%v", i.Value)
}

func (i *Integer) Type() ObjectType {
	return INTEGER_OBJ
}

type Boolean struct {
	Value bool
}

func (b *Boolean) Inspect() string {
	return fmt.Sprintf("%v", b.Value)
}

func (b *Boolean) Type() ObjectType {
	return BOOLEAN_OBJ
}

type ReturnValue struct {
	Value Object
}

func (r *ReturnValue) Inspect() string {
	return fmt.Sprintf("%v", r.Value.Inspect())
}

func (r *ReturnValue) Type() ObjectType {
	return RETURN_VALUE_OBJ
}

type Null struct {
}

func (n *Null) Inspect() string {
	return "null"
}

func (n *Null) Type() ObjectType {
	return NULL_OBJ
}

type Function struct {
	Parameters []*ast.Identifier
	Body       *ast.BlockStatement
	Env        *Environment
}

func (f *Function) Inspect() string {
	var out bytes.Buffer

	params := []string{}
	for _, p := range f.Parameters {
		params = append(params, p.String())
	}

	out.WriteString("fn")
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(") {\n")
	out.WriteString(f.Body.String())
	out.WriteString("\n}")

	return out.String()
}

func (f *Function) Type() ObjectType {
	return FUNCTION_OBJ
}

type Error struct {
	Message string
}

func (e *Error) Inspect() string {
	return "ERROR: " + e.Message
}

func (e *Error) Type() ObjectType {
	return ERROR_OBJ
}
