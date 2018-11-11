package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"text/template"

	"github.com/spf13/pflag"

	"toolman.org/file/viminfo"
	"toolman.org/flags/tristate"
)

func main() {
	if err := run(); err != nil && err != pflag.ErrHelp {
		fmt.Fprintf(os.Stderr, "ERROR: %v\n", err)
		os.Exit(1)
	}
}

func run() error {
	var (
		err error
		i   *vsInspector
	)

	if i, err = inspector(); err != nil {
		return err
	}

	for _, sf := range i.files {
		if err = i.inspect(sf); err != nil {
			if !i.cont {
				return err
			}
			i.warn("%s: %v\n", sf, err)
		}
	}

	return nil
}

type marshaller func(*viminfo.VimInfo) ([]byte, error)

type vsInspector struct {
	user     string
	host     string
	modified tristate.TriState
	running  tristate.TriState
	files    []string
	tmpl     *template.Template
	marshal  marshaller
	thisHost string
	term     byte
	cont     bool
	quiet    bool
}

func inspector() (*vsInspector, error) {
	var (
		base = filepath.Base(os.Args[0])
		i    = new(vsInspector)
		err  error
	)

	fs := pflag.NewFlagSet(base, pflag.ContinueOnError)

	var (
		format = fs.StringP("format", "f", "{{.SwapFile}}", "Output format for VimInfo as a text/template")
		zero   = fs.BoolP("zero", "0", false, "Output records terminated by a null byte instead of a newline")
		rjs    = fs.BoolP("json", "j", false, "Output JSON instead of text as defined by --format")
	)

	fs.StringVarP(&i.user, "user", "u", "", "Only emit output for files edited by `username`")
	fs.StringVarP(&i.host, "host", "h", "", "Only emit output for files edited on `hostname`")

	fs.BoolVarP(&i.cont, "continue", "c", i.cont, "Continue on errors reading swap files")
	fs.BoolVarP(&i.quiet, "quiet", "q", i.quiet, "Emit no file processing errors or warnings")

	tristate.FlagVarPFS(fs, &i.modified, "modified", "m", i.modified, "Only emit output for swap files with a modified flag matching this `TriState` value")
	tristate.FlagVarPFS(fs, &i.running, "running", "r", i.running, "Only emit output if the associated edit session's run state matches this `TriState` value")

	fs.Usage = func() {
		fmt.Fprintf(os.Stderr, "USAGE: %s {flags} swapfile [...]\n\n", base)
		fs.PrintDefaults()
	}

	if err = fs.Parse(os.Args); err != nil {
		return nil, err
	}

	if args := fs.Args(); len(args) < 2 {
		return nil, errors.New("no vim swap files specified")
	} else {
		i.files = args[1:]
	}

	if i.running.IsSet() {
		if i.thisHost, err = os.Hostname(); err != nil {
			return nil, fmt.Errorf("connot determine current hostname: %v", err)
		}
	}

	if *rjs {
		i.marshal = i.marshalJSON
	} else {
		if i.tmpl, err = template.New(base).Parse(*format); err != nil {
			return nil, fmt.Errorf("bad output format: %v", err)
		}
		i.marshal = i.marshalText
	}

	i.term = '\n'

	if *zero {
		i.term = 0
	}

	return i, nil
}

func (i *vsInspector) warn(mesg string, args ...interface{}) {
	if i.quiet {
		return
	}

	fmt.Fprintf(os.Stderr, "WARNING: "+mesg, args...)
}

func (i *vsInspector) inspect(swapfile string) error {
	vi, err := viminfo.Parse(swapfile)
	if err != nil {
		return fmt.Errorf("file %q: %v", swapfile, err)
	}

	return i.emit(vi)
}

func (i *vsInspector) emit(vi *viminfo.VimInfo) error {
	if i.filter(vi) {
		return nil
	}

	b, err := i.marshal(vi)
	if err != nil {
		return err
	}

	b = append(b, i.term)

	_, err = os.Stdout.Write(b)
	return err
}

func (i *vsInspector) filter(vi *viminfo.VimInfo) bool {
	if i.user != "" && vi.User != i.user {
		return true
	}

	if i.host != "" && vi.Hostname != i.host {
		return true
	}

	if !i.modified.Match(vi.Modified, true) {
		return true
	}

	if rp := i.running.Bool(); rp != nil {
		err := vi.Running()
		return (err == nil) == *rp
	}

	return false
}

func (i *vsInspector) marshalJSON(vi *viminfo.VimInfo) ([]byte, error) {
	b, err := json.Marshal(vi)
	if err != nil {
		return nil, fmt.Errorf("marshalling json: %v", err)
	}

	return b, nil
}

func (i *vsInspector) marshalText(vi *viminfo.VimInfo) ([]byte, error) {
	b := new(bytes.Buffer)
	if err := i.tmpl.Execute(b, vi); err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}
