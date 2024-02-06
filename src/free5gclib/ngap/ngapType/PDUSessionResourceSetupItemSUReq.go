package ngapType

import "free5gclib/aper"

// Need to import "free5gclib/aper" if it uses "aper"

type PDUSessionResourceSetupItemSUReq struct {
	PDUSessionID                           PDUSessionID
	PDUSessionNASPDU                       *NASPDU `aper:"optional"`
	SNSSAI                                 SNSSAI  `aper:"valueExt"`
	PDUSessionResourceSetupRequestTransfer aper.OctetString
	IEExtensions                           *ProtocolExtensionContainerPDUSessionResourceSetupItemSUReqExtIEs `aper:"optional"`
}
