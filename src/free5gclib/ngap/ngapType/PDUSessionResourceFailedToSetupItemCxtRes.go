package ngapType

import "free5gclib/aper"

// Need to import "free5gclib/aper" if it uses "aper"

type PDUSessionResourceFailedToSetupItemCxtRes struct {
	PDUSessionID                                PDUSessionID
	PDUSessionResourceSetupUnsuccessfulTransfer aper.OctetString
	IEExtensions                                *ProtocolExtensionContainerPDUSessionResourceFailedToSetupItemCxtResExtIEs `aper:"optional"`
}
