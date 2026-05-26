package parser

import (
	"go/ast"
)

// Function holds the signature of a function with name, Parameters and Return values
type Function struct {
	Name                 string
	Parameters           []*Field
	ReturnValues         []string
	PackageName          string
	FullNameReturnValues []string
}

// SignturesAreEqual Returns true if the two functions have the same signature (parameter names are not checked)
func (f *Function) SignturesAreEqual(function *Function) bool {
	_ = "STUB: not implemented"
	return false
}

// generate and return a function object from the given Functype. The names must be passed to this
// function since the FuncType does not have this information
func getFunction(f *ast.FuncType, name string, aliases map[string]string, packageName string) *Function {
	_ = "STUB: not implemented"
	return nil
}
