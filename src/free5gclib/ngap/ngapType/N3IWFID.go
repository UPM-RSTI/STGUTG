package ngapType

import "free5gclib/aper"

// Need to import "free5gclib/aper" if it uses "aper"

const (
	N3IWFIDPresentNothing int = iota /* No components present */
	N3IWFIDPresentN3IWFID
	N3IWFIDPresentChoiceExtensions
)

type N3IWFID struct {
	Present          int
	N3IWFID          *aper.BitString `aper:"sizeLB:16,sizeUB:16"`
	ChoiceExtensions *ProtocolIESingleContainerN3IWFIDExtIEs
}
