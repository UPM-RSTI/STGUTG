package ngapType

import "free5gclib/aper"

// Need to import "free5gclib/aper" if it uses "aper"

type PDUSessionResourceHandoverItem struct {
	PDUSessionID            PDUSessionID
	HandoverCommandTransfer aper.OctetString
	IEExtensions            *ProtocolExtensionContainerPDUSessionResourceHandoverItemExtIEs `aper:"optional"`
}
