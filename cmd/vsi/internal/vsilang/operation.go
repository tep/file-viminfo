package vsilang

import (
	"toolman.org/file/viminfo"
)

type optype int

const (
	opNone optype = iota
	opNot
	opAnd
	opOr
)

type operation struct {
	typ  optype
	exp1 Expression
	exp2 Expression
}

func (o *operation) Evaluate(vi *viminfo.VimInfo) (bool, error) {
	r1, err := o.exp1.Evaluate(vi)
	if err != nil {
		return false, err
	}

	switch o.typ {
	case opNot:
		return !r1, nil

	case opAnd:
		if r1 {
			return o.exp2.Evaluate(vi)
		}

	case opOr:
		if !r1 {
			return o.exp2.Evaluate(vi)
		}

	default:
		return false, ErrUnknownOp
	}

	return r1, nil
}
