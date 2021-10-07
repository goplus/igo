package cltest

import (
	"fmt"
	"go/token"
	"os"
	"testing"

	"github.com/goplus/igo/internal/ast/asttest"
	"github.com/goplus/igo/internal/cl"
	"github.com/goplus/igo/internal/parser"
	"github.com/qiniu/x/ts"

	exec "github.com/goplus/igo/internal/exec/bytecode"
	_ "github.com/goplus/igo/lib" // libraries
)

// -----------------------------------------------------------------------------

// Expect runs a script and check expected output and if panic or not.
func Expect(t *testing.T, script string, expected string, panicMsg ...interface{}) {
	fset := token.NewFileSet()
	fs := asttest.NewSingleFileFS("/foo", "bar.go", script)
	pkgs, err := parser.ParseFSDir(fset, fs, "/foo", nil, 0)
	if err != nil || len(pkgs) != 1 {
		t.Fatal("ParseFSDir failed:", err, len(pkgs))
	}

	e := ts.StartExpecting(t, ts.CapStdout)
	defer e.Close()

	e.Call(func() {
		bar := pkgs["main"]
		b := exec.NewBuilder(nil)
		pkg, err := cl.NewPackage(b.Interface(), bar, fset, cl.PkgActClMain)
		if err != nil {
			t.Fatal("Compile failed:", err)
		}
		cl.Debug(pkg)
		code := b.Resolve()
		code.Dump(os.Stderr)
		fmt.Fprintln(os.Stderr)
		exec.NewContext(code).Run()
	}).Panic(panicMsg...).Expect(expected)
}

// -----------------------------------------------------------------------------

// Call runs a script and gets the last expression value to check
func Call(t *testing.T, noEntrypoint bool, script string, idx ...int) *ts.TestCase {
	fset := token.NewFileSet()
	fs := asttest.NewSingleFileFS("/foo", "bar.go", script)
	pkgs, err := parser.ParseFSDir(fset, fs, "/foo", nil, 0)
	if err != nil || len(pkgs) != 1 {
		t.Fatal("ParseFSDir failed:", err, len(pkgs))
	}

	return ts.New(t).Call(func() interface{} {
		bar := pkgs["main"]
		if noEntrypoint {
			for _, f := range bar.Files {
				f.Package = token.NoPos // mark noEntrypoint
				break
			}
		}
		b := exec.NewBuilder(nil)
		pkg, err := cl.NewPackage(b.Interface(), bar, fset, cl.PkgActClMain)
		if err != nil {
			t.Fatal("Compile failed:", err)
		}
		cl.Debug(pkg)
		code := b.Resolve()
		ctx := exec.NewContext(code)
		ctx.Run()
		return ctx.Get(append(idx, -1)[0])
	})
}

// -----------------------------------------------------------------------------
