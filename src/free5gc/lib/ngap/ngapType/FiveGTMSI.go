package ngapType

import "free5gclib/aper"

// Need to import "free5gclib/aper" if it uses "aper"

type FiveGTMSI struct {
	Value aper.OctetString `aper:"sizeLB:4,sizeUB:4"`
}
