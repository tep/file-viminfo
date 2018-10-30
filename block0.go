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
	"bytes"
	"encoding/binary"
	"io"
	"os"
)

type block0 []byte

// As defined by MIN_SWAP_PAGE_SIZE in vim.h (as of v8.1.0500)
const block0Size = 1048

func readBlock0(filename string) (block0, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var b0 [block0Size]byte
	// The biggest offset we use is 1008, and read an uint64 from it,
	// so we need at least 1008+8 bytes.
	n, err := io.ReadAtLeast(f, b0[:], 1008+8)
	return b0[:n], err
}

func (b0 block0) frontString(offset, length int) string {
	b := b0[offset : offset+length]
	if i := bytes.IndexByte(b, 0); i >= 0 {
		b = b[:i]
	}
	return string(b)
}

func (b0 block0) backString(offset, length int) string {
	b := b0[offset : offset+length]
	if i := bytes.LastIndexByte(b, 0); i >= 0 {
		b = b[i+1:]
	}
	return string(b)
}

func (b0 block0) uint64At(offset int) uint64 {
	return binary.LittleEndian.Uint64(b0[offset : offset+8])
}

func (b0 block0) uint32At(offset int) uint32 {
	return binary.LittleEndian.Uint32(b0[offset : offset+4])
}

func (b0 block0) byteAt(offset int) byte {
	return b0[offset]
}
