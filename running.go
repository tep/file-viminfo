//
// Copyright 2018 Timothy E. Peoples
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to
// deal in the Software without restriction, including without limitation the
// rights to use, copy, modify, merge, publish, distribute, sublicense, and/or
// sell copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING
// FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS
// IN THE SOFTWARE.
//

package viminfo

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"golang.org/x/sys/unix"
)

var (
	ErrNoPID         = errors.New("missing vim process id")
	ErrNoSuchProcess = errors.New("no such process")
	ErrNotThisHost   = errors.New("vim session not on this host")
	ErrNotVim        = errors.New("process is not vim")

	VimCommands = []string{"ex", "vi", "vim", "gvim", "view", "gview", "vimdiff"}
)

func (vi *VimInfo) Running() error {
	if vi.PID == 0 {
		return ErrNoPID
	}

	if vi.Hostname != "" {
		host, err := os.Hostname()
		if err != nil {
			return err
		}

		if vi.Hostname != host {
			return ErrNotThisHost
		}
	}

	if err := unix.Kill(int(vi.PID), 0); err != nil {
		return ErrNoSuchProcess
	}

	return isVim(vi.PID)
}

func isVim(pid uint32) error {
	cmd, err := pidCommand(pid)
	if err != nil {
		return err
	}

	if ok, err := filepath.Match("vim.*", cmd); err == nil && ok {
		return nil
	}

	for _, vc := range VimCommands {
		if cmd == vc {
			return nil
		}
	}

	return ErrNotVim
}

// TODO(tep): This only works on Linux; make it work elsewhere too.
func pidCommand(pid uint32) (string, error) {
	exe, err := os.Readlink(fmt.Sprintf("/proc/%d/exe", pid))
	if err != nil {
		return "", err
	}

	return filepath.Base(exe), nil
}
