package vsilang

import (
	"errors"

	"toolman.org/file/viminfo"
)

type declaration int

const (
	decNone declaration = iota
	decMissing
	decModified
	decRunning
	decThisHost
	decThisUser
)

func tokenDeclaration(tok int) declaration {
	switch tok {
	case MISSING:
		return decMissing
	case MODIFIED:
		return decModified
	case RUNNING:
		return decRunning
	case THISHOST:
		return decThisHost
	case THISUSER:
		return decThisUser
	default:
		return decNone
	}
}

func (d declaration) Evaluate(vi *viminfo.VimInfo) (bool, error) {
	return false, errors.New("unimplemented")
}
