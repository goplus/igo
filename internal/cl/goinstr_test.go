package cl

/*
import (
	"go/ast"
	"go/token"
	"testing"

	exec "github.com/goplus/igo/internal/exec/bytecode"
	"github.com/qiniu/x/ts"
)

// -----------------------------------------------------------------------------

var instrDeferDiscardsResult = map[string]goInstrInfo{
	"len":     {igoLen},
	"cap":     {igoCap},
	"make":    {igoMake},
	"new":     {igoNew},
	"complex": {igoComplex},
	"real":    {igoReal},
	"imag":    {igoImag},
}

func TestDeferFileNotFound(t *testing.T) {
	b := exec.NewBuilder(nil)
	pkg := &ast.Package{
		Files: map[string]*ast.File{},
	}
	pkgCtx := newPkgCtx(b.Interface(), pkg, token.NewFileSet())
	ctx := newGblBlockCtx(pkgCtx)
	ts.New(t).Call(func() {
		fn := &ast.Ident{Name: "len"}
		expr := &ast.CallExpr{Fun: fn}
		igoLen(ctx, expr, callByDefer)
	}).Panic(
		"pkgCtx.getCodeInfo failed: file not found - \n",
	)
}
*/
// -----------------------------------------------------------------------------
