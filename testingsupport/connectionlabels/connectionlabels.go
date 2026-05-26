package connectionlabels

// AbstractInterface for testing purposes
type AbstractInterface interface {
	interfaceFunction() bool
}

// ImplementsAbstractInterface for testing purposes
type ImplementsAbstractInterface struct {
	AliasOfInt
	PublicUse AbstractInterface
}

func (iai *ImplementsAbstractInterface) interfaceFunction() bool {
	_ = "STUB: not implemented"

	// AliasOfInt for testing purposes
	return false
}

type AliasOfInt int
