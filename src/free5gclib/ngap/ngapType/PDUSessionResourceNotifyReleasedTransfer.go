package ngapType

// Need to import "free5gclib/aper" if it uses "aper"

type PDUSessionResourceNotifyReleasedTransfer struct {
	Cause        Cause                                                                     `aper:"valueLB:0,valueUB:5"`
	IEExtensions *ProtocolExtensionContainerPDUSessionResourceNotifyReleasedTransferExtIEs `aper:"optional"`
}
