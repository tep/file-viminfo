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
	"strings"
)

// FileFormat indicates the EOL token for an edited file
type FileFormat byte

// Supported formats
const (
	FFnone FileFormat = iota // Format is not known
	FFunix                   // Unix Format "\n"
	FFdos                    // DOS Format  "\r\n"
	FFmac                    // MAC Format  "\r"
)

func (ff FileFormat) String() string {
	switch ff {
	case FFunix:
		return "unix"
	case FFdos:
		return "dos"
	case FFmac:
		return "mac"
	default:
		return "unknown"
	}
}

func (ff *FileFormat) MarshalText() ([]byte, error) {
	return []byte(ff.String()), nil
}

var ErrUnknownFileFormat = errors.New("unknown file format")

func (ff *FileFormat) UnmarshalText(data []byte) error {
	switch strings.ToLower(string(data)) {
	case "unix":
		*ff = FFunix
	case "dos":
		*ff = FFdos
	case "mac":
		*ff = FFmac
	default:
		return ErrUnknownFileFormat
	}
	return nil
}
