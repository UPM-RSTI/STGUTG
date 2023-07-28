package ngapType

import "free5gclib/aper"

// Need to import "free5gclib/aper" if it uses "aper"

type PDUSessionResourceFailedToModifyItemModCfm struct {
	PDUSessionID                                           PDUSessionID
	PDUSessionResourceModifyIndicationUnsuccessfulTransfer aper.OctetString
	IEExtensions                                           *ProtocolExtensionContainerPDUSessionResourceFailedToModifyItemModCfmExtIEs `aper:"optional"`
}
