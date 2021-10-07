package interp

import (
	"go/ast"
	"go/token"
	"io"

	"github.com/goplus/igo/internal/cl"
	exec "github.com/goplus/igo/internal/exec/bytecode"
)

func init() {
	cl.CallBuiltinOp = exec.CallBuiltinOp
}

// -----------------------------------------------------------------------------

type Program exec.Code

// CompileAST compiles a package by specified AST.
func CompileAST(pkg *ast.Package, fset *token.FileSet) (app *Program, err error) {
	b := exec.NewBuilder(nil)
	_, err = cl.NewPackage(b.Interface(), pkg, fset, cl.PkgActClMain)
	if err != nil {
		return
	}
	code := b.Resolve()
	return (*Program)(code), nil
}

// Dump dumps code.
func (p *Program) Dump(w io.Writer) {
	((*exec.Code)(p)).Dump(w)
}

// -----------------------------------------------------------------------------

type Runtime struct {
}

func New() *Runtime {
	return &Runtime{}
}

func (p *Runtime) RunProgram(app *Program) {
	ctx := exec.NewContext((*exec.Code)(app))
	ctx.Run()
}

// -----------------------------------------------------------------------------
