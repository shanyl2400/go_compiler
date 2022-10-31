package compiler

import (
	"shanyl2400/go_compiler/ast"
	"shanyl2400/go_compiler/code"
	"testing"

	"github.com/stretchr/testify/assert"
)

type complierTestCase struct {
	input                string
	expectedConstant     []interface{}
	expectedInstructions []code.Instructions
}

func TestIntegerArithmetic(t *testing.T) {
	tests := []complierTestCase{
		{
			input:            "1 + 2",
			expectedConstant: []interface{}{1, 2},
			expectedInstructions: []code.Instructions{
				code.Make(code.OpConstant, 1),
				code.Make(code.OpConstant, 2),
			},
		},
	}
	runCompilerTests(t, tests)
}

func runCompilerTests(t *testing.T, tests []complierTestCase) {
	t.Helper()

	for _, tt := range tests {
		program := parse(tt.input)
		compiler := New()
		err := compiler.Compile(program)
		assert.NoError(t, err)

		byteCode := compiler.ByteCode()
		assert.Equal(t, tt.expectedConstant, byteCode.Constants)
		assert.Equal(t, tt.expectedInstructions, byteCode.Instructions)
	}
}

func parse(input string) ast.Node {
	return nil
}
