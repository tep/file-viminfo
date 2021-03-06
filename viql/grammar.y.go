//line grammar.y:3
package viql

import __yyfmt__ "fmt"

//line grammar.y:3
import "toolman.org/text/scanner"

//line grammar.y:8
type yySymType struct {
	yys int
	lxr *lexer
	pos scanner.Position

	expr  Expression
	oper  *operation
	decl  declaration
	comp  *comparison
	value *value
	fld   field

	cmp comparitor
}

const ERROR = 57346
const CMP = 57347
const EQ = 57348
const NEQ = 57349
const REM = 57350
const NRE = 57351
const AND = 57352
const OR = 57353
const NOT = 57354
const DECL = 57355
const ALL = 57356
const MISSING = 57357
const MODIFIED = 57358
const NONE = 57359
const RUNNING = 57360
const THISHOST = 57361
const THISUSER = 57362
const FIELD = 57363
const CRYPTMETHOD = 57364
const FILEFORMAT = 57365
const FILENAME = 57366
const HOSTNAME = 57367
const INODE = 57368
const PID = 57369
const USER = 57370
const FFDOS = 57371
const FFMAC = 57372
const FFUNIX = 57373
const CMBLOWFISH = 57374
const CMBLOWFISH2 = 57375
const CMPLAINTEXT = 57376
const CMZIP = 57377
const VALUE = 57378

var yyToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"ERROR",
	"CMP",
	"EQ",
	"NEQ",
	"REM",
	"NRE",
	"AND",
	"OR",
	"NOT",
	"DECL",
	"ALL",
	"MISSING",
	"MODIFIED",
	"NONE",
	"RUNNING",
	"THISHOST",
	"THISUSER",
	"FIELD",
	"CRYPTMETHOD",
	"FILEFORMAT",
	"FILENAME",
	"HOSTNAME",
	"INODE",
	"PID",
	"USER",
	"FFDOS",
	"FFMAC",
	"FFUNIX",
	"CMBLOWFISH",
	"CMBLOWFISH2",
	"CMPLAINTEXT",
	"CMZIP",
	"VALUE",
	"'('",
	"')'",
}
var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyInitialStackSize = 16

//line grammar.y:94
//line yacctab:1
var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
}

const yyPrivate = 57344

const yyLast = 30

var yyAct = [...]int{

	12, 13, 22, 2, 8, 9, 12, 13, 13, 17,
	10, 14, 15, 11, 16, 6, 18, 19, 21, 5,
	4, 3, 1, 0, 0, 0, 0, 0, 20, 7,
}
var yyPact = [...]int{

	-8, -1000, -4, -1000, -1000, -1000, -1000, -8, -8, -1000,
	4, -1000, -8, -8, -10, -1000, -34, -1000, -3, -1000,
	-1000, -1000, -1000,
}
var yyPgo = [...]int{

	0, 22, 3, 21, 20, 19, 18, 15, 14, 10,
}
var yyR1 = [...]int{

	0, 1, 2, 2, 2, 2, 3, 4, 4, 4,
	5, 7, 9, 8, 6,
}
var yyR2 = [...]int{

	0, 1, 1, 1, 1, 1, 3, 2, 3, 3,
	1, 3, 1, 1, 1,
}
var yyChk = [...]int{

	-1000, -1, -2, -3, -4, -5, -7, 37, 12, 13,
	-9, 21, 10, 11, -2, -2, -8, 5, -2, -2,
	38, -6, 36,
}
var yyDef = [...]int{

	0, -2, 1, 2, 3, 4, 5, 0, 0, 10,
	0, 12, 0, 0, 0, 7, 0, 13, 8, 9,
	6, 11, 14,
}
var yyTok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	37, 38,
}
var yyTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36,
}
var yyTok3 = [...]int{
	0,
}

var yyErrorMessages = [...]struct {
	state int
	token int
	msg   string
}{}

//line yaccpar:1

/*	parser for yacc output	*/

var (
	yyDebug        = 0
	yyErrorVerbose = false
)

type yyLexer interface {
	Lex(lval *yySymType) int
	Error(s string)
}

type yyParser interface {
	Parse(yyLexer) int
	Lookahead() int
}

type yyParserImpl struct {
	lval  yySymType
	stack [yyInitialStackSize]yySymType
	char  int
}

func (p *yyParserImpl) Lookahead() int {
	return p.char
}

func yyNewParser() yyParser {
	return &yyParserImpl{}
}

const yyFlag = -1000

func yyTokname(c int) string {
	if c >= 1 && c-1 < len(yyToknames) {
		if yyToknames[c-1] != "" {
			return yyToknames[c-1]
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

func yyErrorMessage(state, lookAhead int) string {
	const TOKSTART = 4

	if !yyErrorVerbose {
		return "syntax error"
	}

	for _, e := range yyErrorMessages {
		if e.state == state && e.token == lookAhead {
			return "syntax error: " + e.msg
		}
	}

	res := "syntax error: unexpected " + yyTokname(lookAhead)

	// To match Bison, suggest at most four expected tokens.
	expected := make([]int, 0, 4)

	// Look for shiftable tokens.
	base := yyPact[state]
	for tok := TOKSTART; tok-1 < len(yyToknames); tok++ {
		if n := base + tok; n >= 0 && n < yyLast && yyChk[yyAct[n]] == tok {
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}
	}

	if yyDef[state] == -2 {
		i := 0
		for yyExca[i] != -1 || yyExca[i+1] != state {
			i += 2
		}

		// Look for tokens that we accept or reduce.
		for i += 2; yyExca[i] >= 0; i += 2 {
			tok := yyExca[i]
			if tok < TOKSTART || yyExca[i+1] == 0 {
				continue
			}
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}

		// If the default action is to accept or reduce, give up.
		if yyExca[i+1] != 0 {
			return res
		}
	}

	for i, tok := range expected {
		if i == 0 {
			res += ", expecting "
		} else {
			res += " or "
		}
		res += yyTokname(tok)
	}
	return res
}

func yylex1(lex yyLexer, lval *yySymType) (char, token int) {
	token = 0
	char = lex.Lex(lval)
	if char <= 0 {
		token = yyTok1[0]
		goto out
	}
	if char < len(yyTok1) {
		token = yyTok1[char]
		goto out
	}
	if char >= yyPrivate {
		if char < yyPrivate+len(yyTok2) {
			token = yyTok2[char-yyPrivate]
			goto out
		}
	}
	for i := 0; i < len(yyTok3); i += 2 {
		token = yyTok3[i+0]
		if token == char {
			token = yyTok3[i+1]
			goto out
		}
	}

out:
	if token == 0 {
		token = yyTok2[1] /* unknown char */
	}
	if yyDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", yyTokname(token), uint(char))
	}
	return char, token
}

func yyParse(yylex yyLexer) int {
	return yyNewParser().Parse(yylex)
}

func (yyrcvr *yyParserImpl) Parse(yylex yyLexer) int {
	var yyn int
	var yyVAL yySymType
	var yyDollar []yySymType
	_ = yyDollar // silence set and not used
	yyS := yyrcvr.stack[:]

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	yystate := 0
	yyrcvr.char = -1
	yytoken := -1 // yyrcvr.char translated into internal numbering
	defer func() {
		// Make sure we report no lookahead when not parsing.
		yystate = -1
		yyrcvr.char = -1
		yytoken = -1
	}()
	yyp := -1
	goto yystack

ret0:
	return 0

ret1:
	return 1

yystack:
	/* put a state and value onto the stack */
	if yyDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", yyTokname(yytoken), yyStatname(yystate))
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
	if yyrcvr.char < 0 {
		yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
	}
	yyn += yytoken
	if yyn < 0 || yyn >= yyLast {
		goto yydefault
	}
	yyn = yyAct[yyn]
	if yyChk[yyn] == yytoken { /* valid shift */
		yyrcvr.char = -1
		yytoken = -1
		yyVAL = yyrcvr.lval
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
		if yyrcvr.char < 0 {
			yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
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
			if yyn < 0 || yyn == yytoken {
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
			yylex.Error(yyErrorMessage(yystate, yytoken))
			Nerrs++
			if yyDebug >= 1 {
				__yyfmt__.Printf("%s", yyStatname(yystate))
				__yyfmt__.Printf(" saw %s\n", yyTokname(yytoken))
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
				__yyfmt__.Printf("error recovery discards %s\n", yyTokname(yytoken))
			}
			if yytoken == yyEofCode {
				goto ret1
			}
			yyrcvr.char = -1
			yytoken = -1
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
	// yyp is now the index of $0. Perform the default action. Iff the
	// reduced production is ε, $1 is possibly out of range.
	if yyp+1 >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
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
		yyDollar = yyS[yypt-1 : yypt+1]
		//line grammar.y:55
		{
			yyVAL.lxr.expr = yyDollar[1].expr
		}
	case 3:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line grammar.y:59
		{
			yyVAL.expr = yyDollar[1].oper
		}
	case 4:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line grammar.y:60
		{
			yyVAL.expr = yyDollar[1].decl
		}
	case 5:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line grammar.y:61
		{
			yyVAL.expr = yyDollar[1].comp
		}
	case 6:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line grammar.y:64
		{
			yyVAL.expr = yyDollar[2].expr
		}
	case 7:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line grammar.y:67
		{
			yyVAL.oper = &operation{opNot, yyDollar[2].expr, nil}
		}
	case 8:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line grammar.y:68
		{
			yyVAL.oper = &operation{opAnd, yyDollar[1].expr, yyDollar[3].expr}
		}
	case 9:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line grammar.y:69
		{
			yyVAL.oper = &operation{opOr, yyDollar[1].expr, yyDollar[3].expr}
		}
	case 10:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line grammar.y:72
		{
			yyVAL.decl = yyVAL.decl
		}
	case 11:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line grammar.y:75
		{
			c, err := mkComparison(yyDollar[1].fld, yyDollar[2].cmp, yyDollar[3].value)
			if err != nil {
				yyVAL.lxr.err = err
				return 1
			}
			yyVAL.comp = c
		}
	case 12:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line grammar.y:85
		{
			yyVAL.fld = yyVAL.fld
		}
	case 13:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line grammar.y:88
		{
			yyVAL.cmp = yyVAL.cmp
		}
	case 14:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line grammar.y:91
		{
			yyVAL.value = yyVAL.value
		}
	}
	goto yystack /* stack new state and value */
}
