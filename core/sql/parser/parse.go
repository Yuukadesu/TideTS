package parser

import (
	"strings"

	"github.com/hanami/tidets/antlr/parser"
	commons "github.com/hanami/tidets/commons/errors"
	"github.com/hanami/tidets/core/sql/ast"
	"github.com/hanami/tidets/core/sql/visitor"

	antlr "github.com/antlr4-go/antlr/v4"
)

// Parse 解析单条 SQL 语句为 AST。
func Parse(sqlText string) (ast.Stmt, error) {
	sqlText = strings.TrimSpace(sqlText)
	if sqlText == "" {
		return nil, commons.ErrSQLTextEmpty
	}

	input := antlr.NewInputStream(sqlText)
	lexer := parser.NewTideSQLLexer(input)
	lexer.RemoveErrorListeners()
	lexErr := &errorListener{}
	lexer.AddErrorListener(lexErr)

	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser.NewTideSQLParser(stream)
	p.RemoveErrorListeners()
	p.AddErrorListener(lexErr)

	tree := p.Statement()
	if lexErr.hasErr {
		return nil, lexErr.err
	}
	stmtCtx, ok := tree.(*parser.StatementContext)
	if !ok || stmtCtx == nil {
		return nil, commons.ErrSQLParseFailed
	}
	return visitor.Build(stmtCtx)
}

type errorListener struct {
	antlr.DefaultErrorListener
	hasErr bool
	err    error
}

func (l *errorListener) SyntaxError(_ antlr.Recognizer, _ interface{}, line, column int, msg string, _ antlr.RecognitionException) {
	l.hasErr = true
	if l.err == nil {
		l.err = commons.Errorf("sql", commons.CodeInvalidArgument, "syntax error at %d:%d: %s", line, column, msg)
	}
}
