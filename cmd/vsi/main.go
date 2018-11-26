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
	"toolman.org/file/viminfo/viql"
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
			if viminfo.IsFatal(err) && !i.cont {
				return err
			}
			i.warn("%s: %v\n", sf, err)
		}
	}

	return nil
}

type marshaller func(*viminfo.VimInfo) ([]byte, error)

type vsInspector struct {
	files   []string
	expr    viql.Expression
	tmpl    *template.Template
	marshal marshaller
	term    byte
	cont    bool
	quiet   bool
}

func inspector() (*vsInspector, error) {
	var (
		name = filepath.Base(os.Args[0])
		i    = new(vsInspector)
		err  error
	)

	fs := pflag.NewFlagSet(name, pflag.ContinueOnError)

	var (
		query  = fs.StringP("query", "q", "all", "Query statement used for selecting output")
		format = fs.StringP("format", "f", "{{.SwapFile}}", "Output format for VimInfo as a text/template")
		zero   = fs.BoolP("zero", "0", false, "Output records terminated by a null byte instead of a newline")
		rjs    = fs.BoolP("json", "j", false, "Output JSON instead of text as defined by --format")
		qlh    = fs.Bool("help-viql", false, "Show help on how to write query statements")
	)

	fs.BoolVarP(&i.cont, "continue", "c", i.cont, "Continue on errors")
	fs.BoolVarP(&i.quiet, "silent", "s", i.quiet, "Emit no file processing errors or warnings")

	fs.Usage = func() {
		fmt.Fprintf(os.Stderr, "USAGE: %s [flags] swapfile [...]\n\n", name)
		fs.PrintDefaults()
	}

	if err = fs.Parse(os.Args); err != nil {
		return nil, err
	}

	if *qlh {
		fs.Usage()
		viql.Help(os.Stderr)
		os.Exit(1)
	}

	if args := fs.Args(); len(args) < 2 {
		return nil, errors.New("no vim swap files specified")
	} else {
		i.files = args[1:]
	}

	if i.expr, err = viql.Compile(*query); err != nil {
		return nil, fmt.Errorf("bad query statement: %v", err)
	}

	if *rjs {
		i.marshal = i.marshalJSON
	} else {
		if i.tmpl, err = template.New(name).Parse(*format); err != nil {
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
	if match, err := i.expr.Evaluate(vi); err != nil || !match {
		return err
	}

	b, err := i.marshal(vi)
	if err != nil {
		return err
	}

	b = append(b, i.term)

	_, err = os.Stdout.Write(b)
	return err
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
