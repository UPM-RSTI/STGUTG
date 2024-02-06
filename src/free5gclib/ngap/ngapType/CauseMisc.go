package ngapType

import "free5gclib/aper"

// Need to import "free5gclib/aper" if it uses "aper"

const (
	CauseMiscPresentControlProcessingOverload             aper.Enumerated = 0
	CauseMiscPresentNotEnoughUserPlaneProcessingResources aper.Enumerated = 1
	CauseMiscPresentHardwareFailure                       aper.Enumerated = 2
	CauseMiscPresentOmIntervention                        aper.Enumerated = 3
	CauseMiscPresentUnknownPLMN                           aper.Enumerated = 4
	CauseMiscPresentUnspecified                           aper.Enumerated = 5
)

type CauseMisc struct {
	Value aper.Enumerated `aper:"valueExt,valueLB:0,valueUB:5"`
}
