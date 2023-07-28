package ngapType

import "free5gclib/aper"

// Need to import "free5gclib/aper" if it uses "aper"

const (
	NotificationCausePresentFulfilled    aper.Enumerated = 0
	NotificationCausePresentNotFulfilled aper.Enumerated = 1
)

type NotificationCause struct {
	Value aper.Enumerated `aper:"valueExt,valueLB:0,valueUB:1"`
}
