package viminfo

import (
	"fmt"
	"os"
	"os/user"
	"strconv"
	"syscall"
	"time"
)

// to override for tests
var osStat = os.Stat

type ownerInfo struct {
	owner string
	mtime time.Time
}

func fileOwner(filename string) (*ownerInfo, error) {
	fi, err := osStat(filename)
	if err != nil {
		return nil, err
	}

	var owner string

	switch sys := fi.Sys().(type) {
	case *syscall.Stat_t:
		owner = strconv.Itoa(int(sys.Uid))
		if u, err := user.LookupId(owner); err == nil {
			owner = u.Username
		}

		return &ownerInfo{owner, time.Unix(sys.Mtim.Unix())}, nil

	default:
		return nil, fmt.Errorf("sys is type: %T", sys)
	}
}
