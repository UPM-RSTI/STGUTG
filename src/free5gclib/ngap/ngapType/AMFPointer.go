package ngapType

import "free5gclib/aper"

// Need to import "free5gclib/aper" if it uses "aper"

type AMFPointer struct {
	Value aper.BitString `aper:"sizeLB:6,sizeUB:6"`
}
