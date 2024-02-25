package ast

import "github.com/liamawhite/schedule/pkg/token"

type Node interface {
    TokenLiteral() string
}

type Statement interface {
    Node
}

type Expression interface {
    Node
}



// Due Dates
// In x days
type InStatement struct {
    token token.Token
    value Expression
}

func (i *InStatement) TokenLiteral() string {
    return i.token.Literal
}

// Next monday
type NextStatement struct {
    token token.Token
    value Expression
}

func (n *NextStatement) TokenLiteral() string {
    return n.token.Literal
}

// Monday, feb, etc.
type DateStatement struct {
    token token.Token
    value Expression
}

func (d *DateStatement) TokenLiteral() string {
    return d.token.Literal
}

// Recurrences
// Every x days
type EveryStatement struct {
    token token.Token
    value Expression
}

func (e *EveryStatement) TokenLiteral() string {
    return e.token.Literal
}



