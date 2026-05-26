package parser

// Alias defines a type that is an alias for some other type
type Alias struct {
	Name        string
	PackageName string
	AliasOf     string
}

func getNewAlias(name, packageName, aliasOf string) *Alias { _ = "STUB: not implemented"; return nil }

// AliasSlice implement the sort.Interface interface to allow for proper sorting of an alias slice
type AliasSlice []Alias

// Len is the number of elements in the collection.
func (as AliasSlice) Len() int {
	_ = "STUB: not implemented"

	// Less reports whether the element with
	// index i should sort before the element with index j.
	return 0
}

func (as AliasSlice) Less(i, j int) bool { _ = "STUB: not implemented"; return false }

// Swap swaps the elements with indexes i and j.
func (as AliasSlice) Swap(i, j int) { _ = "STUB: not implemented"; return }
