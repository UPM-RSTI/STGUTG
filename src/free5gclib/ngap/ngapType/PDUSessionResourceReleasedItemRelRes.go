package ngapType

import "free5gclib/aper"

// Need to import "free5gclib/aper" if it uses "aper"

type PDUSessionResourceReleasedItemRelRes struct {
	PDUSessionID                              PDUSessionID
	PDUSessionResourceReleaseResponseTransfer aper.OctetString
	IEExtensions                              *ProtocolExtensionContainerPDUSessionResourceReleasedItemRelResExtIEs `aper:"optional"`
}
