package ngapType

import "free5gclib/aper"

// Need to import "free5gclib/aper" if it uses "aper"

type AMFRegionID struct {
	Value aper.BitString `aper:"sizeLB:8,sizeUB:8"`
}
