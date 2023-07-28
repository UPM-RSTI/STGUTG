package ngapType

import "free5gclib/aper"

// Need to import "free5gclib/aper" if it uses "aper"

type SerialNumber struct {
	Value aper.BitString `aper:"sizeLB:16,sizeUB:16"`
}
