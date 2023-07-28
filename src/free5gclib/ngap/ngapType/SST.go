package ngapType

import "free5gclib/aper"

// Need to import "free5gclib/aper" if it uses "aper"

type SST struct {
	Value aper.OctetString `aper:"sizeLB:1,sizeUB:1"`
}
