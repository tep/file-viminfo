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
	// "os"
	"reflect"
	"testing"
	"time"

	"github.com/kr/pretty"
)

func TestVimInfo(t *testing.T) {
	// defer func(s func() (os.FileInfo, err)) { osStat = s }(osStat)
	file := "testdata/plain.swp"

	want := &VimInfo{
		SwapFile: file,
		Version:  "VIM 7.4",
		LastMod:  time.Unix(1540822648, 0),
		Inode:    272982037,
		PID:      4988,
		User:     "tep",
		Hostname: "droog.toolman.org",
		Filename: "~tep/working/go/src/experiments/read-vim-swapfile/main.go",
		Encoding: "utf-8",
		Crypto:   CMnone,
		Format:   FFunix,
		Modified: false,
		SameDir:  true,
	}

	if got, err := Parse(file); err != nil || !reflect.DeepEqual(got, want) {
		t.Errorf("Parse(%q) := (%#v, %v); Wanted (%#v, %v)", file, got, err, want, nil)
		for _, d := range pretty.Diff(got, want) {
			t.Log(d)
		}
	}
}

// func fakeFileInfo struct {
// }
