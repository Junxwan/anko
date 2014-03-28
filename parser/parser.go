//line parser.go.y:2
package parser

import __yyfmt__ "fmt"

//line parser.go.y:2
import (
	"github.com/mattn/anko/ast"
)

//line parser.go.y:22
type yySymType struct {
	yys          int
	stmt_func    ast.Stmt
	stmt_if      ast.Stmt
	stmt_if_else ast.Stmt
	stmt_for     ast.Stmt
	stmts        []ast.Stmt
	stmt         ast.Stmt
	teim         ast.Expr
	expr         ast.Expr
	tok          Token
	idents       []string
	exprs        []ast.Expr
	pair         *ast.PairExpr
	pairs        []*ast.PairExpr
}

const IDENT = 57346
const NUMBER = 57347
const STRING = 57348
const ARRAY = 57349
const VAR = 57350
const FUNC = 57351
const RETURN = 57352
const IF = 57353
const ELSE = 57354
const FOR = 57355
const IN = 57356
const EQ = 57357
const NE = 57358
const GE = 57359
const LE = 57360
const UNARY = 57361

var yyToknames = []string{
	"IDENT",
	"NUMBER",
	"STRING",
	"ARRAY",
	"VAR",
	"FUNC",
	"RETURN",
	"IF",
	"ELSE",
	"FOR",
	"IN",
	"EQ",
	"NE",
	"GE",
	"LE",
	" +",
	" -",
	" *",
	" /",
	" %",
	"UNARY",
}
var yyStatenames = []string{}

const yyEofCode = 1
const yyErrCode = 2
const yyMaxDepth = 200

//line parser.go.y:247

//line yacctab:1
var yyExca = []int{
	-1, 1,
	1, -1,
	-2, 0,
}

const yyNprod = 43
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 259

var yyAct = []int{

	1, 7, 48, 20, 21, 22, 23, 24, 73, 99,
	69, 39, 45, 70, 101, 86, 97, 44, 87, 46,
	81, 50, 90, 69, 72, 71, 96, 95, 51, 52,
	53, 54, 55, 56, 57, 58, 59, 60, 61, 62,
	88, 65, 43, 66, 41, 46, 32, 33, 35, 37,
	27, 28, 29, 30, 31, 63, 68, 67, 98, 49,
	89, 91, 78, 42, 26, 76, 34, 36, 40, 80,
	38, 82, 77, 47, 83, 84, 32, 33, 35, 37,
	27, 28, 29, 30, 31, 6, 85, 5, 4, 92,
	93, 94, 3, 2, 26, 0, 34, 36, 0, 0,
	100, 32, 33, 35, 37, 27, 28, 29, 30, 31,
	0, 0, 0, 0, 79, 0, 0, 0, 0, 26,
	0, 34, 36, 32, 33, 35, 37, 27, 28, 29,
	30, 31, 32, 33, 35, 37, 27, 28, 29, 30,
	31, 26, 75, 34, 36, 74, 0, 0, 0, 0,
	26, 0, 34, 36, 32, 33, 35, 37, 27, 28,
	29, 30, 31, 0, 64, 0, 0, 32, 33, 35,
	37, 0, 26, 0, 34, 36, 32, 33, 35, 37,
	27, 28, 29, 30, 31, 26, 25, 34, 36, 0,
	0, 0, 0, 0, 26, 0, 34, 36, 32, 33,
	35, 37, 27, 28, 29, 30, 31, 32, 33, 35,
	37, 0, 0, 29, 30, 31, 26, 0, 34, 36,
	0, 14, 13, 16, 0, 26, 0, 34, 36, 14,
	13, 16, 0, 8, 10, 9, 11, 15, 12, 0,
	0, 0, 0, 0, 19, 15, 18, 0, 0, 0,
	17, 0, 19, 0, 18, 0, 0, 0, 17,
}
var yyPact = []int{

	225, -1000, 225, 225, 225, 225, 225, 161, 66, 217,
	64, 17, 59, -1000, 15, 217, -1000, 217, 53, 217,
	-1000, -1000, -1000, -1000, -1000, -1000, 217, 217, 217, 217,
	217, 217, 217, 217, 217, 217, 217, 217, 29, 139,
	14, 217, 43, 217, 152, -21, 183, -6, -1000, -24,
	117, 108, 192, 192, 152, 152, 152, 183, 183, 183,
	183, 183, 183, 217, -1000, 58, 86, 217, -8, 217,
	-1000, 53, -1000, 217, -1000, -1000, 61, -13, -1000, 11,
	31, -1000, 183, -1000, 183, -1000, -7, 57, 225, 225,
	225, -1000, -3, -4, -14, 46, -1000, -1000, -20, 225,
	-16, -1000,
}
var yyPgo = []int{

	0, 0, 93, 92, 88, 87, 85, 1, 12, 2,
	73, 72,
}
var yyR1 = []int{

	0, 1, 1, 1, 1, 1, 1, 2, 2, 2,
	3, 5, 4, 6, 11, 11, 9, 10, 10, 10,
	8, 8, 8, 7, 7, 7, 7, 7, 7, 7,
	7, 7, 7, 7, 7, 7, 7, 7, 7, 7,
	7, 7, 7,
}
var yyR2 = []int{

	0, 0, 2, 2, 2, 2, 2, 2, 5, 3,
	8, 11, 7, 7, 1, 3, 3, 0, 1, 3,
	0, 1, 3, 1, 1, 2, 1, 4, 3, 3,
	4, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3,
}
var yyChk = []int{

	-1000, -1, -2, -3, -4, -5, -6, -7, 8, 10,
	9, 11, 13, 5, 4, 20, 6, 33, 29, 27,
	-1, -1, -1, -1, -1, 25, 33, 19, 20, 21,
	22, 23, 15, 16, 35, 17, 36, 18, 4, -7,
	4, 27, 4, 27, -7, -8, -7, -10, -9, 6,
	-7, -7, -7, -7, -7, -7, -7, -7, -7, -7,
	-7, -7, -7, 26, 25, 27, -7, 14, -8, 31,
	34, 31, 30, 32, 28, 34, -7, -11, 4, 28,
	-7, 28, -7, -9, -7, 25, 28, 31, 29, 29,
	29, 4, -1, -1, -1, 30, 30, 30, 12, 29,
	-1, 30,
}
var yyDef = []int{

	1, -2, 1, 1, 1, 1, 1, 0, 0, 0,
	0, 0, 0, 23, 24, 0, 26, 20, 17, 0,
	2, 3, 4, 5, 6, 7, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 20, 25, 0, 21, 0, 18, 0,
	0, 0, 32, 33, 34, 35, 36, 37, 38, 39,
	40, 41, 42, 0, 9, 0, 0, 0, 0, 0,
	28, 0, 29, 0, 31, 30, 0, 0, 14, 0,
	0, 27, 22, 19, 16, 8, 0, 0, 1, 1,
	1, 15, 0, 0, 0, 12, 13, 10, 0, 1,
	0, 11,
}
var yyTok1 = []int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 23, 3, 3,
	27, 28, 21, 19, 31, 20, 3, 22, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 32, 25,
	36, 26, 35, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 33, 3, 34, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 29, 3, 30,
}
var yyTok2 = []int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 24,
}
var yyTok3 = []int{
	0,
}

//line yaccpar:1

/*	parser for yacc output	*/

var yyDebug = 0

type yyLexer interface {
	Lex(lval *yySymType) int
	Error(s string)
}

const yyFlag = -1000

func yyTokname(c int) string {
	// 4 is TOKSTART above
	if c >= 4 && c-4 < len(yyToknames) {
		if yyToknames[c-4] != "" {
			return yyToknames[c-4]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func yyStatname(s int) string {
	if s >= 0 && s < len(yyStatenames) {
		if yyStatenames[s] != "" {
			return yyStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func yylex1(lex yyLexer, lval *yySymType) int {
	c := 0
	char := lex.Lex(lval)
	if char <= 0 {
		c = yyTok1[0]
		goto out
	}
	if char < len(yyTok1) {
		c = yyTok1[char]
		goto out
	}
	if char >= yyPrivate {
		if char < yyPrivate+len(yyTok2) {
			c = yyTok2[char-yyPrivate]
			goto out
		}
	}
	for i := 0; i < len(yyTok3); i += 2 {
		c = yyTok3[i+0]
		if c == char {
			c = yyTok3[i+1]
			goto out
		}
	}

out:
	if c == 0 {
		c = yyTok2[1] /* unknown char */
	}
	if yyDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", yyTokname(c), uint(char))
	}
	return c
}

func yyParse(yylex yyLexer) int {
	var yyn int
	var yylval yySymType
	var yyVAL yySymType
	yyS := make([]yySymType, yyMaxDepth)

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	yystate := 0
	yychar := -1
	yyp := -1
	goto yystack

ret0:
	return 0

ret1:
	return 1

yystack:
	/* put a state and value onto the stack */
	if yyDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", yyTokname(yychar), yyStatname(yystate))
	}

	yyp++
	if yyp >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyS[yyp] = yyVAL
	yyS[yyp].yys = yystate

yynewstate:
	yyn = yyPact[yystate]
	if yyn <= yyFlag {
		goto yydefault /* simple state */
	}
	if yychar < 0 {
		yychar = yylex1(yylex, &yylval)
	}
	yyn += yychar
	if yyn < 0 || yyn >= yyLast {
		goto yydefault
	}
	yyn = yyAct[yyn]
	if yyChk[yyn] == yychar { /* valid shift */
		yychar = -1
		yyVAL = yylval
		yystate = yyn
		if Errflag > 0 {
			Errflag--
		}
		goto yystack
	}

yydefault:
	/* default state action */
	yyn = yyDef[yystate]
	if yyn == -2 {
		if yychar < 0 {
			yychar = yylex1(yylex, &yylval)
		}

		/* look through exception table */
		xi := 0
		for {
			if yyExca[xi+0] == -1 && yyExca[xi+1] == yystate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			yyn = yyExca[xi+0]
			if yyn < 0 || yyn == yychar {
				break
			}
		}
		yyn = yyExca[xi+1]
		if yyn < 0 {
			goto ret0
		}
	}
	if yyn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			yylex.Error("syntax error")
			Nerrs++
			if yyDebug >= 1 {
				__yyfmt__.Printf("%s", yyStatname(yystate))
				__yyfmt__.Printf(" saw %s\n", yyTokname(yychar))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for yyp >= 0 {
				yyn = yyPact[yyS[yyp].yys] + yyErrCode
				if yyn >= 0 && yyn < yyLast {
					yystate = yyAct[yyn] /* simulate a shift of "error" */
					if yyChk[yystate] == yyErrCode {
						goto yystack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if yyDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", yyS[yyp].yys)
				}
				yyp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if yyDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", yyTokname(yychar))
			}
			if yychar == yyEofCode {
				goto ret1
			}
			yychar = -1
			goto yynewstate /* try again in the same state */
		}
	}

	/* reduction by production yyn */
	if yyDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", yyn, yyStatname(yystate))
	}

	yynt := yyn
	yypt := yyp
	_ = yypt // guard against "declared and not used"

	yyp -= yyR2[yyn]
	yyVAL = yyS[yyp+1]

	/* consult goto table to find next state */
	yyn = yyR1[yyn]
	yyg := yyPgo[yyn]
	yyj := yyg + yyS[yyp].yys + 1

	if yyj >= yyLast {
		yystate = yyAct[yyg]
	} else {
		yystate = yyAct[yyj]
		if yyChk[yystate] != -yyn {
			yystate = yyAct[yyg]
		}
	}
	// dummy call; replaced with literal code
	switch yynt {

	case 1:
		//line parser.go.y:47
		{
			yyVAL.stmts = nil
			if l, ok := yylex.(*Lexer); ok {
				l.stmts = yyVAL.stmts
			}
		}
	case 2:
		//line parser.go.y:54
		{
			yyVAL.stmts = append([]ast.Stmt{yyS[yypt-1].stmt}, yyS[yypt-0].stmts...)
			if l, ok := yylex.(*Lexer); ok {
				l.stmts = yyVAL.stmts
			}
		}
	case 3:
		//line parser.go.y:61
		{
			yyVAL.stmts = append([]ast.Stmt{yyS[yypt-1].stmt_func}, yyS[yypt-0].stmts...)
			if l, ok := yylex.(*Lexer); ok {
				l.stmts = yyVAL.stmts
			}
		}
	case 4:
		//line parser.go.y:68
		{
			yyVAL.stmts = append([]ast.Stmt{yyS[yypt-1].stmt_if}, yyS[yypt-0].stmts...)
			if l, ok := yylex.(*Lexer); ok {
				l.stmts = yyVAL.stmts
			}
		}
	case 5:
		//line parser.go.y:75
		{
			yyVAL.stmts = append([]ast.Stmt{yyS[yypt-1].stmt_if_else}, yyS[yypt-0].stmts...)
			if l, ok := yylex.(*Lexer); ok {
				l.stmts = yyVAL.stmts
			}
		}
	case 6:
		//line parser.go.y:82
		{
			yyVAL.stmts = append([]ast.Stmt{yyS[yypt-1].stmt_for}, yyS[yypt-0].stmts...)
			if l, ok := yylex.(*Lexer); ok {
				l.stmts = yyVAL.stmts
			}
		}
	case 7:
		//line parser.go.y:90
		{
			yyVAL.stmt = &ast.ExprStmt{Expr: yyS[yypt-1].expr}
		}
	case 8:
		//line parser.go.y:94
		{
			yyVAL.stmt = &ast.VarStmt{Name: yyS[yypt-3].tok.lit, Expr: yyS[yypt-1].expr}
		}
	case 9:
		//line parser.go.y:98
		{
			yyVAL.stmt = &ast.ReturnStmt{Expr: yyS[yypt-1].expr}
		}
	case 10:
		//line parser.go.y:103
		{
			yyVAL.stmt_func = &ast.FuncStmt{Name: yyS[yypt-6].tok.lit, Args: yyS[yypt-4].idents, Stmts: yyS[yypt-1].stmts}
		}
	case 11:
		//line parser.go.y:108
		{
			yyVAL.stmt_if_else = &ast.IfStmt{Expr: yyS[yypt-8].expr, ThenStmts: yyS[yypt-5].stmts, ElseStmts: yyS[yypt-1].stmts}
		}
	case 12:
		//line parser.go.y:113
		{
			yyVAL.stmt_if = &ast.IfStmt{Expr: yyS[yypt-4].expr, ThenStmts: yyS[yypt-1].stmts}
		}
	case 13:
		//line parser.go.y:118
		{
			yyVAL.stmt_for = &ast.ForStmt{Var: yyS[yypt-5].tok.lit, Value: yyS[yypt-3].expr, Stmts: yyS[yypt-1].stmts}
		}
	case 14:
		//line parser.go.y:123
		{
			yyVAL.idents = []string{yyS[yypt-0].tok.lit}
		}
	case 15:
		//line parser.go.y:127
		{
			yyVAL.idents = append(yyS[yypt-2].idents, yyS[yypt-0].tok.lit)
		}
	case 16:
		//line parser.go.y:132
		{
			yyVAL.pair = &ast.PairExpr{Key: yyS[yypt-2].tok.lit, Value: yyS[yypt-0].expr}
		}
	case 17:
		//line parser.go.y:137
		{
			yyVAL.pairs = []*ast.PairExpr{}
		}
	case 18:
		//line parser.go.y:141
		{
			yyVAL.pairs = []*ast.PairExpr{yyS[yypt-0].pair}
		}
	case 19:
		//line parser.go.y:145
		{
			yyVAL.pairs = append(yyS[yypt-2].pairs, yyS[yypt-0].pair)
		}
	case 20:
		//line parser.go.y:150
		{
			yyVAL.exprs = []ast.Expr{}
		}
	case 21:
		//line parser.go.y:154
		{
			yyVAL.exprs = []ast.Expr{yyS[yypt-0].expr}
		}
	case 22:
		//line parser.go.y:158
		{
			yyVAL.exprs = append(yyS[yypt-2].exprs, yyS[yypt-0].expr)
		}
	case 23:
		//line parser.go.y:163
		{
			yyVAL.expr = &ast.NumberExpr{Lit: yyS[yypt-0].tok.lit}
		}
	case 24:
		//line parser.go.y:167
		{
			yyVAL.expr = &ast.IdentExpr{Lit: yyS[yypt-0].tok.lit}
		}
	case 25:
		//line parser.go.y:171
		{
			yyVAL.expr = &ast.UnaryMinusExpr{SubExpr: yyS[yypt-0].expr}
		}
	case 26:
		//line parser.go.y:175
		{
			yyVAL.expr = &ast.StringExpr{Lit: yyS[yypt-0].tok.lit}
		}
	case 27:
		//line parser.go.y:179
		{
			yyVAL.expr = &ast.CallExpr{Name: yyS[yypt-3].tok.lit, SubExprs: yyS[yypt-1].exprs}
		}
	case 28:
		//line parser.go.y:183
		{
			yyVAL.expr = &ast.ArrayExpr{Exprs: yyS[yypt-1].exprs}
		}
	case 29:
		//line parser.go.y:187
		{
			mapExpr := make(map[string]ast.Expr)
			for _, v := range yyS[yypt-1].pairs {
				mapExpr[v.Key] = v.Value
			}
			yyVAL.expr = &ast.MapExpr{MapExpr: mapExpr}
		}
	case 30:
		//line parser.go.y:195
		{
			yyVAL.expr = &ast.ItemExpr{Value: yyS[yypt-3].expr, Index: yyS[yypt-1].expr}
		}
	case 31:
		//line parser.go.y:199
		{
			yyVAL.expr = &ast.ParenExpr{SubExpr: yyS[yypt-1].expr}
		}
	case 32:
		//line parser.go.y:203
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "+", Rhs: yyS[yypt-0].expr}
		}
	case 33:
		//line parser.go.y:207
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "-", Rhs: yyS[yypt-0].expr}
		}
	case 34:
		//line parser.go.y:211
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "*", Rhs: yyS[yypt-0].expr}
		}
	case 35:
		//line parser.go.y:215
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "/", Rhs: yyS[yypt-0].expr}
		}
	case 36:
		//line parser.go.y:219
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "%", Rhs: yyS[yypt-0].expr}
		}
	case 37:
		//line parser.go.y:223
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "==", Rhs: yyS[yypt-0].expr}
		}
	case 38:
		//line parser.go.y:227
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "!=", Rhs: yyS[yypt-0].expr}
		}
	case 39:
		//line parser.go.y:231
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: ">", Rhs: yyS[yypt-0].expr}
		}
	case 40:
		//line parser.go.y:235
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: ">=", Rhs: yyS[yypt-0].expr}
		}
	case 41:
		//line parser.go.y:239
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "<", Rhs: yyS[yypt-0].expr}
		}
	case 42:
		//line parser.go.y:243
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyS[yypt-2].expr, Operator: "<=", Rhs: yyS[yypt-0].expr}
		}
	}
	goto yystack /* stack new state and value */
}
