package ngapType

import "free5gclib/aper"

// Need to import "free5gclib/aper" if it uses "aper"

type EUTRACellIdentity struct {
	Value aper.BitString `aper:"sizeLB:28,sizeUB:28"`
}
