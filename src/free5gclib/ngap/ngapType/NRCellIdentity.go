package ngapType

import "free5gclib/aper"

// Need to import "free5gclib/aper" if it uses "aper"

type NRCellIdentity struct {
	Value aper.BitString `aper:"sizeLB:36,sizeUB:36"`
}
