package ngapType

// Need to import "free5gclib/aper" if it uses "aper"

type RepetitionPeriod struct {
	Value int64 `aper:"valueLB:0,valueUB:131071"`
}
