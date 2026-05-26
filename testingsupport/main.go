package testingsupport

import (
	"strings"
)

func (t *test) test() { _ = "STUB: not implemented"; return }

type test struct {
	field  int
	field2 TestComplicatedAlias
}

type myInt int

var globalVariable int

// TestComplicatedAlias for testing purposes only
type TestComplicatedAlias func(strings.Builder) bool
