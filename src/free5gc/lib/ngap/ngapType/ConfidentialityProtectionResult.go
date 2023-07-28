package ngapType

import "free5gclib/aper"

// Need to import "free5gclib/aper" if it uses "aper"

const (
	ConfidentialityProtectionResultPresentPerformed    aper.Enumerated = 0
	ConfidentialityProtectionResultPresentNotPerformed aper.Enumerated = 1
)

type ConfidentialityProtectionResult struct {
	Value aper.Enumerated `aper:"valueExt,valueLB:0,valueUB:1"`
}
