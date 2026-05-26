package parser

import (
	"go/ast"
)

// Struct represent a struct in golang, it can be of Type "class" or "interface" and can be associated
// with other structs via Composition and Extends
type Struct struct {
	PackageName         string
	Functions           []*Function
	Fields              []*Field
	Type                string
	Composition         map[string]struct{}
	Extends             map[string]struct{}
	Aggregations        map[string]struct{}
	PrivateAggregations map[string]struct{}
	TypeParameters      []TypeParameter
}

// ImplementsInterface returns true if the struct st conforms ot the given interface
func (st *Struct) ImplementsInterface(inter *Struct) bool { _ = "STUB: not implemented"; return false }

// AddToComposition adds the composition relation to the structure. We want to make sure that *ExampleStruct
// gets added as ExampleStruct so that we can properly build the relation later to the
// class identifier
func (st *Struct) AddToComposition(fType string) { _ = "STUB: not implemented"; return }

// AddToExtends Adds an extends relationship to this struct. We want to make sure that *ExampleStruct
// gets added as ExampleStruct so that we can properly build the relation later to the
// class identifier
func (st *Struct) AddToExtends(fType string) { _ = "STUB: not implemented"; return }

// AddToAggregation adds an aggregation type to the list of aggregations
func (st *Struct) AddToAggregation(fType string) { _ = "STUB: not implemented"; return }

// addToPrivateAggregation adds an aggregation type to the list of aggregations for private members
func (st *Struct) addToPrivateAggregation(fType string) { _ = "STUB: not implemented"; return }

// AddField adds a field into this structure. It parses the ast.Field and extract all
// needed information
func (st *Struct) AddField(field *ast.Field, aliases map[string]string) {
	_ = "STUB: not implemented"
	return
}

// TypeParameter represents a generic type parameter and its constraint
type TypeParameter struct {
	Name        string
	Constraints string
}

// isGenericParamType reports whether the given fundamental type refers to a type parameter of this struct
func (st *Struct) isGenericParamType(fundamentalType string) bool {
	_ = "STUB: not implemented"
	return false
}

// AddMethod Parse the Field and if it is an ast.FuncType, then add the methods into the structure
func (st *Struct) AddMethod(method *ast.Field, aliases map[string]string) {
	_ = "STUB: not implemented"
	return
}
