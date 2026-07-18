// Code generated from TqParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package tqgrammar // TqParser
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

type TqParser struct {
	*antlr.BaseParser
}

var TqParserParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func tqparserParserInit() {
	staticData := &TqParserParserStaticData
	staticData.LiteralNames = []string{
		"", "'select'", "'drop'", "'where'", "'derive'", "'rename'", "'sort'",
		"'distinct'", "'limit'", "'offset'", "'group'", "'as'", "'{'", "'}'",
		"", "", "'>='", "'<='", "'<>'", "'>'", "'<'", "'TRUE'", "'FALSE'", "",
		"'='", "'('", "')'", "':'", "','", "'$'", "'*'", "'+'", "'-'", "'/'",
		"'%'", "'^'", "'&'", "'!'", "'|'",
	}
	staticData.SymbolicNames = []string{
		"", "SELECT", "DROP", "WHERE", "DERIVE", "RENAME", "SORT", "DISTINCT",
		"LIMIT", "OFFSET", "GROUP", "AS", "LBRACE", "RBRACE", "COLUMN", "WS",
		"GE", "LE", "NE", "GT", "LT", "TRUE", "FALSE", "ERRORCONST", "EQ", "LPAREN",
		"RPAREN", "COLON", "COMMA", "DOLLAR", "STAR", "PLUS", "DASH", "SLASH",
		"PERCENT", "CARET", "AMP", "BANG", "PIPE", "NUMBER", "COL", "NAME",
		"STRING",
	}
	staticData.RuleNames = []string{
		"program", "stage", "assignment", "renamePair", "sortKey", "columnList",
		"columnItem", "columnName", "bareName", "verbKeyword", "tqExpr", "tqFunctionCall",
		"tqArgList", "expression", "functionCall", "argList", "reference", "sheetQualifier",
		"cellRef",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 42, 280, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 1, 0, 1, 0, 1, 0, 5, 0, 42, 8,
		0, 10, 0, 12, 0, 45, 9, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 5, 1, 59, 8, 1, 10, 1, 12, 1, 62, 9, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 5, 1, 68, 8, 1, 10, 1, 12, 1, 71, 9, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 5, 1, 77, 8, 1, 10, 1, 12, 1, 80, 9, 1, 1, 1, 1, 1, 3, 1, 84,
		8, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 5, 1,
		96, 8, 1, 10, 1, 12, 1, 99, 9, 1, 1, 1, 1, 1, 3, 1, 103, 8, 1, 1, 2, 1,
		2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 4, 3, 4, 114, 8, 4, 1, 4, 1,
		4, 1, 5, 1, 5, 1, 5, 5, 5, 121, 8, 5, 10, 5, 12, 5, 124, 9, 5, 1, 6, 1,
		6, 1, 6, 3, 6, 129, 8, 6, 1, 7, 1, 7, 3, 7, 133, 8, 7, 1, 8, 1, 8, 1, 8,
		3, 8, 138, 8, 8, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10,
		1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 3, 10, 156, 8,
		10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10,
		1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 5, 10, 175, 8, 10, 10,
		10, 12, 10, 178, 9, 10, 1, 11, 1, 11, 1, 11, 3, 11, 183, 8, 11, 1, 11,
		3, 11, 186, 8, 11, 1, 11, 1, 11, 3, 11, 190, 8, 11, 1, 11, 1, 11, 1, 12,
		1, 12, 1, 12, 5, 12, 197, 8, 12, 10, 12, 12, 12, 200, 9, 12, 1, 13, 1,
		13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13,
		1, 13, 3, 13, 215, 8, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1,
		13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13,
		1, 13, 1, 13, 1, 13, 5, 13, 237, 8, 13, 10, 13, 12, 13, 240, 9, 13, 1,
		14, 1, 14, 3, 14, 244, 8, 14, 1, 14, 1, 14, 3, 14, 248, 8, 14, 1, 14, 1,
		14, 1, 15, 1, 15, 1, 15, 5, 15, 255, 8, 15, 10, 15, 12, 15, 258, 9, 15,
		1, 16, 3, 16, 261, 8, 16, 1, 16, 1, 16, 1, 16, 3, 16, 266, 8, 16, 1, 17,
		1, 17, 1, 17, 1, 18, 3, 18, 272, 8, 18, 1, 18, 1, 18, 3, 18, 276, 8, 18,
		1, 18, 1, 18, 1, 18, 0, 2, 20, 26, 19, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18,
		20, 22, 24, 26, 28, 30, 32, 34, 36, 0, 6, 1, 0, 1, 11, 1, 0, 31, 32, 1,
		0, 21, 22, 2, 0, 30, 30, 33, 33, 2, 0, 16, 20, 24, 24, 1, 0, 40, 41, 322,
		0, 38, 1, 0, 0, 0, 2, 102, 1, 0, 0, 0, 4, 104, 1, 0, 0, 0, 6, 108, 1, 0,
		0, 0, 8, 113, 1, 0, 0, 0, 10, 117, 1, 0, 0, 0, 12, 128, 1, 0, 0, 0, 14,
		132, 1, 0, 0, 0, 16, 137, 1, 0, 0, 0, 18, 139, 1, 0, 0, 0, 20, 155, 1,
		0, 0, 0, 22, 182, 1, 0, 0, 0, 24, 193, 1, 0, 0, 0, 26, 214, 1, 0, 0, 0,
		28, 241, 1, 0, 0, 0, 30, 251, 1, 0, 0, 0, 32, 260, 1, 0, 0, 0, 34, 267,
		1, 0, 0, 0, 36, 271, 1, 0, 0, 0, 38, 43, 3, 2, 1, 0, 39, 40, 5, 38, 0,
		0, 40, 42, 3, 2, 1, 0, 41, 39, 1, 0, 0, 0, 42, 45, 1, 0, 0, 0, 43, 41,
		1, 0, 0, 0, 43, 44, 1, 0, 0, 0, 44, 46, 1, 0, 0, 0, 45, 43, 1, 0, 0, 0,
		46, 47, 5, 0, 0, 1, 47, 1, 1, 0, 0, 0, 48, 49, 5, 1, 0, 0, 49, 103, 3,
		10, 5, 0, 50, 51, 5, 2, 0, 0, 51, 103, 3, 10, 5, 0, 52, 53, 5, 3, 0, 0,
		53, 103, 3, 20, 10, 0, 54, 55, 5, 4, 0, 0, 55, 60, 3, 4, 2, 0, 56, 57,
		5, 28, 0, 0, 57, 59, 3, 4, 2, 0, 58, 56, 1, 0, 0, 0, 59, 62, 1, 0, 0, 0,
		60, 58, 1, 0, 0, 0, 60, 61, 1, 0, 0, 0, 61, 103, 1, 0, 0, 0, 62, 60, 1,
		0, 0, 0, 63, 64, 5, 5, 0, 0, 64, 69, 3, 6, 3, 0, 65, 66, 5, 28, 0, 0, 66,
		68, 3, 6, 3, 0, 67, 65, 1, 0, 0, 0, 68, 71, 1, 0, 0, 0, 69, 67, 1, 0, 0,
		0, 69, 70, 1, 0, 0, 0, 70, 103, 1, 0, 0, 0, 71, 69, 1, 0, 0, 0, 72, 73,
		5, 6, 0, 0, 73, 78, 3, 8, 4, 0, 74, 75, 5, 28, 0, 0, 75, 77, 3, 8, 4, 0,
		76, 74, 1, 0, 0, 0, 77, 80, 1, 0, 0, 0, 78, 76, 1, 0, 0, 0, 78, 79, 1,
		0, 0, 0, 79, 103, 1, 0, 0, 0, 80, 78, 1, 0, 0, 0, 81, 83, 5, 7, 0, 0, 82,
		84, 3, 10, 5, 0, 83, 82, 1, 0, 0, 0, 83, 84, 1, 0, 0, 0, 84, 103, 1, 0,
		0, 0, 85, 86, 5, 8, 0, 0, 86, 103, 5, 39, 0, 0, 87, 88, 5, 9, 0, 0, 88,
		103, 5, 39, 0, 0, 89, 90, 5, 10, 0, 0, 90, 91, 3, 10, 5, 0, 91, 92, 5,
		12, 0, 0, 92, 97, 3, 4, 2, 0, 93, 94, 5, 28, 0, 0, 94, 96, 3, 4, 2, 0,
		95, 93, 1, 0, 0, 0, 96, 99, 1, 0, 0, 0, 97, 95, 1, 0, 0, 0, 97, 98, 1,
		0, 0, 0, 98, 100, 1, 0, 0, 0, 99, 97, 1, 0, 0, 0, 100, 101, 5, 13, 0, 0,
		101, 103, 1, 0, 0, 0, 102, 48, 1, 0, 0, 0, 102, 50, 1, 0, 0, 0, 102, 52,
		1, 0, 0, 0, 102, 54, 1, 0, 0, 0, 102, 63, 1, 0, 0, 0, 102, 72, 1, 0, 0,
		0, 102, 81, 1, 0, 0, 0, 102, 85, 1, 0, 0, 0, 102, 87, 1, 0, 0, 0, 102,
		89, 1, 0, 0, 0, 103, 3, 1, 0, 0, 0, 104, 105, 3, 14, 7, 0, 105, 106, 5,
		24, 0, 0, 106, 107, 3, 20, 10, 0, 107, 5, 1, 0, 0, 0, 108, 109, 3, 12,
		6, 0, 109, 110, 5, 11, 0, 0, 110, 111, 3, 14, 7, 0, 111, 7, 1, 0, 0, 0,
		112, 114, 5, 32, 0, 0, 113, 112, 1, 0, 0, 0, 113, 114, 1, 0, 0, 0, 114,
		115, 1, 0, 0, 0, 115, 116, 3, 12, 6, 0, 116, 9, 1, 0, 0, 0, 117, 122, 3,
		12, 6, 0, 118, 119, 5, 28, 0, 0, 119, 121, 3, 12, 6, 0, 120, 118, 1, 0,
		0, 0, 121, 124, 1, 0, 0, 0, 122, 120, 1, 0, 0, 0, 122, 123, 1, 0, 0, 0,
		123, 11, 1, 0, 0, 0, 124, 122, 1, 0, 0, 0, 125, 129, 5, 14, 0, 0, 126,
		129, 3, 16, 8, 0, 127, 129, 5, 39, 0, 0, 128, 125, 1, 0, 0, 0, 128, 126,
		1, 0, 0, 0, 128, 127, 1, 0, 0, 0, 129, 13, 1, 0, 0, 0, 130, 133, 5, 14,
		0, 0, 131, 133, 3, 16, 8, 0, 132, 130, 1, 0, 0, 0, 132, 131, 1, 0, 0, 0,
		133, 15, 1, 0, 0, 0, 134, 138, 5, 41, 0, 0, 135, 138, 5, 40, 0, 0, 136,
		138, 3, 18, 9, 0, 137, 134, 1, 0, 0, 0, 137, 135, 1, 0, 0, 0, 137, 136,
		1, 0, 0, 0, 138, 17, 1, 0, 0, 0, 139, 140, 7, 0, 0, 0, 140, 19, 1, 0, 0,
		0, 141, 142, 6, 10, -1, 0, 142, 143, 5, 25, 0, 0, 143, 144, 3, 20, 10,
		0, 144, 145, 5, 26, 0, 0, 145, 156, 1, 0, 0, 0, 146, 147, 7, 1, 0, 0, 147,
		156, 3, 20, 10, 12, 148, 156, 3, 22, 11, 0, 149, 156, 5, 14, 0, 0, 150,
		156, 3, 32, 16, 0, 151, 156, 5, 39, 0, 0, 152, 156, 5, 42, 0, 0, 153, 156,
		7, 2, 0, 0, 154, 156, 5, 23, 0, 0, 155, 141, 1, 0, 0, 0, 155, 146, 1, 0,
		0, 0, 155, 148, 1, 0, 0, 0, 155, 149, 1, 0, 0, 0, 155, 150, 1, 0, 0, 0,
		155, 151, 1, 0, 0, 0, 155, 152, 1, 0, 0, 0, 155, 153, 1, 0, 0, 0, 155,
		154, 1, 0, 0, 0, 156, 176, 1, 0, 0, 0, 157, 158, 10, 13, 0, 0, 158, 159,
		5, 35, 0, 0, 159, 175, 3, 20, 10, 13, 160, 161, 10, 11, 0, 0, 161, 162,
		7, 3, 0, 0, 162, 175, 3, 20, 10, 12, 163, 164, 10, 10, 0, 0, 164, 165,
		7, 1, 0, 0, 165, 175, 3, 20, 10, 11, 166, 167, 10, 9, 0, 0, 167, 168, 5,
		36, 0, 0, 168, 175, 3, 20, 10, 10, 169, 170, 10, 8, 0, 0, 170, 171, 7,
		4, 0, 0, 171, 175, 3, 20, 10, 9, 172, 173, 10, 14, 0, 0, 173, 175, 5, 34,
		0, 0, 174, 157, 1, 0, 0, 0, 174, 160, 1, 0, 0, 0, 174, 163, 1, 0, 0, 0,
		174, 166, 1, 0, 0, 0, 174, 169, 1, 0, 0, 0, 174, 172, 1, 0, 0, 0, 175,
		178, 1, 0, 0, 0, 176, 174, 1, 0, 0, 0, 176, 177, 1, 0, 0, 0, 177, 21, 1,
		0, 0, 0, 178, 176, 1, 0, 0, 0, 179, 183, 5, 41, 0, 0, 180, 183, 5, 40,
		0, 0, 181, 183, 3, 18, 9, 0, 182, 179, 1, 0, 0, 0, 182, 180, 1, 0, 0, 0,
		182, 181, 1, 0, 0, 0, 183, 185, 1, 0, 0, 0, 184, 186, 5, 39, 0, 0, 185,
		184, 1, 0, 0, 0, 185, 186, 1, 0, 0, 0, 186, 187, 1, 0, 0, 0, 187, 189,
		5, 25, 0, 0, 188, 190, 3, 24, 12, 0, 189, 188, 1, 0, 0, 0, 189, 190, 1,
		0, 0, 0, 190, 191, 1, 0, 0, 0, 191, 192, 5, 26, 0, 0, 192, 23, 1, 0, 0,
		0, 193, 198, 3, 20, 10, 0, 194, 195, 5, 28, 0, 0, 195, 197, 3, 20, 10,
		0, 196, 194, 1, 0, 0, 0, 197, 200, 1, 0, 0, 0, 198, 196, 1, 0, 0, 0, 198,
		199, 1, 0, 0, 0, 199, 25, 1, 0, 0, 0, 200, 198, 1, 0, 0, 0, 201, 202, 6,
		13, -1, 0, 202, 203, 5, 25, 0, 0, 203, 204, 3, 26, 13, 0, 204, 205, 5,
		26, 0, 0, 205, 215, 1, 0, 0, 0, 206, 207, 7, 1, 0, 0, 207, 215, 3, 26,
		13, 12, 208, 215, 3, 28, 14, 0, 209, 215, 3, 32, 16, 0, 210, 215, 5, 39,
		0, 0, 211, 215, 5, 42, 0, 0, 212, 215, 7, 2, 0, 0, 213, 215, 5, 23, 0,
		0, 214, 201, 1, 0, 0, 0, 214, 206, 1, 0, 0, 0, 214, 208, 1, 0, 0, 0, 214,
		209, 1, 0, 0, 0, 214, 210, 1, 0, 0, 0, 214, 211, 1, 0, 0, 0, 214, 212,
		1, 0, 0, 0, 214, 213, 1, 0, 0, 0, 215, 238, 1, 0, 0, 0, 216, 217, 10, 13,
		0, 0, 217, 218, 5, 35, 0, 0, 218, 237, 3, 26, 13, 13, 219, 220, 10, 11,
		0, 0, 220, 221, 7, 3, 0, 0, 221, 237, 3, 26, 13, 12, 222, 223, 10, 10,
		0, 0, 223, 224, 7, 1, 0, 0, 224, 237, 3, 26, 13, 11, 225, 226, 10, 9, 0,
		0, 226, 227, 5, 36, 0, 0, 227, 237, 3, 26, 13, 10, 228, 229, 10, 8, 0,
		0, 229, 230, 7, 4, 0, 0, 230, 237, 3, 26, 13, 9, 231, 232, 10, 14, 0, 0,
		232, 237, 5, 34, 0, 0, 233, 234, 10, 7, 0, 0, 234, 235, 5, 38, 0, 0, 235,
		237, 3, 28, 14, 0, 236, 216, 1, 0, 0, 0, 236, 219, 1, 0, 0, 0, 236, 222,
		1, 0, 0, 0, 236, 225, 1, 0, 0, 0, 236, 228, 1, 0, 0, 0, 236, 231, 1, 0,
		0, 0, 236, 233, 1, 0, 0, 0, 237, 240, 1, 0, 0, 0, 238, 236, 1, 0, 0, 0,
		238, 239, 1, 0, 0, 0, 239, 27, 1, 0, 0, 0, 240, 238, 1, 0, 0, 0, 241, 243,
		7, 5, 0, 0, 242, 244, 5, 39, 0, 0, 243, 242, 1, 0, 0, 0, 243, 244, 1, 0,
		0, 0, 244, 245, 1, 0, 0, 0, 245, 247, 5, 25, 0, 0, 246, 248, 3, 30, 15,
		0, 247, 246, 1, 0, 0, 0, 247, 248, 1, 0, 0, 0, 248, 249, 1, 0, 0, 0, 249,
		250, 5, 26, 0, 0, 250, 29, 1, 0, 0, 0, 251, 256, 3, 26, 13, 0, 252, 253,
		5, 28, 0, 0, 253, 255, 3, 26, 13, 0, 254, 252, 1, 0, 0, 0, 255, 258, 1,
		0, 0, 0, 256, 254, 1, 0, 0, 0, 256, 257, 1, 0, 0, 0, 257, 31, 1, 0, 0,
		0, 258, 256, 1, 0, 0, 0, 259, 261, 3, 34, 17, 0, 260, 259, 1, 0, 0, 0,
		260, 261, 1, 0, 0, 0, 261, 262, 1, 0, 0, 0, 262, 265, 3, 36, 18, 0, 263,
		264, 5, 27, 0, 0, 264, 266, 3, 36, 18, 0, 265, 263, 1, 0, 0, 0, 265, 266,
		1, 0, 0, 0, 266, 33, 1, 0, 0, 0, 267, 268, 5, 42, 0, 0, 268, 269, 5, 37,
		0, 0, 269, 35, 1, 0, 0, 0, 270, 272, 5, 29, 0, 0, 271, 270, 1, 0, 0, 0,
		271, 272, 1, 0, 0, 0, 272, 273, 1, 0, 0, 0, 273, 275, 5, 40, 0, 0, 274,
		276, 5, 29, 0, 0, 275, 274, 1, 0, 0, 0, 275, 276, 1, 0, 0, 0, 276, 277,
		1, 0, 0, 0, 277, 278, 5, 39, 0, 0, 278, 37, 1, 0, 0, 0, 29, 43, 60, 69,
		78, 83, 97, 102, 113, 122, 128, 132, 137, 155, 174, 176, 182, 185, 189,
		198, 214, 236, 238, 243, 247, 256, 260, 265, 271, 275,
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

// TqParserInit initializes any static state used to implement TqParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewTqParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func TqParserInit() {
	staticData := &TqParserParserStaticData
	staticData.once.Do(tqparserParserInit)
}

// NewTqParser produces a new parser instance for the optional input antlr.TokenStream.
func NewTqParser(input antlr.TokenStream) *TqParser {
	TqParserInit()
	this := new(TqParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &TqParserParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "TqParser.g4"

	return this
}

// TqParser tokens.
const (
	TqParserEOF        = antlr.TokenEOF
	TqParserSELECT     = 1
	TqParserDROP       = 2
	TqParserWHERE      = 3
	TqParserDERIVE     = 4
	TqParserRENAME     = 5
	TqParserSORT       = 6
	TqParserDISTINCT   = 7
	TqParserLIMIT      = 8
	TqParserOFFSET     = 9
	TqParserGROUP      = 10
	TqParserAS         = 11
	TqParserLBRACE     = 12
	TqParserRBRACE     = 13
	TqParserCOLUMN     = 14
	TqParserWS         = 15
	TqParserGE         = 16
	TqParserLE         = 17
	TqParserNE         = 18
	TqParserGT         = 19
	TqParserLT         = 20
	TqParserTRUE       = 21
	TqParserFALSE      = 22
	TqParserERRORCONST = 23
	TqParserEQ         = 24
	TqParserLPAREN     = 25
	TqParserRPAREN     = 26
	TqParserCOLON      = 27
	TqParserCOMMA      = 28
	TqParserDOLLAR     = 29
	TqParserSTAR       = 30
	TqParserPLUS       = 31
	TqParserDASH       = 32
	TqParserSLASH      = 33
	TqParserPERCENT    = 34
	TqParserCARET      = 35
	TqParserAMP        = 36
	TqParserBANG       = 37
	TqParserPIPE       = 38
	TqParserNUMBER     = 39
	TqParserCOL        = 40
	TqParserNAME       = 41
	TqParserSTRING     = 42
)

// TqParser rules.
const (
	TqParserRULE_program        = 0
	TqParserRULE_stage          = 1
	TqParserRULE_assignment     = 2
	TqParserRULE_renamePair     = 3
	TqParserRULE_sortKey        = 4
	TqParserRULE_columnList     = 5
	TqParserRULE_columnItem     = 6
	TqParserRULE_columnName     = 7
	TqParserRULE_bareName       = 8
	TqParserRULE_verbKeyword    = 9
	TqParserRULE_tqExpr         = 10
	TqParserRULE_tqFunctionCall = 11
	TqParserRULE_tqArgList      = 12
	TqParserRULE_expression     = 13
	TqParserRULE_functionCall   = 14
	TqParserRULE_argList        = 15
	TqParserRULE_reference      = 16
	TqParserRULE_sheetQualifier = 17
	TqParserRULE_cellRef        = 18
)

// IProgramContext is an interface to support dynamic dispatch.
type IProgramContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllStage() []IStageContext
	Stage(i int) IStageContext
	EOF() antlr.TerminalNode
	AllPIPE() []antlr.TerminalNode
	PIPE(i int) antlr.TerminalNode

	// IsProgramContext differentiates from other interfaces.
	IsProgramContext()
}

type ProgramContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgramContext() *ProgramContext {
	var p = new(ProgramContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TqParserRULE_program
	return p
}

func InitEmptyProgramContext(p *ProgramContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TqParserRULE_program
}

func (*ProgramContext) IsProgramContext() {}

func NewProgramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgramContext {
	var p = new(ProgramContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TqParserRULE_program

	return p
}

func (s *ProgramContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgramContext) AllStage() []IStageContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStageContext); ok {
			len++
		}
	}

	tst := make([]IStageContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStageContext); ok {
			tst[i] = t.(IStageContext)
			i++
		}
	}

	return tst
}

func (s *ProgramContext) Stage(i int) IStageContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStageContext); ok {
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

	return t.(IStageContext)
}

func (s *ProgramContext) EOF() antlr.TerminalNode {
	return s.GetToken(TqParserEOF, 0)
}

func (s *ProgramContext) AllPIPE() []antlr.TerminalNode {
	return s.GetTokens(TqParserPIPE)
}

func (s *ProgramContext) PIPE(i int) antlr.TerminalNode {
	return s.GetToken(TqParserPIPE, i)
}

func (s *ProgramContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgramContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgramContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TqParserListener); ok {
		listenerT.EnterProgram(s)
	}
}

func (s *ProgramContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TqParserListener); ok {
		listenerT.ExitProgram(s)
	}
}

func (s *ProgramContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TqParserVisitor:
		return t.VisitProgram(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TqParser) Program() (localctx IProgramContext) {
	localctx = NewProgramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, TqParserRULE_program)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(38)
		p.Stage()
	}
	p.SetState(43)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == TqParserPIPE {
		{
			p.SetState(39)
			p.Match(TqParserPIPE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(40)
			p.Stage()
		}

		p.SetState(45)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(46)
		p.Match(TqParserEOF)
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

// IStageContext is an interface to support dynamic dispatch.
type IStageContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsStageContext differentiates from other interfaces.
	IsStageContext()
}

type StageContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStageContext() *StageContext {
	var p = new(StageContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TqParserRULE_stage
	return p
}

func InitEmptyStageContext(p *StageContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TqParserRULE_stage
}

func (*StageContext) IsStageContext() {}

func NewStageContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StageContext {
	var p = new(StageContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TqParserRULE_stage

	return p
}

func (s *StageContext) GetParser() antlr.Parser { return s.parser }

func (s *StageContext) CopyAll(ctx *StageContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *StageContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StageContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type SortStageContext struct {
	StageContext
}

func NewSortStageContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SortStageContext {
	var p = new(SortStageContext)

	InitEmptyStageContext(&p.StageContext)
	p.parser = parser
	p.CopyAll(ctx.(*StageContext))

	return p
}

func (s *SortStageContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SortStageContext) SORT() antlr.TerminalNode {
	return s.GetToken(TqParserSORT, 0)
}

func (s *SortStageContext) AllSortKey() []ISortKeyContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISortKeyContext); ok {
			len++
		}
	}

	tst := make([]ISortKeyContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISortKeyContext); ok {
			tst[i] = t.(ISortKeyContext)
			i++
		}
	}

	return tst
}

func (s *SortStageContext) SortKey(i int) ISortKeyContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISortKeyContext); ok {
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

	return t.(ISortKeyContext)
}

func (s *SortStageContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(TqParserCOMMA)
}

func (s *SortStageContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(TqParserCOMMA, i)
}

func (s *SortStageContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TqParserListener); ok {
		listenerT.EnterSortStage(s)
	}
}

func (s *SortStageContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TqParserListener); ok {
		listenerT.ExitSortStage(s)
	}
}

func (s *SortStageContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TqParserVisitor:
		return t.VisitSortStage(s)

	default:
		return t.VisitChildren(s)
	}
}

type DeriveStageContext struct {
	StageContext
}

func NewDeriveStageContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DeriveStageContext {
	var p = new(DeriveStageContext)

	InitEmptyStageContext(&p.StageContext)
	p.parser = parser
	p.CopyAll(ctx.(*StageContext))

	return p
}

func (s *DeriveStageContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeriveStageContext) DERIVE() antlr.TerminalNode {
	return s.GetToken(TqParserDERIVE, 0)
}

func (s *DeriveStageContext) AllAssignment() []IAssignmentContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IAssignmentContext); ok {
			len++
		}
	}

	tst := make([]IAssignmentContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IAssignmentContext); ok {
			tst[i] = t.(IAssignmentContext)
			i++
		}
	}

	return tst
}

func (s *DeriveStageContext) Assignment(i int) IAssignmentContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssignmentContext); ok {
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

	return t.(IAssignmentContext)
}

func (s *DeriveStageContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(TqParserCOMMA)
}

func (s *DeriveStageContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(TqParserCOMMA, i)
}

func (s *DeriveStageContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TqParserListener); ok {
		listenerT.EnterDeriveStage(s)
	}
}

func (s *DeriveStageContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TqParserListener); ok {
		listenerT.ExitDeriveStage(s)
	}
}

func (s *DeriveStageContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TqParserVisitor:
		return t.VisitDeriveStage(s)

	default:
		return t.VisitChildren(s)
	}
}

type RenameStageContext struct {
	StageContext
}

func NewRenameStageContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *RenameStageContext {
	var p = new(RenameStageContext)

	InitEmptyStageContext(&p.StageContext)
	p.parser = parser
	p.CopyAll(ctx.(*StageContext))

	return p
}

func (s *RenameStageContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RenameStageContext) RENAME() antlr.TerminalNode {
	return s.GetToken(TqParserRENAME, 0)
}

func (s *RenameStageContext) AllRenamePair() []IRenamePairContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IRenamePairContext); ok {
			len++
		}
	}

	tst := make([]IRenamePairContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IRenamePairContext); ok {
			tst[i] = t.(IRenamePairContext)
			i++
		}
	}

	return tst
}

func (s *RenameStageContext) RenamePair(i int) IRenamePairContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRenamePairContext); ok {
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

	return t.(IRenamePairContext)
}

func (s *RenameStageContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(TqParserCOMMA)
}

func (s *RenameStageContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(TqParserCOMMA, i)
}

func (s *RenameStageContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TqParserListener); ok {
		listenerT.EnterRenameStage(s)
	}
}

func (s *RenameStageContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TqParserListener); ok {
		listenerT.ExitRenameStage(s)
	}
}

func (s *RenameStageContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TqParserVisitor:
		return t.VisitRenameStage(s)

	default:
		return t.VisitChildren(s)
	}
}

type DropStageContext struct {
	StageContext
}

func NewDropStageContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DropStageContext {
	var p = new(DropStageContext)

	InitEmptyStageContext(&p.StageContext)
	p.parser = parser
	p.CopyAll(ctx.(*StageContext))

	return p
}

func (s *DropStageContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DropStageContext) DROP() antlr.TerminalNode {
	return s.GetToken(TqParserDROP, 0)
}

func (s *DropStageContext) ColumnList() IColumnListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IColumnListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IColumnListContext)
}

func (s *DropStageContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TqParserListener); ok {
		listenerT.EnterDropStage(s)
	}
}

func (s *DropStageContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TqParserListener); ok {
		listenerT.ExitDropStage(s)
	}
}

func (s *DropStageContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TqParserVisitor:
		return t.VisitDropStage(s)

	default:
		return t.VisitChildren(s)
	}
}

type DistinctStageContext struct {
	StageContext
}

func NewDistinctStageContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DistinctStageContext {
	var p = new(DistinctStageContext)

	InitEmptyStageContext(&p.StageContext)
	p.parser = parser
	p.CopyAll(ctx.(*StageContext))

	return p
}

func (s *DistinctStageContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DistinctStageContext) DISTINCT() antlr.TerminalNode {
	return s.GetToken(TqParserDISTINCT, 0)
}

func (s *DistinctStageContext) ColumnList() IColumnListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IColumnListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IColumnListContext)
}

func (s *DistinctStageContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TqParserListener); ok {
		listenerT.EnterDistinctStage(s)
	}
}

func (s *DistinctStageContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TqParserListener); ok {
		listenerT.ExitDistinctStage(s)
	}
}

func (s *DistinctStageContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TqParserVisitor:
		return t.VisitDistinctStage(s)

	default:
		return t.VisitChildren(s)
	}
}

type LimitStageContext struct {
	StageContext
}

func NewLimitStageContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LimitStageContext {
	var p = new(LimitStageContext)

	InitEmptyStageContext(&p.StageContext)
	p.parser = parser
	p.CopyAll(ctx.(*StageContext))

	return p
}

func (s *LimitStageContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LimitStageContext) LIMIT() antlr.TerminalNode {
	return s.GetToken(TqParserLIMIT, 0)
}

func (s *LimitStageContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(TqParserNUMBER, 0)
}

func (s *LimitStageContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TqParserListener); ok {
		listenerT.EnterLimitStage(s)
	}
}

func (s *LimitStageContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TqParserListener); ok {
		listenerT.ExitLimitStage(s)
	}
}

func (s *LimitStageContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TqParserVisitor:
		return t.VisitLimitStage(s)

	default:
		return t.VisitChildren(s)
	}
}

type GroupStageContext struct {
	StageContext
}

func NewGroupStageContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *GroupStageContext {
	var p = new(GroupStageContext)

	InitEmptyStageContext(&p.StageContext)
	p.parser = parser
	p.CopyAll(ctx.(*StageContext))

	return p
}

func (s *GroupStageContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GroupStageContext) GROUP() antlr.TerminalNode {
	return s.GetToken(TqParserGROUP, 0)
}

func (s *GroupStageContext) ColumnList() IColumnListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IColumnListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IColumnListContext)
}

func (s *GroupStageContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(TqParserLBRACE, 0)
}

func (s *GroupStageContext) AllAssignment() []IAssignmentContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IAssignmentContext); ok {
			len++
		}
	}

	tst := make([]IAssignmentContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IAssignmentContext); ok {
			tst[i] = t.(IAssignmentContext)
			i++
		}
	}

	return tst
}

func (s *GroupStageContext) Assignment(i int) IAssignmentContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssignmentContext); ok {
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

	return t.(IAssignmentContext)
}

func (s *GroupStageContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(TqParserRBRACE, 0)
}

func (s *GroupStageContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(TqParserCOMMA)
}

func (s *GroupStageContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(TqParserCOMMA, i)
}

func (s *GroupStageContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TqParserListener); ok {
		listenerT.EnterGroupStage(s)
	}
}

func (s *GroupStageContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TqParserListener); ok {
		listenerT.ExitGroupStage(s)
	}
}

func (s *GroupStageContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TqParserVisitor:
		return t.VisitGroupStage(s)

	default:
		return t.VisitChildren(s)
	}
}

type WhereStageContext struct {
	StageContext
}

func NewWhereStageContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *WhereStageContext {
	var p = new(WhereStageContext)

	InitEmptyStageContext(&p.StageContext)
	p.parser = parser
	p.CopyAll(ctx.(*StageContext))

	return p
}

func (s *WhereStageContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WhereStageContext) WHERE() antlr.TerminalNode {
	return s.GetToken(TqParserWHERE, 0)
}

func (s *WhereStageContext) TqExpr() ITqExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITqExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITqExprContext)
}

func (s *WhereStageContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TqParserListener); ok {
		listenerT.EnterWhereStage(s)
	}
}

func (s *WhereStageContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TqParserListener); ok {
		listenerT.ExitWhereStage(s)
	}
}

func (s *WhereStageContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TqParserVisitor:
		return t.VisitWhereStage(s)

	default:
		return t.VisitChildren(s)
	}
}

type SelectStageContext struct {
	StageContext
}

func NewSelectStageContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SelectStageContext {
	var p = new(SelectStageContext)

	InitEmptyStageContext(&p.StageContext)
	p.parser = parser
	p.CopyAll(ctx.(*StageContext))

	return p
}

func (s *SelectStageContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SelectStageContext) SELECT() antlr.TerminalNode {
	return s.GetToken(TqParserSELECT, 0)
}

func (s *SelectStageContext) ColumnList() IColumnListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IColumnListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IColumnListContext)
}

func (s *SelectStageContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TqParserListener); ok {
		listenerT.EnterSelectStage(s)
	}
}

func (s *SelectStageContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TqParserListener); ok {
		listenerT.ExitSelectStage(s)
	}
}

func (s *SelectStageContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TqParserVisitor:
		return t.VisitSelectStage(s)

	default:
		return t.VisitChildren(s)
	}
}

type OffsetStageContext struct {
	StageContext
}

func NewOffsetStageContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *OffsetStageContext {
	var p = new(OffsetStageContext)

	InitEmptyStageContext(&p.StageContext)
	p.parser = parser
	p.CopyAll(ctx.(*StageContext))

	return p
}

func (s *OffsetStageContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OffsetStageContext) OFFSET() antlr.TerminalNode {
	return s.GetToken(TqParserOFFSET, 0)
}

func (s *OffsetStageContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(TqParserNUMBER, 0)
}

func (s *OffsetStageContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TqParserListener); ok {
		listenerT.EnterOffsetStage(s)
	}
}

func (s *OffsetStageContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TqParserListener); ok {
		listenerT.ExitOffsetStage(s)
	}
}

func (s *OffsetStageContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TqParserVisitor:
		return t.VisitOffsetStage(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TqParser) Stage() (localctx IStageContext) {
	localctx = NewStageContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, TqParserRULE_stage)
	var _la int

	p.SetState(102)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case TqParserSELECT:
		localctx = NewSelectStageContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(48)
			p.Match(TqParserSELECT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(49)
			p.ColumnList()
		}

	case TqParserDROP:
		localctx = NewDropStageContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(50)
			p.Match(TqParserDROP)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(51)
			p.ColumnList()
		}

	case TqParserWHERE:
		localctx = NewWhereStageContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(52)
			p.Match(TqParserWHERE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(53)
			p.tqExpr(0)
		}

	case TqParserDERIVE:
		localctx = NewDeriveStageContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(54)
			p.Match(TqParserDERIVE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(55)
			p.Assignment()
		}
		p.SetState(60)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == TqParserCOMMA {
			{
				p.SetState(56)
				p.Match(TqParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(57)
				p.Assignment()
			}

			p.SetState(62)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	case TqParserRENAME:
		localctx = NewRenameStageContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(63)
			p.Match(TqParserRENAME)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(64)
			p.RenamePair()
		}
		p.SetState(69)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == TqParserCOMMA {
			{
				p.SetState(65)
				p.Match(TqParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(66)
				p.RenamePair()
			}

			p.SetState(71)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	case TqParserSORT:
		localctx = NewSortStageContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(72)
			p.Match(TqParserSORT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(73)
			p.SortKey()
		}
		p.SetState(78)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == TqParserCOMMA {
			{
				p.SetState(74)
				p.Match(TqParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(75)
				p.SortKey()
			}

			p.SetState(80)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	case TqParserDISTINCT:
		localctx = NewDistinctStageContext(p, localctx)
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(81)
			p.Match(TqParserDISTINCT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(83)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&3848290717694) != 0 {
			{
				p.SetState(82)
				p.ColumnList()
			}

		}

	case TqParserLIMIT:
		localctx = NewLimitStageContext(p, localctx)
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(85)
			p.Match(TqParserLIMIT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(86)
			p.Match(TqParserNUMBER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case TqParserOFFSET:
		localctx = NewOffsetStageContext(p, localctx)
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(87)
			p.Match(TqParserOFFSET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(88)
			p.Match(TqParserNUMBER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case TqParserGROUP:
		localctx = NewGroupStageContext(p, localctx)
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(89)
			p.Match(TqParserGROUP)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(90)
			p.ColumnList()
		}
		{
			p.SetState(91)
			p.Match(TqParserLBRACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(92)
			p.Assignment()
		}
		p.SetState(97)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == TqParserCOMMA {
			{
				p.SetState(93)
				p.Match(TqParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(94)
				p.Assignment()
			}

			p.SetState(99)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(100)
			p.Match(TqParserRBRACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
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

// IAssignmentContext is an interface to support dynamic dispatch.
type IAssignmentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ColumnName() IColumnNameContext
	EQ() antlr.TerminalNode
	TqExpr() ITqExprContext

	// IsAssignmentContext differentiates from other interfaces.
	IsAssignmentContext()
}

type AssignmentContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssignmentContext() *AssignmentContext {
	var p = new(AssignmentContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TqParserRULE_assignment
	return p
}

func InitEmptyAssignmentContext(p *AssignmentContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TqParserRULE_assignment
}

func (*AssignmentContext) IsAssignmentContext() {}

func NewAssignmentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssignmentContext {
	var p = new(AssignmentContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TqParserRULE_assignment

	return p
}

func (s *AssignmentContext) GetParser() antlr.Parser { return s.parser }

func (s *AssignmentContext) ColumnName() IColumnNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IColumnNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IColumnNameContext)
}

func (s *AssignmentContext) EQ() antlr.TerminalNode {
	return s.GetToken(TqParserEQ, 0)
}

func (s *AssignmentContext) TqExpr() ITqExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITqExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITqExprContext)
}

func (s *AssignmentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignmentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AssignmentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TqParserListener); ok {
		listenerT.EnterAssignment(s)
	}
}

func (s *AssignmentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TqParserListener); ok {
		listenerT.ExitAssignment(s)
	}
}

func (s *AssignmentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TqParserVisitor:
		return t.VisitAssignment(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TqParser) Assignment() (localctx IAssignmentContext) {
	localctx = NewAssignmentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, TqParserRULE_assignment)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(104)
		p.ColumnName()
	}
	{
		p.SetState(105)
		p.Match(TqParserEQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(106)
		p.tqExpr(0)
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

// IRenamePairContext is an interface to support dynamic dispatch.
type IRenamePairContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ColumnItem() IColumnItemContext
	AS() antlr.TerminalNode
	ColumnName() IColumnNameContext

	// IsRenamePairContext differentiates from other interfaces.
	IsRenamePairContext()
}

type RenamePairContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRenamePairContext() *RenamePairContext {
	var p = new(RenamePairContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TqParserRULE_renamePair
	return p
}

func InitEmptyRenamePairContext(p *RenamePairContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TqParserRULE_renamePair
}

func (*RenamePairContext) IsRenamePairContext() {}

func NewRenamePairContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RenamePairContext {
	var p = new(RenamePairContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TqParserRULE_renamePair

	return p
}

func (s *RenamePairContext) GetParser() antlr.Parser { return s.parser }

func (s *RenamePairContext) ColumnItem() IColumnItemContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IColumnItemContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IColumnItemContext)
}

func (s *RenamePairContext) AS() antlr.TerminalNode {
	return s.GetToken(TqParserAS, 0)
}

func (s *RenamePairContext) ColumnName() IColumnNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IColumnNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IColumnNameContext)
}

func (s *RenamePairContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RenamePairContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RenamePairContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TqParserListener); ok {
		listenerT.EnterRenamePair(s)
	}
}

func (s *RenamePairContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TqParserListener); ok {
		listenerT.ExitRenamePair(s)
	}
}

func (s *RenamePairContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TqParserVisitor:
		return t.VisitRenamePair(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TqParser) RenamePair() (localctx IRenamePairContext) {
	localctx = NewRenamePairContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, TqParserRULE_renamePair)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(108)
		p.ColumnItem()
	}
	{
		p.SetState(109)
		p.Match(TqParserAS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(110)
		p.ColumnName()
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

// ISortKeyContext is an interface to support dynamic dispatch.
type ISortKeyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ColumnItem() IColumnItemContext
	DASH() antlr.TerminalNode

	// IsSortKeyContext differentiates from other interfaces.
	IsSortKeyContext()
}

type SortKeyContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySortKeyContext() *SortKeyContext {
	var p = new(SortKeyContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TqParserRULE_sortKey
	return p
}

func InitEmptySortKeyContext(p *SortKeyContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TqParserRULE_sortKey
}

func (*SortKeyContext) IsSortKeyContext() {}

func NewSortKeyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SortKeyContext {
	var p = new(SortKeyContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TqParserRULE_sortKey

	return p
}

func (s *SortKeyContext) GetParser() antlr.Parser { return s.parser }

func (s *SortKeyContext) ColumnItem() IColumnItemContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IColumnItemContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IColumnItemContext)
}

func (s *SortKeyContext) DASH() antlr.TerminalNode {
	return s.GetToken(TqParserDASH, 0)
}

func (s *SortKeyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SortKeyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SortKeyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TqParserListener); ok {
		listenerT.EnterSortKey(s)
	}
}

func (s *SortKeyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TqParserListener); ok {
		listenerT.ExitSortKey(s)
	}
}

func (s *SortKeyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TqParserVisitor:
		return t.VisitSortKey(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TqParser) SortKey() (localctx ISortKeyContext) {
	localctx = NewSortKeyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, TqParserRULE_sortKey)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(113)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == TqParserDASH {
		{
			p.SetState(112)
			p.Match(TqParserDASH)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(115)
		p.ColumnItem()
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

// IColumnListContext is an interface to support dynamic dispatch.
type IColumnListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllColumnItem() []IColumnItemContext
	ColumnItem(i int) IColumnItemContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsColumnListContext differentiates from other interfaces.
	IsColumnListContext()
}

type ColumnListContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyColumnListContext() *ColumnListContext {
	var p = new(ColumnListContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TqParserRULE_columnList
	return p
}

func InitEmptyColumnListContext(p *ColumnListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TqParserRULE_columnList
}

func (*ColumnListContext) IsColumnListContext() {}

func NewColumnListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ColumnListContext {
	var p = new(ColumnListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TqParserRULE_columnList

	return p
}

func (s *ColumnListContext) GetParser() antlr.Parser { return s.parser }

func (s *ColumnListContext) AllColumnItem() []IColumnItemContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IColumnItemContext); ok {
			len++
		}
	}

	tst := make([]IColumnItemContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IColumnItemContext); ok {
			tst[i] = t.(IColumnItemContext)
			i++
		}
	}

	return tst
}

func (s *ColumnListContext) ColumnItem(i int) IColumnItemContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IColumnItemContext); ok {
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

	return t.(IColumnItemContext)
}

func (s *ColumnListContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(TqParserCOMMA)
}

func (s *ColumnListContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(TqParserCOMMA, i)
}

func (s *ColumnListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ColumnListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ColumnListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TqParserListener); ok {
		listenerT.EnterColumnList(s)
	}
}

func (s *ColumnListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TqParserListener); ok {
		listenerT.ExitColumnList(s)
	}
}

func (s *ColumnListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TqParserVisitor:
		return t.VisitColumnList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TqParser) ColumnList() (localctx IColumnListContext) {
	localctx = NewColumnListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, TqParserRULE_columnList)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(117)
		p.ColumnItem()
	}
	p.SetState(122)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == TqParserCOMMA {
		{
			p.SetState(118)
			p.Match(TqParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(119)
			p.ColumnItem()
		}

		p.SetState(124)
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

// IColumnItemContext is an interface to support dynamic dispatch.
type IColumnItemContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	COLUMN() antlr.TerminalNode
	BareName() IBareNameContext
	NUMBER() antlr.TerminalNode

	// IsColumnItemContext differentiates from other interfaces.
	IsColumnItemContext()
}

type ColumnItemContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyColumnItemContext() *ColumnItemContext {
	var p = new(ColumnItemContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TqParserRULE_columnItem
	return p
}

func InitEmptyColumnItemContext(p *ColumnItemContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TqParserRULE_columnItem
}

func (*ColumnItemContext) IsColumnItemContext() {}

func NewColumnItemContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ColumnItemContext {
	var p = new(ColumnItemContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TqParserRULE_columnItem

	return p
}

func (s *ColumnItemContext) GetParser() antlr.Parser { return s.parser }

func (s *ColumnItemContext) COLUMN() antlr.TerminalNode {
	return s.GetToken(TqParserCOLUMN, 0)
}

func (s *ColumnItemContext) BareName() IBareNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBareNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBareNameContext)
}

func (s *ColumnItemContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(TqParserNUMBER, 0)
}

func (s *ColumnItemContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ColumnItemContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ColumnItemContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TqParserListener); ok {
		listenerT.EnterColumnItem(s)
	}
}

func (s *ColumnItemContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TqParserListener); ok {
		listenerT.ExitColumnItem(s)
	}
}

func (s *ColumnItemContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TqParserVisitor:
		return t.VisitColumnItem(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TqParser) ColumnItem() (localctx IColumnItemContext) {
	localctx = NewColumnItemContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, TqParserRULE_columnItem)
	p.SetState(128)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case TqParserCOLUMN:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(125)
			p.Match(TqParserCOLUMN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case TqParserSELECT, TqParserDROP, TqParserWHERE, TqParserDERIVE, TqParserRENAME, TqParserSORT, TqParserDISTINCT, TqParserLIMIT, TqParserOFFSET, TqParserGROUP, TqParserAS, TqParserCOL, TqParserNAME:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(126)
			p.BareName()
		}

	case TqParserNUMBER:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(127)
			p.Match(TqParserNUMBER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
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

// IColumnNameContext is an interface to support dynamic dispatch.
type IColumnNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	COLUMN() antlr.TerminalNode
	BareName() IBareNameContext

	// IsColumnNameContext differentiates from other interfaces.
	IsColumnNameContext()
}

type ColumnNameContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyColumnNameContext() *ColumnNameContext {
	var p = new(ColumnNameContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TqParserRULE_columnName
	return p
}

func InitEmptyColumnNameContext(p *ColumnNameContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TqParserRULE_columnName
}

func (*ColumnNameContext) IsColumnNameContext() {}

func NewColumnNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ColumnNameContext {
	var p = new(ColumnNameContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TqParserRULE_columnName

	return p
}

func (s *ColumnNameContext) GetParser() antlr.Parser { return s.parser }

func (s *ColumnNameContext) COLUMN() antlr.TerminalNode {
	return s.GetToken(TqParserCOLUMN, 0)
}

func (s *ColumnNameContext) BareName() IBareNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBareNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBareNameContext)
}

func (s *ColumnNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ColumnNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ColumnNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TqParserListener); ok {
		listenerT.EnterColumnName(s)
	}
}

func (s *ColumnNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TqParserListener); ok {
		listenerT.ExitColumnName(s)
	}
}

func (s *ColumnNameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TqParserVisitor:
		return t.VisitColumnName(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TqParser) ColumnName() (localctx IColumnNameContext) {
	localctx = NewColumnNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, TqParserRULE_columnName)
	p.SetState(132)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case TqParserCOLUMN:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(130)
			p.Match(TqParserCOLUMN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case TqParserSELECT, TqParserDROP, TqParserWHERE, TqParserDERIVE, TqParserRENAME, TqParserSORT, TqParserDISTINCT, TqParserLIMIT, TqParserOFFSET, TqParserGROUP, TqParserAS, TqParserCOL, TqParserNAME:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(131)
			p.BareName()
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

// IBareNameContext is an interface to support dynamic dispatch.
type IBareNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	NAME() antlr.TerminalNode
	COL() antlr.TerminalNode
	VerbKeyword() IVerbKeywordContext

	// IsBareNameContext differentiates from other interfaces.
	IsBareNameContext()
}

type BareNameContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBareNameContext() *BareNameContext {
	var p = new(BareNameContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TqParserRULE_bareName
	return p
}

func InitEmptyBareNameContext(p *BareNameContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TqParserRULE_bareName
}

func (*BareNameContext) IsBareNameContext() {}

func NewBareNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BareNameContext {
	var p = new(BareNameContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TqParserRULE_bareName

	return p
}

func (s *BareNameContext) GetParser() antlr.Parser { return s.parser }

func (s *BareNameContext) NAME() antlr.TerminalNode {
	return s.GetToken(TqParserNAME, 0)
}

func (s *BareNameContext) COL() antlr.TerminalNode {
	return s.GetToken(TqParserCOL, 0)
}

func (s *BareNameContext) VerbKeyword() IVerbKeywordContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVerbKeywordContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVerbKeywordContext)
}

func (s *BareNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BareNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BareNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TqParserListener); ok {
		listenerT.EnterBareName(s)
	}
}

func (s *BareNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TqParserListener); ok {
		listenerT.ExitBareName(s)
	}
}

func (s *BareNameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TqParserVisitor:
		return t.VisitBareName(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TqParser) BareName() (localctx IBareNameContext) {
	localctx = NewBareNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, TqParserRULE_bareName)
	p.SetState(137)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case TqParserNAME:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(134)
			p.Match(TqParserNAME)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case TqParserCOL:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(135)
			p.Match(TqParserCOL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case TqParserSELECT, TqParserDROP, TqParserWHERE, TqParserDERIVE, TqParserRENAME, TqParserSORT, TqParserDISTINCT, TqParserLIMIT, TqParserOFFSET, TqParserGROUP, TqParserAS:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(136)
			p.VerbKeyword()
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

// IVerbKeywordContext is an interface to support dynamic dispatch.
type IVerbKeywordContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SELECT() antlr.TerminalNode
	DROP() antlr.TerminalNode
	WHERE() antlr.TerminalNode
	DERIVE() antlr.TerminalNode
	RENAME() antlr.TerminalNode
	SORT() antlr.TerminalNode
	DISTINCT() antlr.TerminalNode
	LIMIT() antlr.TerminalNode
	OFFSET() antlr.TerminalNode
	GROUP() antlr.TerminalNode
	AS() antlr.TerminalNode

	// IsVerbKeywordContext differentiates from other interfaces.
	IsVerbKeywordContext()
}

type VerbKeywordContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVerbKeywordContext() *VerbKeywordContext {
	var p = new(VerbKeywordContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TqParserRULE_verbKeyword
	return p
}

func InitEmptyVerbKeywordContext(p *VerbKeywordContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TqParserRULE_verbKeyword
}

func (*VerbKeywordContext) IsVerbKeywordContext() {}

func NewVerbKeywordContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VerbKeywordContext {
	var p = new(VerbKeywordContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TqParserRULE_verbKeyword

	return p
}

func (s *VerbKeywordContext) GetParser() antlr.Parser { return s.parser }

func (s *VerbKeywordContext) SELECT() antlr.TerminalNode {
	return s.GetToken(TqParserSELECT, 0)
}

func (s *VerbKeywordContext) DROP() antlr.TerminalNode {
	return s.GetToken(TqParserDROP, 0)
}

func (s *VerbKeywordContext) WHERE() antlr.TerminalNode {
	return s.GetToken(TqParserWHERE, 0)
}

func (s *VerbKeywordContext) DERIVE() antlr.TerminalNode {
	return s.GetToken(TqParserDERIVE, 0)
}

func (s *VerbKeywordContext) RENAME() antlr.TerminalNode {
	return s.GetToken(TqParserRENAME, 0)
}

func (s *VerbKeywordContext) SORT() antlr.TerminalNode {
	return s.GetToken(TqParserSORT, 0)
}

func (s *VerbKeywordContext) DISTINCT() antlr.TerminalNode {
	return s.GetToken(TqParserDISTINCT, 0)
}

func (s *VerbKeywordContext) LIMIT() antlr.TerminalNode {
	return s.GetToken(TqParserLIMIT, 0)
}

func (s *VerbKeywordContext) OFFSET() antlr.TerminalNode {
	return s.GetToken(TqParserOFFSET, 0)
}

func (s *VerbKeywordContext) GROUP() antlr.TerminalNode {
	return s.GetToken(TqParserGROUP, 0)
}

func (s *VerbKeywordContext) AS() antlr.TerminalNode {
	return s.GetToken(TqParserAS, 0)
}

func (s *VerbKeywordContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VerbKeywordContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VerbKeywordContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TqParserListener); ok {
		listenerT.EnterVerbKeyword(s)
	}
}

func (s *VerbKeywordContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TqParserListener); ok {
		listenerT.ExitVerbKeyword(s)
	}
}

func (s *VerbKeywordContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TqParserVisitor:
		return t.VisitVerbKeyword(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TqParser) VerbKeyword() (localctx IVerbKeywordContext) {
	localctx = NewVerbKeywordContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, TqParserRULE_verbKeyword)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(139)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&4094) != 0) {
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

// ITqExprContext is an interface to support dynamic dispatch.
type ITqExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsTqExprContext differentiates from other interfaces.
	IsTqExprContext()
}

type TqExprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTqExprContext() *TqExprContext {
	var p = new(TqExprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TqParserRULE_tqExpr
	return p
}

func InitEmptyTqExprContext(p *TqExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TqParserRULE_tqExpr
}

func (*TqExprContext) IsTqExprContext() {}

func NewTqExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TqExprContext {
	var p = new(TqExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TqParserRULE_tqExpr

	return p
}

func (s *TqExprContext) GetParser() antlr.Parser { return s.parser }

func (s *TqExprContext) CopyAll(ctx *TqExprContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *TqExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TqExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type TqPowExprContext struct {
	TqExprContext
}

func NewTqPowExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TqPowExprContext {
	var p = new(TqPowExprContext)

	InitEmptyTqExprContext(&p.TqExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*TqExprContext))

	return p
}

func (s *TqPowExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TqPowExprContext) AllTqExpr() []ITqExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITqExprContext); ok {
			len++
		}
	}

	tst := make([]ITqExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITqExprContext); ok {
			tst[i] = t.(ITqExprContext)
			i++
		}
	}

	return tst
}

func (s *TqPowExprContext) TqExpr(i int) ITqExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITqExprContext); ok {
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

	return t.(ITqExprContext)
}

func (s *TqPowExprContext) CARET() antlr.TerminalNode {
	return s.GetToken(TqParserCARET, 0)
}

func (s *TqPowExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TqParserListener); ok {
		listenerT.EnterTqPowExpr(s)
	}
}

func (s *TqPowExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TqParserListener); ok {
		listenerT.ExitTqPowExpr(s)
	}
}

func (s *TqPowExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TqParserVisitor:
		return t.VisitTqPowExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type TqCompareExprContext struct {
	TqExprContext
	op antlr.Token
}

func NewTqCompareExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TqCompareExprContext {
	var p = new(TqCompareExprContext)

	InitEmptyTqExprContext(&p.TqExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*TqExprContext))

	return p
}

func (s *TqCompareExprContext) GetOp() antlr.Token { return s.op }

func (s *TqCompareExprContext) SetOp(v antlr.Token) { s.op = v }

func (s *TqCompareExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TqCompareExprContext) AllTqExpr() []ITqExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITqExprContext); ok {
			len++
		}
	}

	tst := make([]ITqExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITqExprContext); ok {
			tst[i] = t.(ITqExprContext)
			i++
		}
	}

	return tst
}

func (s *TqCompareExprContext) TqExpr(i int) ITqExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITqExprContext); ok {
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

	return t.(ITqExprContext)
}

func (s *TqCompareExprContext) EQ() antlr.TerminalNode {
	return s.GetToken(TqParserEQ, 0)
}

func (s *TqCompareExprContext) NE() antlr.TerminalNode {
	return s.GetToken(TqParserNE, 0)
}

func (s *TqCompareExprContext) LT() antlr.TerminalNode {
	return s.GetToken(TqParserLT, 0)
}

func (s *TqCompareExprContext) LE() antlr.TerminalNode {
	return s.GetToken(TqParserLE, 0)
}

func (s *TqCompareExprContext) GT() antlr.TerminalNode {
	return s.GetToken(TqParserGT, 0)
}

func (s *TqCompareExprContext) GE() antlr.TerminalNode {
	return s.GetToken(TqParserGE, 0)
}

func (s *TqCompareExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TqParserListener); ok {
		listenerT.EnterTqCompareExpr(s)
	}
}

func (s *TqCompareExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TqParserListener); ok {
		listenerT.ExitTqCompareExpr(s)
	}
}

func (s *TqCompareExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TqParserVisitor:
		return t.VisitTqCompareExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type TqBoolExprContext struct {
	TqExprContext
}

func NewTqBoolExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TqBoolExprContext {
	var p = new(TqBoolExprContext)

	InitEmptyTqExprContext(&p.TqExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*TqExprContext))

	return p
}

func (s *TqBoolExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TqBoolExprContext) TRUE() antlr.TerminalNode {
	return s.GetToken(TqParserTRUE, 0)
}

func (s *TqBoolExprContext) FALSE() antlr.TerminalNode {
	return s.GetToken(TqParserFALSE, 0)
}

func (s *TqBoolExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TqParserListener); ok {
		listenerT.EnterTqBoolExpr(s)
	}
}

func (s *TqBoolExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TqParserListener); ok {
		listenerT.ExitTqBoolExpr(s)
	}
}

func (s *TqBoolExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TqParserVisitor:
		return t.VisitTqBoolExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type TqNumberExprContext struct {
	TqExprContext
}

func NewTqNumberExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TqNumberExprContext {
	var p = new(TqNumberExprContext)

	InitEmptyTqExprContext(&p.TqExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*TqExprContext))

	return p
}

func (s *TqNumberExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TqNumberExprContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(TqParserNUMBER, 0)
}

func (s *TqNumberExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TqParserListener); ok {
		listenerT.EnterTqNumberExpr(s)
	}
}

func (s *TqNumberExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TqParserListener); ok {
		listenerT.ExitTqNumberExpr(s)
	}
}

func (s *TqNumberExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TqParserVisitor:
		return t.VisitTqNumberExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type TqPercentExprContext struct {
	TqExprContext
}

func NewTqPercentExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TqPercentExprContext {
	var p = new(TqPercentExprContext)

	InitEmptyTqExprContext(&p.TqExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*TqExprContext))

	return p
}

func (s *TqPercentExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TqPercentExprContext) TqExpr() ITqExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITqExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITqExprContext)
}

func (s *TqPercentExprContext) PERCENT() antlr.TerminalNode {
	return s.GetToken(TqParserPERCENT, 0)
}

func (s *TqPercentExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TqParserListener); ok {
		listenerT.EnterTqPercentExpr(s)
	}
}

func (s *TqPercentExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TqParserListener); ok {
		listenerT.ExitTqPercentExpr(s)
	}
}

func (s *TqPercentExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TqParserVisitor:
		return t.VisitTqPercentExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type TqConcatExprContext struct {
	TqExprContext
}

func NewTqConcatExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TqConcatExprContext {
	var p = new(TqConcatExprContext)

	InitEmptyTqExprContext(&p.TqExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*TqExprContext))

	return p
}

func (s *TqConcatExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TqConcatExprContext) AllTqExpr() []ITqExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITqExprContext); ok {
			len++
		}
	}

	tst := make([]ITqExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITqExprContext); ok {
			tst[i] = t.(ITqExprContext)
			i++
		}
	}

	return tst
}

func (s *TqConcatExprContext) TqExpr(i int) ITqExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITqExprContext); ok {
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

	return t.(ITqExprContext)
}

func (s *TqConcatExprContext) AMP() antlr.TerminalNode {
	return s.GetToken(TqParserAMP, 0)
}

func (s *TqConcatExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TqParserListener); ok {
		listenerT.EnterTqConcatExpr(s)
	}
}

func (s *TqConcatExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TqParserListener); ok {
		listenerT.ExitTqConcatExpr(s)
	}
}

func (s *TqConcatExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TqParserVisitor:
		return t.VisitTqConcatExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type TqParenExprContext struct {
	TqExprContext
}

func NewTqParenExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TqParenExprContext {
	var p = new(TqParenExprContext)

	InitEmptyTqExprContext(&p.TqExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*TqExprContext))

	return p
}

func (s *TqParenExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TqParenExprContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(TqParserLPAREN, 0)
}

func (s *TqParenExprContext) TqExpr() ITqExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITqExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITqExprContext)
}

func (s *TqParenExprContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(TqParserRPAREN, 0)
}

func (s *TqParenExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TqParserListener); ok {
		listenerT.EnterTqParenExpr(s)
	}
}

func (s *TqParenExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TqParserListener); ok {
		listenerT.ExitTqParenExpr(s)
	}
}

func (s *TqParenExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TqParserVisitor:
		return t.VisitTqParenExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type TqUnaryExprContext struct {
	TqExprContext
	op antlr.Token
}

func NewTqUnaryExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TqUnaryExprContext {
	var p = new(TqUnaryExprContext)

	InitEmptyTqExprContext(&p.TqExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*TqExprContext))

	return p
}

func (s *TqUnaryExprContext) GetOp() antlr.Token { return s.op }

func (s *TqUnaryExprContext) SetOp(v antlr.Token) { s.op = v }

func (s *TqUnaryExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TqUnaryExprContext) TqExpr() ITqExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITqExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITqExprContext)
}

func (s *TqUnaryExprContext) PLUS() antlr.TerminalNode {
	return s.GetToken(TqParserPLUS, 0)
}

func (s *TqUnaryExprContext) DASH() antlr.TerminalNode {
	return s.GetToken(TqParserDASH, 0)
}

func (s *TqUnaryExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TqParserListener); ok {
		listenerT.EnterTqUnaryExpr(s)
	}
}

func (s *TqUnaryExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TqParserListener); ok {
		listenerT.ExitTqUnaryExpr(s)
	}
}

func (s *TqUnaryExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TqParserVisitor:
		return t.VisitTqUnaryExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type TqCallExprContext struct {
	TqExprContext
}

func NewTqCallExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TqCallExprContext {
	var p = new(TqCallExprContext)

	InitEmptyTqExprContext(&p.TqExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*TqExprContext))

	return p
}

func (s *TqCallExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TqCallExprContext) TqFunctionCall() ITqFunctionCallContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITqFunctionCallContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITqFunctionCallContext)
}

func (s *TqCallExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TqParserListener); ok {
		listenerT.EnterTqCallExpr(s)
	}
}

func (s *TqCallExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TqParserListener); ok {
		listenerT.ExitTqCallExpr(s)
	}
}

func (s *TqCallExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TqParserVisitor:
		return t.VisitTqCallExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type TqRefExprContext struct {
	TqExprContext
}

func NewTqRefExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TqRefExprContext {
	var p = new(TqRefExprContext)

	InitEmptyTqExprContext(&p.TqExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*TqExprContext))

	return p
}

func (s *TqRefExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TqRefExprContext) Reference() IReferenceContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IReferenceContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IReferenceContext)
}

func (s *TqRefExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TqParserListener); ok {
		listenerT.EnterTqRefExpr(s)
	}
}

func (s *TqRefExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TqParserListener); ok {
		listenerT.ExitTqRefExpr(s)
	}
}

func (s *TqRefExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TqParserVisitor:
		return t.VisitTqRefExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type TqColumnExprContext struct {
	TqExprContext
}

func NewTqColumnExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TqColumnExprContext {
	var p = new(TqColumnExprContext)

	InitEmptyTqExprContext(&p.TqExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*TqExprContext))

	return p
}

func (s *TqColumnExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TqColumnExprContext) COLUMN() antlr.TerminalNode {
	return s.GetToken(TqParserCOLUMN, 0)
}

func (s *TqColumnExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TqParserListener); ok {
		listenerT.EnterTqColumnExpr(s)
	}
}

func (s *TqColumnExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TqParserListener); ok {
		listenerT.ExitTqColumnExpr(s)
	}
}

func (s *TqColumnExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TqParserVisitor:
		return t.VisitTqColumnExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type TqStringExprContext struct {
	TqExprContext
}

func NewTqStringExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TqStringExprContext {
	var p = new(TqStringExprContext)

	InitEmptyTqExprContext(&p.TqExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*TqExprContext))

	return p
}

func (s *TqStringExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TqStringExprContext) STRING() antlr.TerminalNode {
	return s.GetToken(TqParserSTRING, 0)
}

func (s *TqStringExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TqParserListener); ok {
		listenerT.EnterTqStringExpr(s)
	}
}

func (s *TqStringExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TqParserListener); ok {
		listenerT.ExitTqStringExpr(s)
	}
}

func (s *TqStringExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TqParserVisitor:
		return t.VisitTqStringExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type TqAddExprContext struct {
	TqExprContext
	op antlr.Token
}

func NewTqAddExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TqAddExprContext {
	var p = new(TqAddExprContext)

	InitEmptyTqExprContext(&p.TqExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*TqExprContext))

	return p
}

func (s *TqAddExprContext) GetOp() antlr.Token { return s.op }

func (s *TqAddExprContext) SetOp(v antlr.Token) { s.op = v }

func (s *TqAddExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TqAddExprContext) AllTqExpr() []ITqExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITqExprContext); ok {
			len++
		}
	}

	tst := make([]ITqExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITqExprContext); ok {
			tst[i] = t.(ITqExprContext)
			i++
		}
	}

	return tst
}

func (s *TqAddExprContext) TqExpr(i int) ITqExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITqExprContext); ok {
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

	return t.(ITqExprContext)
}

func (s *TqAddExprContext) PLUS() antlr.TerminalNode {
	return s.GetToken(TqParserPLUS, 0)
}

func (s *TqAddExprContext) DASH() antlr.TerminalNode {
	return s.GetToken(TqParserDASH, 0)
}

func (s *TqAddExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TqParserListener); ok {
		listenerT.EnterTqAddExpr(s)
	}
}

func (s *TqAddExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TqParserListener); ok {
		listenerT.ExitTqAddExpr(s)
	}
}

func (s *TqAddExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TqParserVisitor:
		return t.VisitTqAddExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type TqErrorExprContext struct {
	TqExprContext
}

func NewTqErrorExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TqErrorExprContext {
	var p = new(TqErrorExprContext)

	InitEmptyTqExprContext(&p.TqExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*TqExprContext))

	return p
}

func (s *TqErrorExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TqErrorExprContext) ERRORCONST() antlr.TerminalNode {
	return s.GetToken(TqParserERRORCONST, 0)
}

func (s *TqErrorExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TqParserListener); ok {
		listenerT.EnterTqErrorExpr(s)
	}
}

func (s *TqErrorExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TqParserListener); ok {
		listenerT.ExitTqErrorExpr(s)
	}
}

func (s *TqErrorExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TqParserVisitor:
		return t.VisitTqErrorExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type TqMulExprContext struct {
	TqExprContext
	op antlr.Token
}

func NewTqMulExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TqMulExprContext {
	var p = new(TqMulExprContext)

	InitEmptyTqExprContext(&p.TqExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*TqExprContext))

	return p
}

func (s *TqMulExprContext) GetOp() antlr.Token { return s.op }

func (s *TqMulExprContext) SetOp(v antlr.Token) { s.op = v }

func (s *TqMulExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TqMulExprContext) AllTqExpr() []ITqExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITqExprContext); ok {
			len++
		}
	}

	tst := make([]ITqExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITqExprContext); ok {
			tst[i] = t.(ITqExprContext)
			i++
		}
	}

	return tst
}

func (s *TqMulExprContext) TqExpr(i int) ITqExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITqExprContext); ok {
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

	return t.(ITqExprContext)
}

func (s *TqMulExprContext) STAR() antlr.TerminalNode {
	return s.GetToken(TqParserSTAR, 0)
}

func (s *TqMulExprContext) SLASH() antlr.TerminalNode {
	return s.GetToken(TqParserSLASH, 0)
}

func (s *TqMulExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TqParserListener); ok {
		listenerT.EnterTqMulExpr(s)
	}
}

func (s *TqMulExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TqParserListener); ok {
		listenerT.ExitTqMulExpr(s)
	}
}

func (s *TqMulExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TqParserVisitor:
		return t.VisitTqMulExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TqParser) TqExpr() (localctx ITqExprContext) {
	return p.tqExpr(0)
}

func (p *TqParser) tqExpr(_p int) (localctx ITqExprContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewTqExprContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx ITqExprContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 20
	p.EnterRecursionRule(localctx, 20, TqParserRULE_tqExpr, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(155)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 12, p.GetParserRuleContext()) {
	case 1:
		localctx = NewTqParenExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(142)
			p.Match(TqParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(143)
			p.tqExpr(0)
		}
		{
			p.SetState(144)
			p.Match(TqParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		localctx = NewTqUnaryExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(146)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*TqUnaryExprContext).op = _lt

			_la = p.GetTokenStream().LA(1)

			if !(_la == TqParserPLUS || _la == TqParserDASH) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*TqUnaryExprContext).op = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(147)
			p.tqExpr(12)
		}

	case 3:
		localctx = NewTqCallExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(148)
			p.TqFunctionCall()
		}

	case 4:
		localctx = NewTqColumnExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(149)
			p.Match(TqParserCOLUMN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 5:
		localctx = NewTqRefExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(150)
			p.Reference()
		}

	case 6:
		localctx = NewTqNumberExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(151)
			p.Match(TqParserNUMBER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 7:
		localctx = NewTqStringExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(152)
			p.Match(TqParserSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 8:
		localctx = NewTqBoolExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(153)
			_la = p.GetTokenStream().LA(1)

			if !(_la == TqParserTRUE || _la == TqParserFALSE) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	case 9:
		localctx = NewTqErrorExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(154)
			p.Match(TqParserERRORCONST)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(176)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 14, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(174)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 13, p.GetParserRuleContext()) {
			case 1:
				localctx = NewTqPowExprContext(p, NewTqExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, TqParserRULE_tqExpr)
				p.SetState(157)

				if !(p.Precpred(p.GetParserRuleContext(), 13)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 13)", ""))
					goto errorExit
				}
				{
					p.SetState(158)
					p.Match(TqParserCARET)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(159)
					p.tqExpr(13)
				}

			case 2:
				localctx = NewTqMulExprContext(p, NewTqExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, TqParserRULE_tqExpr)
				p.SetState(160)

				if !(p.Precpred(p.GetParserRuleContext(), 11)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 11)", ""))
					goto errorExit
				}
				{
					p.SetState(161)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*TqMulExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == TqParserSTAR || _la == TqParserSLASH) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*TqMulExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(162)
					p.tqExpr(12)
				}

			case 3:
				localctx = NewTqAddExprContext(p, NewTqExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, TqParserRULE_tqExpr)
				p.SetState(163)

				if !(p.Precpred(p.GetParserRuleContext(), 10)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 10)", ""))
					goto errorExit
				}
				{
					p.SetState(164)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*TqAddExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == TqParserPLUS || _la == TqParserDASH) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*TqAddExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(165)
					p.tqExpr(11)
				}

			case 4:
				localctx = NewTqConcatExprContext(p, NewTqExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, TqParserRULE_tqExpr)
				p.SetState(166)

				if !(p.Precpred(p.GetParserRuleContext(), 9)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 9)", ""))
					goto errorExit
				}
				{
					p.SetState(167)
					p.Match(TqParserAMP)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(168)
					p.tqExpr(10)
				}

			case 5:
				localctx = NewTqCompareExprContext(p, NewTqExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, TqParserRULE_tqExpr)
				p.SetState(169)

				if !(p.Precpred(p.GetParserRuleContext(), 8)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 8)", ""))
					goto errorExit
				}
				{
					p.SetState(170)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*TqCompareExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&18808832) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*TqCompareExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(171)
					p.tqExpr(9)
				}

			case 6:
				localctx = NewTqPercentExprContext(p, NewTqExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, TqParserRULE_tqExpr)
				p.SetState(172)

				if !(p.Precpred(p.GetParserRuleContext(), 14)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 14)", ""))
					goto errorExit
				}
				{
					p.SetState(173)
					p.Match(TqParserPERCENT)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(178)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 14, p.GetParserRuleContext())
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
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITqFunctionCallContext is an interface to support dynamic dispatch.
type ITqFunctionCallContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LPAREN() antlr.TerminalNode
	RPAREN() antlr.TerminalNode
	NAME() antlr.TerminalNode
	COL() antlr.TerminalNode
	VerbKeyword() IVerbKeywordContext
	NUMBER() antlr.TerminalNode
	TqArgList() ITqArgListContext

	// IsTqFunctionCallContext differentiates from other interfaces.
	IsTqFunctionCallContext()
}

type TqFunctionCallContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTqFunctionCallContext() *TqFunctionCallContext {
	var p = new(TqFunctionCallContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TqParserRULE_tqFunctionCall
	return p
}

func InitEmptyTqFunctionCallContext(p *TqFunctionCallContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TqParserRULE_tqFunctionCall
}

func (*TqFunctionCallContext) IsTqFunctionCallContext() {}

func NewTqFunctionCallContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TqFunctionCallContext {
	var p = new(TqFunctionCallContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TqParserRULE_tqFunctionCall

	return p
}

func (s *TqFunctionCallContext) GetParser() antlr.Parser { return s.parser }

func (s *TqFunctionCallContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(TqParserLPAREN, 0)
}

func (s *TqFunctionCallContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(TqParserRPAREN, 0)
}

func (s *TqFunctionCallContext) NAME() antlr.TerminalNode {
	return s.GetToken(TqParserNAME, 0)
}

func (s *TqFunctionCallContext) COL() antlr.TerminalNode {
	return s.GetToken(TqParserCOL, 0)
}

func (s *TqFunctionCallContext) VerbKeyword() IVerbKeywordContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVerbKeywordContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVerbKeywordContext)
}

func (s *TqFunctionCallContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(TqParserNUMBER, 0)
}

func (s *TqFunctionCallContext) TqArgList() ITqArgListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITqArgListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITqArgListContext)
}

func (s *TqFunctionCallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TqFunctionCallContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TqFunctionCallContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TqParserListener); ok {
		listenerT.EnterTqFunctionCall(s)
	}
}

func (s *TqFunctionCallContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TqParserListener); ok {
		listenerT.ExitTqFunctionCall(s)
	}
}

func (s *TqFunctionCallContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TqParserVisitor:
		return t.VisitTqFunctionCall(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TqParser) TqFunctionCall() (localctx ITqFunctionCallContext) {
	localctx = NewTqFunctionCallContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, TqParserRULE_tqFunctionCall)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(182)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case TqParserNAME:
		{
			p.SetState(179)
			p.Match(TqParserNAME)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case TqParserCOL:
		{
			p.SetState(180)
			p.Match(TqParserCOL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case TqParserSELECT, TqParserDROP, TqParserWHERE, TqParserDERIVE, TqParserRENAME, TqParserSORT, TqParserDISTINCT, TqParserLIMIT, TqParserOFFSET, TqParserGROUP, TqParserAS:
		{
			p.SetState(181)
			p.VerbKeyword()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	p.SetState(185)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == TqParserNUMBER {
		{
			p.SetState(184)
			p.Match(TqParserNUMBER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(187)
		p.Match(TqParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(189)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&8253364785150) != 0 {
		{
			p.SetState(188)
			p.TqArgList()
		}

	}
	{
		p.SetState(191)
		p.Match(TqParserRPAREN)
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

// ITqArgListContext is an interface to support dynamic dispatch.
type ITqArgListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllTqExpr() []ITqExprContext
	TqExpr(i int) ITqExprContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsTqArgListContext differentiates from other interfaces.
	IsTqArgListContext()
}

type TqArgListContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTqArgListContext() *TqArgListContext {
	var p = new(TqArgListContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TqParserRULE_tqArgList
	return p
}

func InitEmptyTqArgListContext(p *TqArgListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TqParserRULE_tqArgList
}

func (*TqArgListContext) IsTqArgListContext() {}

func NewTqArgListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TqArgListContext {
	var p = new(TqArgListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TqParserRULE_tqArgList

	return p
}

func (s *TqArgListContext) GetParser() antlr.Parser { return s.parser }

func (s *TqArgListContext) AllTqExpr() []ITqExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITqExprContext); ok {
			len++
		}
	}

	tst := make([]ITqExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITqExprContext); ok {
			tst[i] = t.(ITqExprContext)
			i++
		}
	}

	return tst
}

func (s *TqArgListContext) TqExpr(i int) ITqExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITqExprContext); ok {
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

	return t.(ITqExprContext)
}

func (s *TqArgListContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(TqParserCOMMA)
}

func (s *TqArgListContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(TqParserCOMMA, i)
}

func (s *TqArgListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TqArgListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TqArgListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TqParserListener); ok {
		listenerT.EnterTqArgList(s)
	}
}

func (s *TqArgListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TqParserListener); ok {
		listenerT.ExitTqArgList(s)
	}
}

func (s *TqArgListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TqParserVisitor:
		return t.VisitTqArgList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TqParser) TqArgList() (localctx ITqArgListContext) {
	localctx = NewTqArgListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, TqParserRULE_tqArgList)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(193)
		p.tqExpr(0)
	}
	p.SetState(198)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == TqParserCOMMA {
		{
			p.SetState(194)
			p.Match(TqParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(195)
			p.tqExpr(0)
		}

		p.SetState(200)
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

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TqParserRULE_expression
	return p
}

func InitEmptyExpressionContext(p *ExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TqParserRULE_expression
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TqParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) CopyAll(ctx *ExpressionContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ErrorExprContext struct {
	ExpressionContext
}

func NewErrorExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ErrorExprContext {
	var p = new(ErrorExprContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *ErrorExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ErrorExprContext) ERRORCONST() antlr.TerminalNode {
	return s.GetToken(TqParserERRORCONST, 0)
}

func (s *ErrorExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TqParserListener); ok {
		listenerT.EnterErrorExpr(s)
	}
}

func (s *ErrorExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TqParserListener); ok {
		listenerT.ExitErrorExpr(s)
	}
}

func (s *ErrorExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TqParserVisitor:
		return t.VisitErrorExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type PipeExprContext struct {
	ExpressionContext
}

func NewPipeExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PipeExprContext {
	var p = new(PipeExprContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *PipeExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PipeExprContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *PipeExprContext) PIPE() antlr.TerminalNode {
	return s.GetToken(TqParserPIPE, 0)
}

func (s *PipeExprContext) FunctionCall() IFunctionCallContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunctionCallContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunctionCallContext)
}

func (s *PipeExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TqParserListener); ok {
		listenerT.EnterPipeExpr(s)
	}
}

func (s *PipeExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TqParserListener); ok {
		listenerT.ExitPipeExpr(s)
	}
}

func (s *PipeExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TqParserVisitor:
		return t.VisitPipeExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type NumberExprContext struct {
	ExpressionContext
}

func NewNumberExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NumberExprContext {
	var p = new(NumberExprContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *NumberExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumberExprContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(TqParserNUMBER, 0)
}

func (s *NumberExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TqParserListener); ok {
		listenerT.EnterNumberExpr(s)
	}
}

func (s *NumberExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TqParserListener); ok {
		listenerT.ExitNumberExpr(s)
	}
}

func (s *NumberExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TqParserVisitor:
		return t.VisitNumberExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type ParenExprContext struct {
	ExpressionContext
}

func NewParenExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ParenExprContext {
	var p = new(ParenExprContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *ParenExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParenExprContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(TqParserLPAREN, 0)
}

func (s *ParenExprContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ParenExprContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(TqParserRPAREN, 0)
}

func (s *ParenExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TqParserListener); ok {
		listenerT.EnterParenExpr(s)
	}
}

func (s *ParenExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TqParserListener); ok {
		listenerT.ExitParenExpr(s)
	}
}

func (s *ParenExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TqParserVisitor:
		return t.VisitParenExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type ConcatExprContext struct {
	ExpressionContext
}

func NewConcatExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ConcatExprContext {
	var p = new(ConcatExprContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *ConcatExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConcatExprContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *ConcatExprContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
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

	return t.(IExpressionContext)
}

func (s *ConcatExprContext) AMP() antlr.TerminalNode {
	return s.GetToken(TqParserAMP, 0)
}

func (s *ConcatExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TqParserListener); ok {
		listenerT.EnterConcatExpr(s)
	}
}

func (s *ConcatExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TqParserListener); ok {
		listenerT.ExitConcatExpr(s)
	}
}

func (s *ConcatExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TqParserVisitor:
		return t.VisitConcatExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type StringExprContext struct {
	ExpressionContext
}

func NewStringExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StringExprContext {
	var p = new(StringExprContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *StringExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringExprContext) STRING() antlr.TerminalNode {
	return s.GetToken(TqParserSTRING, 0)
}

func (s *StringExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TqParserListener); ok {
		listenerT.EnterStringExpr(s)
	}
}

func (s *StringExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TqParserListener); ok {
		listenerT.ExitStringExpr(s)
	}
}

func (s *StringExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TqParserVisitor:
		return t.VisitStringExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type UnaryExprContext struct {
	ExpressionContext
	op antlr.Token
}

func NewUnaryExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *UnaryExprContext {
	var p = new(UnaryExprContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *UnaryExprContext) GetOp() antlr.Token { return s.op }

func (s *UnaryExprContext) SetOp(v antlr.Token) { s.op = v }

func (s *UnaryExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnaryExprContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *UnaryExprContext) PLUS() antlr.TerminalNode {
	return s.GetToken(TqParserPLUS, 0)
}

func (s *UnaryExprContext) DASH() antlr.TerminalNode {
	return s.GetToken(TqParserDASH, 0)
}

func (s *UnaryExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TqParserListener); ok {
		listenerT.EnterUnaryExpr(s)
	}
}

func (s *UnaryExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TqParserListener); ok {
		listenerT.ExitUnaryExpr(s)
	}
}

func (s *UnaryExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TqParserVisitor:
		return t.VisitUnaryExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type AddExprContext struct {
	ExpressionContext
	op antlr.Token
}

func NewAddExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AddExprContext {
	var p = new(AddExprContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *AddExprContext) GetOp() antlr.Token { return s.op }

func (s *AddExprContext) SetOp(v antlr.Token) { s.op = v }

func (s *AddExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AddExprContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *AddExprContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
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

	return t.(IExpressionContext)
}

func (s *AddExprContext) PLUS() antlr.TerminalNode {
	return s.GetToken(TqParserPLUS, 0)
}

func (s *AddExprContext) DASH() antlr.TerminalNode {
	return s.GetToken(TqParserDASH, 0)
}

func (s *AddExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TqParserListener); ok {
		listenerT.EnterAddExpr(s)
	}
}

func (s *AddExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TqParserListener); ok {
		listenerT.ExitAddExpr(s)
	}
}

func (s *AddExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TqParserVisitor:
		return t.VisitAddExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type RefExprContext struct {
	ExpressionContext
}

func NewRefExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *RefExprContext {
	var p = new(RefExprContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *RefExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RefExprContext) Reference() IReferenceContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IReferenceContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IReferenceContext)
}

func (s *RefExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TqParserListener); ok {
		listenerT.EnterRefExpr(s)
	}
}

func (s *RefExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TqParserListener); ok {
		listenerT.ExitRefExpr(s)
	}
}

func (s *RefExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TqParserVisitor:
		return t.VisitRefExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type MulExprContext struct {
	ExpressionContext
	op antlr.Token
}

func NewMulExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MulExprContext {
	var p = new(MulExprContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *MulExprContext) GetOp() antlr.Token { return s.op }

func (s *MulExprContext) SetOp(v antlr.Token) { s.op = v }

func (s *MulExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MulExprContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *MulExprContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
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

	return t.(IExpressionContext)
}

func (s *MulExprContext) STAR() antlr.TerminalNode {
	return s.GetToken(TqParserSTAR, 0)
}

func (s *MulExprContext) SLASH() antlr.TerminalNode {
	return s.GetToken(TqParserSLASH, 0)
}

func (s *MulExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TqParserListener); ok {
		listenerT.EnterMulExpr(s)
	}
}

func (s *MulExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TqParserListener); ok {
		listenerT.ExitMulExpr(s)
	}
}

func (s *MulExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TqParserVisitor:
		return t.VisitMulExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type PercentExprContext struct {
	ExpressionContext
}

func NewPercentExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PercentExprContext {
	var p = new(PercentExprContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *PercentExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PercentExprContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *PercentExprContext) PERCENT() antlr.TerminalNode {
	return s.GetToken(TqParserPERCENT, 0)
}

func (s *PercentExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TqParserListener); ok {
		listenerT.EnterPercentExpr(s)
	}
}

func (s *PercentExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TqParserListener); ok {
		listenerT.ExitPercentExpr(s)
	}
}

func (s *PercentExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TqParserVisitor:
		return t.VisitPercentExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type CallExprContext struct {
	ExpressionContext
}

func NewCallExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CallExprContext {
	var p = new(CallExprContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *CallExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CallExprContext) FunctionCall() IFunctionCallContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunctionCallContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunctionCallContext)
}

func (s *CallExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TqParserListener); ok {
		listenerT.EnterCallExpr(s)
	}
}

func (s *CallExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TqParserListener); ok {
		listenerT.ExitCallExpr(s)
	}
}

func (s *CallExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TqParserVisitor:
		return t.VisitCallExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type BoolExprContext struct {
	ExpressionContext
}

func NewBoolExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BoolExprContext {
	var p = new(BoolExprContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *BoolExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BoolExprContext) TRUE() antlr.TerminalNode {
	return s.GetToken(TqParserTRUE, 0)
}

func (s *BoolExprContext) FALSE() antlr.TerminalNode {
	return s.GetToken(TqParserFALSE, 0)
}

func (s *BoolExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TqParserListener); ok {
		listenerT.EnterBoolExpr(s)
	}
}

func (s *BoolExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TqParserListener); ok {
		listenerT.ExitBoolExpr(s)
	}
}

func (s *BoolExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TqParserVisitor:
		return t.VisitBoolExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type PowExprContext struct {
	ExpressionContext
}

func NewPowExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PowExprContext {
	var p = new(PowExprContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *PowExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PowExprContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *PowExprContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
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

	return t.(IExpressionContext)
}

func (s *PowExprContext) CARET() antlr.TerminalNode {
	return s.GetToken(TqParserCARET, 0)
}

func (s *PowExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TqParserListener); ok {
		listenerT.EnterPowExpr(s)
	}
}

func (s *PowExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TqParserListener); ok {
		listenerT.ExitPowExpr(s)
	}
}

func (s *PowExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TqParserVisitor:
		return t.VisitPowExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type CompareExprContext struct {
	ExpressionContext
	op antlr.Token
}

func NewCompareExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CompareExprContext {
	var p = new(CompareExprContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *CompareExprContext) GetOp() antlr.Token { return s.op }

func (s *CompareExprContext) SetOp(v antlr.Token) { s.op = v }

func (s *CompareExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CompareExprContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *CompareExprContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
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

	return t.(IExpressionContext)
}

func (s *CompareExprContext) EQ() antlr.TerminalNode {
	return s.GetToken(TqParserEQ, 0)
}

func (s *CompareExprContext) NE() antlr.TerminalNode {
	return s.GetToken(TqParserNE, 0)
}

func (s *CompareExprContext) LT() antlr.TerminalNode {
	return s.GetToken(TqParserLT, 0)
}

func (s *CompareExprContext) LE() antlr.TerminalNode {
	return s.GetToken(TqParserLE, 0)
}

func (s *CompareExprContext) GT() antlr.TerminalNode {
	return s.GetToken(TqParserGT, 0)
}

func (s *CompareExprContext) GE() antlr.TerminalNode {
	return s.GetToken(TqParserGE, 0)
}

func (s *CompareExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TqParserListener); ok {
		listenerT.EnterCompareExpr(s)
	}
}

func (s *CompareExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TqParserListener); ok {
		listenerT.ExitCompareExpr(s)
	}
}

func (s *CompareExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TqParserVisitor:
		return t.VisitCompareExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TqParser) Expression() (localctx IExpressionContext) {
	return p.expression(0)
}

func (p *TqParser) expression(_p int) (localctx IExpressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 26
	p.EnterRecursionRule(localctx, 26, TqParserRULE_expression, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(214)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 19, p.GetParserRuleContext()) {
	case 1:
		localctx = NewParenExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(202)
			p.Match(TqParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(203)
			p.expression(0)
		}
		{
			p.SetState(204)
			p.Match(TqParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		localctx = NewUnaryExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(206)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*UnaryExprContext).op = _lt

			_la = p.GetTokenStream().LA(1)

			if !(_la == TqParserPLUS || _la == TqParserDASH) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*UnaryExprContext).op = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(207)
			p.expression(12)
		}

	case 3:
		localctx = NewCallExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(208)
			p.FunctionCall()
		}

	case 4:
		localctx = NewRefExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(209)
			p.Reference()
		}

	case 5:
		localctx = NewNumberExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(210)
			p.Match(TqParserNUMBER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 6:
		localctx = NewStringExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(211)
			p.Match(TqParserSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 7:
		localctx = NewBoolExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(212)
			_la = p.GetTokenStream().LA(1)

			if !(_la == TqParserTRUE || _la == TqParserFALSE) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	case 8:
		localctx = NewErrorExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(213)
			p.Match(TqParserERRORCONST)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(238)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 21, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(236)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 20, p.GetParserRuleContext()) {
			case 1:
				localctx = NewPowExprContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, TqParserRULE_expression)
				p.SetState(216)

				if !(p.Precpred(p.GetParserRuleContext(), 13)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 13)", ""))
					goto errorExit
				}
				{
					p.SetState(217)
					p.Match(TqParserCARET)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(218)
					p.expression(13)
				}

			case 2:
				localctx = NewMulExprContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, TqParserRULE_expression)
				p.SetState(219)

				if !(p.Precpred(p.GetParserRuleContext(), 11)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 11)", ""))
					goto errorExit
				}
				{
					p.SetState(220)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*MulExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == TqParserSTAR || _la == TqParserSLASH) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*MulExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(221)
					p.expression(12)
				}

			case 3:
				localctx = NewAddExprContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, TqParserRULE_expression)
				p.SetState(222)

				if !(p.Precpred(p.GetParserRuleContext(), 10)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 10)", ""))
					goto errorExit
				}
				{
					p.SetState(223)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*AddExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == TqParserPLUS || _la == TqParserDASH) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*AddExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(224)
					p.expression(11)
				}

			case 4:
				localctx = NewConcatExprContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, TqParserRULE_expression)
				p.SetState(225)

				if !(p.Precpred(p.GetParserRuleContext(), 9)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 9)", ""))
					goto errorExit
				}
				{
					p.SetState(226)
					p.Match(TqParserAMP)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(227)
					p.expression(10)
				}

			case 5:
				localctx = NewCompareExprContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, TqParserRULE_expression)
				p.SetState(228)

				if !(p.Precpred(p.GetParserRuleContext(), 8)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 8)", ""))
					goto errorExit
				}
				{
					p.SetState(229)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*CompareExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&18808832) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*CompareExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(230)
					p.expression(9)
				}

			case 6:
				localctx = NewPercentExprContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, TqParserRULE_expression)
				p.SetState(231)

				if !(p.Precpred(p.GetParserRuleContext(), 14)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 14)", ""))
					goto errorExit
				}
				{
					p.SetState(232)
					p.Match(TqParserPERCENT)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			case 7:
				localctx = NewPipeExprContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, TqParserRULE_expression)
				p.SetState(233)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
					goto errorExit
				}
				{
					p.SetState(234)
					p.Match(TqParserPIPE)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(235)
					p.FunctionCall()
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(240)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 21, p.GetParserRuleContext())
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
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFunctionCallContext is an interface to support dynamic dispatch.
type IFunctionCallContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LPAREN() antlr.TerminalNode
	RPAREN() antlr.TerminalNode
	NAME() antlr.TerminalNode
	COL() antlr.TerminalNode
	NUMBER() antlr.TerminalNode
	ArgList() IArgListContext

	// IsFunctionCallContext differentiates from other interfaces.
	IsFunctionCallContext()
}

type FunctionCallContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionCallContext() *FunctionCallContext {
	var p = new(FunctionCallContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TqParserRULE_functionCall
	return p
}

func InitEmptyFunctionCallContext(p *FunctionCallContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TqParserRULE_functionCall
}

func (*FunctionCallContext) IsFunctionCallContext() {}

func NewFunctionCallContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionCallContext {
	var p = new(FunctionCallContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TqParserRULE_functionCall

	return p
}

func (s *FunctionCallContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionCallContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(TqParserLPAREN, 0)
}

func (s *FunctionCallContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(TqParserRPAREN, 0)
}

func (s *FunctionCallContext) NAME() antlr.TerminalNode {
	return s.GetToken(TqParserNAME, 0)
}

func (s *FunctionCallContext) COL() antlr.TerminalNode {
	return s.GetToken(TqParserCOL, 0)
}

func (s *FunctionCallContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(TqParserNUMBER, 0)
}

func (s *FunctionCallContext) ArgList() IArgListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArgListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArgListContext)
}

func (s *FunctionCallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionCallContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionCallContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TqParserListener); ok {
		listenerT.EnterFunctionCall(s)
	}
}

func (s *FunctionCallContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TqParserListener); ok {
		listenerT.ExitFunctionCall(s)
	}
}

func (s *FunctionCallContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TqParserVisitor:
		return t.VisitFunctionCall(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TqParser) FunctionCall() (localctx IFunctionCallContext) {
	localctx = NewFunctionCallContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, TqParserRULE_functionCall)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(241)
		_la = p.GetTokenStream().LA(1)

		if !(_la == TqParserCOL || _la == TqParserNAME) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	p.SetState(243)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == TqParserNUMBER {
		{
			p.SetState(242)
			p.Match(TqParserNUMBER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(245)
		p.Match(TqParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(247)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&8253364764672) != 0 {
		{
			p.SetState(246)
			p.ArgList()
		}

	}
	{
		p.SetState(249)
		p.Match(TqParserRPAREN)
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

// IArgListContext is an interface to support dynamic dispatch.
type IArgListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllExpression() []IExpressionContext
	Expression(i int) IExpressionContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsArgListContext differentiates from other interfaces.
	IsArgListContext()
}

type ArgListContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArgListContext() *ArgListContext {
	var p = new(ArgListContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TqParserRULE_argList
	return p
}

func InitEmptyArgListContext(p *ArgListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TqParserRULE_argList
}

func (*ArgListContext) IsArgListContext() {}

func NewArgListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArgListContext {
	var p = new(ArgListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TqParserRULE_argList

	return p
}

func (s *ArgListContext) GetParser() antlr.Parser { return s.parser }

func (s *ArgListContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *ArgListContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
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

	return t.(IExpressionContext)
}

func (s *ArgListContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(TqParserCOMMA)
}

func (s *ArgListContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(TqParserCOMMA, i)
}

func (s *ArgListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArgListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArgListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TqParserListener); ok {
		listenerT.EnterArgList(s)
	}
}

func (s *ArgListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TqParserListener); ok {
		listenerT.ExitArgList(s)
	}
}

func (s *ArgListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TqParserVisitor:
		return t.VisitArgList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TqParser) ArgList() (localctx IArgListContext) {
	localctx = NewArgListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, TqParserRULE_argList)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(251)
		p.expression(0)
	}
	p.SetState(256)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == TqParserCOMMA {
		{
			p.SetState(252)
			p.Match(TqParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(253)
			p.expression(0)
		}

		p.SetState(258)
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

// IReferenceContext is an interface to support dynamic dispatch.
type IReferenceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllCellRef() []ICellRefContext
	CellRef(i int) ICellRefContext
	SheetQualifier() ISheetQualifierContext
	COLON() antlr.TerminalNode

	// IsReferenceContext differentiates from other interfaces.
	IsReferenceContext()
}

type ReferenceContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReferenceContext() *ReferenceContext {
	var p = new(ReferenceContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TqParserRULE_reference
	return p
}

func InitEmptyReferenceContext(p *ReferenceContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TqParserRULE_reference
}

func (*ReferenceContext) IsReferenceContext() {}

func NewReferenceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReferenceContext {
	var p = new(ReferenceContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TqParserRULE_reference

	return p
}

func (s *ReferenceContext) GetParser() antlr.Parser { return s.parser }

func (s *ReferenceContext) AllCellRef() []ICellRefContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ICellRefContext); ok {
			len++
		}
	}

	tst := make([]ICellRefContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ICellRefContext); ok {
			tst[i] = t.(ICellRefContext)
			i++
		}
	}

	return tst
}

func (s *ReferenceContext) CellRef(i int) ICellRefContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICellRefContext); ok {
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

	return t.(ICellRefContext)
}

func (s *ReferenceContext) SheetQualifier() ISheetQualifierContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISheetQualifierContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISheetQualifierContext)
}

func (s *ReferenceContext) COLON() antlr.TerminalNode {
	return s.GetToken(TqParserCOLON, 0)
}

func (s *ReferenceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReferenceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ReferenceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TqParserListener); ok {
		listenerT.EnterReference(s)
	}
}

func (s *ReferenceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TqParserListener); ok {
		listenerT.ExitReference(s)
	}
}

func (s *ReferenceContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TqParserVisitor:
		return t.VisitReference(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TqParser) Reference() (localctx IReferenceContext) {
	localctx = NewReferenceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, TqParserRULE_reference)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(260)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == TqParserSTRING {
		{
			p.SetState(259)
			p.SheetQualifier()
		}

	}
	{
		p.SetState(262)
		p.CellRef()
	}
	p.SetState(265)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 26, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(263)
			p.Match(TqParserCOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(264)
			p.CellRef()
		}

	} else if p.HasError() { // JIM
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

// ISheetQualifierContext is an interface to support dynamic dispatch.
type ISheetQualifierContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	STRING() antlr.TerminalNode
	BANG() antlr.TerminalNode

	// IsSheetQualifierContext differentiates from other interfaces.
	IsSheetQualifierContext()
}

type SheetQualifierContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySheetQualifierContext() *SheetQualifierContext {
	var p = new(SheetQualifierContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TqParserRULE_sheetQualifier
	return p
}

func InitEmptySheetQualifierContext(p *SheetQualifierContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TqParserRULE_sheetQualifier
}

func (*SheetQualifierContext) IsSheetQualifierContext() {}

func NewSheetQualifierContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SheetQualifierContext {
	var p = new(SheetQualifierContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TqParserRULE_sheetQualifier

	return p
}

func (s *SheetQualifierContext) GetParser() antlr.Parser { return s.parser }

func (s *SheetQualifierContext) STRING() antlr.TerminalNode {
	return s.GetToken(TqParserSTRING, 0)
}

func (s *SheetQualifierContext) BANG() antlr.TerminalNode {
	return s.GetToken(TqParserBANG, 0)
}

func (s *SheetQualifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SheetQualifierContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SheetQualifierContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TqParserListener); ok {
		listenerT.EnterSheetQualifier(s)
	}
}

func (s *SheetQualifierContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TqParserListener); ok {
		listenerT.ExitSheetQualifier(s)
	}
}

func (s *SheetQualifierContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TqParserVisitor:
		return t.VisitSheetQualifier(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TqParser) SheetQualifier() (localctx ISheetQualifierContext) {
	localctx = NewSheetQualifierContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, TqParserRULE_sheetQualifier)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(267)
		p.Match(TqParserSTRING)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(268)
		p.Match(TqParserBANG)
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

// ICellRefContext is an interface to support dynamic dispatch.
type ICellRefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	COL() antlr.TerminalNode
	NUMBER() antlr.TerminalNode
	AllDOLLAR() []antlr.TerminalNode
	DOLLAR(i int) antlr.TerminalNode

	// IsCellRefContext differentiates from other interfaces.
	IsCellRefContext()
}

type CellRefContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCellRefContext() *CellRefContext {
	var p = new(CellRefContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TqParserRULE_cellRef
	return p
}

func InitEmptyCellRefContext(p *CellRefContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TqParserRULE_cellRef
}

func (*CellRefContext) IsCellRefContext() {}

func NewCellRefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CellRefContext {
	var p = new(CellRefContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TqParserRULE_cellRef

	return p
}

func (s *CellRefContext) GetParser() antlr.Parser { return s.parser }

func (s *CellRefContext) COL() antlr.TerminalNode {
	return s.GetToken(TqParserCOL, 0)
}

func (s *CellRefContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(TqParserNUMBER, 0)
}

func (s *CellRefContext) AllDOLLAR() []antlr.TerminalNode {
	return s.GetTokens(TqParserDOLLAR)
}

func (s *CellRefContext) DOLLAR(i int) antlr.TerminalNode {
	return s.GetToken(TqParserDOLLAR, i)
}

func (s *CellRefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CellRefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CellRefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TqParserListener); ok {
		listenerT.EnterCellRef(s)
	}
}

func (s *CellRefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TqParserListener); ok {
		listenerT.ExitCellRef(s)
	}
}

func (s *CellRefContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TqParserVisitor:
		return t.VisitCellRef(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TqParser) CellRef() (localctx ICellRefContext) {
	localctx = NewCellRefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, TqParserRULE_cellRef)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(271)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == TqParserDOLLAR {
		{
			p.SetState(270)
			p.Match(TqParserDOLLAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(273)
		p.Match(TqParserCOL)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(275)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == TqParserDOLLAR {
		{
			p.SetState(274)
			p.Match(TqParserDOLLAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(277)
		p.Match(TqParserNUMBER)
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

func (p *TqParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 10:
		var t *TqExprContext = nil
		if localctx != nil {
			t = localctx.(*TqExprContext)
		}
		return p.TqExpr_Sempred(t, predIndex)

	case 13:
		var t *ExpressionContext = nil
		if localctx != nil {
			t = localctx.(*ExpressionContext)
		}
		return p.Expression_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *TqParser) TqExpr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 13)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 11)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 10)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 9)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 8)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 14)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *TqParser) Expression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 6:
		return p.Precpred(p.GetParserRuleContext(), 13)

	case 7:
		return p.Precpred(p.GetParserRuleContext(), 11)

	case 8:
		return p.Precpred(p.GetParserRuleContext(), 10)

	case 9:
		return p.Precpred(p.GetParserRuleContext(), 9)

	case 10:
		return p.Precpred(p.GetParserRuleContext(), 8)

	case 11:
		return p.Precpred(p.GetParserRuleContext(), 14)

	case 12:
		return p.Precpred(p.GetParserRuleContext(), 7)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
