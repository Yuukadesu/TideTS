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
		"", "", "", "", "", "", "", "", "", "", "", "'>='", "'<='", "'>'", "'<'",
		"'='", "'('", "')'", "','", "'.'", "';'", "'*'",
	}
	staticData.SymbolicNames = []string{
		"", "INSERT", "DELETE", "INTO", "SELECT", "COUNT", "FROM", "WHERE",
		"AND", "VALUES", "LIMIT", "TIME", "CREATE", "TIMESERIES", "WITH", "DATATYPE",
		"SHOW", "DEVICES", "USER", "ALTER", "SET", "PASSWORD", "GRANT", "REVOKE",
		"READ", "WRITE", "ON", "TO", "GTE", "LTE", "GT", "LT", "EQ", "LPAREN",
		"RPAREN", "COMMA", "DOT", "SEMI", "STAR", "BOOLEAN", "INTEGER", "FLOAT",
		"STRING", "IDENTIFIER", "WS",
	}
	staticData.RuleNames = []string{
		"statement", "insertStmt", "valueRow", "selectStmt", "deleteStmt", "createTimeseriesStmt",
		"showDevicesStmt", "showTimeseriesStmt", "createUserStmt", "alterUserStmt",
		"grantStmt", "revokeStmt", "privList", "priv", "grantPath", "showPattern",
		"whereClause", "timePredicate", "cmpOp", "limitClause", "path", "measurement",
		"dataTypeName", "timestamp", "value",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 44, 222, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 1, 0, 1, 0, 1, 0,
		1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 3, 0, 61, 8, 0, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 5, 1, 73, 8, 1, 10, 1, 12,
		1, 76, 9, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 3,
		1, 3, 3, 3, 89, 8, 3, 1, 3, 3, 3, 92, 8, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1,
		3, 1, 3, 1, 3, 1, 3, 3, 3, 102, 8, 3, 3, 3, 104, 8, 3, 1, 4, 1, 4, 1, 4,
		1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5,
		1, 5, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 3, 6, 128, 8, 6, 1, 7, 1, 7,
		1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9,
		1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 11, 1,
		11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 12, 1, 12, 1, 12, 5, 12, 163,
		8, 12, 10, 12, 12, 12, 166, 9, 12, 1, 13, 1, 13, 1, 14, 1, 14, 1, 14, 1,
		14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 3, 14, 180, 8, 14, 1, 15,
		1, 15, 1, 15, 1, 15, 3, 15, 186, 8, 15, 1, 16, 1, 16, 1, 16, 1, 16, 5,
		16, 192, 8, 16, 10, 16, 12, 16, 195, 9, 16, 1, 17, 1, 17, 1, 17, 1, 17,
		1, 18, 1, 18, 1, 19, 1, 19, 1, 19, 1, 20, 1, 20, 1, 20, 5, 20, 209, 8,
		20, 10, 20, 12, 20, 212, 9, 20, 1, 21, 1, 21, 1, 22, 1, 22, 1, 23, 1, 23,
		1, 24, 1, 24, 1, 24, 0, 0, 25, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22,
		24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 0, 3, 1, 0, 24, 25,
		1, 0, 28, 32, 1, 0, 39, 42, 217, 0, 60, 1, 0, 0, 0, 2, 62, 1, 0, 0, 0,
		4, 77, 1, 0, 0, 0, 6, 103, 1, 0, 0, 0, 8, 105, 1, 0, 0, 0, 10, 113, 1,
		0, 0, 0, 12, 124, 1, 0, 0, 0, 14, 129, 1, 0, 0, 0, 16, 133, 1, 0, 0, 0,
		18, 138, 1, 0, 0, 0, 20, 145, 1, 0, 0, 0, 22, 152, 1, 0, 0, 0, 24, 159,
		1, 0, 0, 0, 26, 167, 1, 0, 0, 0, 28, 179, 1, 0, 0, 0, 30, 181, 1, 0, 0,
		0, 32, 187, 1, 0, 0, 0, 34, 196, 1, 0, 0, 0, 36, 200, 1, 0, 0, 0, 38, 202,
		1, 0, 0, 0, 40, 205, 1, 0, 0, 0, 42, 213, 1, 0, 0, 0, 44, 215, 1, 0, 0,
		0, 46, 217, 1, 0, 0, 0, 48, 219, 1, 0, 0, 0, 50, 61, 3, 2, 1, 0, 51, 61,
		3, 6, 3, 0, 52, 61, 3, 8, 4, 0, 53, 61, 3, 10, 5, 0, 54, 61, 3, 12, 6,
		0, 55, 61, 3, 14, 7, 0, 56, 61, 3, 16, 8, 0, 57, 61, 3, 18, 9, 0, 58, 61,
		3, 20, 10, 0, 59, 61, 3, 22, 11, 0, 60, 50, 1, 0, 0, 0, 60, 51, 1, 0, 0,
		0, 60, 52, 1, 0, 0, 0, 60, 53, 1, 0, 0, 0, 60, 54, 1, 0, 0, 0, 60, 55,
		1, 0, 0, 0, 60, 56, 1, 0, 0, 0, 60, 57, 1, 0, 0, 0, 60, 58, 1, 0, 0, 0,
		60, 59, 1, 0, 0, 0, 61, 1, 1, 0, 0, 0, 62, 63, 5, 1, 0, 0, 63, 64, 5, 3,
		0, 0, 64, 65, 3, 40, 20, 0, 65, 66, 5, 33, 0, 0, 66, 67, 3, 42, 21, 0,
		67, 68, 5, 34, 0, 0, 68, 69, 5, 9, 0, 0, 69, 74, 3, 4, 2, 0, 70, 71, 5,
		35, 0, 0, 71, 73, 3, 4, 2, 0, 72, 70, 1, 0, 0, 0, 73, 76, 1, 0, 0, 0, 74,
		72, 1, 0, 0, 0, 74, 75, 1, 0, 0, 0, 75, 3, 1, 0, 0, 0, 76, 74, 1, 0, 0,
		0, 77, 78, 5, 33, 0, 0, 78, 79, 3, 46, 23, 0, 79, 80, 5, 35, 0, 0, 80,
		81, 3, 48, 24, 0, 81, 82, 5, 34, 0, 0, 82, 5, 1, 0, 0, 0, 83, 84, 5, 4,
		0, 0, 84, 85, 3, 42, 21, 0, 85, 86, 5, 6, 0, 0, 86, 88, 3, 40, 20, 0, 87,
		89, 3, 32, 16, 0, 88, 87, 1, 0, 0, 0, 88, 89, 1, 0, 0, 0, 89, 91, 1, 0,
		0, 0, 90, 92, 3, 38, 19, 0, 91, 90, 1, 0, 0, 0, 91, 92, 1, 0, 0, 0, 92,
		104, 1, 0, 0, 0, 93, 94, 5, 4, 0, 0, 94, 95, 5, 5, 0, 0, 95, 96, 5, 33,
		0, 0, 96, 97, 3, 42, 21, 0, 97, 98, 5, 34, 0, 0, 98, 99, 5, 6, 0, 0, 99,
		101, 3, 40, 20, 0, 100, 102, 3, 32, 16, 0, 101, 100, 1, 0, 0, 0, 101, 102,
		1, 0, 0, 0, 102, 104, 1, 0, 0, 0, 103, 83, 1, 0, 0, 0, 103, 93, 1, 0, 0,
		0, 104, 7, 1, 0, 0, 0, 105, 106, 5, 2, 0, 0, 106, 107, 5, 6, 0, 0, 107,
		108, 3, 40, 20, 0, 108, 109, 5, 33, 0, 0, 109, 110, 3, 42, 21, 0, 110,
		111, 5, 34, 0, 0, 111, 112, 3, 32, 16, 0, 112, 9, 1, 0, 0, 0, 113, 114,
		5, 12, 0, 0, 114, 115, 5, 13, 0, 0, 115, 116, 3, 40, 20, 0, 116, 117, 5,
		33, 0, 0, 117, 118, 3, 42, 21, 0, 118, 119, 5, 34, 0, 0, 119, 120, 5, 14,
		0, 0, 120, 121, 5, 15, 0, 0, 121, 122, 5, 32, 0, 0, 122, 123, 3, 44, 22,
		0, 123, 11, 1, 0, 0, 0, 124, 125, 5, 16, 0, 0, 125, 127, 5, 17, 0, 0, 126,
		128, 3, 30, 15, 0, 127, 126, 1, 0, 0, 0, 127, 128, 1, 0, 0, 0, 128, 13,
		1, 0, 0, 0, 129, 130, 5, 16, 0, 0, 130, 131, 5, 13, 0, 0, 131, 132, 3,
		40, 20, 0, 132, 15, 1, 0, 0, 0, 133, 134, 5, 12, 0, 0, 134, 135, 5, 18,
		0, 0, 135, 136, 5, 43, 0, 0, 136, 137, 5, 42, 0, 0, 137, 17, 1, 0, 0, 0,
		138, 139, 5, 19, 0, 0, 139, 140, 5, 18, 0, 0, 140, 141, 5, 43, 0, 0, 141,
		142, 5, 20, 0, 0, 142, 143, 5, 21, 0, 0, 143, 144, 5, 42, 0, 0, 144, 19,
		1, 0, 0, 0, 145, 146, 5, 22, 0, 0, 146, 147, 3, 24, 12, 0, 147, 148, 5,
		26, 0, 0, 148, 149, 3, 28, 14, 0, 149, 150, 5, 27, 0, 0, 150, 151, 5, 43,
		0, 0, 151, 21, 1, 0, 0, 0, 152, 153, 5, 23, 0, 0, 153, 154, 3, 24, 12,
		0, 154, 155, 5, 26, 0, 0, 155, 156, 3, 28, 14, 0, 156, 157, 5, 6, 0, 0,
		157, 158, 5, 43, 0, 0, 158, 23, 1, 0, 0, 0, 159, 164, 3, 26, 13, 0, 160,
		161, 5, 35, 0, 0, 161, 163, 3, 26, 13, 0, 162, 160, 1, 0, 0, 0, 163, 166,
		1, 0, 0, 0, 164, 162, 1, 0, 0, 0, 164, 165, 1, 0, 0, 0, 165, 25, 1, 0,
		0, 0, 166, 164, 1, 0, 0, 0, 167, 168, 7, 0, 0, 0, 168, 27, 1, 0, 0, 0,
		169, 170, 3, 40, 20, 0, 170, 171, 5, 36, 0, 0, 171, 172, 5, 38, 0, 0, 172,
		173, 5, 38, 0, 0, 173, 180, 1, 0, 0, 0, 174, 175, 3, 40, 20, 0, 175, 176,
		5, 36, 0, 0, 176, 177, 5, 38, 0, 0, 177, 180, 1, 0, 0, 0, 178, 180, 3,
		40, 20, 0, 179, 169, 1, 0, 0, 0, 179, 174, 1, 0, 0, 0, 179, 178, 1, 0,
		0, 0, 180, 29, 1, 0, 0, 0, 181, 185, 3, 40, 20, 0, 182, 183, 5, 36, 0,
		0, 183, 184, 5, 38, 0, 0, 184, 186, 5, 38, 0, 0, 185, 182, 1, 0, 0, 0,
		185, 186, 1, 0, 0, 0, 186, 31, 1, 0, 0, 0, 187, 188, 5, 7, 0, 0, 188, 193,
		3, 34, 17, 0, 189, 190, 5, 8, 0, 0, 190, 192, 3, 34, 17, 0, 191, 189, 1,
		0, 0, 0, 192, 195, 1, 0, 0, 0, 193, 191, 1, 0, 0, 0, 193, 194, 1, 0, 0,
		0, 194, 33, 1, 0, 0, 0, 195, 193, 1, 0, 0, 0, 196, 197, 5, 11, 0, 0, 197,
		198, 3, 36, 18, 0, 198, 199, 5, 40, 0, 0, 199, 35, 1, 0, 0, 0, 200, 201,
		7, 1, 0, 0, 201, 37, 1, 0, 0, 0, 202, 203, 5, 10, 0, 0, 203, 204, 5, 40,
		0, 0, 204, 39, 1, 0, 0, 0, 205, 210, 5, 43, 0, 0, 206, 207, 5, 36, 0, 0,
		207, 209, 5, 43, 0, 0, 208, 206, 1, 0, 0, 0, 209, 212, 1, 0, 0, 0, 210,
		208, 1, 0, 0, 0, 210, 211, 1, 0, 0, 0, 211, 41, 1, 0, 0, 0, 212, 210, 1,
		0, 0, 0, 213, 214, 5, 43, 0, 0, 214, 43, 1, 0, 0, 0, 215, 216, 5, 43, 0,
		0, 216, 45, 1, 0, 0, 0, 217, 218, 5, 40, 0, 0, 218, 47, 1, 0, 0, 0, 219,
		220, 7, 2, 0, 0, 220, 49, 1, 0, 0, 0, 12, 60, 74, 88, 91, 101, 103, 127,
		164, 179, 185, 193, 210,
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
	TideSQLParserUSER       = 18
	TideSQLParserALTER      = 19
	TideSQLParserSET        = 20
	TideSQLParserPASSWORD   = 21
	TideSQLParserGRANT      = 22
	TideSQLParserREVOKE     = 23
	TideSQLParserREAD       = 24
	TideSQLParserWRITE      = 25
	TideSQLParserON         = 26
	TideSQLParserTO         = 27
	TideSQLParserGTE        = 28
	TideSQLParserLTE        = 29
	TideSQLParserGT         = 30
	TideSQLParserLT         = 31
	TideSQLParserEQ         = 32
	TideSQLParserLPAREN     = 33
	TideSQLParserRPAREN     = 34
	TideSQLParserCOMMA      = 35
	TideSQLParserDOT        = 36
	TideSQLParserSEMI       = 37
	TideSQLParserSTAR       = 38
	TideSQLParserBOOLEAN    = 39
	TideSQLParserINTEGER    = 40
	TideSQLParserFLOAT      = 41
	TideSQLParserSTRING     = 42
	TideSQLParserIDENTIFIER = 43
	TideSQLParserWS         = 44
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
	TideSQLParserRULE_createUserStmt       = 8
	TideSQLParserRULE_alterUserStmt        = 9
	TideSQLParserRULE_grantStmt            = 10
	TideSQLParserRULE_revokeStmt           = 11
	TideSQLParserRULE_privList             = 12
	TideSQLParserRULE_priv                 = 13
	TideSQLParserRULE_grantPath            = 14
	TideSQLParserRULE_showPattern          = 15
	TideSQLParserRULE_whereClause          = 16
	TideSQLParserRULE_timePredicate        = 17
	TideSQLParserRULE_cmpOp                = 18
	TideSQLParserRULE_limitClause          = 19
	TideSQLParserRULE_path                 = 20
	TideSQLParserRULE_measurement          = 21
	TideSQLParserRULE_dataTypeName         = 22
	TideSQLParserRULE_timestamp            = 23
	TideSQLParserRULE_value                = 24
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
	CreateUserStmt() ICreateUserStmtContext
	AlterUserStmt() IAlterUserStmtContext
	GrantStmt() IGrantStmtContext
	RevokeStmt() IRevokeStmtContext

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

func (s *StatementContext) CreateUserStmt() ICreateUserStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICreateUserStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICreateUserStmtContext)
}

func (s *StatementContext) AlterUserStmt() IAlterUserStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAlterUserStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAlterUserStmtContext)
}

func (s *StatementContext) GrantStmt() IGrantStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IGrantStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IGrantStmtContext)
}

func (s *StatementContext) RevokeStmt() IRevokeStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRevokeStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRevokeStmtContext)
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
	p.SetState(60)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 0, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(50)
			p.InsertStmt()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(51)
			p.SelectStmt()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(52)
			p.DeleteStmt()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(53)
			p.CreateTimeseriesStmt()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(54)
			p.ShowDevicesStmt()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(55)
			p.ShowTimeseriesStmt()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(56)
			p.CreateUserStmt()
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(57)
			p.AlterUserStmt()
		}

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(58)
			p.GrantStmt()
		}

	case 10:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(59)
			p.RevokeStmt()
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
		p.SetState(62)
		p.Match(TideSQLParserINSERT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(63)
		p.Match(TideSQLParserINTO)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(64)
		p.Path()
	}
	{
		p.SetState(65)
		p.Match(TideSQLParserLPAREN)
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
		p.Match(TideSQLParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(68)
		p.Match(TideSQLParserVALUES)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(69)
		p.ValueRow()
	}
	p.SetState(74)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == TideSQLParserCOMMA {
		{
			p.SetState(70)
			p.Match(TideSQLParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(71)
			p.ValueRow()
		}

		p.SetState(76)
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
		p.SetState(77)
		p.Match(TideSQLParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(78)
		p.Timestamp()
	}
	{
		p.SetState(79)
		p.Match(TideSQLParserCOMMA)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(80)
		p.Value()
	}
	{
		p.SetState(81)
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

	p.SetState(103)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 5, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(83)
			p.Match(TideSQLParserSELECT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(84)
			p.Measurement()
		}
		{
			p.SetState(85)
			p.Match(TideSQLParserFROM)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(86)
			p.Path()
		}
		p.SetState(88)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == TideSQLParserWHERE {
			{
				p.SetState(87)
				p.WhereClause()
			}

		}
		p.SetState(91)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == TideSQLParserLIMIT {
			{
				p.SetState(90)
				p.LimitClause()
			}

		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(93)
			p.Match(TideSQLParserSELECT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(94)
			p.Match(TideSQLParserCOUNT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(95)
			p.Match(TideSQLParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(96)
			p.Measurement()
		}
		{
			p.SetState(97)
			p.Match(TideSQLParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(98)
			p.Match(TideSQLParserFROM)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(99)
			p.Path()
		}
		p.SetState(101)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == TideSQLParserWHERE {
			{
				p.SetState(100)
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
		p.SetState(105)
		p.Match(TideSQLParserDELETE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(106)
		p.Match(TideSQLParserFROM)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(107)
		p.Path()
	}
	{
		p.SetState(108)
		p.Match(TideSQLParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(109)
		p.Measurement()
	}
	{
		p.SetState(110)
		p.Match(TideSQLParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(111)
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
		p.SetState(113)
		p.Match(TideSQLParserCREATE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(114)
		p.Match(TideSQLParserTIMESERIES)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(115)
		p.Path()
	}
	{
		p.SetState(116)
		p.Match(TideSQLParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(117)
		p.Measurement()
	}
	{
		p.SetState(118)
		p.Match(TideSQLParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(119)
		p.Match(TideSQLParserWITH)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(120)
		p.Match(TideSQLParserDATATYPE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(121)
		p.Match(TideSQLParserEQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(122)
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
		p.SetState(124)
		p.Match(TideSQLParserSHOW)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(125)
		p.Match(TideSQLParserDEVICES)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(127)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == TideSQLParserIDENTIFIER {
		{
			p.SetState(126)
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
		p.SetState(129)
		p.Match(TideSQLParserSHOW)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(130)
		p.Match(TideSQLParserTIMESERIES)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(131)
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

// ICreateUserStmtContext is an interface to support dynamic dispatch.
type ICreateUserStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	CREATE() antlr.TerminalNode
	USER() antlr.TerminalNode
	IDENTIFIER() antlr.TerminalNode
	STRING() antlr.TerminalNode

	// IsCreateUserStmtContext differentiates from other interfaces.
	IsCreateUserStmtContext()
}

type CreateUserStmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCreateUserStmtContext() *CreateUserStmtContext {
	var p = new(CreateUserStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TideSQLParserRULE_createUserStmt
	return p
}

func InitEmptyCreateUserStmtContext(p *CreateUserStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TideSQLParserRULE_createUserStmt
}

func (*CreateUserStmtContext) IsCreateUserStmtContext() {}

func NewCreateUserStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CreateUserStmtContext {
	var p = new(CreateUserStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TideSQLParserRULE_createUserStmt

	return p
}

func (s *CreateUserStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *CreateUserStmtContext) CREATE() antlr.TerminalNode {
	return s.GetToken(TideSQLParserCREATE, 0)
}

func (s *CreateUserStmtContext) USER() antlr.TerminalNode {
	return s.GetToken(TideSQLParserUSER, 0)
}

func (s *CreateUserStmtContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(TideSQLParserIDENTIFIER, 0)
}

func (s *CreateUserStmtContext) STRING() antlr.TerminalNode {
	return s.GetToken(TideSQLParserSTRING, 0)
}

func (s *CreateUserStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CreateUserStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CreateUserStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TideSQLListener); ok {
		listenerT.EnterCreateUserStmt(s)
	}
}

func (s *CreateUserStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TideSQLListener); ok {
		listenerT.ExitCreateUserStmt(s)
	}
}

func (s *CreateUserStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TideSQLVisitor:
		return t.VisitCreateUserStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TideSQLParser) CreateUserStmt() (localctx ICreateUserStmtContext) {
	localctx = NewCreateUserStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, TideSQLParserRULE_createUserStmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(133)
		p.Match(TideSQLParserCREATE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(134)
		p.Match(TideSQLParserUSER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(135)
		p.Match(TideSQLParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(136)
		p.Match(TideSQLParserSTRING)
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

// IAlterUserStmtContext is an interface to support dynamic dispatch.
type IAlterUserStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ALTER() antlr.TerminalNode
	USER() antlr.TerminalNode
	IDENTIFIER() antlr.TerminalNode
	SET() antlr.TerminalNode
	PASSWORD() antlr.TerminalNode
	STRING() antlr.TerminalNode

	// IsAlterUserStmtContext differentiates from other interfaces.
	IsAlterUserStmtContext()
}

type AlterUserStmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAlterUserStmtContext() *AlterUserStmtContext {
	var p = new(AlterUserStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TideSQLParserRULE_alterUserStmt
	return p
}

func InitEmptyAlterUserStmtContext(p *AlterUserStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TideSQLParserRULE_alterUserStmt
}

func (*AlterUserStmtContext) IsAlterUserStmtContext() {}

func NewAlterUserStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AlterUserStmtContext {
	var p = new(AlterUserStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TideSQLParserRULE_alterUserStmt

	return p
}

func (s *AlterUserStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *AlterUserStmtContext) ALTER() antlr.TerminalNode {
	return s.GetToken(TideSQLParserALTER, 0)
}

func (s *AlterUserStmtContext) USER() antlr.TerminalNode {
	return s.GetToken(TideSQLParserUSER, 0)
}

func (s *AlterUserStmtContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(TideSQLParserIDENTIFIER, 0)
}

func (s *AlterUserStmtContext) SET() antlr.TerminalNode {
	return s.GetToken(TideSQLParserSET, 0)
}

func (s *AlterUserStmtContext) PASSWORD() antlr.TerminalNode {
	return s.GetToken(TideSQLParserPASSWORD, 0)
}

func (s *AlterUserStmtContext) STRING() antlr.TerminalNode {
	return s.GetToken(TideSQLParserSTRING, 0)
}

func (s *AlterUserStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AlterUserStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AlterUserStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TideSQLListener); ok {
		listenerT.EnterAlterUserStmt(s)
	}
}

func (s *AlterUserStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TideSQLListener); ok {
		listenerT.ExitAlterUserStmt(s)
	}
}

func (s *AlterUserStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TideSQLVisitor:
		return t.VisitAlterUserStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TideSQLParser) AlterUserStmt() (localctx IAlterUserStmtContext) {
	localctx = NewAlterUserStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, TideSQLParserRULE_alterUserStmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(138)
		p.Match(TideSQLParserALTER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(139)
		p.Match(TideSQLParserUSER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(140)
		p.Match(TideSQLParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(141)
		p.Match(TideSQLParserSET)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(142)
		p.Match(TideSQLParserPASSWORD)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(143)
		p.Match(TideSQLParserSTRING)
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

// IGrantStmtContext is an interface to support dynamic dispatch.
type IGrantStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	GRANT() antlr.TerminalNode
	PrivList() IPrivListContext
	ON() antlr.TerminalNode
	GrantPath() IGrantPathContext
	TO() antlr.TerminalNode
	IDENTIFIER() antlr.TerminalNode

	// IsGrantStmtContext differentiates from other interfaces.
	IsGrantStmtContext()
}

type GrantStmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGrantStmtContext() *GrantStmtContext {
	var p = new(GrantStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TideSQLParserRULE_grantStmt
	return p
}

func InitEmptyGrantStmtContext(p *GrantStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TideSQLParserRULE_grantStmt
}

func (*GrantStmtContext) IsGrantStmtContext() {}

func NewGrantStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GrantStmtContext {
	var p = new(GrantStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TideSQLParserRULE_grantStmt

	return p
}

func (s *GrantStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *GrantStmtContext) GRANT() antlr.TerminalNode {
	return s.GetToken(TideSQLParserGRANT, 0)
}

func (s *GrantStmtContext) PrivList() IPrivListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrivListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrivListContext)
}

func (s *GrantStmtContext) ON() antlr.TerminalNode {
	return s.GetToken(TideSQLParserON, 0)
}

func (s *GrantStmtContext) GrantPath() IGrantPathContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IGrantPathContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IGrantPathContext)
}

func (s *GrantStmtContext) TO() antlr.TerminalNode {
	return s.GetToken(TideSQLParserTO, 0)
}

func (s *GrantStmtContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(TideSQLParserIDENTIFIER, 0)
}

func (s *GrantStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GrantStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GrantStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TideSQLListener); ok {
		listenerT.EnterGrantStmt(s)
	}
}

func (s *GrantStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TideSQLListener); ok {
		listenerT.ExitGrantStmt(s)
	}
}

func (s *GrantStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TideSQLVisitor:
		return t.VisitGrantStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TideSQLParser) GrantStmt() (localctx IGrantStmtContext) {
	localctx = NewGrantStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, TideSQLParserRULE_grantStmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(145)
		p.Match(TideSQLParserGRANT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(146)
		p.PrivList()
	}
	{
		p.SetState(147)
		p.Match(TideSQLParserON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(148)
		p.GrantPath()
	}
	{
		p.SetState(149)
		p.Match(TideSQLParserTO)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(150)
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

// IRevokeStmtContext is an interface to support dynamic dispatch.
type IRevokeStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	REVOKE() antlr.TerminalNode
	PrivList() IPrivListContext
	ON() antlr.TerminalNode
	GrantPath() IGrantPathContext
	FROM() antlr.TerminalNode
	IDENTIFIER() antlr.TerminalNode

	// IsRevokeStmtContext differentiates from other interfaces.
	IsRevokeStmtContext()
}

type RevokeStmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRevokeStmtContext() *RevokeStmtContext {
	var p = new(RevokeStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TideSQLParserRULE_revokeStmt
	return p
}

func InitEmptyRevokeStmtContext(p *RevokeStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TideSQLParserRULE_revokeStmt
}

func (*RevokeStmtContext) IsRevokeStmtContext() {}

func NewRevokeStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RevokeStmtContext {
	var p = new(RevokeStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TideSQLParserRULE_revokeStmt

	return p
}

func (s *RevokeStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *RevokeStmtContext) REVOKE() antlr.TerminalNode {
	return s.GetToken(TideSQLParserREVOKE, 0)
}

func (s *RevokeStmtContext) PrivList() IPrivListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrivListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrivListContext)
}

func (s *RevokeStmtContext) ON() antlr.TerminalNode {
	return s.GetToken(TideSQLParserON, 0)
}

func (s *RevokeStmtContext) GrantPath() IGrantPathContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IGrantPathContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IGrantPathContext)
}

func (s *RevokeStmtContext) FROM() antlr.TerminalNode {
	return s.GetToken(TideSQLParserFROM, 0)
}

func (s *RevokeStmtContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(TideSQLParserIDENTIFIER, 0)
}

func (s *RevokeStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RevokeStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RevokeStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TideSQLListener); ok {
		listenerT.EnterRevokeStmt(s)
	}
}

func (s *RevokeStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TideSQLListener); ok {
		listenerT.ExitRevokeStmt(s)
	}
}

func (s *RevokeStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TideSQLVisitor:
		return t.VisitRevokeStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TideSQLParser) RevokeStmt() (localctx IRevokeStmtContext) {
	localctx = NewRevokeStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, TideSQLParserRULE_revokeStmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(152)
		p.Match(TideSQLParserREVOKE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(153)
		p.PrivList()
	}
	{
		p.SetState(154)
		p.Match(TideSQLParserON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(155)
		p.GrantPath()
	}
	{
		p.SetState(156)
		p.Match(TideSQLParserFROM)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(157)
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

// IPrivListContext is an interface to support dynamic dispatch.
type IPrivListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllPriv() []IPrivContext
	Priv(i int) IPrivContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsPrivListContext differentiates from other interfaces.
	IsPrivListContext()
}

type PrivListContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrivListContext() *PrivListContext {
	var p = new(PrivListContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TideSQLParserRULE_privList
	return p
}

func InitEmptyPrivListContext(p *PrivListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TideSQLParserRULE_privList
}

func (*PrivListContext) IsPrivListContext() {}

func NewPrivListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrivListContext {
	var p = new(PrivListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TideSQLParserRULE_privList

	return p
}

func (s *PrivListContext) GetParser() antlr.Parser { return s.parser }

func (s *PrivListContext) AllPriv() []IPrivContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IPrivContext); ok {
			len++
		}
	}

	tst := make([]IPrivContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IPrivContext); ok {
			tst[i] = t.(IPrivContext)
			i++
		}
	}

	return tst
}

func (s *PrivListContext) Priv(i int) IPrivContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrivContext); ok {
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

	return t.(IPrivContext)
}

func (s *PrivListContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(TideSQLParserCOMMA)
}

func (s *PrivListContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(TideSQLParserCOMMA, i)
}

func (s *PrivListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrivListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrivListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TideSQLListener); ok {
		listenerT.EnterPrivList(s)
	}
}

func (s *PrivListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TideSQLListener); ok {
		listenerT.ExitPrivList(s)
	}
}

func (s *PrivListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TideSQLVisitor:
		return t.VisitPrivList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TideSQLParser) PrivList() (localctx IPrivListContext) {
	localctx = NewPrivListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, TideSQLParserRULE_privList)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(159)
		p.Priv()
	}
	p.SetState(164)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == TideSQLParserCOMMA {
		{
			p.SetState(160)
			p.Match(TideSQLParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(161)
			p.Priv()
		}

		p.SetState(166)
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

// IPrivContext is an interface to support dynamic dispatch.
type IPrivContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	READ() antlr.TerminalNode
	WRITE() antlr.TerminalNode

	// IsPrivContext differentiates from other interfaces.
	IsPrivContext()
}

type PrivContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrivContext() *PrivContext {
	var p = new(PrivContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TideSQLParserRULE_priv
	return p
}

func InitEmptyPrivContext(p *PrivContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TideSQLParserRULE_priv
}

func (*PrivContext) IsPrivContext() {}

func NewPrivContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrivContext {
	var p = new(PrivContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TideSQLParserRULE_priv

	return p
}

func (s *PrivContext) GetParser() antlr.Parser { return s.parser }

func (s *PrivContext) READ() antlr.TerminalNode {
	return s.GetToken(TideSQLParserREAD, 0)
}

func (s *PrivContext) WRITE() antlr.TerminalNode {
	return s.GetToken(TideSQLParserWRITE, 0)
}

func (s *PrivContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrivContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrivContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TideSQLListener); ok {
		listenerT.EnterPriv(s)
	}
}

func (s *PrivContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TideSQLListener); ok {
		listenerT.ExitPriv(s)
	}
}

func (s *PrivContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TideSQLVisitor:
		return t.VisitPriv(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TideSQLParser) Priv() (localctx IPrivContext) {
	localctx = NewPrivContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, TideSQLParserRULE_priv)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(167)
		_la = p.GetTokenStream().LA(1)

		if !(_la == TideSQLParserREAD || _la == TideSQLParserWRITE) {
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

// IGrantPathContext is an interface to support dynamic dispatch.
type IGrantPathContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Path() IPathContext
	DOT() antlr.TerminalNode
	AllSTAR() []antlr.TerminalNode
	STAR(i int) antlr.TerminalNode

	// IsGrantPathContext differentiates from other interfaces.
	IsGrantPathContext()
}

type GrantPathContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGrantPathContext() *GrantPathContext {
	var p = new(GrantPathContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TideSQLParserRULE_grantPath
	return p
}

func InitEmptyGrantPathContext(p *GrantPathContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TideSQLParserRULE_grantPath
}

func (*GrantPathContext) IsGrantPathContext() {}

func NewGrantPathContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GrantPathContext {
	var p = new(GrantPathContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TideSQLParserRULE_grantPath

	return p
}

func (s *GrantPathContext) GetParser() antlr.Parser { return s.parser }

func (s *GrantPathContext) Path() IPathContext {
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

func (s *GrantPathContext) DOT() antlr.TerminalNode {
	return s.GetToken(TideSQLParserDOT, 0)
}

func (s *GrantPathContext) AllSTAR() []antlr.TerminalNode {
	return s.GetTokens(TideSQLParserSTAR)
}

func (s *GrantPathContext) STAR(i int) antlr.TerminalNode {
	return s.GetToken(TideSQLParserSTAR, i)
}

func (s *GrantPathContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GrantPathContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GrantPathContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TideSQLListener); ok {
		listenerT.EnterGrantPath(s)
	}
}

func (s *GrantPathContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TideSQLListener); ok {
		listenerT.ExitGrantPath(s)
	}
}

func (s *GrantPathContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TideSQLVisitor:
		return t.VisitGrantPath(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TideSQLParser) GrantPath() (localctx IGrantPathContext) {
	localctx = NewGrantPathContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, TideSQLParserRULE_grantPath)
	p.SetState(179)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 8, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(169)
			p.Path()
		}
		{
			p.SetState(170)
			p.Match(TideSQLParserDOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(171)
			p.Match(TideSQLParserSTAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(172)
			p.Match(TideSQLParserSTAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(174)
			p.Path()
		}
		{
			p.SetState(175)
			p.Match(TideSQLParserDOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(176)
			p.Match(TideSQLParserSTAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(178)
			p.Path()
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
	p.EnterRule(localctx, 30, TideSQLParserRULE_showPattern)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(181)
		p.Path()
	}
	p.SetState(185)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == TideSQLParserDOT {
		{
			p.SetState(182)
			p.Match(TideSQLParserDOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(183)
			p.Match(TideSQLParserSTAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(184)
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
	p.EnterRule(localctx, 32, TideSQLParserRULE_whereClause)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(187)
		p.Match(TideSQLParserWHERE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(188)
		p.TimePredicate()
	}
	p.SetState(193)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == TideSQLParserAND {
		{
			p.SetState(189)
			p.Match(TideSQLParserAND)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(190)
			p.TimePredicate()
		}

		p.SetState(195)
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
	p.EnterRule(localctx, 34, TideSQLParserRULE_timePredicate)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(196)
		p.Match(TideSQLParserTIME)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(197)
		p.CmpOp()
	}
	{
		p.SetState(198)
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
	p.EnterRule(localctx, 36, TideSQLParserRULE_cmpOp)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(200)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&8321499136) != 0) {
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
	p.EnterRule(localctx, 38, TideSQLParserRULE_limitClause)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(202)
		p.Match(TideSQLParserLIMIT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(203)
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
	p.EnterRule(localctx, 40, TideSQLParserRULE_path)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(205)
		p.Match(TideSQLParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(210)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 11, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(206)
				p.Match(TideSQLParserDOT)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(207)
				p.Match(TideSQLParserIDENTIFIER)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		p.SetState(212)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 11, p.GetParserRuleContext())
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
	p.EnterRule(localctx, 42, TideSQLParserRULE_measurement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(213)
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
	p.EnterRule(localctx, 44, TideSQLParserRULE_dataTypeName)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(215)
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
	p.EnterRule(localctx, 46, TideSQLParserRULE_timestamp)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(217)
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
	p.EnterRule(localctx, 48, TideSQLParserRULE_value)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(219)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&8246337208320) != 0) {
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
