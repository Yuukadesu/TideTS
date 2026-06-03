// Code generated from TideSQL.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // TideSQL
import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type TideSQLParser struct {
	*antlr.BaseParser
}

var TideSQLParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func tidesqlParserInit() {
	staticData := &TideSQLParserStaticData
	staticData.LiteralNames = []string{
		"", "", "", "", "", "", "", "", "", "", "'>='", "'<='", "'>'", "'<'",
		"'='", "'('", "')'", "','", "'.'", "';'",
	}
	staticData.SymbolicNames = []string{
		"", "INSERT", "INTO", "SELECT", "FROM", "WHERE", "AND", "VALUES", "LIMIT",
		"TIME", "GTE", "LTE", "GT", "LT", "EQ", "LPAREN", "RPAREN", "COMMA",
		"DOT", "SEMI", "BOOLEAN", "INTEGER", "FLOAT", "STRING", "IDENTIFIER",
		"WS",
	}
	staticData.RuleNames = []string{
		"statement", "insertStmt", "selectStmt", "whereClause", "timePredicate",
		"cmpOp", "limitClause", "path", "measurement", "timestamp", "value",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 25, 82, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 1, 0, 1, 0, 3, 0, 25, 8, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 3,
		2, 45, 8, 2, 1, 2, 3, 2, 48, 8, 2, 1, 3, 1, 3, 1, 3, 1, 3, 5, 3, 54, 8,
		3, 10, 3, 12, 3, 57, 9, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 6, 1,
		6, 1, 6, 1, 7, 1, 7, 1, 7, 5, 7, 71, 8, 7, 10, 7, 12, 7, 74, 9, 7, 1, 8,
		1, 8, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 0, 0, 11, 0, 2, 4, 6, 8, 10, 12,
		14, 16, 18, 20, 0, 2, 1, 0, 10, 14, 1, 0, 20, 23, 75, 0, 24, 1, 0, 0, 0,
		2, 26, 1, 0, 0, 0, 4, 39, 1, 0, 0, 0, 6, 49, 1, 0, 0, 0, 8, 58, 1, 0, 0,
		0, 10, 62, 1, 0, 0, 0, 12, 64, 1, 0, 0, 0, 14, 67, 1, 0, 0, 0, 16, 75,
		1, 0, 0, 0, 18, 77, 1, 0, 0, 0, 20, 79, 1, 0, 0, 0, 22, 25, 3, 2, 1, 0,
		23, 25, 3, 4, 2, 0, 24, 22, 1, 0, 0, 0, 24, 23, 1, 0, 0, 0, 25, 1, 1, 0,
		0, 0, 26, 27, 5, 1, 0, 0, 27, 28, 5, 2, 0, 0, 28, 29, 3, 14, 7, 0, 29,
		30, 5, 15, 0, 0, 30, 31, 3, 16, 8, 0, 31, 32, 5, 16, 0, 0, 32, 33, 5, 7,
		0, 0, 33, 34, 5, 15, 0, 0, 34, 35, 3, 18, 9, 0, 35, 36, 5, 17, 0, 0, 36,
		37, 3, 20, 10, 0, 37, 38, 5, 16, 0, 0, 38, 3, 1, 0, 0, 0, 39, 40, 5, 3,
		0, 0, 40, 41, 3, 16, 8, 0, 41, 42, 5, 4, 0, 0, 42, 44, 3, 14, 7, 0, 43,
		45, 3, 6, 3, 0, 44, 43, 1, 0, 0, 0, 44, 45, 1, 0, 0, 0, 45, 47, 1, 0, 0,
		0, 46, 48, 3, 12, 6, 0, 47, 46, 1, 0, 0, 0, 47, 48, 1, 0, 0, 0, 48, 5,
		1, 0, 0, 0, 49, 50, 5, 5, 0, 0, 50, 55, 3, 8, 4, 0, 51, 52, 5, 6, 0, 0,
		52, 54, 3, 8, 4, 0, 53, 51, 1, 0, 0, 0, 54, 57, 1, 0, 0, 0, 55, 53, 1,
		0, 0, 0, 55, 56, 1, 0, 0, 0, 56, 7, 1, 0, 0, 0, 57, 55, 1, 0, 0, 0, 58,
		59, 5, 9, 0, 0, 59, 60, 3, 10, 5, 0, 60, 61, 5, 21, 0, 0, 61, 9, 1, 0,
		0, 0, 62, 63, 7, 0, 0, 0, 63, 11, 1, 0, 0, 0, 64, 65, 5, 8, 0, 0, 65, 66,
		5, 21, 0, 0, 66, 13, 1, 0, 0, 0, 67, 72, 5, 24, 0, 0, 68, 69, 5, 18, 0,
		0, 69, 71, 5, 24, 0, 0, 70, 68, 1, 0, 0, 0, 71, 74, 1, 0, 0, 0, 72, 70,
		1, 0, 0, 0, 72, 73, 1, 0, 0, 0, 73, 15, 1, 0, 0, 0, 74, 72, 1, 0, 0, 0,
		75, 76, 5, 24, 0, 0, 76, 17, 1, 0, 0, 0, 77, 78, 5, 21, 0, 0, 78, 19, 1,
		0, 0, 0, 79, 80, 7, 1, 0, 0, 80, 21, 1, 0, 0, 0, 5, 24, 44, 47, 55, 72,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// TideSQLParserInit initializes any static state used to implement TideSQLParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewTideSQLParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func TideSQLParserInit() {
	staticData := &TideSQLParserStaticData
	staticData.once.Do(tidesqlParserInit)
}

// NewTideSQLParser produces a new parser instance for the optional input antlr.TokenStream.
func NewTideSQLParser(input antlr.TokenStream) *TideSQLParser {
	TideSQLParserInit()
	this := new(TideSQLParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &TideSQLParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "TideSQL.g4"

	return this
}

// TideSQLParser tokens.
const (
	TideSQLParserEOF        = antlr.TokenEOF
	TideSQLParserINSERT     = 1
	TideSQLParserINTO       = 2
	TideSQLParserSELECT     = 3
	TideSQLParserFROM       = 4
	TideSQLParserWHERE      = 5
	TideSQLParserAND        = 6
	TideSQLParserVALUES     = 7
	TideSQLParserLIMIT      = 8
	TideSQLParserTIME       = 9
	TideSQLParserGTE        = 10
	TideSQLParserLTE        = 11
	TideSQLParserGT         = 12
	TideSQLParserLT         = 13
	TideSQLParserEQ         = 14
	TideSQLParserLPAREN     = 15
	TideSQLParserRPAREN     = 16
	TideSQLParserCOMMA      = 17
	TideSQLParserDOT        = 18
	TideSQLParserSEMI       = 19
	TideSQLParserBOOLEAN    = 20
	TideSQLParserINTEGER    = 21
	TideSQLParserFLOAT      = 22
	TideSQLParserSTRING     = 23
	TideSQLParserIDENTIFIER = 24
	TideSQLParserWS         = 25
)

// TideSQLParser rules.
const (
	TideSQLParserRULE_statement     = 0
	TideSQLParserRULE_insertStmt    = 1
	TideSQLParserRULE_selectStmt    = 2
	TideSQLParserRULE_whereClause   = 3
	TideSQLParserRULE_timePredicate = 4
	TideSQLParserRULE_cmpOp         = 5
	TideSQLParserRULE_limitClause   = 6
	TideSQLParserRULE_path          = 7
	TideSQLParserRULE_measurement   = 8
	TideSQLParserRULE_timestamp     = 9
	TideSQLParserRULE_value         = 10
)

// IStatementContext is an interface to support dynamic dispatch.
type IStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	InsertStmt() IInsertStmtContext
	SelectStmt() ISelectStmtContext

	// IsStatementContext differentiates from other interfaces.
	IsStatementContext()
}

type StatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementContext() *StatementContext {
	var p = new(StatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TideSQLParserRULE_statement
	return p
}

func InitEmptyStatementContext(p *StatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TideSQLParserRULE_statement
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TideSQLParserRULE_statement

	return p
}

func (s *StatementContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementContext) InsertStmt() IInsertStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInsertStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInsertStmtContext)
}

func (s *StatementContext) SelectStmt() ISelectStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISelectStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISelectStmtContext)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TideSQLListener); ok {
		listenerT.EnterStatement(s)
	}
}

func (s *StatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TideSQLListener); ok {
		listenerT.ExitStatement(s)
	}
}

func (s *StatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TideSQLVisitor:
		return t.VisitStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TideSQLParser) Statement() (localctx IStatementContext) {
	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, TideSQLParserRULE_statement)
	p.SetState(24)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case TideSQLParserINSERT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(22)
			p.InsertStmt()
		}

	case TideSQLParserSELECT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(23)
			p.SelectStmt()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IInsertStmtContext is an interface to support dynamic dispatch.
type IInsertStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	INSERT() antlr.TerminalNode
	INTO() antlr.TerminalNode
	Path() IPathContext
	AllLPAREN() []antlr.TerminalNode
	LPAREN(i int) antlr.TerminalNode
	Measurement() IMeasurementContext
	AllRPAREN() []antlr.TerminalNode
	RPAREN(i int) antlr.TerminalNode
	VALUES() antlr.TerminalNode
	Timestamp() ITimestampContext
	COMMA() antlr.TerminalNode
	Value() IValueContext

	// IsInsertStmtContext differentiates from other interfaces.
	IsInsertStmtContext()
}

type InsertStmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInsertStmtContext() *InsertStmtContext {
	var p = new(InsertStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TideSQLParserRULE_insertStmt
	return p
}

func InitEmptyInsertStmtContext(p *InsertStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TideSQLParserRULE_insertStmt
}

func (*InsertStmtContext) IsInsertStmtContext() {}

func NewInsertStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InsertStmtContext {
	var p = new(InsertStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TideSQLParserRULE_insertStmt

	return p
}

func (s *InsertStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *InsertStmtContext) INSERT() antlr.TerminalNode {
	return s.GetToken(TideSQLParserINSERT, 0)
}

func (s *InsertStmtContext) INTO() antlr.TerminalNode {
	return s.GetToken(TideSQLParserINTO, 0)
}

func (s *InsertStmtContext) Path() IPathContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPathContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPathContext)
}

func (s *InsertStmtContext) AllLPAREN() []antlr.TerminalNode {
	return s.GetTokens(TideSQLParserLPAREN)
}

func (s *InsertStmtContext) LPAREN(i int) antlr.TerminalNode {
	return s.GetToken(TideSQLParserLPAREN, i)
}

func (s *InsertStmtContext) Measurement() IMeasurementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMeasurementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMeasurementContext)
}

func (s *InsertStmtContext) AllRPAREN() []antlr.TerminalNode {
	return s.GetTokens(TideSQLParserRPAREN)
}

func (s *InsertStmtContext) RPAREN(i int) antlr.TerminalNode {
	return s.GetToken(TideSQLParserRPAREN, i)
}

func (s *InsertStmtContext) VALUES() antlr.TerminalNode {
	return s.GetToken(TideSQLParserVALUES, 0)
}

func (s *InsertStmtContext) Timestamp() ITimestampContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITimestampContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITimestampContext)
}

func (s *InsertStmtContext) COMMA() antlr.TerminalNode {
	return s.GetToken(TideSQLParserCOMMA, 0)
}

func (s *InsertStmtContext) Value() IValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *InsertStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InsertStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InsertStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TideSQLListener); ok {
		listenerT.EnterInsertStmt(s)
	}
}

func (s *InsertStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TideSQLListener); ok {
		listenerT.ExitInsertStmt(s)
	}
}

func (s *InsertStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TideSQLVisitor:
		return t.VisitInsertStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TideSQLParser) InsertStmt() (localctx IInsertStmtContext) {
	localctx = NewInsertStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, TideSQLParserRULE_insertStmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(26)
		p.Match(TideSQLParserINSERT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(27)
		p.Match(TideSQLParserINTO)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(28)
		p.Path()
	}
	{
		p.SetState(29)
		p.Match(TideSQLParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(30)
		p.Measurement()
	}
	{
		p.SetState(31)
		p.Match(TideSQLParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(32)
		p.Match(TideSQLParserVALUES)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(33)
		p.Match(TideSQLParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(34)
		p.Timestamp()
	}
	{
		p.SetState(35)
		p.Match(TideSQLParserCOMMA)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(36)
		p.Value()
	}
	{
		p.SetState(37)
		p.Match(TideSQLParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISelectStmtContext is an interface to support dynamic dispatch.
type ISelectStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SELECT() antlr.TerminalNode
	Measurement() IMeasurementContext
	FROM() antlr.TerminalNode
	Path() IPathContext
	WhereClause() IWhereClauseContext
	LimitClause() ILimitClauseContext

	// IsSelectStmtContext differentiates from other interfaces.
	IsSelectStmtContext()
}

type SelectStmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySelectStmtContext() *SelectStmtContext {
	var p = new(SelectStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TideSQLParserRULE_selectStmt
	return p
}

func InitEmptySelectStmtContext(p *SelectStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TideSQLParserRULE_selectStmt
}

func (*SelectStmtContext) IsSelectStmtContext() {}

func NewSelectStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SelectStmtContext {
	var p = new(SelectStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TideSQLParserRULE_selectStmt

	return p
}

func (s *SelectStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *SelectStmtContext) SELECT() antlr.TerminalNode {
	return s.GetToken(TideSQLParserSELECT, 0)
}

func (s *SelectStmtContext) Measurement() IMeasurementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMeasurementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMeasurementContext)
}

func (s *SelectStmtContext) FROM() antlr.TerminalNode {
	return s.GetToken(TideSQLParserFROM, 0)
}

func (s *SelectStmtContext) Path() IPathContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPathContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPathContext)
}

func (s *SelectStmtContext) WhereClause() IWhereClauseContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IWhereClauseContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IWhereClauseContext)
}

func (s *SelectStmtContext) LimitClause() ILimitClauseContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILimitClauseContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILimitClauseContext)
}

func (s *SelectStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SelectStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SelectStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TideSQLListener); ok {
		listenerT.EnterSelectStmt(s)
	}
}

func (s *SelectStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TideSQLListener); ok {
		listenerT.ExitSelectStmt(s)
	}
}

func (s *SelectStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TideSQLVisitor:
		return t.VisitSelectStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TideSQLParser) SelectStmt() (localctx ISelectStmtContext) {
	localctx = NewSelectStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, TideSQLParserRULE_selectStmt)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(39)
		p.Match(TideSQLParserSELECT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(40)
		p.Measurement()
	}
	{
		p.SetState(41)
		p.Match(TideSQLParserFROM)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(42)
		p.Path()
	}
	p.SetState(44)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == TideSQLParserWHERE {
		{
			p.SetState(43)
			p.WhereClause()
		}

	}
	p.SetState(47)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == TideSQLParserLIMIT {
		{
			p.SetState(46)
			p.LimitClause()
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IWhereClauseContext is an interface to support dynamic dispatch.
type IWhereClauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	WHERE() antlr.TerminalNode
	AllTimePredicate() []ITimePredicateContext
	TimePredicate(i int) ITimePredicateContext
	AllAND() []antlr.TerminalNode
	AND(i int) antlr.TerminalNode

	// IsWhereClauseContext differentiates from other interfaces.
	IsWhereClauseContext()
}

type WhereClauseContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWhereClauseContext() *WhereClauseContext {
	var p = new(WhereClauseContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TideSQLParserRULE_whereClause
	return p
}

func InitEmptyWhereClauseContext(p *WhereClauseContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TideSQLParserRULE_whereClause
}

func (*WhereClauseContext) IsWhereClauseContext() {}

func NewWhereClauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WhereClauseContext {
	var p = new(WhereClauseContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TideSQLParserRULE_whereClause

	return p
}

func (s *WhereClauseContext) GetParser() antlr.Parser { return s.parser }

func (s *WhereClauseContext) WHERE() antlr.TerminalNode {
	return s.GetToken(TideSQLParserWHERE, 0)
}

func (s *WhereClauseContext) AllTimePredicate() []ITimePredicateContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITimePredicateContext); ok {
			len++
		}
	}

	tst := make([]ITimePredicateContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITimePredicateContext); ok {
			tst[i] = t.(ITimePredicateContext)
			i++
		}
	}

	return tst
}

func (s *WhereClauseContext) TimePredicate(i int) ITimePredicateContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITimePredicateContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITimePredicateContext)
}

func (s *WhereClauseContext) AllAND() []antlr.TerminalNode {
	return s.GetTokens(TideSQLParserAND)
}

func (s *WhereClauseContext) AND(i int) antlr.TerminalNode {
	return s.GetToken(TideSQLParserAND, i)
}

func (s *WhereClauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WhereClauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *WhereClauseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TideSQLListener); ok {
		listenerT.EnterWhereClause(s)
	}
}

func (s *WhereClauseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TideSQLListener); ok {
		listenerT.ExitWhereClause(s)
	}
}

func (s *WhereClauseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TideSQLVisitor:
		return t.VisitWhereClause(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TideSQLParser) WhereClause() (localctx IWhereClauseContext) {
	localctx = NewWhereClauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, TideSQLParserRULE_whereClause)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(49)
		p.Match(TideSQLParserWHERE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(50)
		p.TimePredicate()
	}
	p.SetState(55)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == TideSQLParserAND {
		{
			p.SetState(51)
			p.Match(TideSQLParserAND)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(52)
			p.TimePredicate()
		}

		p.SetState(57)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITimePredicateContext is an interface to support dynamic dispatch.
type ITimePredicateContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	TIME() antlr.TerminalNode
	CmpOp() ICmpOpContext
	INTEGER() antlr.TerminalNode

	// IsTimePredicateContext differentiates from other interfaces.
	IsTimePredicateContext()
}

type TimePredicateContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTimePredicateContext() *TimePredicateContext {
	var p = new(TimePredicateContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TideSQLParserRULE_timePredicate
	return p
}

func InitEmptyTimePredicateContext(p *TimePredicateContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TideSQLParserRULE_timePredicate
}

func (*TimePredicateContext) IsTimePredicateContext() {}

func NewTimePredicateContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TimePredicateContext {
	var p = new(TimePredicateContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TideSQLParserRULE_timePredicate

	return p
}

func (s *TimePredicateContext) GetParser() antlr.Parser { return s.parser }

func (s *TimePredicateContext) TIME() antlr.TerminalNode {
	return s.GetToken(TideSQLParserTIME, 0)
}

func (s *TimePredicateContext) CmpOp() ICmpOpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICmpOpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICmpOpContext)
}

func (s *TimePredicateContext) INTEGER() antlr.TerminalNode {
	return s.GetToken(TideSQLParserINTEGER, 0)
}

func (s *TimePredicateContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TimePredicateContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TimePredicateContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TideSQLListener); ok {
		listenerT.EnterTimePredicate(s)
	}
}

func (s *TimePredicateContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TideSQLListener); ok {
		listenerT.ExitTimePredicate(s)
	}
}

func (s *TimePredicateContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TideSQLVisitor:
		return t.VisitTimePredicate(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TideSQLParser) TimePredicate() (localctx ITimePredicateContext) {
	localctx = NewTimePredicateContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, TideSQLParserRULE_timePredicate)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(58)
		p.Match(TideSQLParserTIME)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(59)
		p.CmpOp()
	}
	{
		p.SetState(60)
		p.Match(TideSQLParserINTEGER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICmpOpContext is an interface to support dynamic dispatch.
type ICmpOpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	GTE() antlr.TerminalNode
	LTE() antlr.TerminalNode
	GT() antlr.TerminalNode
	LT() antlr.TerminalNode
	EQ() antlr.TerminalNode

	// IsCmpOpContext differentiates from other interfaces.
	IsCmpOpContext()
}

type CmpOpContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCmpOpContext() *CmpOpContext {
	var p = new(CmpOpContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TideSQLParserRULE_cmpOp
	return p
}

func InitEmptyCmpOpContext(p *CmpOpContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TideSQLParserRULE_cmpOp
}

func (*CmpOpContext) IsCmpOpContext() {}

func NewCmpOpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CmpOpContext {
	var p = new(CmpOpContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TideSQLParserRULE_cmpOp

	return p
}

func (s *CmpOpContext) GetParser() antlr.Parser { return s.parser }

func (s *CmpOpContext) GTE() antlr.TerminalNode {
	return s.GetToken(TideSQLParserGTE, 0)
}

func (s *CmpOpContext) LTE() antlr.TerminalNode {
	return s.GetToken(TideSQLParserLTE, 0)
}

func (s *CmpOpContext) GT() antlr.TerminalNode {
	return s.GetToken(TideSQLParserGT, 0)
}

func (s *CmpOpContext) LT() antlr.TerminalNode {
	return s.GetToken(TideSQLParserLT, 0)
}

func (s *CmpOpContext) EQ() antlr.TerminalNode {
	return s.GetToken(TideSQLParserEQ, 0)
}

func (s *CmpOpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CmpOpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CmpOpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TideSQLListener); ok {
		listenerT.EnterCmpOp(s)
	}
}

func (s *CmpOpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TideSQLListener); ok {
		listenerT.ExitCmpOp(s)
	}
}

func (s *CmpOpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TideSQLVisitor:
		return t.VisitCmpOp(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TideSQLParser) CmpOp() (localctx ICmpOpContext) {
	localctx = NewCmpOpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, TideSQLParserRULE_cmpOp)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(62)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&31744) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILimitClauseContext is an interface to support dynamic dispatch.
type ILimitClauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LIMIT() antlr.TerminalNode
	INTEGER() antlr.TerminalNode

	// IsLimitClauseContext differentiates from other interfaces.
	IsLimitClauseContext()
}

type LimitClauseContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLimitClauseContext() *LimitClauseContext {
	var p = new(LimitClauseContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TideSQLParserRULE_limitClause
	return p
}

func InitEmptyLimitClauseContext(p *LimitClauseContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TideSQLParserRULE_limitClause
}

func (*LimitClauseContext) IsLimitClauseContext() {}

func NewLimitClauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LimitClauseContext {
	var p = new(LimitClauseContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TideSQLParserRULE_limitClause

	return p
}

func (s *LimitClauseContext) GetParser() antlr.Parser { return s.parser }

func (s *LimitClauseContext) LIMIT() antlr.TerminalNode {
	return s.GetToken(TideSQLParserLIMIT, 0)
}

func (s *LimitClauseContext) INTEGER() antlr.TerminalNode {
	return s.GetToken(TideSQLParserINTEGER, 0)
}

func (s *LimitClauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LimitClauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LimitClauseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TideSQLListener); ok {
		listenerT.EnterLimitClause(s)
	}
}

func (s *LimitClauseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TideSQLListener); ok {
		listenerT.ExitLimitClause(s)
	}
}

func (s *LimitClauseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TideSQLVisitor:
		return t.VisitLimitClause(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TideSQLParser) LimitClause() (localctx ILimitClauseContext) {
	localctx = NewLimitClauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, TideSQLParserRULE_limitClause)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(64)
		p.Match(TideSQLParserLIMIT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(65)
		p.Match(TideSQLParserINTEGER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPathContext is an interface to support dynamic dispatch.
type IPathContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllIDENTIFIER() []antlr.TerminalNode
	IDENTIFIER(i int) antlr.TerminalNode
	AllDOT() []antlr.TerminalNode
	DOT(i int) antlr.TerminalNode

	// IsPathContext differentiates from other interfaces.
	IsPathContext()
}

type PathContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPathContext() *PathContext {
	var p = new(PathContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TideSQLParserRULE_path
	return p
}

func InitEmptyPathContext(p *PathContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TideSQLParserRULE_path
}

func (*PathContext) IsPathContext() {}

func NewPathContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PathContext {
	var p = new(PathContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TideSQLParserRULE_path

	return p
}

func (s *PathContext) GetParser() antlr.Parser { return s.parser }

func (s *PathContext) AllIDENTIFIER() []antlr.TerminalNode {
	return s.GetTokens(TideSQLParserIDENTIFIER)
}

func (s *PathContext) IDENTIFIER(i int) antlr.TerminalNode {
	return s.GetToken(TideSQLParserIDENTIFIER, i)
}

func (s *PathContext) AllDOT() []antlr.TerminalNode {
	return s.GetTokens(TideSQLParserDOT)
}

func (s *PathContext) DOT(i int) antlr.TerminalNode {
	return s.GetToken(TideSQLParserDOT, i)
}

func (s *PathContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PathContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PathContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TideSQLListener); ok {
		listenerT.EnterPath(s)
	}
}

func (s *PathContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TideSQLListener); ok {
		listenerT.ExitPath(s)
	}
}

func (s *PathContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TideSQLVisitor:
		return t.VisitPath(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TideSQLParser) Path() (localctx IPathContext) {
	localctx = NewPathContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, TideSQLParserRULE_path)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(67)
		p.Match(TideSQLParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(72)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == TideSQLParserDOT {
		{
			p.SetState(68)
			p.Match(TideSQLParserDOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(69)
			p.Match(TideSQLParserIDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(74)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IMeasurementContext is an interface to support dynamic dispatch.
type IMeasurementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode

	// IsMeasurementContext differentiates from other interfaces.
	IsMeasurementContext()
}

type MeasurementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMeasurementContext() *MeasurementContext {
	var p = new(MeasurementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TideSQLParserRULE_measurement
	return p
}

func InitEmptyMeasurementContext(p *MeasurementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TideSQLParserRULE_measurement
}

func (*MeasurementContext) IsMeasurementContext() {}

func NewMeasurementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MeasurementContext {
	var p = new(MeasurementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TideSQLParserRULE_measurement

	return p
}

func (s *MeasurementContext) GetParser() antlr.Parser { return s.parser }

func (s *MeasurementContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(TideSQLParserIDENTIFIER, 0)
}

func (s *MeasurementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MeasurementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MeasurementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TideSQLListener); ok {
		listenerT.EnterMeasurement(s)
	}
}

func (s *MeasurementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TideSQLListener); ok {
		listenerT.ExitMeasurement(s)
	}
}

func (s *MeasurementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TideSQLVisitor:
		return t.VisitMeasurement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TideSQLParser) Measurement() (localctx IMeasurementContext) {
	localctx = NewMeasurementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, TideSQLParserRULE_measurement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(75)
		p.Match(TideSQLParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITimestampContext is an interface to support dynamic dispatch.
type ITimestampContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	INTEGER() antlr.TerminalNode

	// IsTimestampContext differentiates from other interfaces.
	IsTimestampContext()
}

type TimestampContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTimestampContext() *TimestampContext {
	var p = new(TimestampContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TideSQLParserRULE_timestamp
	return p
}

func InitEmptyTimestampContext(p *TimestampContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TideSQLParserRULE_timestamp
}

func (*TimestampContext) IsTimestampContext() {}

func NewTimestampContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TimestampContext {
	var p = new(TimestampContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TideSQLParserRULE_timestamp

	return p
}

func (s *TimestampContext) GetParser() antlr.Parser { return s.parser }

func (s *TimestampContext) INTEGER() antlr.TerminalNode {
	return s.GetToken(TideSQLParserINTEGER, 0)
}

func (s *TimestampContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TimestampContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TimestampContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TideSQLListener); ok {
		listenerT.EnterTimestamp(s)
	}
}

func (s *TimestampContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TideSQLListener); ok {
		listenerT.ExitTimestamp(s)
	}
}

func (s *TimestampContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TideSQLVisitor:
		return t.VisitTimestamp(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TideSQLParser) Timestamp() (localctx ITimestampContext) {
	localctx = NewTimestampContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, TideSQLParserRULE_timestamp)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(77)
		p.Match(TideSQLParserINTEGER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IValueContext is an interface to support dynamic dispatch.
type IValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	BOOLEAN() antlr.TerminalNode
	INTEGER() antlr.TerminalNode
	FLOAT() antlr.TerminalNode
	STRING() antlr.TerminalNode

	// IsValueContext differentiates from other interfaces.
	IsValueContext()
}

type ValueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValueContext() *ValueContext {
	var p = new(ValueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TideSQLParserRULE_value
	return p
}

func InitEmptyValueContext(p *ValueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TideSQLParserRULE_value
}

func (*ValueContext) IsValueContext() {}

func NewValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValueContext {
	var p = new(ValueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TideSQLParserRULE_value

	return p
}

func (s *ValueContext) GetParser() antlr.Parser { return s.parser }

func (s *ValueContext) BOOLEAN() antlr.TerminalNode {
	return s.GetToken(TideSQLParserBOOLEAN, 0)
}

func (s *ValueContext) INTEGER() antlr.TerminalNode {
	return s.GetToken(TideSQLParserINTEGER, 0)
}

func (s *ValueContext) FLOAT() antlr.TerminalNode {
	return s.GetToken(TideSQLParserFLOAT, 0)
}

func (s *ValueContext) STRING() antlr.TerminalNode {
	return s.GetToken(TideSQLParserSTRING, 0)
}

func (s *ValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TideSQLListener); ok {
		listenerT.EnterValue(s)
	}
}

func (s *ValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TideSQLListener); ok {
		listenerT.ExitValue(s)
	}
}

func (s *ValueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TideSQLVisitor:
		return t.VisitValue(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TideSQLParser) Value() (localctx IValueContext) {
	localctx = NewValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, TideSQLParserRULE_value)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(79)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&15728640) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}
