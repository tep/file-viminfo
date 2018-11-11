package vsilang

import (
	"fmt"
	"strconv"
	"strings"

	"toolman.org/text/scanner"
)

var keywords = scanner.KeywordMap{
	// Operators
	"and": AND,
	"or":  OR,
	"not": NOT,

	// Declarations
	"changed":  MODIFIED,
	"currhost": THISHOST,
	"curruser": THISUSER,
	"missing":  MISSING,
	"modified": MODIFIED,
	"running":  RUNNING,
	"thishost": THISHOST,
	"thisuser": THISUSER,

	// Fields
	"cm":          CRYPTMETHOD,
	"cryptmethod": CRYPTMETHOD,
	"filename":    FILENAME,
	"fileformat":  FILEFORMAT,
	"format":      FILEFORMAT,
	"host":        HOSTNAME,
	"hostname":    HOSTNAME,
	"inode":       INODE,
	"pid":         PID,
	"user":        USER,

	// File Formats
	"dos":  FFDOS,
	"mac":  FFMAC,
	"unix": FFUNIX,

	// Crypt Methods
	"blowfish":  CMBLOWFISH,
	"blowfish2": CMBLOWFISH2,
	"plaintext": CMPLAINTEXT,
	"zip":       CMZIP,
}

var pairs = scanner.RunePairs{
	{'=', '=', EQ},
	{'!', '=', NEQ},
	{'=', '~', REM},
	{'!', '~', NRE},
}

const eof = 0

type lexer struct {
	scnr *scanner.Scanner
	expr Expression
	err  error
}

func newLexer(args []string) *lexer {
	var (
		lx   = new(lexer)
		src  = &cliSource{strings.NewReader(strings.Join(args, " "))}
		errf = scanner.ErrFunc(lx.scanErr)
	)

	lx.scnr = scanner.New(src, errf, keywords, pairs, scanner.ScanRegexen, -scanner.ScanComments, -scanner.ScanFloats, -scanner.ScanRawStrings)

	return lx
}

func (lx *lexer) scanErr(msg string) {
	lx.err = fmt.Errorf("%s: %s", lx.scnr.Position(), msg)
}

func (lx *lexer) Lex(lval *yySymType) int {
	lval.lxr = lx

	tok := lx.scnr.Scan()
	if lx.err != nil {
		return ERROR
	}

	if tok == scanner.EOF {
		return eof
	}

	lval.pos = lx.scnr.Position()

	fmt.Printf("##### pos:%v tok:%d\n", lval.pos, tok)

	switch tok {
	case scanner.KeyWord:
		kwtok := lx.scnr.Token()
		switch kwtok {
		case AND, OR, NOT:
			return kwtok

		case MISSING, MODIFIED, RUNNING, THISHOST, THISUSER:
			lval.decl = tokenDeclaration(int(kwtok))
			return DECL

		case CRYPTMETHOD, FILEFORMAT, FILENAME, HOSTNAME, INODE, PID, USER:
			lval.fld = tokenField(int(kwtok))
			return FIELD

		case FFDOS, FFMAC, FFUNIX:
			lval.value = ffValue(int(kwtok))
			return VALUE

		case CMBLOWFISH, CMBLOWFISH2, CMPLAINTEXT, CMZIP:
			lval.value = cmValue(int(kwtok))
			return VALUE
		default:
			lx.err = fmt.Errorf("%s: unsupported keyword (this is a bug): %v %q", lx.scnr.Position(), lx.scnr.TokenString(rune(kwtok)), lx.scnr.Text())
			return ERROR
		}

	case scanner.Ident:
		lval.value = strValue(lx.scnr.Text())
		return VALUE

	case scanner.Int:
		var i int
		if i, lx.err = strconv.Atoi(lx.scnr.Text()); lx.err != nil {
			return ERROR
		}
		lval.value = intValue(i)
		return VALUE

	case scanner.Regex:
		lval.value = reValue(lx.scnr.Regex())
		return VALUE

	case scanner.String:
		var txt string
		if txt, lx.err = strconv.Unquote(lx.scnr.Text()); lx.err != nil {
			return ERROR
		}

		lval.value = strValue(txt)
		return VALUE

	case '=', EQ, NEQ, REM, NRE:
		lval.cmp = tokenComparitor(int(tok))
		return CMP

	default:
		lx.err = fmt.Errorf("%s: unrecognized token: %v %q", lx.scnr.Position(), lx.scnr.TokenString(tok), lx.scnr.Text())
		return ERROR
	}
}

func (lx *lexer) Error(msg string) {
	if lx.err != nil {
		lx.err = fmt.Errorf("%s: %v", msg, lx.err)
	} else {
		lx.err = fmt.Errorf("parse error: %s: %s", msg, lx.scnr.Position())
	}
}

type cliSource struct {
	*strings.Reader
}

func (s *cliSource) Name() string {
	return "cmdline"
}
