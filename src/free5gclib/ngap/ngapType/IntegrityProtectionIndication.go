package ngapType

import "free5gclib/aper"

// Need to import "free5gclib/aper" if it uses "aper"

const (
	IntegrityProtectionIndicationPresentRequired  aper.Enumerated = 0
	IntegrityProtectionIndicationPresentPreferred aper.Enumerated = 1
	IntegrityProtectionIndicationPresentNotNeeded aper.Enumerated = 2
)

type IntegrityProtectionIndication struct {
	Value aper.Enumerated `aper:"valueExt,valueLB:0,valueUB:2"`
}
