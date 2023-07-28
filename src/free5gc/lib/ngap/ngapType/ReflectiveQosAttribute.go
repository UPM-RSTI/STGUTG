package ngapType

import "free5gclib/aper"

// Need to import "free5gclib/aper" if it uses "aper"

const (
	ReflectiveQosAttributePresentSubjectTo aper.Enumerated = 0
)

type ReflectiveQosAttribute struct {
	Value aper.Enumerated `aper:"valueExt,valueLB:0,valueUB:0"`
}
