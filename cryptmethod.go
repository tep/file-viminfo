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

// CryptMethod is a vim crypto method
type CryptMethod byte

// Supported Crypto Methods
const (
	CMnone      CryptMethod = '0' // No encryption
	CMzip       CryptMethod = 'c' // Default encryption
	CMblowfish  CryptMethod = 'C' // Blowfish encryption
	CMblowfish2 CryptMethod = 'd' // Blowfish2 encryption
)

func (cm CryptMethod) String() string {
	switch cm {
	case CMnone:
		return "plaintext"
	case CMzip:
		return "zip"
	case CMblowfish:
		return "blowfish"
	case CMblowfish2:
		return "blowfish2"
	default:
		return ""
	}
}

func (cm *CryptMethod) MarshalText() ([]byte, error) {
	return []byte(cm.String()), nil
}

var ErrUnknownCryptMethod = errors.New("unknown crypt method")

func (cm *CryptMethod) UnmarshalText(data []byte) error {
	switch strings.ToLower(string(data)) {
	case "plaintext":
		*cm = CMnone
	case "zip":
		*cm = CMzip
	case "blowfish":
		*cm = CMblowfish
	case "blowfish2":
		*cm = CMblowfish2
	default:
		return ErrUnknownCryptMethod
	}
	return nil
}
