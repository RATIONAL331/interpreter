package evaluator

import (
	"interpreter/object"
	"testing"
)

func TestQuote(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{`quote(5)`, `5`},
		{`quote(foobar)`, `foobar`},
		{`quote(foobar + barfoo)`, `(foobar + barfoo)`},
	}

	for _, test := range tests {
		eval := testEval(test.input)
		quote, ok := eval.(*object.Quote)
		if !ok {
			t.Fatalf("expected *object.Quote. got=%T (+%v)", eval, eval)
		}

		if quote.Node == nil {
			t.Fatalf("quote.Node is nil")
		}

		if quote.Node.String() != test.expected {
			t.Errorf("not equal. got=%q, want=%q", quote.Node.String(), test.expected)
		}
	}
}
