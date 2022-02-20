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

func ExpandMacros(program *ast.Program, env *object.Environment) ast.Node {
	return ast.Modify(program, func(node ast.Node) ast.Node {
		expression, ok := node.(*ast.CallExpression)
		if !ok {
			return node
		}

		macro, ok := isMacroCall(expression, env)
		if !ok {
			return node
		}

		args := quoteArgs(expression)
		evalEnv := extendMacroEnv(macro, args)

		evaluated := Eval(macro.Body, evalEnv)

		quote, ok := evaluated.(*object.Quote)
		if !ok {
			panic("we only support returning AST-nodes from macros")
		}

		return quote.Node
	})
}

func extendMacroEnv(macro *object.Macro, args []*object.Quote) *object.Environment {
	extended := object.NewEnclosedEnvironment(macro.Env)

	for idx, param := range macro.Parameters {
		extended.Set(param.Value, args[idx])
	}

	return extended
}

func quoteArgs(expression *ast.CallExpression) []*object.Quote {
	var args []*object.Quote

	for _, arg := range expression.Arguments {
		args = append(args, &object.Quote{Node: arg})
	}

	return args
}

func isMacroCall(expression *ast.CallExpression, env *object.Environment) (*object.Macro, bool) {
	identifier, ok := expression.Function.(*ast.Identifier)
	if !ok {
		return nil, false
	}

	obj, ok := env.Get(identifier.Value)
	if !ok {
		return nil, false
	}

	macro, ok := obj.(*object.Macro)
	if !ok {
		return nil, false
	}

	return macro, true
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
