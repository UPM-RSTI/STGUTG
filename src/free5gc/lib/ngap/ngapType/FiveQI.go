package ngapType

// Need to import "free5gclib/aper" if it uses "aper"

type FiveQI struct {
	Value int64 `aper:"valueExt,valueLB:0,valueUB:255"`
}
