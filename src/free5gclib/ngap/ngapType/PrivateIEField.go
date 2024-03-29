package ngapType

// Need to import "free5gclib/aper" if it uses "aper"

type PrivateMessageIEs struct {
	Id          PrivateIEID
	Criticality Criticality
	Value       PrivateMessageIEsValue `aper:"openType,referenceFieldName:Id"`
}

const (
	PrivateMessageIEsPresentNothing int = iota /* No components present */
)

type PrivateMessageIEsValue struct {
	Present int
}
