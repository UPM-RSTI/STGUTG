package ngapType

import "free5gclib/aper"

// Need to import "free5gclib/aper" if it uses "aper"

const (
	ConfidentialityProtectionIndicationPresentRequired  aper.Enumerated = 0
	ConfidentialityProtectionIndicationPresentPreferred aper.Enumerated = 1
	ConfidentialityProtectionIndicationPresentNotNeeded aper.Enumerated = 2
)

type ConfidentialityProtectionIndication struct {
	Value aper.Enumerated `aper:"valueExt,valueLB:0,valueUB:2"`
}
