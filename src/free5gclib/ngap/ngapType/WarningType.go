package ngapType

import "free5gclib/aper"

// Need to import "free5gclib/aper" if it uses "aper"

type WarningType struct {
	Value aper.OctetString `aper:"sizeLB:2,sizeUB:2"`
}
