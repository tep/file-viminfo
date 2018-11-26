package viql

import (
	"errors"

	"toolman.org/file/viminfo"
)

var (
	ErrBadComp      = errors.New("comparitor does not match value type")
	ErrTypeMismatch = errors.New("type mismatch: field is not comparable to value")
	ErrUnknownField = errors.New("unknown field: this is a bug")
	ErrUnknownOp    = errors.New("unknown operation (this is a bug)")
)

type Expression interface {
	Evaluate(vi *viminfo.VimInfo) (bool, error)
}

func Compile(args ...string) (Expression, error) {
	lx := newLexer(args)

	yyParse(lx)

	if lx.err != nil {
		return nil, lx.err
	}

	return lx.expr, nil
}
