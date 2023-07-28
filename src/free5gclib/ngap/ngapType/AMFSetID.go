package ngapType

import "free5gclib/aper"

// Need to import "free5gclib/aper" if it uses "aper"

type AMFSetID struct {
	Value aper.BitString `aper:"sizeLB:10,sizeUB:10"`
}
