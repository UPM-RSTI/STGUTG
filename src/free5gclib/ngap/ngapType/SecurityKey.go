package ngapType

import "free5gclib/aper"

// Need to import "free5gclib/aper" if it uses "aper"

type SecurityKey struct {
	Value aper.BitString `aper:"sizeLB:256,sizeUB:256"`
}
