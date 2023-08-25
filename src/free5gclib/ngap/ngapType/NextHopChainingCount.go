package ngapType

// Need to import "free5gclib/aper" if it uses "aper"

type NextHopChainingCount struct {
	Value int64 `aper:"valueLB:0,valueUB:7"`
}
