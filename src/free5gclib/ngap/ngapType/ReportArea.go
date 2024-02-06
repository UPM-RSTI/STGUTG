package ngapType

import "free5gclib/aper"

// Need to import "free5gclib/aper" if it uses "aper"

const (
	ReportAreaPresentCell aper.Enumerated = 0
)

type ReportArea struct {
	Value aper.Enumerated `aper:"valueExt,valueLB:0,valueUB:0"`
}
