package ngapType

// Need to import "free5gclib/aper" if it uses "aper"

type ReferenceID struct {
	Value int64 `aper:"valueExt,valueLB:1,valueUB:64"`
}
