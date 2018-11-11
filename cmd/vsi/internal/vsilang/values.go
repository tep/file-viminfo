package vsilang

import (
	"fmt"
	"regexp"

	"toolman.org/file/viminfo"
)

type valtype int

const (
	vtNone valtype = iota
	vtCryptMethod
	vtFileFormat
	vtInt
	vtRegex
	vtString
)

type value struct {
	typ    valtype
	cmval  viminfo.CryptMethod
	ffval  viminfo.FileFormat
	intval int
	reval  *regexp.Regexp
	strval string
}

func (v *value) String() string {
	var t, s string
	switch v.typ {
	case vtCryptMethod:
		t = "CryptMethod"
		s = v.cmval.String()
	case vtFileFormat:
		t = "FileFormat"
		s = v.ffval.String()
	case vtInt:
		t = "Integer"
		s = fmt.Sprintf("%d", v.intval)
	case vtRegex:
		t = "Regexp"
		s = v.reval.String()
	case vtString:
		t = "String"
		s = v.strval
	}
	return fmt.Sprintf("&value{typ:%q val:%q", t, s)
}

func cmValue(tok int) *value {
	var cm viminfo.CryptMethod
	switch tok {
	case CMBLOWFISH:
		cm = viminfo.CMblowfish
	case CMBLOWFISH2:
		cm = viminfo.CMblowfish2
	case CMPLAINTEXT:
		cm = viminfo.CMnone
	case CMZIP:
		cm = viminfo.CMzip
	}
	return &value{typ: vtCryptMethod, cmval: cm}
}

func ffValue(tok int) *value {
	var ff viminfo.FileFormat
	switch tok {
	case FFDOS:
		ff = viminfo.FFdos
	case FFMAC:
		ff = viminfo.FFmac
	case FFUNIX:
		ff = viminfo.FFunix
	}
	return &value{typ: vtFileFormat, ffval: ff}
}

func intValue(i int) *value {
	return &value{typ: vtInt, intval: i}
}

func reValue(re *regexp.Regexp) *value {
	return &value{typ: vtRegex, reval: re}
}

func strValue(s string) *value {
	return &value{typ: vtString, strval: s}
}

// XXX Might not need this
func (v *value) value() interface{} {
	switch v.typ {
	case vtCryptMethod:
		return v.cmval
	case vtFileFormat:
		return v.ffval
	case vtInt:
		return v.intval
	case vtRegex:
		return v.reval
	case vtString:
		return v.strval
	default:
		return nil
	}
}

func (v *value) eqString(s string) (bool, error) {
	if v.typ != vtString {
		return false, ErrTypeMismatch
	}

	return v.strval == s, nil
}

func (v *value) eqInt(i int) (bool, error) {
	if v.typ != vtInt {
		return false, ErrTypeMismatch
	}

	return v.intval == i, nil
}
