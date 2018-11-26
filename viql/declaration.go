package viql

import (
	"errors"
	"fmt"
	"os"
	"os/user"

	"toolman.org/file/viminfo"
)

type declaration int

const (
	decNone declaration = iota
	decAll
	decMissing
	decModified
	decRunning
	decThisHost
	decThisUser
)

func tokenDeclaration(tok int) declaration {
	switch tok {
	case ALL:
		return decAll
	case MISSING:
		return decMissing
	case MODIFIED:
		return decModified
	case NONE:
		return decNone
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
	switch d {
	case decAll:
		return true, nil

	case decMissing:
		// TODO(tep): Implement logic to detect a missing edit file
		//            n.b. This may belong as a VimInfo method (similar to Running)
		return false, errors.New("not implemented")

	case decModified:
		return vi.Modified, nil

	case decNone:
		return false, nil

	case decRunning:
		err := vi.Running()
		switch err {
		case nil:
			return true, nil
		case viminfo.ErrNoSuchProcess:
			return false, nil
		default:
			return false, err
		}

	case decThisHost:
		h, err := os.Hostname()
		if err != nil {
			return false, err
		}
		return h == vi.Hostname, nil

	case decThisUser:
		u, err := user.Current()
		if err != nil {
			return false, err
		}
		return u.Username == vi.User, nil

	default:
		return false, fmt.Errorf("unknown declaration type %d: this is a bug", d)
	}
}
