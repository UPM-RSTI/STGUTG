package ngapType

import "free5gclib/aper"

// Need to import "free5gclib/aper" if it uses "aper"

type PDUSessionResourceToReleaseItemRelCmd struct {
	PDUSessionID                             PDUSessionID
	PDUSessionResourceReleaseCommandTransfer aper.OctetString
	IEExtensions                             *ProtocolExtensionContainerPDUSessionResourceToReleaseItemRelCmdExtIEs `aper:"optional"`
}
