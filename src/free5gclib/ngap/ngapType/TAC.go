package ngapType

import "free5gclib/aper"

// Need to import "free5gclib/aper" if it uses "aper"

type TAC struct {
	Value aper.OctetString `aper:"sizeLB:3,sizeUB:3"`
}
