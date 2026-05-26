/*
Package parser generates PlantUml http://plantuml.com/ Class diagrams for your golang projects
The main structure is the ClassParser which you can generate by calling the NewClassDiagram(dir)
function.

Pass the directory where the .go files are and the parser will analyze the code and build a structure
containing the information it needs to Render the class diagram.

call the Render() function and this will return a string with the class diagram.

See github.com/jfeliu007/goplantuml/cmd/goplantuml/main.go for a command that uses this functions and outputs the text to
the console.
*/
package parser

import (
	"go/ast"
	"os"
	"strings"

	"github.com/spf13/afero"
)

// LineStringBuilder extends the strings.Builder and adds functionality to build a string with tabs and
// adding new lines
type LineStringBuilder struct {
	strings.Builder
}

const tab = "    "
const builtinPackageName = "__builtin__"
const implements = `"implements"`
const extends = `"extends"`
const aggregates = `"uses"`
const aliasOf = `"alias of"`

// WriteLineWithDepth will write the given text with added tabs at the beginning into the string builder.
func (lsb *LineStringBuilder) WriteLineWithDepth(depth int, str string) {
	_ = "STUB: not implemented"
	return
}

// ClassDiagramOptions will provide a way for callers of the NewClassDiagramFs() function to pass all the necessary arguments.
type ClassDiagramOptions struct {
	FileSystem         afero.Fs
	Directories        []string
	IgnoredDirectories []string
	RenderingOptions   map[RenderingOption]interface{}
	Recursive          bool
	MaxDepth           int // Maximum nesting depth for packages (0 = unlimited)
}

// RenderingOptions will allow the class parser to optionally enebale or disable the things to render.
type RenderingOptions struct {
	Title                   string
	Notes                   string
	Aggregations            bool
	Fields                  bool
	Methods                 bool
	Compositions            bool
	Implementations         bool
	Aliases                 bool
	ConnectionLabels        bool
	AggregatePrivateMembers bool
	PrivateMembers          bool
}

const aliasComplexNameComment = "'This class was created so that we can correctly have an alias pointing to this name. Since it contains dots that can break namespaces"

const (
	// RenderAggregations is to be used in the SetRenderingOptions argument as the key to the map, when value is true, it will set the parser to render aggregations
	RenderAggregations RenderingOption = iota

	// RenderCompositions is to be used in the SetRenderingOptions argument as the key to the map, when value is true, it will set the parser to render compositions
	RenderCompositions

	// RenderImplementations is to be used in the SetRenderingOptions argument as the key to the map, when value is true, it will set the parser to render implementations
	RenderImplementations

	// RenderAliases is to be used in the SetRenderingOptions argument as the key to the map, when value is true, it will set the parser to render aliases
	RenderAliases

	// RenderFields is to be used in the SetRenderingOptions argument as the key to the map, when value is true, it will set the parser to render fields
	RenderFields

	// RenderMethods is to be used in the SetRenderingOptions argument as the key to the map, when value is true, it will set the parser to render methods
	RenderMethods

	// RenderConnectionLabels is to be used in the SetRenderingOptions argument as the key to the map, when value is true, it will set the parser to render the connection labels
	RenderConnectionLabels

	// RenderTitle is the options for the Title of the diagram. The value of this will be rendered as a title unless empty
	RenderTitle

	// RenderNotes contains a list of notes to be rendered in the class diagram
	RenderNotes

	// AggregatePrivateMembers is to be used in the SetRenderingOptions argument as the key to the map, when value is true, it will connect aggregations with private members
	AggregatePrivateMembers

	// RenderPrivateMembers is used if private members (fields, methods) should be rendered
	RenderPrivateMembers
)

// RenderingOption is an alias for an it so it is easier to use it as options in a map (see SetRenderingOptions(map[RenderingOption]bool) error)
type RenderingOption int

// ClassParser contains the structure of the parsed files. The structure is a map of package_names that contains
// a map of structure_names -> Structs
type ClassParser struct {
	renderingOptions   *RenderingOptions
	structure          map[string]map[string]*Struct
	currentPackageName string
	currentDirPath     string   // Current directory being parsed
	rootDirectories    []string // Root directories being processed
	allInterfaces      map[string]struct{}
	allStructs         map[string]struct{}
	allImports         map[string]string
	allAliases         map[string]*Alias
	allRenamedStructs  map[string]map[string]string
	maxDepth           int
	packageHierarchy   map[string]*PackageNode // Maps package full path to PackageNode
}

// PackageNode represents a package in the hierarchy
type PackageNode struct {
	Name       string                  // Short name (e.g., "subfolder")
	FullPath   string                  // Full path (e.g., "testingsupport.subfolder")
	Parent     *PackageNode            // Parent package
	Children   map[string]*PackageNode // Child packages
	Structures map[string]*Struct      // Structures in this package
	Depth      int                     // Depth in hierarchy
}

// NewClassDiagramWithOptions returns a new classParser with which can Render the class diagram of
// files in the given directory passed in the ClassDiargamOptions. This will also alow for different types of FileSystems
// Passed since it is part of the ClassDiagramOptions as well.
func NewClassDiagramWithOptions(options *ClassDiagramOptions) (*ClassParser, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// --- helpers ---

func newClassParser(options *ClassDiagramOptions) *ClassParser {
	_ = "STUB: not implemented"
	return nil
}

func buildIgnoreDirectoryMap(ignored []string) map[string]struct{} {
	_ = "STUB: not implemented"
	return nil
}

func processDirectories(classParser *ClassParser, options *ClassDiagramOptions, ignoreDirs map[string]struct{}) error {
	_ = "STUB: not implemented"
	return nil
}

func walkDirectory(fs afero.Fs, root string, ignoreDirs map[string]struct{}, classParser *ClassParser) error {
	_ = "STUB: not implemented"
	return nil
}

func shouldSkipDir(path string, info os.FileInfo, ignoreDirs map[string]struct{}) bool {
	_ = "STUB: not implemented"
	return false
}

func populateInterfaceImplementations(classParser *ClassParser) { _ = "STUB: not implemented"; return }

// NewClassDiagram returns a new classParser with which can Render the class diagram of
// files in the given directory
func NewClassDiagram(directoryPaths []string, ignoreDirectories []string, recursive bool) (*ClassParser, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NewClassDiagramWithMaxDepth returns a new classParser with which can Render the class diagram of
// files in the given directory with a maximum nesting depth
func NewClassDiagramWithMaxDepth(directoryPaths []string, ignoreDirectories []string, recursive bool, maxDepth int) (*ClassParser, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// getOrCreatePackageNode creates or retrieves a package node in the hierarchy
func (p *ClassParser) getOrCreatePackageNode(dirPath string) *PackageNode {
	_ = "STUB: not implemented"
	// Calculate the package path relative to the root directories
	return nil
}

// Create new package node
// Use the last component of the package path as the display name

// Check depth limit

// Establish parent-child relationships

// establishParentChildRelationships sets up parent-child relationships for a package node
func (p *ClassParser) establishParentChildRelationships(node *PackageNode) {
	_ = "STUB: not implemented"
	return
}

// Root package, no parent

// Find parent path by removing the last component

// Get or create parent node

// Create parent node if it doesn't exist

// getParentPath returns the parent path of a given package path
func (p *ClassParser) getParentPath(packagePath string) string {
	_ = "STUB: not implemented"
	return ""
}

// No parent

// getDirectoryForPackagePath returns the directory path for a given package path
func (p *ClassParser) getDirectoryForPackagePath(packagePath string) string {
	_ = "STUB: not implemented"
	// Convert package path back to directory path
	// For example: "cmd.goplantuml" -> "cmd/goplantuml"
	return ""
}

// Check if this directory exists relative to any of our root directories

// calculatePackagePath determines the package path from directory path
func (p *ClassParser) calculatePackagePath(dirPath string) string {
	_ = "STUB: not implemented"
	return ""
}

// Find the shortest root directory that contains this path

// Get relative path from root

// Convert path separators to dots for package naming

// Check if we're at the project root level (no nesting)
// If the relative path doesn't contain separators, we're at the top level

// Special case: if we're processing the project root (current directory)
// and the path contains testingsupport or cmd, we want to preserve the nesting
// This handles the case where these are subdirectories of the project

// Special case: if we're processing cmd/goplantuml, it should be treated as cmd.goplantuml
// not as a separate root package

// Special case: if we're processing cmd directory, it should be treated as cmd

// Only prepend root directory name if we're not at the root level
// and if the root directory is not "." (current directory)

// calculateDepth calculates the nesting depth of a package path
func (p *ClassParser) calculateDepth(packagePath string) int { _ = "STUB: not implemented"; return 0 }

// getRootDirectories returns the root directories being processed
func (p *ClassParser) getRootDirectories() []string { _ = "STUB: not implemented"; return nil }

// parse the given ast.Package into the ClassParser structure
func (p *ClassParser) parsePackage(node ast.Node) { _ = "STUB: not implemented"; return }

// Create package node for this directory

// Skip if depth limit exceeded

// Use the hierarchical package name for the structure map

// Initialize structure maps

func (p *ClassParser) parseImports(impt *ast.ImportSpec) { _ = "STUB: not implemented"; return }

func (p *ClassParser) parseDirectory(directoryPath string) error {
	_ = "STUB: not implemented"
	return nil
}

// parse the given declaration looking for classes, interfaces, or member functions
func (p *ClassParser) parseFileDeclarations(node ast.Decl) { _ = "STUB: not implemented"; return }

func (p *ClassParser) handleFuncDecl(decl *ast.FuncDecl) { _ = "STUB: not implemented"; return }

// Only get in when the function is defined for a structure. Global functions are not needed for class diagram

// Skip functions with empty or invalid receiver types (e.g., malformed syntax)

func handleGenDecStructType(p *ClassParser, typeName string, c *ast.StructType) {
	_ = "STUB: not implemented"
	return
}

func handleGenDecInterfaceType(p *ClassParser, typeName string, c *ast.InterfaceType) {
	_ = "STUB: not implemented"
	return
}

func (p *ClassParser) handleGenDecl(decl *ast.GenDecl) { _ = "STUB: not implemented"; return }

// This might be a type of General Declaration we do not know how to handle.

func (p *ClassParser) processSpec(spec ast.Spec) { _ = "STUB: not implemented"; return }

// Not needed for class diagrams (Imports, global variables, regular functions, etc)

// processTypeSpec handles the processing of a TypeSpec and returns type name, declaration type, and alias
func (p *ClassParser) processTypeSpec(typeSpec *ast.TypeSpec) (string, string, *Alias) {
	_ = "STUB: not implemented"
	return "", "", nil
}

// For aliases, we need to use the full name for the structure

// processStructType handles struct type processing including generic parameters
func (p *ClassParser) processStructType(typeSpec *ast.TypeSpec, typeName string, structType *ast.StructType) {
	_ = "STUB: not implemented"
	return
}

// processInterfaceType handles interface type processing including generic parameters
func (p *ClassParser) processInterfaceType(typeSpec *ast.TypeSpec, typeName string, interfaceType *ast.InterfaceType) {
	_ = "STUB: not implemented"
	return
}

// processAliasType handles alias type processing
func (p *ClassParser) processAliasType(typeSpec *ast.TypeSpec, typeName string, typeExpr ast.Expr) *Alias {
	_ = "STUB: not implemented"
	return nil
}

// For aliases, we need to create the full name with package

// parseGenericTypeParameters extracts and sets type parameters for a type spec
func (p *ClassParser) parseGenericTypeParameters(typeSpec *ast.TypeSpec, typeName string) {
	_ = "STUB: not implemented"
	return
}

// registerDeclaration registers the declaration in the appropriate collections
func (p *ClassParser) registerDeclaration(typeName, declarationType string, alias *Alias) {
	_ = "STUB: not implemented"
	return
}

// registerAlias handles alias registration including renamed structs
func (p *ClassParser) registerAlias(typeName string, alias *Alias) {
	_ = "STUB: not implemented"
	// Use the full name from the alias for registration
	return
}

// parseTypeParameters parses ast.FieldList of type parameters into []TypeParameter
func parseTypeParameters(fl *ast.FieldList, aliases map[string]string) []TypeParameter {
	_ = "STUB: not implemented"
	return nil
}

// Each field may declare multiple identifiers with the same constraint

// stringifyConstraint produces a string for a type constraint expression
func stringifyConstraint(expr ast.Expr, aliases map[string]string) string {
	_ = "STUB: not implemented"
	return ""
}

// any, comparable, or named constraints

// Could be interface{ ~int | ~string | Method() }
// For simplicity, reuse getInterfaceType to render methods, but try to flatten embedded types
// Prefer union-like listing when possible

// Union of terms: X | Y

// Tilde type terms: ~int

// Fallback to existing field type stringifier

// renderInterfaceConstraint renders interface type constraints similarly to getInterfaceType but returns body only
func renderInterfaceConstraint(v *ast.InterfaceType, aliases map[string]string) string {
	_ = "STUB: not implemented"
	return ""
}

// If method has a name, keep signature; otherwise it may be embedded type or union-like element

// Embedded constraint type

// If this element is an array or a pointer, this function will return the type that is closer to these
// two definitions. For example []***map[int] string will return map[int]string
func getBasicType(theType ast.Expr) ast.Expr { _ = "STUB: not implemented"; return *new(ast.Expr) }

// Render returns a string of the class diagram that this parser has generated.
func (p *ClassParser) Render() string { _ = "STUB: not implemented"; return "" }

// Create builders for relationships

// Deduplication sets for generic outputs

// Render hierarchical packages

// Render aliases

// Render all relationships collected during package rendering

// Always render generic parameter relationships at the end

// Wrapper to pass dedupe maps
func (p *ClassParser) renderHierarchicalPackagesWithGenerics(str *LineStringBuilder, composition *LineStringBuilder, extends *LineStringBuilder, aggregations *LineStringBuilder, params *LineStringBuilder, emittedTypeParamClass map[string]struct{}, emittedParamLink map[string]struct{}) {
	_ = "STUB: not implemented"
	// Find root packages (packages with no parent)
	return
}

// Sort root packages by name

// Render each root package and its children

// renderHierarchicalPackages renders packages in a hierarchical structure
func (p *ClassParser) renderHierarchicalPackages(str *LineStringBuilder, composition *LineStringBuilder, extends *LineStringBuilder, aggregations *LineStringBuilder, params *LineStringBuilder) {
	_ = "STUB: not implemented"
	// Find root packages (packages with no parent)
	return
}

// Sort root packages by name

// Render each root package and its children

// renderPackageNode renders a package node and its children recursively
func (p *ClassParser) renderPackageNode(node *PackageNode, str *LineStringBuilder, composition *LineStringBuilder, extends *LineStringBuilder, aggregations *LineStringBuilder, params *LineStringBuilder, depth int) {
	_ = "STUB: not implemented"
	return
}

// Render this package's namespace using the short name

// Render structures in this package using the full path

// Render child packages

// Close namespace

// Generic-aware variant
func (p *ClassParser) renderPackageNodeWithGenerics(node *PackageNode, str *LineStringBuilder, composition *LineStringBuilder, extends *LineStringBuilder, aggregations *LineStringBuilder, params *LineStringBuilder, emittedTypeParamClass map[string]struct{}, emittedParamLink map[string]struct{}, depth int) {
	_ = "STUB: not implemented"
	return
}

// Render this package's namespace using the short name

// Render structures in this package using the full path

// Render child packages

// Close namespace

// renderStructuresInPackage renders structures within a package namespace
func (p *ClassParser) renderStructuresInPackage(pack string, structures map[string]*Struct, str *LineStringBuilder, depth int, composition *LineStringBuilder, extends *LineStringBuilder, aggregations *LineStringBuilder, params *LineStringBuilder) {
	_ = "STUB: not implemented"
	return
}

// Render renamed structs if any

// Generic-aware variant
func (p *ClassParser) renderStructuresInPackageWithGenerics(pack string, structures map[string]*Struct, str *LineStringBuilder, depth int, composition *LineStringBuilder, extends *LineStringBuilder, aggregations *LineStringBuilder, params *LineStringBuilder, emittedTypeParamClass map[string]struct{}, emittedParamLink map[string]struct{}) {
	_ = "STUB: not implemented"
	return
}

// Render renamed structs if any

func (p *ClassParser) renderStructures(pack string, structures map[string]*Struct, str *LineStringBuilder) {
	_ = "STUB: not implemented"
	return
}

// Always render generic parameter relationships at the end

func (p *ClassParser) renderAliases(str *LineStringBuilder) { _ = "STUB: not implemented"; return }

func (p *ClassParser) renderStructure(structure *Struct, pack string, name string, str *LineStringBuilder, composition *LineStringBuilder, extends *LineStringBuilder, aggregations *LineStringBuilder, params *LineStringBuilder) {
	_ = "STUB: not implemented"
	return
}

// Build display name including type parameters if present

// Generic stereotype

// Render type parameter classes and enqueue relationships

// Type parameter class

// Relationship: "T" <-- "param" "Model[T]"

// Bridge: generic-aware structure renderer using the existing renderStructure plus dedupe for params/classes
func (p *ClassParser) renderStructureWithGenerics(structure *Struct, pack string, name string, str *LineStringBuilder, composition *LineStringBuilder, extends *LineStringBuilder, aggregations *LineStringBuilder, params *LineStringBuilder, emittedTypeParamClass map[string]struct{}, emittedParamLink map[string]struct{}) {
	_ = "STUB: not implemented"
	// Render the main structure and its standard relations
	return
}

// Build display name and alias for generic classes

// Create unique alias for generic class

// Use type parameters in stereotype instead of brackets in name

// Render class with alias if generic

// Render type parameter classes and enqueue relationships (deduped)

// Use alias name for connections if available

func (p *ClassParser) renderCompositions(structure *Struct, name string, composition *LineStringBuilder) {
	_ = "STUB: not implemented"
	return
}

func (p *ClassParser) renderAggregations(structure *Struct, name string, aggregations *LineStringBuilder) {
	_ = "STUB: not implemented"
	return
}

func (p *ClassParser) updatePrivateAggregations(structure *Struct, aggregationsMap map[string]struct{}) {
	_ = "STUB: not implemented"
	return
}

func (p *ClassParser) renderAggregationMap(aggregationMap map[string]struct{}, structure *Struct, aggregations *LineStringBuilder, name string) {
	_ = "STUB: not implemented"
	return
}

func (p *ClassParser) getPackageName(t string, st *Struct) string {
	_ = "STUB: not implemented"
	return ""
}

func (p *ClassParser) renderExtends(structure *Struct, name string, extends *LineStringBuilder) {
	_ = "STUB: not implemented"
	return
}

func (p *ClassParser) renderStructMethods(structure *Struct, privateMethods *LineStringBuilder, publicMethods *LineStringBuilder) {
	_ = "STUB: not implemented"
	return
}

func (p *ClassParser) renderStructFields(structure *Struct, privateFields *LineStringBuilder, publicFields *LineStringBuilder) {
	_ = "STUB: not implemented"
	return
}

// Returns an initialized struct of the given name or returns the existing one if it was already created
func (p *ClassParser) getOrCreateStruct(name string) *Struct {
	_ = "STUB: not implemented"
	// Skip empty or invalid struct names to prevent PlantUML syntax errors
	return nil
}

// Returns an existing struct only if it was created. nil otherwhise
func (p *ClassParser) getStruct(structName string) *Struct { _ = "STUB: not implemented"; return nil }

// SetRenderingOptions Sets the rendering options for the Render() Function
func (p *ClassParser) SetRenderingOptions(ro map[RenderingOption]interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func generateRenamedStructName(currentName string) string { _ = "STUB: not implemented"; return "" }
