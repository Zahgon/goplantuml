package subfolder2

// Subfolder2 structure for testing purpose only
type Subfolder2 struct {
}

// SubfolderFunction is for testing purposes
func (s *Subfolder2) SubfolderFunction(b bool, i int) bool { _ = "STUB: not implemented"; return false }

func (s *Subfolder2) SubfolderFunctionWithReturnListParametrized() (a, b, c []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil, nil, nil
}
