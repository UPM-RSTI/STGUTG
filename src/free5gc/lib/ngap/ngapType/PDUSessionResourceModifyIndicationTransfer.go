package ngapType

// Need to import "free5gclib/aper" if it uses "aper"

type PDUSessionResourceModifyIndicationTransfer struct {
	DLUPTNLInformation *UPTNLInformation                                                           `aper:"valueLB:0,valueUB:2,optional"`
	IEExtensions       *ProtocolExtensionContainerPDUSessionResourceModifyIndicationTransferExtIEs `aper:"optional"`
}
