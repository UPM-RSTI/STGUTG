package ngapType

import "free5gclib/aper"

// Need to import "free5gclib/aper" if it uses "aper"

const (
	OverloadActionPresentRejectNonEmergencyMoDt                                    aper.Enumerated = 0
	OverloadActionPresentRejectRrcCrSignalling                                     aper.Enumerated = 1
	OverloadActionPresentPermitEmergencySessionsAndMobileTerminatedServicesOnly    aper.Enumerated = 2
	OverloadActionPresentPermitHighPrioritySessionsAndMobileTerminatedServicesOnly aper.Enumerated = 3
)

type OverloadAction struct {
	Value aper.Enumerated `aper:"valueExt,valueLB:0,valueUB:3"`
}
