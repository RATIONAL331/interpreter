package evaluator

import (
	"interpreter/ast"
	"interpreter/object"
)

func DefineMacros(program *ast.Program, env *object.Environment) {
	var definitions []int

	for i, statement := range program.Statements {
		if isMacroDefinition(statement) {
			addMacro(statement, env)
			definitions = append(definitions, i)
		}
	}

	for i := len(definitions) - 1; i >= 0; i = i - 1 {
		definitionIndex := definitions[i]
		program.Statements = append(program.Statements[:definitionIndex], program.Statements[definitionIndex+1:]...)
	}
}

func addMacro(statement ast.Statement, env *object.Environment) {
	letStatement := statement.(*ast.LetStatement)
	literal := letStatement.Value.(*ast.MacroLiteral)

	macro := &object.Macro{
		Parameters: literal.Parameters,
		Env:        env,
		Body:       literal.Body,
	}

	env.Set(letStatement.Name.Value, macro)
}

func isMacroDefinition(statement ast.Statement) bool {
	letStatement, ok := statement.(*ast.LetStatement)
	if !ok {
		return false
	}

	_, ok = letStatement.Value.(*ast.MacroLiteral)
	if !ok {
		return false
	}

	return true
}
