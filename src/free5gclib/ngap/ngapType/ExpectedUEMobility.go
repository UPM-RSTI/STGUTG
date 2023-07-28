package ngapType

import "free5gclib/aper"

// Need to import "free5gclib/aper" if it uses "aper"

const (
	ExpectedUEMobilityPresentStationary aper.Enumerated = 0
	ExpectedUEMobilityPresentMobile     aper.Enumerated = 1
)

type ExpectedUEMobility struct {
	Value aper.Enumerated `aper:"valueExt,valueLB:0,valueUB:1"`
}
