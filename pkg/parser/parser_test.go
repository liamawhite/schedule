package parser_test

import (
	"reflect"
	"testing"

	"github.com/liamawhite/schedule/pkg/ast"
	"github.com/liamawhite/schedule/pkg/lexer"
	"github.com/liamawhite/schedule/pkg/parser"
)

// in 6 days
// in 1 day
// in 2 weeks
// in 1 week
// in 3 months
// in 1 month
// in 1 year
// in 2 years


func TestInStatement(t *testing.T) {
    tests := []struct {
        input string
        want *ast.InStatement
    }{
        {
            input: "in 6 days",
            want: &ast.InStatement{
                Token: token.Token{Type: token.IN, Literal: "in"},
            },
        },

    }

    for _, tt := range tests {
        p := parser.New(lexer.New(tt.input))
        got := p.Parse()

        if !reflect.DeepEqual(tt.want, got) {
            t.Errorf("got %+v, want %+v", got, tt.want)
        }
    }
}
