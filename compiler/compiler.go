package compiler

import (
	"shanyl2400/go_compiler/ast"
	"shanyl2400/go_compiler/code"
	"shanyl2400/go_compiler/object"
)

type Compiler struct {
	instructions code.Instructions
	constants    []object.Object
}

func (c *Compiler) Compile(node ast.Node) error {
	return nil
}

func (c *Compiler) ByteCode() *ByteCode {
	return &ByteCode{
		Instructions: c.instructions,
		Constants:    c.constants,
	}
}

func New() *Compiler {
	return &Compiler{
		instructions: code.Instructions{},
		constants:    []object.Object{},
	}
}

type ByteCode struct {
	Instructions code.Instructions
	Constants    []object.Object
}
