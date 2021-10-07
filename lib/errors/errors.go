package errors

import (
	"reflect"

	"github.com/goplus/igo"
	"github.com/qiniu/x/errors"
)

// NewFrame creates a new error frame.
func execNewFrame(arity int, p *igo.Context) {
	args := p.GetArgs(arity)
	err := errors.NewFrame(
		igo.ToError(args[0]),
		args[1].(string), args[2].(string), args[3].(int),
		args[4].(string), args[5].(string), args[6:]...,
	)
	p.Ret(arity, err)
}

func execIs(_ int, p *igo.Context) {
	args := p.GetArgs(2)
	is := errors.Is(igo.ToError(args[0]), igo.ToError(args[1]))
	p.Ret(2, is)
}

// -----------------------------------------------------------------------------

// I is a Go package instance.
var I = igo.NewGoPackage("github.com/qiniu/x/errors")

func init() {
	I.RegisterFuncvs(
		I.Funcv("NewFrame", errors.NewFrame, execNewFrame),
	)
	I.RegisterFuncs(
		I.Func("Is", errors.Is, execIs),
	)
	I.RegisterTypes(
		I.Type("Frame", reflect.TypeOf(errors.Frame{})),
	)
}

// -----------------------------------------------------------------------------
