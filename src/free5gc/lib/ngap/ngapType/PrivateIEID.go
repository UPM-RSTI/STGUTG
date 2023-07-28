package ngapType

import "free5gclib/aper"

// Need to import "free5gclib/aper" if it uses "aper"

const (
	PrivateIEIDPresentNothing int = iota /* No components present */
	PrivateIEIDPresentLocal
	PrivateIEIDPresentGlobal
)

type PrivateIEID struct {
	Present int
	Local   *int64 `aper:"valueLB:0,valueUB:65535"`
	Global  *aper.ObjectIdentifier
}
