package ngapType

import "free5gclib/aper"

// Need to import "free5gclib/aper" if it uses "aper"

type WarningSecurityInfo struct {
	Value aper.OctetString `aper:"sizeLB:50,sizeUB:50"`
}
