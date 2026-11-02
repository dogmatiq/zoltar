package zoltar_test

import (
	"go/ast"
	"go/importer"
	"go/parser"
	"go/token"
	"go/types"

	"golang.org/x/tools/go/ssa"
	"golang.org/x/tools/go/ssa/ssautil"
)

// buildPackage parses src as a single-file package, and builds it in SSA
// form, which is the form of a program that the predictor reads.
//
// Tools usually load packages with [golang.org/x/tools/go/packages] instead,
// and build them with [ssautil.AllPackages].
func buildPackage(src string) *ssa.Package {
	fset := token.NewFileSet()

	file, err := parser.ParseFile(fset, "example.go", src, 0)
	if err != nil {
		panic(err)
	}

	pkg, _, err := ssautil.BuildPackage(
		&types.Config{Importer: importer.Default()},
		fset,
		types.NewPackage("example.com/"+file.Name.Name, ""),
		[]*ast.File{file},
		ssa.InstantiateGenerics,
	)
	if err != nil {
		panic(err)
	}

	return pkg
}

// callTo returns the first call within fn to a function with the given name.
func callTo(fn *ssa.Function, name string) *ssa.Call {
	for _, b := range fn.Blocks {
		for _, instr := range b.Instrs {
			if call, ok := instr.(*ssa.Call); ok {
				if callee := call.Call.StaticCallee(); callee != nil && callee.Name() == name {
					return call
				}
			}
		}
	}

	panic("no call to " + name)
}
