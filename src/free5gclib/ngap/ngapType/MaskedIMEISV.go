package ngapType

import "free5gclib/aper"

// Need to import "free5gclib/aper" if it uses "aper"

type MaskedIMEISV struct {
	Value aper.BitString `aper:"sizeLB:64,sizeUB:64"`
}
