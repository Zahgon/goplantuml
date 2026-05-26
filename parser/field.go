package parser

import (
	"go/ast"
)

const packageConstant = "{packageName}"

// Field can hold the name and type of any field
type Field struct {
	Name     string
	Type     string
	FullType string
}

// Returns a string representation of the given expression if it was recognized.
// Refer to the implementation to see the different string representations.
func getFieldType(exp ast.Expr, aliases map[string]string) (string, []string) {
	_ = "STUB: not implemented"
	return "", nil
}

// Generic instantiation like Foo[T] or pkg.Foo[T]; we only care about the base type name

// Multi-parameter instantiation like Foo[T, U]

func getIdent(v *ast.Ident, aliases map[string]string) (string, []string) {
	_ = "STUB: not implemented"
	return "", nil
}

func getArrayType(v *ast.ArrayType, aliases map[string]string) (string, []string) {
	_ = "STUB: not implemented"
	return "", nil
}

func getSelectorExp(v *ast.SelectorExpr, aliases map[string]string) (string, []string) {
	_ = "STUB: not implemented"
	return "", nil
}

func getMapType(v *ast.MapType, aliases map[string]string) (string, []string) {
	_ = "STUB: not implemented"
	return "", nil
}

func getStarExp(v *ast.StarExpr, aliases map[string]string) (string, []string) {
	_ = "STUB: not implemented"
	return "", nil
}

func getChanType(v *ast.ChanType, aliases map[string]string) (string, []string) {
	_ = "STUB: not implemented"
	return "", nil
}

func getStructType(v *ast.StructType, aliases map[string]string) (string, []string) {
	_ = "STUB: not implemented"
	return "", nil
}

func getInterfaceType(v *ast.InterfaceType, aliases map[string]string) (string, []string) {
	_ = "STUB: not implemented"
	return "", nil
}

func getFuncType(v *ast.FuncType, aliases map[string]string) (string, []string) {
	_ = "STUB: not implemented"
	return "", nil
}

func getEllipsis(v *ast.Ellipsis, aliases map[string]string) (string, []string) {
	_ = "STUB: not implemented"
	return "", nil
}

var globalPrimitives = map[string]struct{}{
	"bool":        {},
	"string":      {},
	"int":         {},
	"int8":        {},
	"int16":       {},
	"int32":       {},
	"int64":       {},
	"uint":        {},
	"uint8":       {},
	"uint16":      {},
	"uint32":      {},
	"uint64":      {},
	"uintptr":     {},
	"byte":        {},
	"rune":        {},
	"float32":     {},
	"float64":     {},
	"complex64":   {},
	"complex128":  {},
	"error":       {},
	"*bool":       {},
	"*string":     {},
	"*int":        {},
	"*int8":       {},
	"*int16":      {},
	"*int32":      {},
	"*int64":      {},
	"*uint":       {},
	"*uint8":      {},
	"*uint16":     {},
	"*uint32":     {},
	"*uint64":     {},
	"*uintptr":    {},
	"*byte":       {},
	"*rune":       {},
	"*float32":    {},
	"*float64":    {},
	"*complex64":  {},
	"*complex128": {},
	"*error":      {},
}

func isPrimitive(ty *ast.Ident) bool { _ = "STUB: not implemented"; return false }

func isPrimitiveString(t string) bool { _ = "STUB: not implemented"; return false }

func replacePackageConstant(field, packageName string) string {
	_ = "STUB: not implemented"
	// Don't replace package constants for external packages
	return ""
}

// isExternalPackage checks if a field type represents an external package
func isExternalPackage(field string) bool {
	_ = "STUB: not implemented"
	// External packages contain dots and don't start with packageConstant
	return false
}
