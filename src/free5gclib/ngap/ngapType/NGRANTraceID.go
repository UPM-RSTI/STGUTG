package ngapType

import "free5gclib/aper"

// Need to import "free5gclib/aper" if it uses "aper"

type NGRANTraceID struct {
	Value aper.OctetString `aper:"sizeLB:8,sizeUB:8"`
}
