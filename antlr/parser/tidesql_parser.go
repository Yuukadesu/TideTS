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
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "'>='", "'<='", "'>'", "'<'", "'='", "'('", "')'", "','", "'.'",
		"';'", "'*'",
	}
	staticData.SymbolicNames = []string{
		"", "INSERT", "DELETE", "INTO", "SELECT", "COUNT", "FROM", "WHERE",
		"AND", "VALUES", "LIMIT", "TIME", "CREATE", "TIMESERIES", "WITH", "DATATYPE",
		"SHOW", "DEVICES", "GTE", "LTE", "GT", "LT", "EQ", "LPAREN", "RPAREN",
		"COMMA", "DOT", "SEMI", "STAR", "BOOLEAN", "INTEGER", "FLOAT", "STRING",
		"IDENTIFIER", "WS",
	}
	staticData.RuleNames = []string{
		"statement", "insertStmt", "valueRow", "selectStmt", "deleteStmt", "createTimeseriesStmt",
		"showDevicesStmt", "showTimeseriesStmt", "showPattern", "whereClause",
		"timePredicate", "cmpOp", "limitClause", "path", "measurement", "dataTypeName",
		"timestamp", "value",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 34, 156, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 3, 0, 43,
		8, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 5, 1,
		55, 8, 1, 10, 1, 12, 1, 58, 9, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1,
		3, 1, 3, 1, 3, 1, 3, 1, 3, 3, 3, 71, 8, 3, 1, 3, 3, 3, 74, 8, 3, 1, 3,
		1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 3, 3, 84, 8, 3, 3, 3, 86, 8,
		3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1,
		5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 3, 6, 110,
		8, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 8, 3, 8, 120, 8, 8,
		1, 9, 1, 9, 1, 9, 1, 9, 5, 9, 126, 8, 9, 10, 9, 12, 9, 129, 9, 9, 1, 10,
		1, 10, 1, 10, 1, 10, 1, 11, 1, 11, 1, 12, 1, 12, 1, 12, 1, 13, 1, 13, 1,
		13, 5, 13, 143, 8, 13, 10, 13, 12, 13, 146, 9, 13, 1, 14, 1, 14, 1, 15,
		1, 15, 1, 16, 1, 16, 1, 17, 1, 17, 1, 17, 0, 0, 18, 0, 2, 4, 6, 8, 10,
		12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 0, 2, 1, 0, 18, 22, 1,
		0, 29, 32, 151, 0, 42, 1, 0, 0, 0, 2, 44, 1, 0, 0, 0, 4, 59, 1, 0, 0, 0,
		6, 85, 1, 0, 0, 0, 8, 87, 1, 0, 0, 0, 10, 95, 1, 0, 0, 0, 12, 106, 1, 0,
		0, 0, 14, 111, 1, 0, 0, 0, 16, 115, 1, 0, 0, 0, 18, 121, 1, 0, 0, 0, 20,
		130, 1, 0, 0, 0, 22, 134, 1, 0, 0, 0, 24, 136, 1, 0, 0, 0, 26, 139, 1,
		0, 0, 0, 28, 147, 1, 0, 0, 0, 30, 149, 1, 0, 0, 0, 32, 151, 1, 0, 0, 0,
		34, 153, 1, 0, 0, 0, 36, 43, 3, 2, 1, 0, 37, 43, 3, 6, 3, 0, 38, 43, 3,
		8, 4, 0, 39, 43, 3, 10, 5, 0, 40, 43, 3, 12, 6, 0, 41, 43, 3, 14, 7, 0,
		42, 36, 1, 0, 0, 0, 42, 37, 1, 0, 0, 0, 42, 38, 1, 0, 0, 0, 42, 39, 1,
		0, 0, 0, 42, 40, 1, 0, 0, 0, 42, 41, 1, 0, 0, 0, 43, 1, 1, 0, 0, 0, 44,
		45, 5, 1, 0, 0, 45, 46, 5, 3, 0, 0, 46, 47, 3, 26, 13, 0, 47, 48, 5, 23,
		0, 0, 48, 49, 3, 28, 14, 0, 49, 50, 5, 24, 0, 0, 50, 51, 5, 9, 0, 0, 51,
		56, 3, 4, 2, 0, 52, 53, 5, 25, 0, 0, 53, 55, 3, 4, 2, 0, 54, 52, 1, 0,
		0, 0, 55, 58, 1, 0, 0, 0, 56, 54, 1, 0, 0, 0, 56, 57, 1, 0, 0, 0, 57, 3,
		1, 0, 0, 0, 58, 56, 1, 0, 0, 0, 59, 60, 5, 23, 0, 0, 60, 61, 3, 32, 16,
		0, 61, 62, 5, 25, 0, 0, 62, 63, 3, 34, 17, 0, 63, 64, 5, 24, 0, 0, 64,
		5, 1, 0, 0, 0, 65, 66, 5, 4, 0, 0, 66, 67, 3, 28, 14, 0, 67, 68, 5, 6,
		0, 0, 68, 70, 3, 26, 13, 0, 69, 71, 3, 18, 9, 0, 70, 69, 1, 0, 0, 0, 70,
		71, 1, 0, 0, 0, 71, 73, 1, 0, 0, 0, 72, 74, 3, 24, 12, 0, 73, 72, 1, 0,
		0, 0, 73, 74, 1, 0, 0, 0, 74, 86, 1, 0, 0, 0, 75, 76, 5, 4, 0, 0, 76, 77,
		5, 5, 0, 0, 77, 78, 5, 23, 0, 0, 78, 79, 3, 28, 14, 0, 79, 80, 5, 24, 0,
		0, 80, 81, 5, 6, 0, 0, 81, 83, 3, 26, 13, 0, 82, 84, 3, 18, 9, 0, 83, 82,
		1, 0, 0, 0, 83, 84, 1, 0, 0, 0, 84, 86, 1, 0, 0, 0, 85, 65, 1, 0, 0, 0,
		85, 75, 1, 0, 0, 0, 86, 7, 1, 0, 0, 0, 87, 88, 5, 2, 0, 0, 88, 89, 5, 6,
		0, 0, 89, 90, 3, 26, 13, 0, 90, 91, 5, 23, 0, 0, 91, 92, 3, 28, 14, 0,
		92, 93, 5, 24, 0, 0, 93, 94, 3, 18, 9, 0, 94, 9, 1, 0, 0, 0, 95, 96, 5,
		12, 0, 0, 96, 97, 5, 13, 0, 0, 97, 98, 3, 26, 13, 0, 98, 99, 5, 23, 0,
		0, 99, 100, 3, 28, 14, 0, 100, 101, 5, 24, 0, 0, 101, 102, 5, 14, 0, 0,
		102, 103, 5, 15, 0, 0, 103, 104, 5, 22, 0, 0, 104, 105, 3, 30, 15, 0, 105,
		11, 1, 0, 0, 0, 106, 107, 5, 16, 0, 0, 107, 109, 5, 17, 0, 0, 108, 110,
		3, 16, 8, 0, 109, 108, 1, 0, 0, 0, 109, 110, 1, 0, 0, 0, 110, 13, 1, 0,
		0, 0, 111, 112, 5, 16, 0, 0, 112, 113, 5, 13, 0, 0, 113, 114, 3, 26, 13,
		0, 114, 15, 1, 0, 0, 0, 115, 119, 3, 26, 13, 0, 116, 117, 5, 26, 0, 0,
		117, 118, 5, 28, 0, 0, 118, 120, 5, 28, 0, 0, 119, 116, 1, 0, 0, 0, 119,
		120, 1, 0, 0, 0, 120, 17, 1, 0, 0, 0, 121, 122, 5, 7, 0, 0, 122, 127, 3,
		20, 10, 0, 123, 124, 5, 8, 0, 0, 124, 126, 3, 20, 10, 0, 125, 123, 1, 0,
		0, 0, 126, 129, 1, 0, 0, 0, 127, 125, 1, 0, 0, 0, 127, 128, 1, 0, 0, 0,
		128, 19, 1, 0, 0, 0, 129, 127, 1, 0, 0, 0, 130, 131, 5, 11, 0, 0, 131,
		132, 3, 22, 11, 0, 132, 133, 5, 30, 0, 0, 133, 21, 1, 0, 0, 0, 134, 135,
		7, 0, 0, 0, 135, 23, 1, 0, 0, 0, 136, 137, 5, 10, 0, 0, 137, 138, 5, 30,
		0, 0, 138, 25, 1, 0, 0, 0, 139, 144, 5, 33, 0, 0, 140, 141, 5, 26, 0, 0,
		141, 143, 5, 33, 0, 0, 142, 140, 1, 0, 0, 0, 143, 146, 1, 0, 0, 0, 144,
		142, 1, 0, 0, 0, 144, 145, 1, 0, 0, 0, 145, 27, 1, 0, 0, 0, 146, 144, 1,
		0, 0, 0, 147, 148, 5, 33, 0, 0, 148, 29, 1, 0, 0, 0, 149, 150, 5, 33, 0,
		0, 150, 31, 1, 0, 0, 0, 151, 152, 5, 30, 0, 0, 152, 33, 1, 0, 0, 0, 153,
		154, 7, 1, 0, 0, 154, 35, 1, 0, 0, 0, 10, 42, 56, 70, 73, 83, 85, 109,
		119, 127, 144,
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
	TideSQLParserDELETE     = 2
	TideSQLParserINTO       = 3
	TideSQLParserSELECT     = 4
	TideSQLParserCOUNT      = 5
	TideSQLParserFROM       = 6
	TideSQLParserWHERE      = 7
	TideSQLParserAND        = 8
	TideSQLParserVALUES     = 9
	TideSQLParserLIMIT      = 10
	TideSQLParserTIME       = 11
	TideSQLParserCREATE     = 12
	TideSQLParserTIMESERIES = 13
	TideSQLParserWITH       = 14
	TideSQLParserDATATYPE   = 15
	TideSQLParserSHOW       = 16
	TideSQLParserDEVICES    = 17
	TideSQLParserGTE        = 18
	TideSQLParserLTE        = 19
	TideSQLParserGT         = 20
	TideSQLParserLT         = 21
	TideSQLParserEQ         = 22
	TideSQLParserLPAREN     = 23
	TideSQLParserRPAREN     = 24
	TideSQLParserCOMMA      = 25
	TideSQLParserDOT        = 26
	TideSQLParserSEMI       = 27
	TideSQLParserSTAR       = 28
	TideSQLParserBOOLEAN    = 29
	TideSQLParserINTEGER    = 30
	TideSQLParserFLOAT      = 31
	TideSQLParserSTRING     = 32
	TideSQLParserIDENTIFIER = 33
	TideSQLParserWS         = 34
)

// TideSQLParser rules.
const (
	TideSQLParserRULE_statement            = 0
	TideSQLParserRULE_insertStmt           = 1
	TideSQLParserRULE_valueRow             = 2
	TideSQLParserRULE_selectStmt           = 3
	TideSQLParserRULE_deleteStmt           = 4
	TideSQLParserRULE_createTimeseriesStmt = 5
	TideSQLParserRULE_showDevicesStmt      = 6
	TideSQLParserRULE_showTimeseriesStmt   = 7
	TideSQLParserRULE_showPattern          = 8
	TideSQLParserRULE_whereClause          = 9
	TideSQLParserRULE_timePredicate        = 10
	TideSQLParserRULE_cmpOp                = 11
	TideSQLParserRULE_limitClause          = 12
	TideSQLParserRULE_path                 = 13
	TideSQLParserRULE_measurement          = 14
	TideSQLParserRULE_dataTypeName         = 15
	TideSQLParserRULE_timestamp            = 16
	TideSQLParserRULE_value                = 17
)

// IStatementContext is an interface to support dynamic dispatch.
type IStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	InsertStmt() IInsertStmtContext
	SelectStmt() ISelectStmtContext
	DeleteStmt() IDeleteStmtContext
	CreateTimeseriesStmt() ICreateTimeseriesStmtContext
	ShowDevicesStmt() IShowDevicesStmtContext
	ShowTimeseriesStmt() IShowTimeseriesStmtContext

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

func (s *StatementContext) DeleteStmt() IDeleteStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeleteStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeleteStmtContext)
}

func (s *StatementContext) CreateTimeseriesStmt() ICreateTimeseriesStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICreateTimeseriesStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICreateTimeseriesStmtContext)
}

func (s *StatementContext) ShowDevicesStmt() IShowDevicesStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IShowDevicesStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IShowDevicesStmtContext)
}

func (s *StatementContext) ShowTimeseriesStmt() IShowTimeseriesStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IShowTimeseriesStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IShowTimeseriesStmtContext)
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
	p.SetState(42)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 0, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(36)
			p.InsertStmt()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(37)
			p.SelectStmt()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(38)
			p.DeleteStmt()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(39)
			p.CreateTimeseriesStmt()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(40)
			p.ShowDevicesStmt()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(41)
			p.ShowTimeseriesStmt()
		}

	case antlr.ATNInvalidAltNumber:
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
	LPAREN() antlr.TerminalNode
	Measurement() IMeasurementContext
	RPAREN() antlr.TerminalNode
	VALUES() antlr.TerminalNode
	AllValueRow() []IValueRowContext
	ValueRow(i int) IValueRowContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

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

func (s *InsertStmtContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(TideSQLParserLPAREN, 0)
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

func (s *InsertStmtContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(TideSQLParserRPAREN, 0)
}

func (s *InsertStmtContext) VALUES() antlr.TerminalNode {
	return s.GetToken(TideSQLParserVALUES, 0)
}

func (s *InsertStmtContext) AllValueRow() []IValueRowContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IValueRowContext); ok {
			len++
		}
	}

	tst := make([]IValueRowContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IValueRowContext); ok {
			tst[i] = t.(IValueRowContext)
			i++
		}
	}

	return tst
}

func (s *InsertStmtContext) ValueRow(i int) IValueRowContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueRowContext); ok {
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

	return t.(IValueRowContext)
}

func (s *InsertStmtContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(TideSQLParserCOMMA)
}

func (s *InsertStmtContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(TideSQLParserCOMMA, i)
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
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(44)
		p.Match(TideSQLParserINSERT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(45)
		p.Match(TideSQLParserINTO)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(46)
		p.Path()
	}
	{
		p.SetState(47)
		p.Match(TideSQLParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(48)
		p.Measurement()
	}
	{
		p.SetState(49)
		p.Match(TideSQLParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(50)
		p.Match(TideSQLParserVALUES)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(51)
		p.ValueRow()
	}
	p.SetState(56)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == TideSQLParserCOMMA {
		{
			p.SetState(52)
			p.Match(TideSQLParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(53)
			p.ValueRow()
		}

		p.SetState(58)
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

// IValueRowContext is an interface to support dynamic dispatch.
type IValueRowContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LPAREN() antlr.TerminalNode
	Timestamp() ITimestampContext
	COMMA() antlr.TerminalNode
	Value() IValueContext
	RPAREN() antlr.TerminalNode

	// IsValueRowContext differentiates from other interfaces.
	IsValueRowContext()
}

type ValueRowContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValueRowContext() *ValueRowContext {
	var p = new(ValueRowContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TideSQLParserRULE_valueRow
	return p
}

func InitEmptyValueRowContext(p *ValueRowContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TideSQLParserRULE_valueRow
}

func (*ValueRowContext) IsValueRowContext() {}

func NewValueRowContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValueRowContext {
	var p = new(ValueRowContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TideSQLParserRULE_valueRow

	return p
}

func (s *ValueRowContext) GetParser() antlr.Parser { return s.parser }

func (s *ValueRowContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(TideSQLParserLPAREN, 0)
}

func (s *ValueRowContext) Timestamp() ITimestampContext {
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

func (s *ValueRowContext) COMMA() antlr.TerminalNode {
	return s.GetToken(TideSQLParserCOMMA, 0)
}

func (s *ValueRowContext) Value() IValueContext {
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

func (s *ValueRowContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(TideSQLParserRPAREN, 0)
}

func (s *ValueRowContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueRowContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ValueRowContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TideSQLListener); ok {
		listenerT.EnterValueRow(s)
	}
}

func (s *ValueRowContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TideSQLListener); ok {
		listenerT.ExitValueRow(s)
	}
}

func (s *ValueRowContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TideSQLVisitor:
		return t.VisitValueRow(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TideSQLParser) ValueRow() (localctx IValueRowContext) {
	localctx = NewValueRowContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, TideSQLParserRULE_valueRow)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(59)
		p.Match(TideSQLParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(60)
		p.Timestamp()
	}
	{
		p.SetState(61)
		p.Match(TideSQLParserCOMMA)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(62)
		p.Value()
	}
	{
		p.SetState(63)
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
	COUNT() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	RPAREN() antlr.TerminalNode

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

func (s *SelectStmtContext) COUNT() antlr.TerminalNode {
	return s.GetToken(TideSQLParserCOUNT, 0)
}

func (s *SelectStmtContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(TideSQLParserLPAREN, 0)
}

func (s *SelectStmtContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(TideSQLParserRPAREN, 0)
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
	p.EnterRule(localctx, 6, TideSQLParserRULE_selectStmt)
	var _la int

	p.SetState(85)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 5, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(65)
			p.Match(TideSQLParserSELECT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(66)
			p.Measurement()
		}
		{
			p.SetState(67)
			p.Match(TideSQLParserFROM)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(68)
			p.Path()
		}
		p.SetState(70)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == TideSQLParserWHERE {
			{
				p.SetState(69)
				p.WhereClause()
			}

		}
		p.SetState(73)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == TideSQLParserLIMIT {
			{
				p.SetState(72)
				p.LimitClause()
			}

		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(75)
			p.Match(TideSQLParserSELECT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(76)
			p.Match(TideSQLParserCOUNT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(77)
			p.Match(TideSQLParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(78)
			p.Measurement()
		}
		{
			p.SetState(79)
			p.Match(TideSQLParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(80)
			p.Match(TideSQLParserFROM)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(81)
			p.Path()
		}
		p.SetState(83)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == TideSQLParserWHERE {
			{
				p.SetState(82)
				p.WhereClause()
			}

		}

	case antlr.ATNInvalidAltNumber:
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

// IDeleteStmtContext is an interface to support dynamic dispatch.
type IDeleteStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	DELETE() antlr.TerminalNode
	FROM() antlr.TerminalNode
	Path() IPathContext
	LPAREN() antlr.TerminalNode
	Measurement() IMeasurementContext
	RPAREN() antlr.TerminalNode
	WhereClause() IWhereClauseContext

	// IsDeleteStmtContext differentiates from other interfaces.
	IsDeleteStmtContext()
}

type DeleteStmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDeleteStmtContext() *DeleteStmtContext {
	var p = new(DeleteStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TideSQLParserRULE_deleteStmt
	return p
}

func InitEmptyDeleteStmtContext(p *DeleteStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TideSQLParserRULE_deleteStmt
}

func (*DeleteStmtContext) IsDeleteStmtContext() {}

func NewDeleteStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeleteStmtContext {
	var p = new(DeleteStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TideSQLParserRULE_deleteStmt

	return p
}

func (s *DeleteStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *DeleteStmtContext) DELETE() antlr.TerminalNode {
	return s.GetToken(TideSQLParserDELETE, 0)
}

func (s *DeleteStmtContext) FROM() antlr.TerminalNode {
	return s.GetToken(TideSQLParserFROM, 0)
}

func (s *DeleteStmtContext) Path() IPathContext {
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

func (s *DeleteStmtContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(TideSQLParserLPAREN, 0)
}

func (s *DeleteStmtContext) Measurement() IMeasurementContext {
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

func (s *DeleteStmtContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(TideSQLParserRPAREN, 0)
}

func (s *DeleteStmtContext) WhereClause() IWhereClauseContext {
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

func (s *DeleteStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeleteStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DeleteStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TideSQLListener); ok {
		listenerT.EnterDeleteStmt(s)
	}
}

func (s *DeleteStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TideSQLListener); ok {
		listenerT.ExitDeleteStmt(s)
	}
}

func (s *DeleteStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TideSQLVisitor:
		return t.VisitDeleteStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TideSQLParser) DeleteStmt() (localctx IDeleteStmtContext) {
	localctx = NewDeleteStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, TideSQLParserRULE_deleteStmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(87)
		p.Match(TideSQLParserDELETE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(88)
		p.Match(TideSQLParserFROM)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(89)
		p.Path()
	}
	{
		p.SetState(90)
		p.Match(TideSQLParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(91)
		p.Measurement()
	}
	{
		p.SetState(92)
		p.Match(TideSQLParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(93)
		p.WhereClause()
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

// ICreateTimeseriesStmtContext is an interface to support dynamic dispatch.
type ICreateTimeseriesStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	CREATE() antlr.TerminalNode
	TIMESERIES() antlr.TerminalNode
	Path() IPathContext
	LPAREN() antlr.TerminalNode
	Measurement() IMeasurementContext
	RPAREN() antlr.TerminalNode
	WITH() antlr.TerminalNode
	DATATYPE() antlr.TerminalNode
	EQ() antlr.TerminalNode
	DataTypeName() IDataTypeNameContext

	// IsCreateTimeseriesStmtContext differentiates from other interfaces.
	IsCreateTimeseriesStmtContext()
}

type CreateTimeseriesStmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCreateTimeseriesStmtContext() *CreateTimeseriesStmtContext {
	var p = new(CreateTimeseriesStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TideSQLParserRULE_createTimeseriesStmt
	return p
}

func InitEmptyCreateTimeseriesStmtContext(p *CreateTimeseriesStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TideSQLParserRULE_createTimeseriesStmt
}

func (*CreateTimeseriesStmtContext) IsCreateTimeseriesStmtContext() {}

func NewCreateTimeseriesStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CreateTimeseriesStmtContext {
	var p = new(CreateTimeseriesStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TideSQLParserRULE_createTimeseriesStmt

	return p
}

func (s *CreateTimeseriesStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *CreateTimeseriesStmtContext) CREATE() antlr.TerminalNode {
	return s.GetToken(TideSQLParserCREATE, 0)
}

func (s *CreateTimeseriesStmtContext) TIMESERIES() antlr.TerminalNode {
	return s.GetToken(TideSQLParserTIMESERIES, 0)
}

func (s *CreateTimeseriesStmtContext) Path() IPathContext {
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

func (s *CreateTimeseriesStmtContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(TideSQLParserLPAREN, 0)
}

func (s *CreateTimeseriesStmtContext) Measurement() IMeasurementContext {
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

func (s *CreateTimeseriesStmtContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(TideSQLParserRPAREN, 0)
}

func (s *CreateTimeseriesStmtContext) WITH() antlr.TerminalNode {
	return s.GetToken(TideSQLParserWITH, 0)
}

func (s *CreateTimeseriesStmtContext) DATATYPE() antlr.TerminalNode {
	return s.GetToken(TideSQLParserDATATYPE, 0)
}

func (s *CreateTimeseriesStmtContext) EQ() antlr.TerminalNode {
	return s.GetToken(TideSQLParserEQ, 0)
}

func (s *CreateTimeseriesStmtContext) DataTypeName() IDataTypeNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDataTypeNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDataTypeNameContext)
}

func (s *CreateTimeseriesStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CreateTimeseriesStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CreateTimeseriesStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TideSQLListener); ok {
		listenerT.EnterCreateTimeseriesStmt(s)
	}
}

func (s *CreateTimeseriesStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TideSQLListener); ok {
		listenerT.ExitCreateTimeseriesStmt(s)
	}
}

func (s *CreateTimeseriesStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TideSQLVisitor:
		return t.VisitCreateTimeseriesStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TideSQLParser) CreateTimeseriesStmt() (localctx ICreateTimeseriesStmtContext) {
	localctx = NewCreateTimeseriesStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, TideSQLParserRULE_createTimeseriesStmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(95)
		p.Match(TideSQLParserCREATE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(96)
		p.Match(TideSQLParserTIMESERIES)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(97)
		p.Path()
	}
	{
		p.SetState(98)
		p.Match(TideSQLParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(99)
		p.Measurement()
	}
	{
		p.SetState(100)
		p.Match(TideSQLParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(101)
		p.Match(TideSQLParserWITH)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(102)
		p.Match(TideSQLParserDATATYPE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(103)
		p.Match(TideSQLParserEQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(104)
		p.DataTypeName()
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

// IShowDevicesStmtContext is an interface to support dynamic dispatch.
type IShowDevicesStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SHOW() antlr.TerminalNode
	DEVICES() antlr.TerminalNode
	ShowPattern() IShowPatternContext

	// IsShowDevicesStmtContext differentiates from other interfaces.
	IsShowDevicesStmtContext()
}

type ShowDevicesStmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyShowDevicesStmtContext() *ShowDevicesStmtContext {
	var p = new(ShowDevicesStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TideSQLParserRULE_showDevicesStmt
	return p
}

func InitEmptyShowDevicesStmtContext(p *ShowDevicesStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TideSQLParserRULE_showDevicesStmt
}

func (*ShowDevicesStmtContext) IsShowDevicesStmtContext() {}

func NewShowDevicesStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ShowDevicesStmtContext {
	var p = new(ShowDevicesStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TideSQLParserRULE_showDevicesStmt

	return p
}

func (s *ShowDevicesStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *ShowDevicesStmtContext) SHOW() antlr.TerminalNode {
	return s.GetToken(TideSQLParserSHOW, 0)
}

func (s *ShowDevicesStmtContext) DEVICES() antlr.TerminalNode {
	return s.GetToken(TideSQLParserDEVICES, 0)
}

func (s *ShowDevicesStmtContext) ShowPattern() IShowPatternContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IShowPatternContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IShowPatternContext)
}

func (s *ShowDevicesStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ShowDevicesStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ShowDevicesStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TideSQLListener); ok {
		listenerT.EnterShowDevicesStmt(s)
	}
}

func (s *ShowDevicesStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TideSQLListener); ok {
		listenerT.ExitShowDevicesStmt(s)
	}
}

func (s *ShowDevicesStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TideSQLVisitor:
		return t.VisitShowDevicesStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TideSQLParser) ShowDevicesStmt() (localctx IShowDevicesStmtContext) {
	localctx = NewShowDevicesStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, TideSQLParserRULE_showDevicesStmt)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(106)
		p.Match(TideSQLParserSHOW)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(107)
		p.Match(TideSQLParserDEVICES)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(109)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == TideSQLParserIDENTIFIER {
		{
			p.SetState(108)
			p.ShowPattern()
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

// IShowTimeseriesStmtContext is an interface to support dynamic dispatch.
type IShowTimeseriesStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SHOW() antlr.TerminalNode
	TIMESERIES() antlr.TerminalNode
	Path() IPathContext

	// IsShowTimeseriesStmtContext differentiates from other interfaces.
	IsShowTimeseriesStmtContext()
}

type ShowTimeseriesStmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyShowTimeseriesStmtContext() *ShowTimeseriesStmtContext {
	var p = new(ShowTimeseriesStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TideSQLParserRULE_showTimeseriesStmt
	return p
}

func InitEmptyShowTimeseriesStmtContext(p *ShowTimeseriesStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TideSQLParserRULE_showTimeseriesStmt
}

func (*ShowTimeseriesStmtContext) IsShowTimeseriesStmtContext() {}

func NewShowTimeseriesStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ShowTimeseriesStmtContext {
	var p = new(ShowTimeseriesStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TideSQLParserRULE_showTimeseriesStmt

	return p
}

func (s *ShowTimeseriesStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *ShowTimeseriesStmtContext) SHOW() antlr.TerminalNode {
	return s.GetToken(TideSQLParserSHOW, 0)
}

func (s *ShowTimeseriesStmtContext) TIMESERIES() antlr.TerminalNode {
	return s.GetToken(TideSQLParserTIMESERIES, 0)
}

func (s *ShowTimeseriesStmtContext) Path() IPathContext {
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

func (s *ShowTimeseriesStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ShowTimeseriesStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ShowTimeseriesStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TideSQLListener); ok {
		listenerT.EnterShowTimeseriesStmt(s)
	}
}

func (s *ShowTimeseriesStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TideSQLListener); ok {
		listenerT.ExitShowTimeseriesStmt(s)
	}
}

func (s *ShowTimeseriesStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TideSQLVisitor:
		return t.VisitShowTimeseriesStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TideSQLParser) ShowTimeseriesStmt() (localctx IShowTimeseriesStmtContext) {
	localctx = NewShowTimeseriesStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, TideSQLParserRULE_showTimeseriesStmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(111)
		p.Match(TideSQLParserSHOW)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(112)
		p.Match(TideSQLParserTIMESERIES)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(113)
		p.Path()
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

// IShowPatternContext is an interface to support dynamic dispatch.
type IShowPatternContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Path() IPathContext
	DOT() antlr.TerminalNode
	AllSTAR() []antlr.TerminalNode
	STAR(i int) antlr.TerminalNode

	// IsShowPatternContext differentiates from other interfaces.
	IsShowPatternContext()
}

type ShowPatternContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyShowPatternContext() *ShowPatternContext {
	var p = new(ShowPatternContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TideSQLParserRULE_showPattern
	return p
}

func InitEmptyShowPatternContext(p *ShowPatternContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TideSQLParserRULE_showPattern
}

func (*ShowPatternContext) IsShowPatternContext() {}

func NewShowPatternContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ShowPatternContext {
	var p = new(ShowPatternContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TideSQLParserRULE_showPattern

	return p
}

func (s *ShowPatternContext) GetParser() antlr.Parser { return s.parser }

func (s *ShowPatternContext) Path() IPathContext {
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

func (s *ShowPatternContext) DOT() antlr.TerminalNode {
	return s.GetToken(TideSQLParserDOT, 0)
}

func (s *ShowPatternContext) AllSTAR() []antlr.TerminalNode {
	return s.GetTokens(TideSQLParserSTAR)
}

func (s *ShowPatternContext) STAR(i int) antlr.TerminalNode {
	return s.GetToken(TideSQLParserSTAR, i)
}

func (s *ShowPatternContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ShowPatternContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ShowPatternContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TideSQLListener); ok {
		listenerT.EnterShowPattern(s)
	}
}

func (s *ShowPatternContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TideSQLListener); ok {
		listenerT.ExitShowPattern(s)
	}
}

func (s *ShowPatternContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TideSQLVisitor:
		return t.VisitShowPattern(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TideSQLParser) ShowPattern() (localctx IShowPatternContext) {
	localctx = NewShowPatternContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, TideSQLParserRULE_showPattern)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(115)
		p.Path()
	}
	p.SetState(119)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == TideSQLParserDOT {
		{
			p.SetState(116)
			p.Match(TideSQLParserDOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(117)
			p.Match(TideSQLParserSTAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(118)
			p.Match(TideSQLParserSTAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
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
	p.EnterRule(localctx, 18, TideSQLParserRULE_whereClause)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(121)
		p.Match(TideSQLParserWHERE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(122)
		p.TimePredicate()
	}
	p.SetState(127)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == TideSQLParserAND {
		{
			p.SetState(123)
			p.Match(TideSQLParserAND)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(124)
			p.TimePredicate()
		}

		p.SetState(129)
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
	p.EnterRule(localctx, 20, TideSQLParserRULE_timePredicate)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(130)
		p.Match(TideSQLParserTIME)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(131)
		p.CmpOp()
	}
	{
		p.SetState(132)
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
	p.EnterRule(localctx, 22, TideSQLParserRULE_cmpOp)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(134)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&8126464) != 0) {
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
	p.EnterRule(localctx, 24, TideSQLParserRULE_limitClause)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(136)
		p.Match(TideSQLParserLIMIT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(137)
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
	p.EnterRule(localctx, 26, TideSQLParserRULE_path)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(139)
		p.Match(TideSQLParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(144)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 9, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(140)
				p.Match(TideSQLParserDOT)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(141)
				p.Match(TideSQLParserIDENTIFIER)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		p.SetState(146)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 9, p.GetParserRuleContext())
		if p.HasError() {
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
	p.EnterRule(localctx, 28, TideSQLParserRULE_measurement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(147)
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

// IDataTypeNameContext is an interface to support dynamic dispatch.
type IDataTypeNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode

	// IsDataTypeNameContext differentiates from other interfaces.
	IsDataTypeNameContext()
}

type DataTypeNameContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDataTypeNameContext() *DataTypeNameContext {
	var p = new(DataTypeNameContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TideSQLParserRULE_dataTypeName
	return p
}

func InitEmptyDataTypeNameContext(p *DataTypeNameContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TideSQLParserRULE_dataTypeName
}

func (*DataTypeNameContext) IsDataTypeNameContext() {}

func NewDataTypeNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DataTypeNameContext {
	var p = new(DataTypeNameContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TideSQLParserRULE_dataTypeName

	return p
}

func (s *DataTypeNameContext) GetParser() antlr.Parser { return s.parser }

func (s *DataTypeNameContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(TideSQLParserIDENTIFIER, 0)
}

func (s *DataTypeNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DataTypeNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DataTypeNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TideSQLListener); ok {
		listenerT.EnterDataTypeName(s)
	}
}

func (s *DataTypeNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TideSQLListener); ok {
		listenerT.ExitDataTypeName(s)
	}
}

func (s *DataTypeNameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TideSQLVisitor:
		return t.VisitDataTypeName(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TideSQLParser) DataTypeName() (localctx IDataTypeNameContext) {
	localctx = NewDataTypeNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, TideSQLParserRULE_dataTypeName)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(149)
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
	p.EnterRule(localctx, 32, TideSQLParserRULE_timestamp)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(151)
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
	p.EnterRule(localctx, 34, TideSQLParserRULE_value)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(153)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&8053063680) != 0) {
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
