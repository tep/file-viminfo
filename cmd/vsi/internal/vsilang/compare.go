package vsilang

import (
	"fmt"

	"toolman.org/file/viminfo"
)

type comparitor int

const (
	cmpNone comparitor = iota
	cmpEqual
	cmpNotEqual
	cmpREMatch
	cmpNotREMatch
)

func tokenComparitor(tok int) comparitor {
	switch tok {
	case '=', EQ:
		return cmpEqual
	case NEQ:
		return cmpNotEqual
	case REM:
		return cmpREMatch
	case NRE:
		return cmpNotREMatch
	default:
		return cmpNone
	}
}

func (c comparitor) negates() bool {
	switch c {
	case cmpNotEqual, cmpNotREMatch:
		return true
	default:
		return false
	}
}

//--------------------------------------------------

type comparison struct {
	fld field
	cmp comparitor
	val *value
}

func (c *comparison) String() string {
	var f, p string
	switch c.fld {
	case fldCrypto:
		f = "cryptmethod"
	case fldFilename:
		f = "filename"
	case fldPID:
		f = "pid"
	case fldHost:
		f = "hostname"
	case fldUser:
		f = "user"
	case fldInode:
		f = "inode"
	case fldFormat:
		f = "fileformat"
	default:
		f = "UNSET"
	}

	switch c.cmp {
	case cmpEqual:
		p = "=="
	case cmpNotEqual:
		p = "!="
	case cmpREMatch:
		p = "=~"
	case cmpNotREMatch:
		p = "!~"
	}

	return fmt.Sprintf("&comparison{fld:%q, cmp:%q, val:%v}", f, p, c.val)
}

func mkComparison(f field, c comparitor, v *value) (*comparison, error) {
	if v.typ != vtRegex && (c == cmpREMatch || c == cmpNotREMatch) {
		return nil, ErrBadComp
	}

	if v.typ == vtRegex && !(c == cmpREMatch || c == cmpNotREMatch) {
		return nil, ErrBadComp
	}

	if err := f.comparableTo(v); err != nil {
		return nil, err
	}

	return &comparison{fld: f, cmp: c, val: v}, nil
}

func (c *comparison) Evaluate(vi *viminfo.VimInfo) (bool, error) {
	fmt.Printf("##### eval: %v\n", c)
	var (
		str    *string
		result bool
	)

	switch c.fld {
	case fldCrypto:
		result = vi.Crypto == c.val.cmval

	case fldFilename:
		str = &vi.Filename

	case fldFormat:
		result = vi.Format == c.val.ffval

	case fldHost:
		str = &vi.Hostname

	case fldInode:
		result = vi.Inode == uint32(c.val.intval)

	case fldPID:
		result = int(vi.PID) == c.val.intval

	case fldUser:
		str = &vi.User
	}

	if str != nil {
		if c.val.typ == vtRegex {
			result = c.val.reval.Match([]byte(*str))
		} else {
			result = *str == c.val.strval
		}
	}

	if c.cmp.negates() {
		result = !result
	}

	return result, nil
}
