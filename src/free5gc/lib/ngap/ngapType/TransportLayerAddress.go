package ngapType

import "free5gclib/aper"

// Need to import "free5gclib/aper" if it uses "aper"

type TransportLayerAddress struct {
	Value aper.BitString `aper:"sizeExt,sizeLB:1,sizeUB:160"`
}
