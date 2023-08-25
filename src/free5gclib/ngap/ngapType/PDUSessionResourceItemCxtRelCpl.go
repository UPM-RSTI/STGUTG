package ngapType

// Need to import "free5gclib/aper" if it uses "aper"

type PDUSessionResourceItemCxtRelCpl struct {
	PDUSessionID PDUSessionID
	IEExtensions *ProtocolExtensionContainerPDUSessionResourceItemCxtRelCplExtIEs `aper:"optional"`
}
