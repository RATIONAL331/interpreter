package ast

import (
	"reflect"
	"testing"
)

func TestModify(t *testing.T) {
	oneFunc := func() Expression { return &IntegerLiteral{Value: 1} }
	twoFunc := func() Expression { return &IntegerLiteral{Value: 2} }

	turnOneIntoTwo := func(node Node) Node {
		integer, ok := node.(*IntegerLiteral)
		if !ok {
			return node
		}

		if integer.Value != 1 {
			return node
		}

		integer.Value = 2
		return integer
	}

	tests := []struct {
		input    Node
		expected Node
	}{
		{oneFunc(), twoFunc()},
		{
			&Program{Statements: []Statement{&ExpressionStatement{Expression: oneFunc()}}},
			&Program{Statements: []Statement{&ExpressionStatement{Expression: twoFunc()}}},
		},
		{
			&InfixExpression{Left: oneFunc(), Operator: "+", Right: twoFunc()},
			&InfixExpression{Left: twoFunc(), Operator: "+", Right: twoFunc()},
		},
		{
			&InfixExpression{Left: twoFunc(), Operator: "+", Right: oneFunc()},
			&InfixExpression{Left: twoFunc(), Operator: "+", Right: twoFunc()},
		},
		{
			&PrefixExpression{Operator: "-", Right: oneFunc()},
			&PrefixExpression{Operator: "-", Right: twoFunc()},
		},
		{
			&IndexExpression{Left: oneFunc(), Index: oneFunc()},
			&IndexExpression{Left: twoFunc(), Index: twoFunc()},
		},
		{
			&IfExpression{
				Condition:   oneFunc(),
				Consequence: &BlockStatement{Statements: []Statement{&ExpressionStatement{Expression: oneFunc()}}},
				Alternative: &BlockStatement{Statements: []Statement{&ExpressionStatement{Expression: oneFunc()}}},
			},
			&IfExpression{
				Condition:   twoFunc(),
				Consequence: &BlockStatement{Statements: []Statement{&ExpressionStatement{Expression: twoFunc()}}},
				Alternative: &BlockStatement{Statements: []Statement{&ExpressionStatement{Expression: twoFunc()}}},
			},
		},
		{
			&ReturnStatement{ReturnValue: oneFunc()},
			&ReturnStatement{ReturnValue: twoFunc()},
		},
		{
			&LetStatement{Value: oneFunc()},
			&LetStatement{Value: twoFunc()},
		},
		{
			&FunctionLiteral{
				Parameters: []*Identifier{},
				Body:       &BlockStatement{Statements: []Statement{&ExpressionStatement{Expression: oneFunc()}}},
			},
			&FunctionLiteral{
				Parameters: []*Identifier{},
				Body:       &BlockStatement{Statements: []Statement{&ExpressionStatement{Expression: twoFunc()}}},
			},
		},
		{
			&ArrayLiteral{Elements: []Expression{oneFunc(), oneFunc()}},
			&ArrayLiteral{Elements: []Expression{twoFunc(), twoFunc()}},
		},
	}

	for _, test := range tests {
		modified := Modify(test.input, turnOneIntoTwo)

		equal := reflect.DeepEqual(modified, test.expected)
		if !equal {
			t.Errorf("not equal. got=%#v, want=%#v", modified, test.expected)
		}
	}

	hashLiteral := &HashLiteral{
		Pairs: map[Expression]Expression{
			oneFunc(): oneFunc(),
			oneFunc(): oneFunc(),
		},
	}

	Modify(hashLiteral, turnOneIntoTwo)

	for key, val := range hashLiteral.Pairs {
		key, _ := key.(*IntegerLiteral)
		if key.Value != 2 {
			t.Errorf("key is not %d, got=%d", 2, key.Value)
		}

		val, _ := val.(*IntegerLiteral)
		if val.Value != 2 {
			t.Errorf("value is not %d, got=%d", 2, val.Value)
		}
	}
}
