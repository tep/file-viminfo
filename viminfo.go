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

// Package viminfo reads and parses Vim swap files into a well-formed structure.
package viminfo // import "toolman.org/file/viminfo"

import (
	"errors"
	"path/filepath"
	"time"
)

const (
	b0MagicLong = uint64(0x30313233)
)

// VimInfo reflects the meta-data stored in a vim swapfile
type VimInfo struct {
	// SwapFile is the name of the file containing this information
	SwapFile string `json:"swap_file,omitempty"`
	// Version indicates which version of Vim wrote this swap file
	Version string `json:"version,omitempty"`
	// LastMod is the modification time for the file being edited
	LastMod time.Time `json:"last_mod,omitempty"`
	// Inode is the filesystem inode of the file being edited
	Inode uint32 `json:"inode,omitempty"`
	// PID is the process ID for the vim session editing the file
	PID uint32 `json:"pid,omitempty"`
	// User is the username for the vim session's process owner (or, UID of username is unavailable)
	User string `json:"user,omitempty"`
	// Hostname is the hostname where the vim session is/was running
	Hostname string `json:"hostname,omitempty"`
	// Filename reflects the name of the file being edited
	Filename string `json:"filename,omitempty"`
	// Encoding is the file encoding for the file being edited (or, the word "encrypted" if the file is encrypted)
	Encoding string `json:"encoding,omitempty"`
	// Crypto indicates the "cryptmethod" for the file being edited (or, "plaintext" if the file is not encrypted)
	Crypto CryptMethod `json:"crypt_method,omitempty"`
	// Format is the FileFormat for the edited file (e.g. unix, dos, mac)
	Format FileFormat `json:"format,omitempty"`
	// Modified indicates whether the edit session has unsaved changes
	Modified bool `json:"modified"`
	// SameDir indicates whether the edited file is in the same directory as the swap file
	SameDir bool `json:"same_dir"`
}

// Parse reads and parses the vim swapfile specified by filename and returns
// a populated *VimInfo, or an error if the file could not be parsed.
func Parse(filename string) (*VimInfo, error) {
	b0, err := readBlock0(filename)
	if err != nil {
		return nil, err
	}

	cm := CryptMethod(b0[1])

	if b0[0] != 'b' || cm.String() == "" {
		return nil, errors.New("not a vim swap file")
	}

	if ml := b0.uint64At(1008); ml != b0MagicLong {
		return nil, errors.New("cannot read swap file from big-endian system")
	}

	flags := b0[1006]

	var (
		encoding string
		fnSize   int
	)

	switch cm {
	case CMnone:
		fnSize = 898
		encoding = b0.backString(108, fnSize)
	default:
		fnSize = 890
		encoding = "encrypted"
	}

	vi := &VimInfo{
		SwapFile: filepath.Clean(filename),
		Version:  b0.frontString(2, 8),
		LastMod:  time.Unix(int64(b0.uint32At(16)), 0),
		Inode:    b0.uint32At(20),
		PID:      b0.uint32At(24),
		User:     b0.frontString(28, 40),
		Hostname: b0.frontString(68, 40),
		Filename: b0.frontString(108, fnSize),
		Encoding: encoding,
		Crypto:   cm,
		Format:   FileFormat(flags & 0x03),
		Modified: b0[1007] == 0x55,
		SameDir:  flags&0x04 != 0,
	}

	return vi, nil
}
