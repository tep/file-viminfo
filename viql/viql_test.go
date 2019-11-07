package viql

import (
	"testing"

	"toolman.org/file/viminfo"
)

var test_viminfo = &viminfo.VimInfo{
	SwapFile: ".dummy.swp",
	Filename: "/path/to/some/dummy",
	User:     "dood",
	Hostname: "foo.example.com",
	PID:      12345,
	Inode:    2345678,
	Modified: true,
	Format:   viminfo.FFunix,
	Crypto:   viminfo.CMnone,
}

func TestCompile(t *testing.T) {
	// yyDebug = 4
	// yyErrorVerbose = true
	args := []string{`user = dood and host =~ /^foo\\..*/`}
	expr, err := Compile(args...)
	if err != nil {
		t.Fatal(err)
	}

	if got, err := expr.Evaluate(test_viminfo); err != nil || got != true {
		t.Errorf("expr.Evaluate(_) == (%v, %v); Wanted(%v, %v)", got, err, true, nil)
	}
}
