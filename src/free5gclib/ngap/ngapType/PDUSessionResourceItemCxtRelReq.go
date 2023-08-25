package ngapType

// Need to import "free5gclib/aper" if it uses "aper"

type PDUSessionResourceItemCxtRelReq struct {
	PDUSessionID PDUSessionID
	IEExtensions *ProtocolExtensionContainerPDUSessionResourceItemCxtRelReqExtIEs `aper:"optional"`
}
