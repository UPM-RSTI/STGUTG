package ngapType

import "free5gclib/aper"

// Need to import "free5gclib/aper" if it uses "aper"

type PDUSessionResourceItemHORqd struct {
	PDUSessionID             PDUSessionID
	HandoverRequiredTransfer aper.OctetString
	IEExtensions             *ProtocolExtensionContainerPDUSessionResourceItemHORqdExtIEs `aper:"optional"`
}
