package ngapType

import "free5gclib/aper"

// Need to import "free5gclib/aper" if it uses "aper"

const (
	RRCInactiveTransitionReportRequestPresentSubsequentStateTransitionReport aper.Enumerated = 0
	RRCInactiveTransitionReportRequestPresentSingleRrcConnectedStateReport   aper.Enumerated = 1
	RRCInactiveTransitionReportRequestPresentCancelReport                    aper.Enumerated = 2
)

type RRCInactiveTransitionReportRequest struct {
	Value aper.Enumerated `aper:"valueExt,valueLB:0,valueUB:2"`
}
